package strategies

import "github.com/Useurmind/solar-controller/pkg/machine/state"

type Strategy interface {
	GetName() string
	ComputeNextState(oldState *state.State, currentState *state.State) (*state.State, error)
}