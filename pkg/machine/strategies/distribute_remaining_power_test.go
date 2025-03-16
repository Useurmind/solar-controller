package strategies

import (
	"fmt"
	"testing"

	"github.com/Useurmind/solar-controller/pkg/machine/state"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestDistributeRemainingPower(t *testing.T) {
	type testCase struct {
		name          string
		lastState     state.State
		newState      state.State
		expectedState state.State
	}

	testCases := []testCase{
		{
			name: "default",
			lastState: state.State{
				Sources: solarSources(1000, 200),
				Consumer: []*state.Consumer{
					totalPower(500),
					wallbox(800, 500),
				},
			},
			newState: state.State{
				Sources: solarSources(700, 100),
				Consumer: []*state.Consumer{
					totalPower(600),
					wallbox(500, 200),
				},
			},
			expectedState: state.State{
				Sources: solarSources(700, 100),
				Consumer: []*state.Consumer{
					totalPower(600),
					wallbox(500, 200),
				},
			},
		},
	}

	s := NewDistributeRemainingPowerStrategy()

	for _, tc := range testCases {
		t.Run(tc.name, func(t2 *testing.T) {
			updatedState, err := s.ComputeNextState(&tc.lastState, &tc.newState)
			require.Nil(t2, err)

			if !cmp.Equal(tc.expectedState, updatedState) {
				t2.Errorf("update state not equal expected state, diff is: %s", cmp.Diff(tc.expectedState, updatedState))
			}
		})
	}
}

func solarSources(powers ...int) []*state.SolarSource {
	sources := []*state.SolarSource{}

	for i, power := range powers {
		id := fmt.Sprintf("source%d", i)
		sources = append(sources, &state.SolarSource{
			ID:           id,
			Name:         id,
			CurrentPower: power,
		})
	}

	return sources
}

func totalPower(power int) *state.Consumer {
	return &state.Consumer{
		ID:           "total",
		Name:         "Total power consumption",
		IsTotal:      true,
		Controllable: false,
		CurrentPower: 400,
	}
}

func wallbox(currentPower int, targetPower int) *state.Consumer {
	return &state.Consumer{
		ID:           "wallbox",
		Name:         "Wallbox",
		CurrentPower: currentPower,
		Controllable: true,
		TargetPower:  targetPower,
	}
}
