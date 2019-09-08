package stunning_palm_tree

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/hzhyvinskyi/stunneni-palm-tree/api"
	"github.com/hzhyvinskyi/stunneni-palm-tree/api/dal"
	"github.com/hzhyvinskyi/stunneni-palm-tree/api/errors"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateVideo(ctx context.Context, input NewVideo) (*Video, error) {
	newVideo := api.Video{
		URL:		input.URL,
		Name:		input.Name,
		CreatedAt:	time.Now().UTC(),
	}

	rows, err := dal.LogAndQuery(r.db, "INSERT INTO videos (name, url, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id",
		input.Name, input.URL, input.UserID, newVideo.CreatedAt)
	defer rows.Close()

	if err != nil || !rows.Next() {
		return api.Video{}, err
	}

	if err := rows.Scan(&newVideo.ID); err != nil {
		errors.DebugPrintf(err)
		if errors.IsForeignKeyError(err) {
		  return api.Video{}, errors.UserNotExist
		}
		return api.Video{}, errors.InternalServerError
	  }
	  
	  return newVideo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Videos(ctx context.Context, limit *int, offset *int) ([]*Video, error) {
	var (
		video	api.Video
		videos	[]api.Video
	)

	rows, err := dal.LogAndQuery(r.db, "SELECT id, name, url, created_at, user_id FROM videos ORDER BY id DESC LIMIT $1 OFFSET $2", limit, offset)
	defer rows.Close()

	if err != nil {
		errors.DebugPrintf(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		if err := rows.Scan(&video.ID, &video.Name, &video.URL, &video.CreatedAt, &video.UserID); err != nil {
			errors.DebugPrintf(err)
			return nil, errors.InternalServerError
		}
		videos = append(videos, video)
	}

	return videos, nil
}

type videoResolver struct{ *Resolver }

func (r *videoResolver) User(ctx context.Context, obj *api.Video) (api.User, error) {
	rows, _ := dal.LogAndQuery(r.db, "SELECT id, name, email FROM users WHERE id = $1", obj.UserID)
	defer rows.Close()

	if !rows.Next() {
		return api.User{}, nil
	}

	var user api.User
	if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		errors.DebugPrintf(err)
		return api.User{}, errors.InternalServerError
	}

	return user, nil
}
