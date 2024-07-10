package cosa_test

import (
	"testing"

	keepertest "cosa/testutil/keeper"
	"cosa/testutil/nullify"
	cosa "cosa/x/cosa/module"
	"cosa/x/cosa/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosaKeeper(t)
	cosa.InitGenesis(ctx, k, genesisState)
	got := cosa.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
