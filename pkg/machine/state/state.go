package state

type State struct {
	Sources  []*SolarSource
	Consumer []*Consumer
}

// Returns a total power consumer if there is one.
// If there are more than one it fails.
func (s *State) GetTotalPower() (*Consumer, error) {

}
