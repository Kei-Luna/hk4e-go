package constant

import "hk4e/pkg/endec"

var (
	DEFAULT_ABILITY_HASH_CODE []int32
)

func init() {
	defaultAbilityStringList := []string{
		"Avatar_DefaultAbility_VisionReplaceDieInvincible",
		"Avatar_DefaultAbility_AvartarInShaderChange",
		"Avatar_SprintBS_Invincible",
		"Avatar_Freeze_Duration_Reducer",
		"Avatar_Attack_ReviveEnergy",
		"Avatar_Component_Initializer",
		"Avatar_HDMesh_Controller",
		"Avatar_Trampoline_Jump_Controller",
	}
	DEFAULT_ABILITY_HASH_CODE = make([]int32, 0)
	for _, defaultAbilityString := range defaultAbilityStringList {
		DEFAULT_ABILITY_HASH_CODE = append(DEFAULT_ABILITY_HASH_CODE, endec.Hk4eAbilityHashCode(defaultAbilityString))
	}
}
