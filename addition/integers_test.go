package addition

import (
	"testing"

	"github.com/adamluzsi/testcase"
)

func TestName(t *testing.T) {
	s := testcase.NewSpec(t)

	// arrange
	var (
		firstInteger  = testcase.Let[int](s, nil)
		secondInteger = testcase.Let[int](s, nil)

		act = func(t *testcase.T) int {
			return Add(firstInteger.Get(t), secondInteger.Get(t))
		}
	)

	s.When("two numbers", func(s *testcase.Spec) {
		firstInteger.LetValue(s, 2)
		secondInteger.LetValue(s, 4)

		s.Then("you add them together", func(t *testcase.T) {
			t.Must.Equal(6, act(t))
		})
	})

	s.When("some different numbers", func(s *testcase.Spec) {
		firstInteger.LetValue(s, 2)
		secondInteger.LetValue(s, 2)

		s.Then("you add them together", func(t *testcase.T) {
			t.Must.Equal(4, act(t))
		})
	})

}

func Add(x int, y int) int {
	return x + y
}
