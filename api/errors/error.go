package errors

import (
	"errors"
	"fmt"
	"log"
	"runtime"

	"github.com/lib/pq"
)

var (
	ServerError			= generateError("Something went wrong! Please try again")
	UserNotExistError	= generateError("User not exists")
	UnauthorizedError	= generateError("You aren't authorized to perform this action")
	TimeStampError		= generateError("Time should be a UNIX timestamp")
	InternalServerError	= generateError("Internal Server Error")
)

func generateError(err string) error {
	return errors.New(err)
}

func IsForeignKeyError(err error) bool {
	pgErr := err.(*pq.Error)
	if pgErr.Code == "23503" {
		return true
	}
	return false
}

func DebugPrintf(err_ error, args ...interface{}) string {
	programCounter, file, line := runtime.Caller(1)
	fn := runtime.FuncForPC(programCounter)
	msg := fmt.Sprintf("[%s: %s %d] %s %s", file, fn.Name, line, err_, args)
	log.Println(msg)
	return msg
}
