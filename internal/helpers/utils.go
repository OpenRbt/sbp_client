package helpers

import "strings"

func SplitCommaSeparatedStr(commaSeparated string) (result []string) {
	for _, item := range strings.Split(commaSeparated, ",") {
		item = strings.TrimSpace(item)
		if item != "" {
			result = append(result, item)
		}
	}
	return result
}
