package shapes

import (
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

type Square struct {
	SideLength float64
}

func (s Square) Area() float64 {
	return s.SideLength * s.SideLength
}

func (s Square) Perimeter() float64 {
	return 4 * s.SideLength
}

func TestShapes(t *testing.T) {
	s := testcase.NewSpec(t)

	s.Describe("Rectangle", func(s *testcase.Spec) {
		var (
			width   = testcase.LetValue[float64](s, 12.0)
			height  = testcase.LetValue[float64](s, 6.0)
			subject = testcase.Let[Rectangle](s, func(t *testcase.T) Rectangle {
				return Rectangle{Width: width.Get(t), Height: height.Get(t)}
			})
		)

		s.When(".Area", func(s *testcase.Spec) {
			AreaSpec(s, subject, 72)
		})

		s.When(".Perimeter", func(s *testcase.Spec) {
			PerimeterSpec(s, subject, 36)
		})
	})

	s.Describe("Square", func(s *testcase.Spec) {
		var (
			sideLength = testcase.LetValue[float64](s, 5)
			subject    = testcase.Let[Square](s, func(t *testcase.T) Square {
				return Square{SideLength: sideLength.Get(t)}
			})
		)

		s.When(".Area", func(s *testcase.Spec) {
			AreaSpec(s, subject, 25)
		})

		s.When(".Perimeter", func(s *testcase.Spec) {
			PerimeterSpec(s, subject, 20)
		})
	})

}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func AreaSpec[T Shape](s *testcase.Spec, shape testcase.Var[T], want float64) {
	var act = func(t *testcase.T) float64 {
		return shape.Get(t).Area()
	}

	s.Then(`it has an area of`, func(t *testcase.T) {
		assert.Must(t).Equal(want, act(t))
	})
}

func PerimeterSpec[T Shape](s *testcase.Spec, shape testcase.Var[T], want float64) {
	var act = func(t *testcase.T) float64 {
		return shape.Get(t).Perimeter()
	}

	s.Then(`it has a perimeter of`, func(t *testcase.T) {
		assert.Must(t).Equal(want, act(t))
	})
}
