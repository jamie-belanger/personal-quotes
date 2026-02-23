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
	return
}

/*
	Sanitizes the given HTML input to ensure users can't insert anything
	malicious.

	# Parameters

	- input string = HTML to sanitize

	# Returns

	- string = safe HTML

	- error = if anything went wrong

	# Remarks

	Yes I know there are other sanitizers out there, but I really only want to support
	a very small subset of HTML with this application. That means <b>, <i>, and <br>
	... nothing else is valid here. At least for now.
*/
func sanitizeHTML(input string) (output string, err error) {
	if len(input) > 0 {
		output = input

		// TODO: implement this
	}

	return
}

