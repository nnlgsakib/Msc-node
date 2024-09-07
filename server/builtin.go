package server

import (
	"github.com/Mind-chain/mind/chain"
	"github.com/Mind-chain/mind/consensus"
	consensusIBFT "github.com/Mind-chain/mind/consensus/NLG-ibft"
	consensusDev "github.com/Mind-chain/mind/consensus/dev"
	consensusDummy "github.com/Mind-chain/mind/consensus/dummy"

	//consensusPolyBFT
	"github.com/Mind-chain/mind/forkmanager"
	"github.com/Mind-chain/mind/secrets"
	"github.com/Mind-chain/mind/secrets/awsssm"
	"github.com/Mind-chain/mind/secrets/gcpssm"
	"github.com/Mind-chain/mind/secrets/hashicorpvault"
	"github.com/Mind-chain/mind/secrets/local"
	"github.com/Mind-chain/mind/state"
)

type GenesisFactoryHook func(config *chain.Chain, engineName string) func(*state.Transition) error

type ConsensusType string

type ForkManagerFactory func(forks *chain.Forks) error

type ForkManagerInitialParamsFactory func(config *chain.Chain) (*forkmanager.ForkParams, error)

const (
	DevConsensus  ConsensusType = "dev"
	IBFTConsensus ConsensusType = "ibft"
	//PolyBFTConsensus ConsensusType = consensusPolyBFT.ConsensusName
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:  consensusDev.Factory,
	IBFTConsensus: consensusIBFT.Factory,
	//PolyBFTConsensus: consensusPolyBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

var genesisCreationFactory = map[ConsensusType]GenesisFactoryHook{}

var forkManagerInitialParamsFactory = map[ConsensusType]ForkManagerInitialParamsFactory{
	// PolyBFTConsensus: consensusPolyBFT.ForkManagerInitialParamsFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
