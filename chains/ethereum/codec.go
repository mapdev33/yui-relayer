package ethereum

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/mapdev33/yui-relayer/core"
)

// RegisterInterfaces register the module interfaces to protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*core.ChainConfigI)(nil),
		&ChainConfig{},
	)
}
