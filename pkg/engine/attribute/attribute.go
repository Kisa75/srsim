// package attribute provides attribute service which is used to keep track of
// character stats
package attribute

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
)

type Getter interface {
	Stats(target key.TargetID) *info.Stats
	Stance(target key.TargetID) float64
	Energy(target key.TargetID) float64
	MaxEnergy(target key.TargetID) float64
	EnergyRatio(target key.TargetID) float64
	HPRatio(target key.TargetID) float64
	IsAlive(target key.TargetID) bool
	State(target key.TargetID) info.TargetState
	FullEnergy(target key.TargetID) bool
	LastAttacker(target key.TargetID) key.TargetID
	SP() int
}

type Manager interface {
	Getter

	ModifySP(data info.ModifySP) error

	AddTarget(target key.TargetID, base info.Attributes) error

	SetHP(data info.ModifyAttribute, isDamage bool) error
	ModifyHPByRatio(data info.ModifyHPByRatio, isDamage bool) error
	ModifyHPByAmount(data info.ModifyAttribute, isDamage bool) error

	SetEnergy(data info.ModifyAttribute) error
	ModifyEnergy(data info.ModifyAttribute) error
	ModifyEnergyFixed(data info.ModifyAttribute) error

	SetStance(data info.ModifyAttribute) error
	ModifyStance(data info.ModifyAttribute) error
}

type Service struct {
	event   *event.System
	modEval modifier.Eval
	targets map[key.TargetID]*attrTarget
	sp      int
}

func New(event *event.System, modEval modifier.Eval) Manager {
	return &Service{
		event:   event,
		modEval: modEval,
		targets: make(map[key.TargetID]*attrTarget, 10),
		sp:      3,
	}
}

func (s *Service) Stats(target key.TargetID) *info.Stats {
	mods := s.modEval.EvalModifiers(target)

	var attr *info.Attributes
	if t, ok := s.targets[target]; ok {
		attr = t.attributes
	} else {
		// default attribute instead of returning an error
		attr = new(info.Attributes)
		*attr = info.DefaultAttribute()
	}

	return info.NewStats(target, attr, mods)
}

func (s *Service) HPRatio(target key.TargetID) float64 {
	if t, ok := s.targets[target]; ok {
		return t.attributes.HPRatio
	}
	return 0.0
}

func (s *Service) Energy(target key.TargetID) float64 {
	if t, ok := s.targets[target]; ok {
		return t.attributes.Energy
	}
	return 0.0
}

func (s *Service) FullEnergy(target key.TargetID) bool {
	if t, ok := s.targets[target]; ok {
		return t.attributes.Energy >= t.attributes.MaxEnergy
	}
	return false
}

func (s *Service) MaxEnergy(target key.TargetID) float64 {
	if t, ok := s.targets[target]; ok {
		return t.attributes.MaxEnergy
	}
	return 0.0
}

func (s *Service) EnergyRatio(target key.TargetID) float64 {
	if t, ok := s.targets[target]; ok {
		return t.attributes.Energy / t.attributes.MaxEnergy
	}
	return 0.0
}

func (s *Service) Stance(target key.TargetID) float64 {
	if t, ok := s.targets[target]; ok {
		return t.attributes.Stance
	}
	return 0.0
}

func (s *Service) State(target key.TargetID) info.TargetState {
	if t, ok := s.targets[target]; ok {
		return t.state
	}
	return info.Invalid
}

func (s *Service) IsAlive(target key.TargetID) bool {
	return s.State(target) == info.Alive
}

func (s *Service) LastAttacker(target key.TargetID) key.TargetID {
	if t, ok := s.targets[target]; ok {
		return t.lastAttacker
	}
	return target
}

func (s *Service) SP() int {
	return s.sp
}
