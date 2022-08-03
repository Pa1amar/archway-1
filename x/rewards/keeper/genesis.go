package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/archway-network/archway/pkg"
	"github.com/archway-network/archway/x/rewards/types"
)

// ExportGenesis exports the module genesis for the current block.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	minConsFee, _ := k.state.MinConsensusFee(ctx).GetFee() // default sdk.Coin value is ok

	return types.NewGenesisState(
		k.GetParams(ctx),
		k.state.ContractMetadataState(ctx).Export(),
		k.state.BlockRewardsState(ctx).Export(),
		k.state.TxRewardsState(ctx).Export(),
		minConsFee,
	)
}

// InitGenesis initializes the module genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)
	k.state.ContractMetadataState(ctx).Import(state.ContractsMetadata)
	k.state.BlockRewardsState(ctx).Import(state.BlockRewards)
	k.state.TxRewardsState(ctx).Import(state.TxRewards)

	if !pkg.DecCoinIsZero(state.MinConsensusFee) {
		k.state.MinConsensusFee(ctx).SetFee(state.MinConsensusFee)
	}
}
