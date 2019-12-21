package app

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	sr25519 "github.com/heystraightedge/straightedge/sr25519"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/multisig"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

const (
	// TODO: Use this cost per byte through parameter or overriding NewConsumeGasForTxSizeDecorator
	// which currently defaults at 10, if intended
	// memoCostPerByte     sdk.Gas = 3
	secp256k1VerifyCost uint64 = 21000
	sr25519VerifyCost   uint64 = 21000
)

// NewAnteHandler returns an ante handler responsible for attempting to route an
// Ethereum or SDK transaction to an internal ante handler for performing
// transaction-level processing (e.g. fee payment, signature verification) before
// being passed onto it's respective handler.
func NewAnteHandler(ak auth.AccountKeeper, sk types.SupplyKeeper) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx, sim bool,
	) (newCtx sdk.Context, err error) {

		switch tx.(type) {
		case auth.StdTx:
			stdAnte := sdk.ChainAnteDecorators(
				ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
				ante.NewMempoolFeeDecorator(),
				ante.NewValidateBasicDecorator(),
				ante.NewValidateMemoDecorator(ak),
				ante.NewConsumeGasForTxSizeDecorator(ak),
				ante.NewSetPubKeyDecorator(ak), // SetPubKeyDecorator must be called before all signature verification decorators
				ante.NewValidateSigCountDecorator(ak),
				ante.NewDeductFeeDecorator(ak, sk),
				ante.NewSigGasConsumeDecorator(ak, consumeSigGas),
				ante.NewSigVerificationDecorator(ak),
				ante.NewIncrementSequenceDecorator(ak), // innermost AnteDecorator
			)

			return stdAnte(ctx, tx, sim)

		default:
			return ctx, sdk.ErrInternal(fmt.Sprintf("transaction type invalid: %T", tx))
		}
	}
}

func consumeSigGas(
	meter sdk.GasMeter, sig []byte, pubkey crypto.PubKey, params types.Params,
) error {
	switch pubkey := pubkey.(type) {
	case secp256k1.PubKeySecp256k1:
		meter.ConsumeGas(params.SigVerifyCostSecp256k1, "ante verify: secp256k1")
		return nil
	case sr25519.PubKeySr25519:
		meter.ConsumeGas(params.SigVerifyCostED25519, "ante verify: sr25519")
		return nil
	case multisig.PubKeyMultisigThreshold:
		var multisignature multisig.Multisignature
		codec.Cdc.MustUnmarshalBinaryBare(sig, &multisignature)
		ConsumeMultisignatureVerificationGas(meter, multisignature, pubkey, params)
		return nil
	default:
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "unrecognized public key type: %T", pubkey)
	}
}

// Copied from sdk ante package as it wasn't exposed there
func ConsumeMultisignatureVerificationGas(meter sdk.GasMeter,
	sig multisig.Multisignature, pubkey multisig.PubKeyMultisigThreshold,
	params types.Params) {
	size := sig.BitArray.Size()
	sigIndex := 0
	for i := 0; i < size; i++ {
		if sig.BitArray.GetIndex(i) {
			ante.DefaultSigVerificationGasConsumer(meter, sig.Sigs[sigIndex], pubkey.PubKeys[i], params)
			sigIndex++
		}
	}
}
