package tools

import (
	"fmt"
	"regexp"
)

const URLRegexp = `[(http(s)?):\/\/(www\.)?a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)$`

func IsURL(url string) bool {
	matched, err := regexp.MatchString(URLRegexp, url)
	if err != nil {
		fmt.Printf("Error while validating url %s", url)
		return false
	}
	return matched
}
