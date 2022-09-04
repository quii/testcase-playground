package facts_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testcase-playground/facts"
	"testing"

	"github.com/adamluzsi/testcase"
	"github.com/adamluzsi/testcase/assert"
)

func TestFactClient(t *testing.T) {
	s := testcase.NewSpec(t)

	var (
		url = testcase.LetValue[string](s, "")
		act = func(t *testcase.T) string {
			return facts.Client{BaseURL: url.Get(t)}.Get()
		}
	)

	s.Describe("Client", func(s *testcase.Spec) {
		var (
			status       = testcase.Let[int](s, nil)
			responseBody = testcase.Let[string](s, nil)
		)
		s.Before(func(t *testcase.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(status.Get(t))
				fmt.Fprint(writer, responseBody.Get(t))
			}))
			url.Set(t, svr.URL)
			t.Cleanup(svr.Close)
		})

		s.When("the server is working", func(s *testcase.Spec) {
			status.LetValue(s, http.StatusOK)

			s.And("it has a response body", func(s *testcase.Spec) {
				responseBody.LetValue(s, `[
  {
    "fact": "Four billion pounds of watermelon were grown in the United States in 1999"
  }
]`)

				s.Then("it returns the body", func(t *testcase.T) {
					assert.Equal(t, "Four billion pounds of watermelon were grown in the United States in 1999", act(t))
				})
			})

			s.And("it returns invalid JSON", func(s *testcase.Spec) {
				responseBody.LetValue(s, "bad json")
				ItReturnsAnError(s, act)
			})

			s.And("valid JSON with no facts", func(s *testcase.Spec) {
				responseBody.LetValue(s, "[]")
				ItReturnsAnError(s, act)
			})
		})

		s.When("the server returns a non 200", func(s *testcase.Spec) {
			status.LetValue(s, http.StatusTeapot)
			responseBody.LetValue(s, "Error")
			ItReturnsAnError(s, act)
		})
	})
}

func ItReturnsAnError(s *testcase.Spec, act func(t *testcase.T) string) {
	s.Test("it returns oh no", func(t *testcase.T) {
		assert.Equal(t, facts.ErrFactAPIProblem, act(t))
	})
}
