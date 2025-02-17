// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash = common.HexToHash("0xfcdc9de4af1feb5e727f269cdee56eb631926df5b783c28c66d3ed7b1a2b73ac")
	TestnetGenesisHash = common.HexToHash("0xc5fce6d9b77ff8cbbd973f7a3251d98c7623f5d6951798e9fd3d0ffa7aabc32c")
)

// TrustedCheckpoints associates each known checkpoint with the genesis hash of
// the chain it belongs to.
var TrustedCheckpoints = map[common.Hash]*TrustedCheckpoint{}

// CheckpointOracles associates each known checkpoint oracles with the genesis hash of
// the chain it belongs to.
var CheckpointOracles = map[common.Hash]*CheckpointOracleConfig{}

var (

	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
<<<<<<< HEAD
		ChainID:             big.NewInt(1323),
=======
		ChainID:             big.NewInt(805),
>>>>>>> change ChainID
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		BerlinBlock:         big.NewInt(0),
		CVE_2021_39137Block: big.NewInt(2509228),
		// @cary the block when hardfork happens.
		IshikariBlock:         big.NewInt(11171299),
		IshikariPatch001Block: big.NewInt(11171299),
		IshikariPatch002Block: big.NewInt(11171299),
		AmazonBlock:           big.NewInt(24766075),
		POSA: &POSAConfig{
			Period: 10,
			Epoch:  100,
			IshikariInitialValidators: []common.Address{
				common.HexToAddress("0x64a1851D37B949c6De0711F7acCaE96584A19b70"),
				common.HexToAddress("0x729Bc1Ce2E9D7A0206575C4a54959E4830DA6680"),
				common.HexToAddress("0x4809d3d4E45f9456c23bE0b0A77a7A5461B31D2b"),
				common.HexToAddress("0x4D90C011AE5323F0A42406c76119d57c98A0DBF1"),
				common.HexToAddress("0x35B3230c4A942979Ed95D6E8c6A12d64bd7624A4"),
				common.HexToAddress("0x26Beac5a5ffFd92823ffE96e5D04B7b3BA09d524"),
				common.HexToAddress("0x6A83575762A56F3a1cBf6c595ED14cA6BccDe905"),
				common.HexToAddress("0xD2e3DAF58ce3e4c71dE344a9579530fF0F8623Bd"),
				common.HexToAddress("0x3Fe9CE44f72A779043C96Db5A2900D3D44404506"),
				common.HexToAddress("0x30950ca20714ea6C3360aDad96fEFFC99c7f41e3"),
				common.HexToAddress("0xFF4633495fa2483bB7C7017AF109834b92600CD3"),
				common.HexToAddress("0xAf1F640FB6250493D45B7e6555C214DdD9bE8c43"),
				common.HexToAddress("0xb22af26700f050bC16e9c87Db44f3bb37Ee0a52c"),
				common.HexToAddress("0x9A403D843cc3e4C5b546a63CA3d46c17726e8e43"),
				common.HexToAddress("0x5F2f17b57f5ADc712c305D5b9C197d2000839603"),
				common.HexToAddress("0x1fFFfD5Ad4C7754e092e73Fa3591E15bf056B93C"),
				common.HexToAddress("0xc19638b5A297cdccC99f7257e116543aaB52c3a1"),
				common.HexToAddress("0x50426176C19cD0830153Ba90e6B99f1c126e4d91"),
				common.HexToAddress("0xcEc7691424D946615c54CBFbf26Dbb174a293BeF"),
				common.HexToAddress("0xEEd4b671fE6231077d94E6b0d7c11afe41e60Daa"),
				common.HexToAddress("0xdb8E0591b2A6c2715D467267fae2cDC19E80DcBe"),
			}, // @cary @Junm TODO: Ishikari initial validators

			IshikariInitialManagers: []common.Address{
				common.HexToAddress("0x69A8a67Abf048cf37b47A7DEBEA89be05FbdEd46"),
				common.HexToAddress("0xe33E0aEd67df4d69A369B7C410cec490a1f6a810"),
				common.HexToAddress("0x7E262C5F9D4b7E442e0Fe7131D5DdEC59cAeEc06"),
				common.HexToAddress("0xAD1899E1A3822b223b48188d7D85a214E236f47f"),
				common.HexToAddress("0x8Bd502Ad9aB07Abc84e2dBee8515c0d25E95E2F4"),
				common.HexToAddress("0xef258713f08221b6e46ae5195D3d0525cc522F3E"),
				common.HexToAddress("0xD4A2f4B1b6c422e8C4161959Fda20F44fA4a522A"),
				common.HexToAddress("0xF0fc22081a4dA1B2ab54E44A9E47856cfB337d67"),
				common.HexToAddress("0xe3DBA899bF4C41ca0a652C27eCf8ACf54093A875"),
				common.HexToAddress("0x1cB11f947e2072d767e3C3EaD1b08819b70BcD47"),
				common.HexToAddress("0xB62Fef5da2e837c1266F76A0b5C57e491686f49C"),
				common.HexToAddress("0x7a07677ABA95A498714De60725bfA77fbd5c8b7E"),
				common.HexToAddress("0xa6e38DBDfb261af5aDc931056752F86Bc81E86f7"),
				common.HexToAddress("0x38337C7d2d358Ad6E5d62d0ef447dd1277655Bea"),
				common.HexToAddress("0x896df5B432fA529BaE80D036629C7506BE764E15"),
				common.HexToAddress("0x9a57c02761CdB30525445B12d7fa9b883d9896E3"),
				common.HexToAddress("0xbd70D632156f07c41642036D85b2D65507017313"),
				common.HexToAddress("0xc0019C287110d89219578928FD714F525Ae39524"),
				common.HexToAddress("0x63996c35AFD02A762dd16Bc7C872790B52Da83c3"),
				common.HexToAddress("0x6F0bA105a90f4A6d2AC4605E3407FfEcf406966F"),
				common.HexToAddress("0xFBd2207867Dab0e6d3698cd78cF9E1De4cbdCb65")},
			IshikariAdminMultiSig: common.HexToAddress("0x6E6fafed1DA84Bf19ffdD4fC1D727a986a39a4Bf"),
		},
	}

	// TestnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
	TestnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(585858),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		BerlinBlock:         big.NewInt(0),
		CVE_2021_39137Block: big.NewInt(0),

		// @cary the block when hardfork happens.
		IshikariBlock: big.NewInt(11321699),
		// Ishikari patch 001
		IshikariPatch001Block: big.NewInt(12153317),
		// Ishikari patch 002
		IshikariPatch002Block: big.NewInt(12162886),
		AmazonBlock:           big.NewInt(25306677),
		// Ishikari patch001
		// Fix minor bugs found in testnet

		POSA: &POSAConfig{
			Period: 10,
			Epoch:  30000,
			IshikariInitialValidators: []common.Address{
				common.HexToAddress("0x7e620BCE6142402fb735B6364500DBF831A6AF44"),
				common.HexToAddress("0x31f52c12897A2Db14B46a5564DD62568CF7A40D2"),
				common.HexToAddress("0xcBd5A98657DE2B9009294E6AeD48D95321DF754D"),
				common.HexToAddress("0xb6586572Cc2cdFB139c2A9B96a619544c511d8d4"),
				common.HexToAddress("0x0482A95A47A4d19d0240f9e5759bdE1434A77dC0"),
			}, //@cary: the initial validators for Ishikari hardfork
			IshikariInitialManagers: []common.Address{
				common.HexToAddress("0x3693CFe5FC3b7e40288CEd052f5Ea74F38cCf40E"),
				common.HexToAddress("0x1C3117d2907CbF637989b17CfDDD691936fca58c"),
				common.HexToAddress("0x33D1259B00eB5F5Ae9fa137Ce13f6383C52a7ef2"),
				common.HexToAddress("0x92bCa3159630eAb52d988f616d7d83079a4c56a1"),
				common.HexToAddress("0x9eb7A2526F215c8cd262c4abF7038fd7621286e5"),
			},
			IshikariAdminMultiSig: common.HexToAddress("0xbc3b1999442205539e833085E601BB7F4699aa28"), // @cary: the multisig address for admin
		},
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Ethash consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, false, new(EthashConfig), nil, nil}

	// AllCliqueProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Clique consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllCliqueProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, false, nil, &CliqueConfig{Period: 0, Epoch: 30000}, nil}

	TestChainConfig = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, false, new(EthashConfig), nil, nil}
	TestRules       = TestChainConfig.Rules(new(big.Int), false)
)

