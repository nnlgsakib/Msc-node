package polybft

import (
	"testing"

	"github.com/Mind-chain/mind/consensus/polybft/bitmap"
	bls "github.com/Mind-chain/mind/consensus/polybft/signer"
	"github.com/Mind-chain/mind/consensus/polybft/validator"
	"github.com/Mind-chain/mind/consensus/polybft/wallet"
	"github.com/Mind-chain/mind/types"
	"github.com/stretchr/testify/assert"
)

func Test_setupHeaderHashFunc(t *testing.T) {
	extra := &Extra{
		Validators: &validator.ValidatorSetDelta{Removed: bitmap.Bitmap{1}},
		Parent:     createSignature(t, []*wallet.Account{generateTestAccount(t)}, types.ZeroHash, bls.DomainCheckpointManager),
		Checkpoint: &CheckpointData{},
		Committed:  &Signature{},
	}

	header := &types.Header{
		Number:    2,
		GasLimit:  10000003,
		Timestamp: 18,
	}

	header.ExtraData = extra.MarshalRLPTo(nil)
	notFullExtraHash := types.HeaderHash(header)

	extra.Committed = createSignature(t, []*wallet.Account{generateTestAccount(t)}, types.ZeroHash, bls.DomainCheckpointManager)
	header.ExtraData = extra.MarshalRLPTo(nil)
	fullExtraHash := types.HeaderHash(header)

	assert.Equal(t, notFullExtraHash, fullExtraHash)

	header.ExtraData = []byte{1, 2, 3, 4, 100, 200, 255}
	assert.Equal(t, types.ZeroHash, types.HeaderHash(header)) // to small extra data
}
