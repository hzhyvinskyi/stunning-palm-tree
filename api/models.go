package api

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/hzhyvinskyi/stunning-palm-tree/api/errors"
)

type Video struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserID      int       `json:"-"`
	URL         string    `json:"url"`
	CreatedAt   time.Time `json:"createdAt"`
	Related     []Video
}

// MarshalID redefines base ID type to use and id from external library
func MarshalID(id int) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Itoa(id))
	})
}

func UnmarshalID(v interface{}) (int, error) {
	id, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("ids must be string")
	}
	i, err := strconv.Atoi(id)
	return i, err
}

func MarshalTimestamp(t time.Time) graphql.Marshaler {
	timestamp := t.Unix() * 1000
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}

func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int); !ok {
		return time.Time{}, errors.TimeStampError
	}
	return time.Unix(int64(tmpStr), 0), nil
}