// TrustedCheckpoint represents a set of post-processed trie roots (CHT and
// BloomTrie) associated with the appropriate section index and head hash. It is
// used to start light syncing from this checkpoint and avoid downloading the
// entire header chain while still being able to securely access old headers/logs.
type TrustedCheckpoint struct {
	SectionIndex uint64      `json:"sectionIndex"`
	SectionHead  common.Hash `json:"sectionHead"`
	CHTRoot      common.Hash `json:"chtRoot"`
	BloomRoot    common.Hash `json:"bloomRoot"`
}

// HashEqual returns an indicator comparing the itself hash with given one.
func (c *TrustedCheckpoint) HashEqual(hash common.Hash) bool {
	if c.Empty() {
		return hash == common.Hash{}
	}
	return c.Hash() == hash
}

// Hash returns the hash of checkpoint's four key fields(index, sectionHead, chtRoot and bloomTrieRoot).
func (c *TrustedCheckpoint) Hash() common.Hash {
	buf := make([]byte, 8+3*common.HashLength)
	binary.BigEndian.PutUint64(buf, c.SectionIndex)
	copy(buf[8:], c.SectionHead.Bytes())
	copy(buf[8+common.HashLength:], c.CHTRoot.Bytes())
	copy(buf[8+2*common.HashLength:], c.BloomRoot.Bytes())
	return crypto.Keccak256Hash(buf)
	var sectionIndex [8]byte
	binary.BigEndian.PutUint64(sectionIndex[:], c.SectionIndex)

	w := sha3.NewLegacyKeccak256()
	w.Write(sectionIndex[:])
	w.Write(c.SectionHead[:])
	w.Write(c.CHTRoot[:])
	w.Write(c.BloomRoot[:])

	var h common.Hash
	w.Sum(h[:0])
	return h
}

