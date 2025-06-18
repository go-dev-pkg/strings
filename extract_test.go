package strings_test

import (
	"testing"

	"github.com/go-dev-pkg/strings"
)

func TestExtractChineseCharacters(t *testing.T) {
	str := `s阿西巴·空你几娃·1.one*你@#$%^&*(_+~2（张C!~にほん/한국어巴扎黑`
	s := strings.ExtractChineseCharacters(str)
	if s != "阿西巴空你几娃你张巴扎黑" {
		t.Error("ExtractChineseCharacters")
		return
	}
}
