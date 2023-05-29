package modifier

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type Listeners struct {
	// listeners for modifier processes

	// Called when a new modifier instance is added. Note: if using Replace or ReplaceBySource,
	// this will always be a fresh instance when stacking. If using Merge, OnAdd will be called
	// on the old instance.
	OnAdd func(mod *ModifierInstance)
	// Called when a modifier instance is removed, either forceably or due to the instance expiring.
	OnRemove func(mod *ModifierInstance)
	// Called when the duration for all modifiers instances of this shape are extended.
	OnExtendDuration func(mod *ModifierInstance)
	// Called when the count/stacks for all modifier instances of this shape are extended.
	// Will not be called if OnAdd is called (doesnt call on standard stacking behavior)
	OnExtendCount func(mod *ModifierInstance)
	// Called when any stat changes on the target this modifier is attached to. Will be called if
	// you modify properties within this call, so take care not to create a recursive loop.
	OnPropertyChange func(mod *ModifierInstance)
	// Called at the start of the turn, before the action takes place (used by DoTs).
	OnPhase1 func(mod *ModifierInstance)
	// Called at the end of the turn
	OnPhase2 func(mod *ModifierInstance)

	// character events

	// Called when a new character is added to the simulation (done as part of sim setup)
	OnCharacterAdded func(mod *ModifierInstance, char info.Character)
}

func (mgr *Manager) subscribe() {
	events := mgr.engine.Events()

	events.CharacterAdded.Subscribe(mgr.characterAdded)
}

func (mgr *Manager) emitPropertyChange(target key.TargetID) {
	for _, mod := range mgr.targets[target] {
		f := mod.listeners.OnPropertyChange
		if f != nil {
			f(mod)
		}
	}
}

func (mgr *Manager) emitAdd(target key.TargetID, mod *ModifierInstance, chance float64) {
	f := mod.listeners.OnAdd
	if f != nil {
		f(mod)
	}
	mgr.engine.Events().ModifierAdded.Emit(event.ModifierAddedEvent{
		Target:   target,
		Modifier: mod.ToModel(),
		Chance:   chance,
	})
}

func (mgr *Manager) emitRemove(target key.TargetID, mods []*ModifierInstance) {
	for _, mod := range mods {
		if len(mod.stats) > 0 {
			mgr.emitPropertyChange(target)
		}

		f := mod.listeners.OnRemove
		if f != nil {
			f(mod)
		}
		mgr.engine.Events().ModifierRemoved.Emit(event.ModifierRemovedEvent{
			Target:   target,
			Modifier: mod.ToModel(),
		})
	}
}

func (mgr *Manager) emitExtendDuration(target key.TargetID, mod *ModifierInstance, old int) {
	f := mod.listeners.OnExtendDuration
	if f != nil {
		f(mod)
	}
	mgr.engine.Events().ModifierExtendedDuration.Emit(event.ModifierExtendedDurationEvent{
		Target:   target,
		Modifier: mod.ToModel(),
		OldValue: old,
		NewValue: mod.duration,
	})
}

func (mgr *Manager) emitExtendCount(target key.TargetID, mod *ModifierInstance, old float64) {
	f := mod.listeners.OnExtendCount
	if f != nil {
		f(mod)
	}
	mgr.engine.Events().ModifierExtendedCount.Emit(event.ModifierExtendedCountEvent{
		Target:   target,
		Modifier: mod.ToModel(),
		OldValue: old,
		NewValue: mod.count,
	})
}

func (mgr *Manager) characterAdded(evt event.CharacterAddedEvent) {
	for _, mod := range mgr.targets[evt.Id] {
		f := mod.listeners.OnCharacterAdded
		if f != nil {
			f(mod, evt.Info)
		}
	}
}
