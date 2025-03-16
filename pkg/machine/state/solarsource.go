package state

type SolarSource struct {
	// Unique ID for this source.
	ID           string `json:"id,omitempty"`

	// Name for display purposes.
	Name         string `json:"name,omitempty"`

	// The current power in Watt that this source is producing.
	CurrentPower int `json:"currentPower,omitempty"`
}
