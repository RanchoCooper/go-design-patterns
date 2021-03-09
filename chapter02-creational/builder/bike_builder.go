package builder

type BikeBuilder struct {

}

func (b *BikeBuilder) SetWheels() BuildProcess {
	panic("implement me")
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	panic("implement me")
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	panic("implement me")
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	panic("implement me")
}

func (b *BikeBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}

var _ BuildProcess = &BikeBuilder{}
