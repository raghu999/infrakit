package flavor

import (
	"encoding/json"
	"github.com/docker/infrakit/pkg/plugin/group/types"
	"github.com/docker/infrakit/pkg/spi"
	"github.com/docker/infrakit/pkg/spi/instance"
)

// InterfaceSpec is the current name and version of the Flavor API.
var InterfaceSpec = spi.InterfaceSpec{
	Name:    "Flavor",
	Version: "0.1.0",
}

// Health is an indication of whether the Flavor is functioning properly.
type Health int

const (
	// Unknown indicates that the Health cannot currently be confirmed.
	Unknown Health = iota

	// Healthy indicates that the Flavor is confirmed to be functioning.
	Healthy

	// Unhealthy indicates that the Flavor is confirmed to not be functioning properly.
	Unhealthy
)

// Plugin defines custom behavior for what runs on instances.
type Plugin interface {

	// Validate checks whether the helper can support a configuration.
	Validate(flavorProperties json.RawMessage, allocation types.AllocationMethod) error

	// Prepare allows the Flavor to modify the provisioning instructions for an instance.  For example, a
	// helper could be used to place additional tags on the machine, or generate a specialized Init command based on
	// the flavor configuration.
	Prepare(flavorProperties json.RawMessage, spec instance.Spec, allocation types.AllocationMethod) (instance.Spec, error)

	// Healthy determines the Health of this Flavor on an instance.
	Healthy(flavorProperties json.RawMessage, inst instance.Description) (Health, error)

	// Drain allows the flavor to perform a best-effort cleanup operation before the instance is destroyed.
	Drain(flavorProperties json.RawMessage, inst instance.Description) error
}
