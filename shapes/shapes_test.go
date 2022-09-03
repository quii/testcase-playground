package shapes

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

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func TestShapes(t *testing.T) {
	s := testcase.NewSpec(t)

	s.Describe("Rectangle", func(s *testcase.Spec) {
		var (
			width   = testcase.LetValue[float64](s, 12.0)
			height  = testcase.LetValue[float64](s, 6.0)
			subject = func(t *testcase.T) Rectangle {
				return Rectangle{Width: width.Get(t), Height: height.Get(t)}
			}
		)

		s.When(".Area", func(s *testcase.Spec) {
			act := func(t *testcase.T) float64 {
				return subject(t).Area()
			}
			s.Then("it returns its area", func(t *testcase.T) {
				assert.Equal(t, 72, act(t))
			})
		})

		s.When(".Perimeter", func(s *testcase.Spec) {
			act := func(t *testcase.T) float64 {
				return subject(t).Perimeter()
			}
			s.Then("it returns its perimeter", func(t *testcase.T) {
				assert.Equal(t, 36, act(t))
			})
		})
	})

}
