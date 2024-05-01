package builder

// Car structure
type Car struct {
	Make  string
	Model string
	Color string
	Year  int
}

// CarBuilder is the interface that will be implemented by concrete builders
type CarBuilder interface {
	SetMake(make string) CarBuilder
	SetModel(model string) CarBuilder
	SetColor(color string) CarBuilder
	SetYear(year int) CarBuilder
	Build() Car
}

// carBuilder is a concrete builder for building a Car
type carBuilder struct {
	make  string
	model string
	color string
	year  int
}

// SetMake sets the make of the car
func (c *carBuilder) SetMake(make string) CarBuilder {
	c.make = make
	return c
}

// SetModel sets the model of the car
func (c *carBuilder) SetModel(model string) CarBuilder {
	c.model = model
	return c
}

// SetColor sets the color of the car
func (c *carBuilder) SetColor(color string) CarBuilder {
	c.color = color
	return c
}

// SetYear sets the year of the car
func (c *carBuilder) SetYear(year int) CarBuilder {
	c.year = year
	return c
}

// Build constructs the final Car object
func (c *carBuilder) Build() Car {
	return Car{
		Make:  c.make,
		Model: c.model,
		Color: c.color,
		Year:  c.year,
	}
}

// NewCarBuilder returns a new instance of a carBuilder
func NewCarBuilder() CarBuilder {
	return &carBuilder{}
}
