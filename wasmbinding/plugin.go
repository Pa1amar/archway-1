package wasmbinding

import (
	wasmKeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	"github.com/archway-network/archway/wasmbinding/rewards"
)

// RewardsKeeperExpected is the expected x/rewards keeper.
type RewardsKeeperExpected interface {
	rewards.KeeperWriterExpected
	rewards.KeeperReaderExpected
}

// BuildWasmOptions returns x/wasmd module options to support WASM bindings functionality.
func BuildWasmOptions(rKeeper RewardsKeeperExpected) []wasmKeeper.Option {
	return []wasmKeeper.Option{
		wasmKeeper.WithMessageHandlerDecorator(BuildWasmMsgDecorator(rKeeper)),
		wasmKeeper.WithQueryPlugins(BuildWasmQueryPlugin(rKeeper)),
	}
}

// BuildWasmMsgDecorator returns the Wasm custom message handler decorator.
func BuildWasmMsgDecorator(rKeeper RewardsKeeperExpected) func(old wasmKeeper.Messenger) wasmKeeper.Messenger {
	return func(old wasmKeeper.Messenger) wasmKeeper.Messenger {
		return NewMsgDispatcher(
			old,
			rewards.NewRewardsMsgHandler(rKeeper),
		)
	}
}

// BuildWasmQueryPlugin returns the Wasm custom querier plugin.
func BuildWasmQueryPlugin(rKeeper RewardsKeeperExpected) *wasmKeeper.QueryPlugins {
	return &wasmKeeper.QueryPlugins{
		Custom: NewQueryDispatcher(
			rewards.NewQueryHandler(rKeeper),
		).DispatchQuery,
	}
}
