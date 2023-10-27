package polymer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

const AccountAddressPrefix = "polymer"

func setPrefixes(accountAddressPrefix string) {
	// Set prefixes
	accountPubKeyPrefix := accountAddressPrefix + "pub"
	validatorAddressPrefix := accountAddressPrefix + "valoper"
	validatorPubKeyPrefix := accountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := accountAddressPrefix + "valcons"
	consNodePubKeyPrefix := accountAddressPrefix + "valconspub"

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(accountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.Seal()
}

func init() {
	// Set prefixes
	setPrefixes(AccountAddressPrefix)
}

// CosmosToEvm converts a sdk.AccAddress to an EVM address
func CosmosToEvm(addr sdk.AccAddress) common.Address {
	return common.BytesToAddress(addr.Bytes())
}

// EvmToCosmos converts an EVM address to a sdk.AccAddress
func EvmToCosmos(addr common.Address) sdk.AccAddress {
	return sdk.AccAddress(addr.Bytes())
}
