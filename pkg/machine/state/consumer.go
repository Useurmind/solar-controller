package state

type Consumer struct {
	// Unique ID of the consumer.
	ID string `json:"id,omitempty"`

	// Name of the consumer for display purposes.
	Name string `json:"name,omitempty"`

	// The current power in Watt that this consumer is taking.
	CurrentPower int `json:"currentPower,omitempty"`

	// Is this the total power consumption of your house.
	IsTotal bool `json:"isTotal,omitempty"`

	// Can the consumption of this consumer be controlled.
	Controllable bool `json:"controllable,omitempty"`

	// This is the target power in Watt that this consumer should get.
	TargetPower int `json:"targetPower,omitempty"`
}
