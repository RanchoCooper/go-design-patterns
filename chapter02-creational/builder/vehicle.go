package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {

}

func (f *ManufacturingDirector) Construct() {
	// implement me
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	// implement me
}

type VehicleProduct struct {
	Wheels int
	Seats int
	Structure string
}