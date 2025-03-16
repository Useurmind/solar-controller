package strategies

import "github.com/Useurmind/solar-controller/pkg/machine/state"

type DistributeRemainingPowerStrategy struct {

}

func NewDistributeRemainingPowerStrategy() *DistributeRemainingPowerStrategy{
	return &DistributeRemainingPowerStrategy{}
}

func (s *DistributeRemainingPowerStrategy) GetName() string {
	return "DistributeRemainingPower"
}

func (s *DistributeRemainingPowerStrategy) ComputeNextState(oldState *state.State, currentState *state.State) (*state.State, error) {


	return &state.State{}, nil
}