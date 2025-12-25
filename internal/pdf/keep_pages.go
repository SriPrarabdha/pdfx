package pdf

import (
	"fmt"
	"regexp"
	"strings"
)

// ParseKeepPages parses page ranges like "1,3-5,7"
func ParseKeepPages(pages string) ([]string, error) {
	pages = strings.TrimSpace(pages)
	if pages == "" {
		return nil, fmt.Errorf("page range cannot be empty")
	}

	re := regexp.MustCompile(`^(\d+(-\d+)?)(,(\d+(-\d+)?))*$`)
	if !re.MatchString(pages) {
		return nil, fmt.Errorf("invalid page range format: %s", pages)
	}

	return strings.Split(pages, ","), nil
}
