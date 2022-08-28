package arrays_and_slices

import (
	"testing"

	"github.com/adamluzsi/testcase"
)

func TestName(t *testing.T) {
	s := testcase.NewSpec(t)

	// lay out our inputs and dependencies in a var block
	var (
		numbers = testcase.Let[[]int](s, func(t *testcase.T) []int {
			return []int{}
		})

		act = func(t *testcase.T) int {
			return Sum(numbers.Get(t))
		}
	)

	s.When("there's a array of numbers", func(s *testcase.Spec) {
		// arrange
		numbers.Let(s, func(t *testcase.T) []int {
			return []int{1, 2, 3, 4, 5}
		})

		//assert
		s.Then("it sums them all up into a single value", func(t *testcase.T) {
			t.Must.Equal(15, act(t))
		})
	})

	s.When("a different array of numbers", func(s *testcase.Spec) {
		// arrange
		numbers.Let(s, func(t *testcase.T) []int {
			return []int{1, 2}
		})

		//assert
		s.Then("it sums them all up into a single value", func(t *testcase.T) {
			t.Must.Equal(3, act(t))
		})
	})

	s.When("a slice, with more than 5 numbers", func(s *testcase.Spec) {
		numbers.Let(s, func(t *testcase.T) []int {
			return []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
		})

		s.Then("it sums them all up into a single value", func(t *testcase.T) {
			t.Must.Equal(11, act(t))
		})
	})
}

func Sum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}
