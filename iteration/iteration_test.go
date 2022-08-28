package iteration

import (
	"testing"

	"github.com/adamluzsi/testcase"
)

func TestRepeat(t *testing.T) {
	s := testcase.NewSpec(t)

	// figure out inputs and dependencies
	var (
		character = testcase.Let[string](s, nil)
		act       = func(t *testcase.T) string {
			return Repeat(character.Get(t))
		}
	)

	s.When("you provide a single character", func(s *testcase.Spec) {
		character.LetValue(s, "a")
		s.Then("it repeats it 5 times", func(t *testcase.T) {
			t.Must.Equal("aaaaa", act(t))
		})
	})
}

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}
