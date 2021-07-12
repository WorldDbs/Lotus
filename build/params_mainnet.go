// +build !debug
// +build !2k
// +build !testground
// +build !calibnet
// +build !nerpanet
// +build !butterflynet

package build

import (
	"math"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	UpgradeSmokeHeight: DrandMainnet,
}

const BootstrappersFile = "mainnet.pi"
const GenesisFile = "mainnet.car"

const UpgradeBreezeHeight = -1

const BreezeGasTampingDuration = 0

const UpgradeSmokeHeight = -1

const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3

const UpgradeAssemblyHeight = 10

const UpgradeTapeHeight = -4

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.
// Miners, clients, developers, custodians all need time to prepare.
// We still have upgrades and state changes to do, but can happen after signaling timing here.
const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 15

const UpgradeCalicoHeight = 20
const UpgradePersianHeight = 25

const UpgradeOrangeHeight = 27

// 2020-12-22T02:00:00Z
var UpgradeClausHeight = abi.ChainEpoch(30)

// 2021-03-04T00:00:30Z
const UpgradeTrustHeight = abi.ChainEpoch(230)

// 2021-04-12T22:00:00Z
const UpgradeNorwegianHeight = abi.ChainEpoch(400)

// 2021-04-29T06:00:00Z
const UpgradeTurboHeight = abi.ChainEpoch(650)

// 2021-06-30T22:00:00Z
var UpgradeHyperdriveHeight = abi.ChainEpoch(800)

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(512 << 20))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)

	if os.Getenv("LOTUS_USE_TEST_ADDRESSES") != "1" {
		SetAddressNetwork(address.Mainnet)
	}

	if os.Getenv("LOTUS_DISABLE_HYPERDRIVE") == "1" {
		UpgradeHyperdriveHeight = math.MaxInt64
	}

	Devnet = false

	BuildType = BuildMainnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

// we skip checks on message validity in this block to sidestep the zero-bls signature
var WhitelistedBlock = MustParseCid("bafy2bzaceapyg2uyzk7vueh3xccxkuwbz3nxewjyguoxvhx77malc2lzn2ybi")
