package models

// System represents single system entity
type System struct {
	ID       int
	Type     SystemType
	UUID     string
	Address  string
	Username string
	Password string
}

// SystemType determiantes which type system is
type SystemType string

const (
	// SystemTypeSimulator means that system is simulator
	SystemTypeSimulator SystemType = "Simulator"
	// SystemTypeHardware means that system is real hardware
	SystemTypeHardware SystemType = "Hardware"
)

// MapToID returns slice of systems' IDs
func (collection SystemCollection) MapToID() []int {
	systemIDs := make([]int, 0, len(collection))
	for _, system := range collection {
		systemIDs = append(systemIDs, system.ID)
	}
	return systemIDs
}

// SystemIDFilter returns filter to get only systems with given ID
func SystemIDFilter(ID int) SystemCollectionFilter {
	return func(system *System) bool {
		return system.ID == ID
	}
}

// SystemTypeFilter returns filter to get only systems with given type
func SystemTypeFilter(systemType SystemType) SystemCollectionFilter {
	return func(system *System) bool {
		return system.Type == systemType
	}
}

// SystemAllFilter returns filter which do nothing
func SystemAllFilter() SystemCollectionFilter {
	return func(*System) bool {
		return true
	}
}
