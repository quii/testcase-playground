package facts_test

import (
	"testcase-playground/facts"
	"testing"

	"github.com/adamluzsi/testcase"
	"github.com/adamluzsi/testcase/assert"
)

type StubRepo struct {
	CannedFact string
}

func (s StubRepo) Get() string {
	return s.CannedFact
}

func TestFactService(t *testing.T) {
	s := testcase.NewSpec(t)

	s.Describe("FactService", func(s *testcase.Spec) {
		var (
			stubRepo = testcase.Let[StubRepo](s, nil)
			act      = func(t *testcase.T) string {
				return facts.Service{Repository: stubRepo.Get(t)}.RandomFact()
			}
		)
		s.When(".Fact()", func(s *testcase.Spec) {
			s.And("a fact repository supplies a fact", func(s *testcase.Spec) {
				stubRepo.LetValue(s, StubRepo{CannedFact: "Cats are nice"})

				s.Then("it is returned", func(t *testcase.T) {
					assert.Equal(t, "Cats are nice", act(t))
				})
			})
			s.And("a fact about dogs is supplied", func(s *testcase.Spec) {
				stubRepo.LetValue(s, StubRepo{CannedFact: "Dogs are ok"})

				s.Then("it is returned with a dog emojii suffixed", func(t *testcase.T) {
					assert.Equal(t, "Dogs are ok 🐶", act(t))
				})
			})
			s.And("a fact about England is supplied", func(s *testcase.Spec) {
				stubRepo.LetValue(s, StubRepo{CannedFact: "In England, it is frequently cloudy"})

				s.Then("it is returned with the English flag suffixed", func(t *testcase.T) {
					assert.Equal(t, "In England, it is frequently cloudy 🏴󠁧󠁢󠁥󠁮󠁧󠁿", act(t))
				})
			})
		})
	})
}
