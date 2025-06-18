package strings

import (
	"regexp"
	"strings"
)

// ExtractChineseCharacters 只提取汉字
func ExtractChineseCharacters(str string) string {
	re := regexp.MustCompile(`[\p{Han}]+`)
	s := re.FindAllString(str, -1)
	return strings.Join(s, "")
}
