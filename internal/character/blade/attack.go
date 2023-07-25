package blade

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Normal key.Attack = "blade-normal"
const EnhancedNormal key.Attack = "blade-enhanced-normal"

var attackHits = []float64{0.5, 0.5}

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	if !c.engine.HasModifier(c.id, Hellscape) {
		for i, hitRatio := range attackHits {
			c.engine.Attack(info.Attack{
				Key:        Normal,
				HitIndex:   i,
				Source:     c.id,
				Targets:    []key.TargetID{target},
				DamageType: model.DamageType_WIND,
				AttackType: model.AttackType_NORMAL,
				BaseDamage: info.DamageMap{
					model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()],
				},
				StanceDamage: 30.0,
				EnergyGain:   20.0,
				HitRatio:     hitRatio,
			})
		}
	} else {
		// Primary Target
		for i, hitRatio := range attackHits {
			c.engine.Attack(info.Attack{
				Key:        EnhancedNormal,
				HitIndex:   i,
				Source:     c.id,
				Targets:    []key.TargetID{target},
				DamageType: model.DamageType_WIND,
				AttackType: model.AttackType_NORMAL,
				BaseDamage: info.DamageMap{
					model.DamageFormula_BY_ATK:    enhancedBasicSingleAtk[c.info.AttackLevelIndex()],
					model.DamageFormula_BY_MAX_HP: enhancedBasicSingleHP[c.info.AttackLevelIndex()],
				},
				StanceDamage: 60.0,
				EnergyGain:   30.0,
				HitRatio:     hitRatio,
			})
		}

		// Adjacent Targets
		c.engine.Attack(info.Attack{
			Key:        EnhancedNormal,
			Source:     c.id,
			Targets:    c.engine.AdjacentTo(target),
			DamageType: model.DamageType_WIND,
			AttackType: model.AttackType_NORMAL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK:    enhancedBasicBlastAtk[c.info.AttackLevelIndex()],
				model.DamageFormula_BY_MAX_HP: enhancedBasicBlastHP[c.info.AttackLevelIndex()],
			},
			StanceDamage: 30.0,
		})
	}
}
