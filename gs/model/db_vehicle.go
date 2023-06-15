package model

type DbVehicle struct {
	VehicleId uint32  // 载具Id
	OwnerUid  uint32  // 所有者Id
	Pos       *Vector // 位置
	Rot       *Vector // 旋转
}

func NewDbVehicle(vehicleId uint32, ownerId uint32, pos *Vector, rot *Vector) *DbVehicle {
	return &DbVehicle{
		VehicleId: vehicleId,
		OwnerUid:  ownerId,
		Pos:       pos,
		Rot:       rot,
	}
}
