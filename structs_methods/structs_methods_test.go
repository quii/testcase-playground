package structs_methods

import (
	"math"
	"testing"

	"github.com/adamluzsi/testcase"
	"github.com/adamluzsi/testcase/assert"
)

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func TestArea(t *testing.T) {
	s := testcase.NewSpec(t)

	s.When("rectangles", func(s *testcase.Spec) {
		var (
			width  = testcase.Let[float64](s, nil)
			height = testcase.Let[float64](s, nil)

			act = func(t *testcase.T) float64 {
				return Rectangle{Width: width.Get(t), Height: height.Get(t)}.Area()
			}
		)

		s.When("that have some height and width", func(s *testcase.Spec) {
			width.LetValue(s, 12.0)
			height.LetValue(s, 6.0)

			s.Then("it has an area", func(t *testcase.T) {
				got := act(t)
				assert.Equal(t, 72.0, got)
			})
		})
	})

	s.When("circles", func(s *testcase.Spec) {
		var (
			radius = testcase.Let[float64](s, nil)

			act = func(t *testcase.T) float64 {
				return Circle{Radius: radius.Get(t)}.Area()
			}
		)

		s.When("and some kind of radius", func(s *testcase.Spec) {
			radius.LetValue(s, 10)

			s.Then("it has an area", func(t *testcase.T) {
				got := act(t)
				assert.Equal(t, 314.1592653589793, got)
			})
		})
	})

}
