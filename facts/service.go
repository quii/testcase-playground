package facts

import (
	"fmt"
	"strings"
)

type Repository interface {
	Get() string
}

type Service struct {
	Repository Repository
}

func (s Service) RandomFact() string {
	fact := s.Repository.Get()

	for word, emoji := range emojiMappings {
		if strings.Contains(fact, word) {
			return fmt.Sprintf("%s %s", fact, emoji)
		}
	}

	return fact
}

var (
	emojiMappings = map[string]string{
		"Dog":     "🐶",
		"England": "🏴󠁧󠁢󠁥󠁮󠁧󠁿",
	}
)
