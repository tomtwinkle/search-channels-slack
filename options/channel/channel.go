package channel

import "github.com/tomtwinkle/search-channels-slack/types"

type Option func(*ChannelOption)
type ChannelOption struct {
	Types []types.ChannelType
}

func Types(types []types.ChannelType) Option {
	return func(ops *ChannelOption) {
		ops.Types = types
	}
}
