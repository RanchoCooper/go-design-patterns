package builder

/**
 * @author Rancho
 * @date 2021/12/26
 */

type BuildProcess interface {
    SetWheels() BuildProcess
    SetSeats() BuildProcess
    SetStructure() BuildProcess
    GetVehicle() VehicleProduct
}

type VehicleProduct struct {
    Wheels    int
    Seats     int
    Structure string
}

type ManufacturingDirector struct {
    builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
    f.builder.SetWheels().SetSeats().SetStructure()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
    f.builder = b
}

type CarBuilder struct {
    v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
    c.v.Wheels = 4
    return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
    c.v.Seats = 5
    return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
    c.v.Structure = "Car"
    return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
    return c.v
}

var _ BuildProcess = &CarBuilder{}

type BikeBuilder struct {
    v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
    b.v.Wheels = 2
    return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
    b.v.Seats = 2
    return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
    b.v.Structure = "Motorbike"
    return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
    return b.v
}

var _ BuildProcess = &BikeBuilder{}

type BusBuilder struct {
    v VehicleProduct
}

func (b *BusBuilder) SetWheels() BuildProcess {
    b.v.Wheels = 8
    return b
}

func (b *BusBuilder) SetSeats() BuildProcess {
    b.v.Seats = 30
    return b
}

func (b *BusBuilder) SetStructure() BuildProcess {
    b.v.Structure = "Bus"
    return b
}

func (b *BusBuilder) GetVehicle() VehicleProduct {
    return b.v
}

var _ BuildProcess = &BusBuilder{}
