package handlers

import (
	"errors"
	"strconv"
)

/*
	Tries to parse the given path value as an integer identifier
*/
func tryParseId(pathValue string) (id int64, err error) {
	if "" == pathValue {
		err = errors.New("Identifier missing")
		return
	}

	id, err = strconv.ParseInt(pathValue, 10, 64)
	if id < 0 {
		err = errors.New("Identifiers cannot be negative")
	}
	return
}
