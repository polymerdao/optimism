package polymer

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// Test cosmos/evm address conversion
func TestAccountConversion(t *testing.T) {
	cosmAddrStr := "polymer1cjxgqn2wxvxlxu5stqzxumyflk4qlwwl2cku48"
	evmAddrStr := "0xC48c804d4E330Df3729058046e6C89FDaA0fb9DF"
	cosmAddr, err := sdk.AccAddressFromBech32(cosmAddrStr)
	require.NoError(t, err)
	evmAddr := common.HexToAddress(evmAddrStr)

	// Test cosmos -> evm
	require.Equal(t, evmAddr.Bytes(), CosmosToEvm(cosmAddr).Bytes())
	// Test evm -> cosmos
	require.Equal(t, cosmAddr.Bytes(), EvmToCosmos(evmAddr).Bytes())
}
