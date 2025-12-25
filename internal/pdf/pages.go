package pdf

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// ParseDeletePages parses user input like "1,3-5,n,7-n"
// and returns a set of pages to DELETE.
func ParseDeletePages(input string, total int) (map[int]bool, error) {
	result := make(map[int]bool)

	parts := strings.Split(input, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)

		// range case
		if strings.Contains(part, "-") {
			ends := strings.Split(part, "-")
			if len(ends) != 2 {
				return nil, fmt.Errorf("invalid range: %s", part)
			}

			start, err := parsePage(ends[0], total)
			if err != nil {
				return nil, err
			}
			end, err := parsePage(ends[1], total)
			if err != nil {
				return nil, err
			}

			if start > end {
				return nil, fmt.Errorf("invalid range %s: start > end", part)
			}

			for i := start; i <= end; i++ {
				result[i] = true
			}
			continue
		}

		// single page
		p, err := parsePage(part, total)
		if err != nil {
			return nil, err
		}
		result[p] = true
	}

	return result, nil
}

// parsePage converts "1", "n" into page number
func parsePage(s string, total int) (int, error) {
	if s == "n" {
		return total, nil
	}

	p, err := strconv.Atoi(s)
	if err != nil || p < 1 || p > total {
		return 0, fmt.Errorf("invalid page: %s", s)
	}
	return p, nil
}

// ComplementPages returns pages to KEEP
func ComplementPages(delete map[int]bool, total int) []string {
	var keep []int
	for i := 1; i <= total; i++ {
		if !delete[i] {
			keep = append(keep, i)
		}
	}

	sort.Ints(keep)
	return compressRanges(keep)
}

// compressRanges converts [1,2,3,5,6] → ["1-3","5-6"]
func compressRanges(pages []int) []string {
	var out []string
	if len(pages) == 0 {
		return out
	}

	start := pages[0]
	prev := pages[0]

	for _, p := range pages[1:] {
		if p == prev+1 {
			prev = p
			continue
		}
		out = append(out, formatRange(start, prev))
		start = p
		prev = p
	}
	out = append(out, formatRange(start, prev))
	return out
}

func formatRange(a, b int) string {
	if a == b {
		return strconv.Itoa(a)
	}
	return fmt.Sprintf("%d-%d", a, b)
}
