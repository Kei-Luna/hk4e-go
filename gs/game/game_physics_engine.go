package game

import (
	"math"

	"hk4e/common/constant"
	"hk4e/pkg/alg"
	"hk4e/pkg/logger"
)

const (
	AVATAR_RADIUS = 0.4
	AVATAR_HEIGHT = 2.0
)

// RigidBody 刚体
type RigidBody struct {
	entityId       uint32       // 子弹实体id
	avatarEntityId uint32       // 子弹发射者角色实体id
	sceneId        uint32       // 子弹所在场景id
	position       *alg.Vector3 // 坐标
	velocity       *alg.Vector3 // 速度
	drag           float32      // 阻力参数
	mass           float32      // 质量
}

// PhysicsEngine 物理引擎
type PhysicsEngine struct {
	rigidBodyMap     map[uint32]*RigidBody      // 刚体集合
	acc              float32                    // 重力加速度
	lastUpdateTime   int64                      // 上一次更新时间
	sceneBlockAoiMap map[uint32]*alg.AoiManager // 全局各场景地图的aoi管理器
	world            *World                     // 世界对象
}

func (w *World) NewPhysicsEngine(sceneBlockAoiMap map[uint32]*alg.AoiManager) {
	w.bulletPhysicsEngine = &PhysicsEngine{
		rigidBodyMap:     make(map[uint32]*RigidBody),
		acc:              -5.0,
		sceneBlockAoiMap: sceneBlockAoiMap,
		world:            w,
	}
}

func (p *PhysicsEngine) Update(now int64) []*Entity {
	hitEntityList := make([]*Entity, 0)
	dt := float32(now-p.lastUpdateTime) / 1000.0
	for _, rigidBody := range p.rigidBodyMap {
		aoiManager, exist := p.sceneBlockAoiMap[rigidBody.sceneId]
		if !exist {
			p.DestroyRigidBody(rigidBody.entityId)
			continue
		}
		if !aoiManager.IsValidAoiPos(rigidBody.position.X, rigidBody.position.Y, rigidBody.position.Z) {
			p.DestroyRigidBody(rigidBody.entityId)
			continue
		}
		// 阻力作用于速度
		dvx := rigidBody.drag * rigidBody.velocity.X * dt
		if math.Abs(float64(dvx)) >= math.Abs(float64(rigidBody.velocity.X)) {
			rigidBody.velocity.X = 0.0
		} else {
			rigidBody.velocity.X -= dvx
		}
		dvy := rigidBody.drag * rigidBody.velocity.Y * dt
		if math.Abs(float64(dvy)) >= math.Abs(float64(rigidBody.velocity.Y)) {
			rigidBody.velocity.Y = 0.0
		} else {
			rigidBody.velocity.Y -= dvy
		}
		dvz := rigidBody.drag * rigidBody.velocity.Z * dt
		if math.Abs(float64(dvz)) >= math.Abs(float64(rigidBody.velocity.Z)) {
			rigidBody.velocity.Z = 0.0
		} else {
			rigidBody.velocity.Z -= dvz
		}
		// 重力作用于速度
		rigidBody.velocity.Y += p.acc * dt
		// 速度作用于位移
		oldPos := &alg.Vector3{X: rigidBody.position.X, Y: rigidBody.position.Y, Z: rigidBody.position.Z}
		rigidBody.position.X += rigidBody.velocity.X * dt
		rigidBody.position.Y += rigidBody.velocity.Y * dt
		rigidBody.position.Z += rigidBody.velocity.Z * dt
		newPos := &alg.Vector3{X: rigidBody.position.X, Y: rigidBody.position.Y, Z: rigidBody.position.Z}
		// 碰撞检测
		entity := p.Collision(rigidBody.sceneId, rigidBody.avatarEntityId, oldPos, newPos)
		if entity != nil {
			hitEntityList = append(hitEntityList, entity)
			p.DestroyRigidBody(rigidBody.entityId)
		}
		logger.Debug("[PhysicsEngineUpdate] e: %v, s: %v, p: %v, v: %v", rigidBody.entityId, rigidBody.sceneId, rigidBody.position, rigidBody.velocity)
	}
	p.lastUpdateTime = now
	return hitEntityList
}

