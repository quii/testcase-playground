package hello_world

import (
	"testing"

	"github.com/adamluzsi/testcase"
)

/* notes and questions

the docs show creating a "tc" with testcase.NewT, i _think_ the intention of a tc.T is it's tools to do assertions, but that's not clear from the readme initially

really, you want to make a spec with testcase.NewSpec

testcase values arrange, act, assert convention. Your "act" is the subject doing something, (in plainer terms, the thing you're testing)

The idea in a given suite the "act" is always the same, for simplicity and consistency. The only things that change are how you "arrange" the inputs (including dependencies) and what you expect, i.e the assertions.

So:

Arrange: Change these to test different things
Act: This should stay the same, and if it's difficult to structure this, rethink your design or what you're testing
Assert: Depends on the arrangement

Adam: Test output in intellij is so verbose, is there a way to make it quiet?
*/

func TestName(t *testing.T) {
	s := testcase.NewSpec(t)

	var (
		// at the top of the suite, organise your inputs and dependencies in _one_ place
		name = testcase.Let[string](s, nil)

		// finally show how to "act", make the subject, plug in variables
		act = func(t *testcase.T) string {
			return Greet(name.Get(t))
		}
	)

	s.When("a name is supplied", func(s *testcase.Spec) {
		// arrange your stuff
		name.LetValue(s, "Chris")

		s.Test("the greeting should be 'Hello, Chris'", func(t *testcase.T) {
			got := act(t)
			t.Must.Equal("Hello, Chris", got)
		})
	})

	s.When("the name is empty", func(s *testcase.Spec) {
		name.LetValue(s, "")

		s.Test("we should greet the world", func(t *testcase.T) {
			t.Must.Equal("Hello, World", act(t))
		})
	})
}

func Greet(name string) string {
	if name == "" {
		return "Hello, World"
	}
	return "Hello, " + name
}
