package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[=~*-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	howMany := 0
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, index := range lines {
		if re.MatchString(index) {
			howMany++
		}
	}
	return howMany
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`.end-of-line\d+`)
	return re.ReplaceAllString(text, " ")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)

	for i, content := range lines {
		match := re.FindStringSubmatch(content)
		if len(match) > 0 {
			lines[i] = "[USR] " + match[1] + " " + content
		}

	}
	return lines
}
