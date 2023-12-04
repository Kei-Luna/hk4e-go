package constant

const (
	VISION_LEVEL_NORMAL        = 0
	VISION_LEVEL_LITTLE_REMOTE = 1
	VISION_LEVEL_REMOTE        = 2
	VISION_LEVEL_SUPER         = 3
	VISION_LEVEL_NEARBY        = 4
	VISION_LEVEL_SUPER_NEARBY  = 5
)

type Vision struct {
	VisionRange uint32
	GridWidth   uint32
}

var VISION_LEVEL map[int]*Vision

func init() {
	VISION_LEVEL = map[int]*Vision{
		VISION_LEVEL_NORMAL:        {VisionRange: 80, GridWidth: 20},
		VISION_LEVEL_LITTLE_REMOTE: {VisionRange: 160, GridWidth: 40},
		VISION_LEVEL_REMOTE:        {VisionRange: 1000, GridWidth: 250},
		VISION_LEVEL_SUPER:         {VisionRange: 4000, GridWidth: 1000},
		VISION_LEVEL_NEARBY:        {VisionRange: 40, GridWidth: 20},
		VISION_LEVEL_SUPER_NEARBY:  {VisionRange: 20, GridWidth: 10},
	}
}
