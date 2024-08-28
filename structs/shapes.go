package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter takes width and height of a rectangle and returns the perimeter
func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

// Area takes width and height of a rectangle and returns the area
func Area(width, height float64) float64 {
	return width * height
}

// Takes the width and height of a rectangle and returns the area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Takes radius of a circle and returns the area
func (c Circle) Area() float64 {
	return 3.141592653589793 * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