// Empty returns an indicator whether the checkpoint is regarded as empty.
func (c *TrustedCheckpoint) Empty() bool {
	return c.SectionHead == (common.Hash{}) || c.CHTRoot == (common.Hash{}) || c.BloomRoot == (common.Hash{})
}

// CheckpointOracleConfig represents a set of checkpoint contract(which acts as an oracle)
// config which used for light client checkpoint syncing.
type CheckpointOracleConfig struct {
	Address   common.Address   `json:"address"`
	Signers   []common.Address `json:"signers"`
	Threshold uint64           `json:"threshold"`
}

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock      *big.Int `json:"byzantiumBlock,omitempty"`      // Byzantium switch block (nil = no fork, 0 = already on byzantium)
	ConstantinopleBlock *big.Int `json:"constantinopleBlock,omitempty"` // Constantinople switch block (nil = no fork, 0 = already activated)
	PetersburgBlock     *big.Int `json:"petersburgBlock,omitempty"`     // Petersburg switch block (nil = same as Constantinople)
	IstanbulBlock       *big.Int `json:"istanbulBlock,omitempty"`       // Istanbul switch block (nil = no fork, 0 = already on istanbul)
	MuirGlacierBlock    *big.Int `json:"muirGlacierBlock,omitempty"`    // Eip-2384 (bomb delay) switch block (nil = no fork, 0 = already activated)
	BerlinBlock         *big.Int `json:"berlinBlock,omitempty"`         // Berlin switch block (nil = no fork, 0 = already on berlin)

	// This is a "fake" hardfork to fix an issue in a past block(#2509228).
	// By saying it is "fake", we mean it should behave as if it does not exist.
	// The hardfork should not reflect on the forkid.
	CVE_2021_39137Block *big.Int `json:"cve_2021_39137Block,omitempty"`

	// In Ishikari hardfork
	// introduces a new version of validators contract.
	IshikariBlock *big.Int `json:"ishikariBlock,omitempty"`

	// A patch for Ishikari hardfork
	// Fix minor bugs found on testnet
	IshikariPatch001Block *big.Int `json:"ishikariPatch001Block,omitempty"`

	// A patch for Ishikari hardfork
	// The punishment parameters for mainnet is determined in this hardfork
	IshikariPatch002Block *big.Int `json:"ishikariPatch002Block,omitempty"`
	YoloV3Block           *big.Int `json:"yoloV3Block,omitempty"` // YOLO v3: Gas repricings TODO @holiman add EIP references
	EWASMBlock            *big.Int `json:"ewasmBlock,omitempty"`  // EWASM switch block (nil = no fork, 0 = already activated)

	// @TODO We will never use the London Block
	// Move it to the end
	LondonBlock *big.Int `json:"londonBlock,omitempty"`

	// @TODO We will never use this hardfork
	ArrowGlacierBlock *big.Int `json:"arrowGlacielock,omitempty"`  // Eip-4345 (bomb delay) switch block (nil = no fork, 0 = already activated)
	GrayGlacierBlock  *big.Int `json:"grayGlacierBlock,omitempty"` // Eip-5133 (bomb delay) switch block (nil = no fork, 0 = already activated)
	// @TODO We will never use the merge h
	MergeNetsplitBlock *big.Int `json:"mergeNetsplitBlock,omitempty"` // Virtual fork after The Merge to use as a network splitter

	ShanghaiBlock *big.Int `json:"shanghaiBlock,omitempty"` // Shanghai switch block (nil = no fork, 0 = already on shanghai)
	CancunBlock   *big.Int `json:"cancunBlock,omitempty"`   // Cancun switch block (nil = no fork, 0 = already on cancun)

	// @Amazon Hardfork block
	AmazonBlock *big.Int `json:"amazonBlock,omitempty"` // Amazon switch block (nil = no fork, x = already on amazon)

	// @TODO: We will not enable TTD
	// The hardfork should not reflect on the forkid.
	// TerminalTotalDifficulty is the amount of total difficulty reached by
	// the network that triggers the consensus upgrade.
	TerminalTotalDifficulty *big.Int `json:"terminalTotalDifficulty,omitempty"`

	// @TODO This is always false
	// TerminalTotalDifficultyPassed is a flag specifying that the network already
	// passed the terminal total difficulty. Its purpose is to disable legacy sync
	// even without having seen the TTD locally (safer long term).
	TerminalTotalDifficultyPassed bool `json:"terminalTotalDifficultyPassed,omitempty"`

	// Various consensus engines
	Ethash *EthashConfig `json:"ethash,omitempty"`
	Clique *CliqueConfig `json:"clique,omitempty"`
	POSA   *POSAConfig   `json:"posa,omitempty"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// POSAConfig is the consensus engine configs for proof-of-stake-authority based sealing.
type POSAConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint

	// Ishikari hardfork related configs
	// Ishikari initial validators
	IshikariInitialValidators []common.Address `json:"ishikariInitialValidators"`
	IshikariInitialManagers   []common.Address `json:"ishikariInitialManagers"`
	// Ishikari admin multisig Address
	IshikariAdminMultiSig common.Address `json:"ishikariAdminAddress"`
}

// Validate POSA Contraints
func (c *POSAConfig) Validate(chainCfg *ChainConfig) error {

	if c.Period == 0 {
		return fmt.Errorf("POSAConfig.Period should not be 0")
	}

	if c.Epoch < 2 {
		return fmt.Errorf("POSAConfig.Epoch should be not be less than 2")
	}

	if chainCfg.IshikariBlock == nil {
		// if Ishikari hardfork is not enabled yet,
		// we don't need to verify other fields at this moment.
		return nil
	}

	if len(c.IshikariInitialManagers) < 1 {
		return fmt.Errorf("length of POSAConfig.V2InitialManagers must not be less than 1")
	}

	if len(c.IshikariInitialManagers) != len(c.IshikariInitialValidators) {
		return fmt.Errorf("numbers of initial validators & initial managers do not match (%v!=%v)",
			len(c.IshikariInitialManagers), len(c.IshikariInitialValidators))
	}

	// The hardfork should happen at the last block of some epoch
	if chainCfg.IshikariBlock != nil && ((chainCfg.IshikariBlock.Uint64()+1)%c.Epoch != 0) {
		return fmt.Errorf("IshikariBlock should be the last block of some epoch")
	}

	return nil
}

// String implements the stringer interface, returning the consensus engine details.
func (c *POSAConfig) String() string {
	d, _ := json.Marshal(c)
	return fmt.Sprintf("posa(%v)", string(d))
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Ethash != nil:
		engine = c.Ethash
	case c.Clique != nil:
		engine = c.Clique
	case c.POSA != nil:
		engine = c.POSA
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v Homestead: %v DAO: %v DAOSupport: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v Constantinople: %v Petersburg: %v Istanbul: %v, Muir Glacier: %v, Berlin: %v, cve_2021_39137Block:%v, Ishikari: %v, IshikariPatch: %v, Amazon: %v, Engine: %v}",
		c.ChainID,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		c.ConstantinopleBlock,
		c.PetersburgBlock,
		c.IstanbulBlock,
		c.MuirGlacierBlock,
		c.BerlinBlock,
		c.CVE_2021_39137Block,
		c.IshikariBlock,
		c.IshikariPatch001Block,
		c.AmazonBlock,
		engine,
	)
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isForked(c.HomesteadBlock, num)
}

// IsDAOFork returns whether num is either equal to the DAO fork block or greater.
func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
	return isForked(c.DAOForkBlock, num)
}

// IsEIP150 returns whether num is either equal to the EIP150 fork block or greater.
func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isForked(c.EIP150Block, num)
}

// IsEIP155 returns whether num is either equal to the EIP155 fork block or greater.
func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isForked(c.EIP155Block, num)
}

// IsEIP158 returns whether num is either equal to the EIP158 fork block or greater.
func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isForked(c.EIP158Block, num)
}

// IsByzantium returns whether num is either equal to the Byzantium fork block or greater.
func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	return isForked(c.ByzantiumBlock, num)
}

// IsConstantinople returns whether num is either equal to the Constantinople fork block or greater.
func (c *ChainConfig) IsConstantinople(num *big.Int) bool {
	return isForked(c.ConstantinopleBlock, num)
}

// IsMuirGlacier returns whether num is either equal to the Muir Glacier (EIP-2384) fork block or greater.
func (c *ChainConfig) IsMuirGlacier(num *big.Int) bool {
	return isForked(c.MuirGlacierBlock, num)
}

// IsPetersburg returns whether num is either
// - equal to or greater than the PetersburgBlock fork block,
// - OR is nil, and Constantinople is active
func (c *ChainConfig) IsPetersburg(num *big.Int) bool {
	return isForked(c.PetersburgBlock, num) || c.PetersburgBlock == nil && isForked(c.ConstantinopleBlock, num)
}

// IsIstanbul returns whether num is either equal to the Istanbul fork block or greater.
func (c *ChainConfig) IsIstanbul(num *big.Int) bool {
	return isForked(c.IstanbulBlock, num)
}

// IsBerlin returns whether num is either equal to the Berlin fork block or greater.
func (c *ChainConfig) IsBerlin(num *big.Int) bool {
	return isForked(c.BerlinBlock, num)
}

// IsLondon returns whether num is either equal to the London fork block or greater.
func (c *ChainConfig) IsLondon(num *big.Int) bool {
	return isForked(c.LondonBlock, num)
}

// IsArrowGlacier returns whether num is either equal to the Arrow Glacier (EIP-4345) fork block or greater.
func (c *ChainConfig) IsArrowGlacier(num *big.Int) bool {
	return isForked(c.ArrowGlacierBlock, num)
}

// IsGrayGlacier returns whether num is either equal to the Gray Glacier (EIP-5133) fork block or greater.
func (c *ChainConfig) IsGrayGlacier(num *big.Int) bool {
	return isForked(c.GrayGlacierBlock, num)
}

// IsTerminalPoWBlock returns whether the given block is the last block of PoW stage.
func (c *ChainConfig) IsTerminalPoWBlock(parentTotalDiff *big.Int, totalDiff *big.Int) bool {
	if c.TerminalTotalDifficulty == nil {
		return false
	}
	return parentTotalDiff.Cmp(c.TerminalTotalDifficulty) < 0 && totalDiff.Cmp(c.TerminalTotalDifficulty) >= 0
}

// IsEWASM returns whether num represents a block number after the EWASM fork
func (c *ChainConfig) IsEWASM(num *big.Int) bool {
	return isForked(c.EWASMBlock, num)
}

// IsShanghai returns whether num is either equal to the Shanghai fork block or greater.
func (c *ChainConfig) IsShanghai(num *big.Int) bool {
	return isForked(c.ShanghaiBlock, num)
}

// is Ishikari hardfork enabled ?
func (c *ChainConfig) IsNUSAIshikari(num *big.Int) bool {
	return isForked(c.IshikariBlock, num)
}

// IsCancun returns whether num is either equal to the Cancun fork block or greater.
func (c *ChainConfig) IsCancun(num *big.Int) bool {
	return isForked(c.CancunBlock, num)
}

// is the block number "num" when Ishikari hardfork happens ?
func (c *ChainConfig) IsIshikariHardforkBlock(num *big.Int) bool {
	if num == nil || c.IshikariBlock == nil {
		return false
	}
	return num.Cmp(c.IshikariBlock) == 0
}

// is the block number "num" when IshikariPatch001 hardfork happens ?
func (c *ChainConfig) IsIshikariPatch001HardforkBlock(num *big.Int) bool {
	if num == nil || c.IshikariPatch001Block == nil {
		return false
	}
	return num.Cmp(c.IshikariPatch001Block) == 0
}

// is the block number "num" when IshikariPatch002 hardfork happens ?
func (c *ChainConfig) IsIshikariPatch002HardforkBlock(num *big.Int) bool {
	if num == nil || c.IshikariPatch002Block == nil {
		return false
	}
	return num.Cmp(c.IshikariPatch002Block) == 0
}

// only return true is the block number is exactly the Amazon Switch Block
func (c *ChainConfig) IsAmazonHardforkBlock(num *big.Int) bool {
	if num == nil || c.AmazonBlock == nil {
		return false
	}
	return num.Cmp(c.AmazonBlock) == 0
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

// CheckConfigForkOrder checks that we don't "skip" any forks, geth isn't pluggable enough
// to guarantee that forks can be implemented in a different order than on official networks
func (c *ChainConfig) CheckConfigForkOrder() error {
	type fork struct {
		name     string
		block    *big.Int
		optional bool // if true, the fork may be nil and next fork is still allowed
	}
	var lastFork fork
	for _, cur := range []fork{
		{name: "homesteadBlock", block: c.HomesteadBlock},
		{name: "daoForkBlock", block: c.DAOForkBlock, optional: true},
		{name: "eip150Block", block: c.EIP150Block},
		{name: "eip155Block", block: c.EIP155Block},
		{name: "eip158Block", block: c.EIP158Block},
		{name: "byzantiumBlock", block: c.ByzantiumBlock},
		{name: "constantinopleBlock", block: c.ConstantinopleBlock},
		{name: "petersburgBlock", block: c.PetersburgBlock},
		{name: "istanbulBlock", block: c.IstanbulBlock},
		{name: "muirGlacierBlock", block: c.MuirGlacierBlock, optional: true},
		{name: "berlinBlock", block: c.BerlinBlock},
		{name: "ishikariBlock", block: c.IshikariBlock},
		{name: "ishikariPatch001Block", block: c.IshikariPatch001Block},
		{name: "ishikariPatch002Block", block: c.IshikariPatch002Block},
		{name: "amazonBlock", block: c.AmazonBlock}, // amazon block should come after all Ishikari blocks
	} {
		if lastFork.name != "" {
			// Next one must be higher number
			if lastFork.block == nil && cur.block != nil {
				return fmt.Errorf("unsupported fork ordering: %v not enabled, but %v enabled at %v",
					lastFork.name, cur.name, cur.block)
			}
			if lastFork.block != nil && cur.block != nil {
				if lastFork.block.Cmp(cur.block) > 0 {
					return fmt.Errorf("unsupported fork ordering: %v enabled at %v, but %v enabled at %v",
						lastFork.name, lastFork.block, cur.name, cur.block)
				}
			}
		}
		// If it was optional and not set, then ignore it
		if !cur.optional || cur.block != nil {
			lastFork = cur
		}
	}
	return nil
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if isForkIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, head) {
		return newCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	}
	if isForkIncompatible(c.DAOForkBlock, newcfg.DAOForkBlock, head) {
		return newCompatError("DAO fork block", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if c.IsDAOFork(head) && c.DAOForkSupport != newcfg.DAOForkSupport {
		return newCompatError("DAO fork support flag", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if isForkIncompatible(c.EIP150Block, newcfg.EIP150Block, head) {
		return newCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	}
	if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkIncompatible(c.EIP158Block, newcfg.EIP158Block, head) {
		return newCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	}
	if c.IsEIP158(head) && !configNumEqual(c.ChainID, newcfg.ChainID) {
		return newCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	}
	if isForkIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, head) {
		return newCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	}
	if isForkIncompatible(c.ConstantinopleBlock, newcfg.ConstantinopleBlock, head) {
		return newCompatError("Constantinople fork block", c.ConstantinopleBlock, newcfg.ConstantinopleBlock)
	}
	if isForkIncompatible(c.PetersburgBlock, newcfg.PetersburgBlock, head) {
		// the only case where we allow Petersburg to be set in the past is if it is equal to Constantinople
		// mainly to satisfy fork ordering requirements which state that Petersburg fork be set if Constantinople fork is set
		if isForkIncompatible(c.ConstantinopleBlock, newcfg.PetersburgBlock, head) {
			return newCompatError("Petersburg fork block", c.PetersburgBlock, newcfg.PetersburgBlock)
		}
	}
	if isForkIncompatible(c.IstanbulBlock, newcfg.IstanbulBlock, head) {
		return newCompatError("Istanbul fork block", c.IstanbulBlock, newcfg.IstanbulBlock)
	}
	if isForkIncompatible(c.MuirGlacierBlock, newcfg.MuirGlacierBlock, head) {
		return newCompatError("Muir Glacier fork block", c.MuirGlacierBlock, newcfg.MuirGlacierBlock)
	}
	if isForkIncompatible(c.BerlinBlock, newcfg.BerlinBlock, head) {
		return newCompatError("Berlin fork block", c.BerlinBlock, newcfg.BerlinBlock)
	}
	if isForkIncompatible(c.YoloV3Block, newcfg.YoloV3Block, head) {
		return newCompatError("YOLOv3 fork block", c.YoloV3Block, newcfg.YoloV3Block)
	}
	if isForkIncompatible(c.EWASMBlock, newcfg.EWASMBlock, head) {
		return newCompatError("ewasm fork block", c.EWASMBlock, newcfg.EWASMBlock)
	}

	if isForkIncompatible(c.IshikariBlock, newcfg.IshikariBlock, head) {
		return newCompatError("Ishikari fork block", c.IshikariBlock, newcfg.IshikariBlock)
	}

	if isForkIncompatible(c.IshikariPatch001Block, newcfg.IshikariPatch001Block, head) {
		return newCompatError("IshikariPatch001 fork block", c.IshikariPatch001Block, newcfg.IshikariPatch001Block)
	}

	if isForkIncompatible(c.IshikariPatch002Block, newcfg.IshikariPatch002Block, head) {
		return newCompatError("IshikariPatch002 fork block", c.IshikariPatch002Block, newcfg.IshikariPatch002Block)
	}

	if isForkIncompatible(c.AmazonBlock, newcfg.AmazonBlock, head) {
		return newCompatError("Amazon fork block", c.AmazonBlock, newcfg.AmazonBlock)
	}
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntactic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainID                                                 *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158               bool
	IsByzantium, IsConstantinople, IsPetersburg, IsIstanbul bool
	IsBerlin                                                bool
	IsIshikari                                              bool
	IsCVE_2021_39137BlockPassed                             bool
	IsLondon                                                bool
	IsMerge, IsShanghai, isCancun                           bool
}

// Rules ensures c's ChainID is not nil.
func (c *ChainConfig) Rules(num *big.Int, isMerge bool) Rules {
	chainID := c.ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	return Rules{
		ChainID:                     new(big.Int).Set(chainID),
		IsHomestead:                 c.IsHomestead(num),
		IsEIP150:                    c.IsEIP150(num),
		IsEIP155:                    c.IsEIP155(num),
		IsEIP158:                    c.IsEIP158(num),
		IsByzantium:                 c.IsByzantium(num),
		IsConstantinople:            c.IsConstantinople(num),
		IsPetersburg:                c.IsPetersburg(num),
		IsIstanbul:                  c.IsIstanbul(num),
		IsBerlin:                    c.IsBerlin(num),
		IsIshikari:                  c.IsNUSAIshikari(num),
		IsCVE_2021_39137BlockPassed: c.CVE_2021_39137Block == nil || c.CVE_2021_39137Block.Cmp(num) < 0,
		IsLondon:                    c.IsLondon(num),
		IsMerge:                     isMerge,
		IsShanghai:                  c.IsShanghai(num),
		isCancun:                    c.IsCancun(num),
	}
}
