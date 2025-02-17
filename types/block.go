package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// BlockGasLimit returns the max gas (limit) defined in the block gas meter. If the meter is not
// set, it returns the max gas from the application consensus params.
// NOTE: see https://github.com/cosmos/cosmos-sdk/issues/9514 for full reference
func BlockGasLimit(ctx sdk.Context) uint64 {
	blockGasMeter := ctx.BlockGasMeter()
	if blockGasMeter != nil {
		return blockGasMeter.Limit()
	}

	cp := ctx.ConsensusParams()
	if cp == nil || cp.Block == nil {
		return 0
	}

	maxGas := cp.Block.MaxGas
	if maxGas > 0 {
		return uint64(maxGas)
	}

	return 0
}
