package libdream

import (
	"context"

	commonSpecs "github.com/taubyte/go-specs/common"
)

func init() {
	// Universes
	universes = make(map[string]*Universe)
	multiverseCtx, multiverseCtxC = context.WithCancel(context.Background())
	fixtures = make(map[string]FixtureHandler)

	// Protocols and P2P Client Registry
	Registry = &handlerRegistry{
		registry: make(map[string]*handlers),
	}

	Ports = make(map[string]int)
	lastPort := portStart
	for _, protocol := range commonSpecs.Protocols {
		Registry.registry[protocol] = &handlers{}

		port := lastPort + portBuffer
		Ports["http/"+protocol] = port
		Ports["p2p/"+protocol] = port + 2
		Ports["ipfs/"+protocol] = port + 4
		Ports["dns/"+protocol] = port + 8
		lastPort = port + portBuffer
	}
}
