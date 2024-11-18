package textbasics

import (
	"bufio"
	"fmt"
	"os"
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

func ParseFileWithRegex(filename string) ([]Information, error) {
	var result []Information
	file, err := os.Open(filename)
	if err != nil {
		return []Information{}, fmt.Errorf("unable to open the file %q", filename)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// get out once we get to the end of the line
			// fmt.Printf("End of line reached at line: %q\n", line)
			break
		}
		// now parse the line using the function defined above
		info, err := ParseRegex(line)
		if err != nil {
			return []Information{}, err
		}
		// now we got the info, append it to result
		result = append(result, info)
	}

	return result, nil
}
