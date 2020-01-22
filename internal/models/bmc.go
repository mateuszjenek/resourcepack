package models

// BMC represents single system entity
type BMC struct {
	ID       int
	Type     BMCType
	UUID     string
	Address  string
	Username string
	Password string
}

// BMCType determiantes which type system is
type BMCType string

const (
	// BMCTypeSimulator means that system is simulator
	BMCTypeSimulator BMCType = "Simulator"
	// BMCTypeHardware means that system is real hardware
	BMCTypeHardware BMCType = "Hardware"
)
