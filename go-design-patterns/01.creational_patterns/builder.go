package creational

// 製造工程を表すインターフェース
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// ビルダーの受け入れは、ManufacturingDirector ディレクタ変数が担当する
// Constructメソッドを持ち、Manufacturingに格納されているBuilderを使用し、必要な手順を再現する
// SetBuilderメソッドにより、Manufacturingディレクターで使用されているBuilderを変更することができる
type ManufactureingDirector struct {
	builder BuildProcess
}

func (f *ManufactureingDirector) Construct() {
	f.builder.SetSeats().SetWheels().SetStructure()
}

func (f *ManufactureingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// 製造の結果、最終的に取得するオブジェクトの構造体
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// CarBuilder は、車の製造を担う
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

// BikeBuilderはバイクの製造を担う
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
