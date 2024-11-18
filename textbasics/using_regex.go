package textbasics

import (
	"fmt"
	"regexp"
	"strconv"
)

// John is 23 years old and lives in Sacramento
// Doe is 21 years old and lives in Ohio

var infoRegex *regexp.Regexp = regexp.MustCompile(`([A-Z][a-z]*) is (\d+) years old and lives in ([A-Z][a-z]*)`)

type Information struct {
	name     string
	age      int
	location string
}

// attempts to parse a regex file and returns an error if unsuccessful.
// do this for one line and call this multiple times when parsing a large document.
func ParseRegex(line string) (Information, error) {
	var result Information
	matches := infoRegex.FindStringSubmatch(line) // has groups+1 elements because the whole string matched is the first element.
	if matches == nil {
		// %q allows for strings to have double quotes when printed.
		return Information{}, fmt.Errorf("unable to process line: %q", line)
	}
	result.name = matches[1]
	result.age, _ = strconv.Atoi(matches[2])
	result.location = matches[3]
	return result, nil
}
