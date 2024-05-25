package structs

type Shape interface {
	CalculateArea() float64
	CalculatePerimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func CalculatePerimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

func CalculateArea(r Rectangle) float64 {
	return r.Width * r.Height
}

func (r Rectangle) CalculatePerimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) CalculateArea() float64 {
	return r.Width * r.Height
}

func (c Circle) CalculatePerimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (c Circle) CalculateArea() float64 {
	return 3.14 * c.Radius * c.Radius
}
