package builder

type CarBuilder struct {

}

func (c *CarBuilder) SetWheels() BuildProcess {
	panic("implement me")
}

func (c *CarBuilder) SetSeats() BuildProcess {
	panic("implement me")
}

func (c *CarBuilder) SetStructure() BuildProcess {
	panic("implement me")
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	panic("implement me")
}

func (c *CarBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}

var _ BuildProcess = &CarBuilder{}
