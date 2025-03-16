package machine

import (
	"github.com/Useurmind/solar-controller/pkg/machine/state"
	"github.com/Useurmind/solar-controller/pkg/machine/strategies"
)

type Machine struct {
	Strategy strategies.Strategy
	lastState *state.State
}

func (m *Machine) ComputeNextState(state *state.State) (*state.State, error) {
	newState, err := m.Strategy.ComputeNextState(m.lastState, state)
	if err != nil {
		return nil, err
	}
	m.lastState = newState

	return newState, nil
}