func (p *PhysicsEngine) Collision(sceneId uint32, avatarEntityId uint32, oldPos *alg.Vector3, newPos *alg.Vector3) *Entity {
	scene := p.world.GetSceneById(sceneId)
	for _, entity := range scene.GetAllEntity() {
		if entity.GetEntityType() != constant.ENTITY_TYPE_AVATAR {
			continue
		}
		avatarEntity := entity.GetAvatarEntity()
		player := USER_MANAGER.GetOnlineUser(avatarEntity.GetUid())
		if avatarEntity.GetAvatarId() != p.world.GetPlayerActiveAvatarId(player) {
			continue
		}
		if entity.GetId() == avatarEntityId {
			continue
		}
		avatarPos := entity.GetPos()
		// x轴
		lineMinX := float32(0)
		lineMaxX := float32(0)
		if oldPos.X < newPos.X {
			lineMinX = oldPos.X
			lineMaxX = newPos.X
		} else {
			lineMinX = newPos.X
			lineMaxX = oldPos.X
		}
		shapeMinX := float32(avatarPos.X) - AVATAR_RADIUS
		shapeMaxX := float32(avatarPos.X) + AVATAR_RADIUS
		if lineMaxX < shapeMinX || lineMinX > shapeMaxX {
			continue
		}
		// z轴
		lineMinZ := float32(0)
		lineMaxZ := float32(0)
		if oldPos.Z < newPos.Z {
			lineMinZ = oldPos.Z
			lineMaxZ = newPos.Z
		} else {
			lineMinZ = newPos.Z
			lineMaxZ = oldPos.Z
		}
		shapeMinZ := float32(avatarPos.Z) - AVATAR_RADIUS
		shapeMaxZ := float32(avatarPos.Z) + AVATAR_RADIUS
		if lineMaxZ < shapeMinZ || lineMinZ > shapeMaxZ {
			continue
		}
		// y轴
		lineMinY := float32(0)
		lineMaxY := float32(0)
		if oldPos.Y < newPos.Y {
			lineMinY = oldPos.Y
			lineMaxY = newPos.Y
		} else {
			lineMinY = newPos.Y
			lineMaxY = oldPos.Y
		}
		shapeMinY := float32(avatarPos.Y) - AVATAR_HEIGHT/2.0
		shapeMaxY := float32(avatarPos.Y) + AVATAR_HEIGHT/2.0
		if lineMaxY < shapeMinY || lineMinY > shapeMaxY {
			continue
		}
		return entity
	}
	return nil
}

func (p *PhysicsEngine) IsRigidBody(entityId uint32) bool {
	_, exist := p.rigidBodyMap[entityId]
	return exist
}

func (p *PhysicsEngine) CreateRigidBody(entityId, avatarEntityId, sceneId uint32, px, py, pz float32, vx, vy, vz float32) {
	rigidBody := &RigidBody{
		entityId:       entityId,
		avatarEntityId: avatarEntityId,
		sceneId:        sceneId,
		position:       &alg.Vector3{X: px, Y: py, Z: pz},
		velocity:       &alg.Vector3{X: vx, Y: vy, Z: vz},
		drag:           0.01,
		mass:           1.0,
	}
	logger.Debug("[CreateRigidBody] e: %v, s: %v, p: %v, v: %v", rigidBody.entityId, rigidBody.sceneId, rigidBody.position, rigidBody.velocity)
	p.rigidBodyMap[entityId] = rigidBody
}

func (p *PhysicsEngine) DestroyRigidBody(entityId uint32) {
	if !p.IsRigidBody(entityId) {
		return
	}
	rigidBody := p.rigidBodyMap[entityId]
	logger.Debug("[DestroyRigidBody] e: %v, s: %v, p: %v, v: %v", rigidBody.entityId, rigidBody.sceneId, rigidBody.position, rigidBody.velocity)
	delete(p.rigidBodyMap, entityId)
}
