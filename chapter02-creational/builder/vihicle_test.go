package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	// car
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.Build()
	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}
	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	// bike
}