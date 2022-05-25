// Copyright © 2022 Obol Labs Inc.
//
// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of  MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for
// more details.
//
// You should have received a copy of the GNU General Public License along with
// this program.  If not, see <http://www.gnu.org/licenses/>.

//nolint:deadcode,varcheck // Busy implementing
package compose

const (
	version           = "obol/charon/compose/1.0.0"
	configFile        = "config.json"
	defaultImageTag   = "latest"
	defaultBeaconNode = "mock"
	defaultKeyGen     = keyGenCreate
	defaultNumVals    = 1
	defaultNumNodes   = 4
	defaultThreshold  = 3
	defaultFeatureSet = "alpha"

	charonImage      = "ghcr.io/obolnetwork/charon"
	localBinary      = "/compose/charon"
	containerBinary  = "/usr/local/bin/charon"
	cmdRun           = "run"
	cmdDKG           = "dkg"
	cmdCreateCluster = "[create,cluster]"
	cmdCreateDKG     = "[create,dkg]"
)

// vcType defines a validator client type.
type vcType string

const (
	vcMock       vcType = "mock"
	vcTeku       vcType = "teku"
	vcLighthouse vcType = "lighthouse"
)

// KeyGen defines a key generation process.
type KeyGen string

const (
	keyGenDKG    KeyGen = "dkg"
	keyGenCreate KeyGen = "create"
)

// step defines the current completed compose step.
type step string

const (
	stepNew     step = "new"
	stepDefined step = "defined"
	stepLocked  step = "locked"
)

// Config defines a local compose cluster; including both keygen and running a cluster.
type Config struct {
	// Version defines the compose config version.
	Version string `json:"version"`

	// Step defines the current completed compose step.
	Step step `json:"step"`

	// NumNodes is the number of charon nodes in the cluster.
	NumNodes int `json:"num_nodes"`

	// Threshold required for signature reconstruction. Defaults to safe value for number of nodes/peers.
	Threshold int `json:"threshold"`

	// NumValidators is the number of DVs (n*32ETH) to be created in the cluster lock file.
	NumValidators int `json:"num_validators"`

	// ImageTag defines the charon docker image tag: ghcr.io/obolnetwork/charon:{ImageTag}.
	ImageTag string `json:"image_tag"`

	// BuildLocal enables building a local charon binary from source and using that in the containers.
	BuildLocal bool `json:"build_local"`

	// KeyGen defines the key generation process.
	KeyGen KeyGen `json:"key_gen"`

	// SplitKeysDir directory containing keys to split for keygen==create.
	SplitKeysDir string `json:"split_keys_dir"`

	// BeaconNode url endpoint or "mock" for simnet.
	BeaconNode string `json:"beacon_node"`

	// VCs define the types of validator clients to use.
	VCs []vcType `json:"validator_clients"`

	// FeatureSet defines the minimum feature set to enable.
	FeatureSet string `json:"feature_set"`
}

// entrypoint returns the path to the charon binary based on the BuildLocal field.
func (c Config) entrypoint() string {
	if c.BuildLocal {
		return localBinary
	}

	return containerBinary
}

// NewDefaultConfig returns a new default config.
func NewDefaultConfig() Config {
	return Config{
		Version:       version,
		NumNodes:      defaultNumNodes,
		Threshold:     defaultThreshold,
		NumValidators: defaultNumVals,
		ImageTag:      defaultImageTag,
		VCs:           []vcType{vcTeku, vcLighthouse, vcMock},
		KeyGen:        defaultKeyGen,
		BeaconNode:    defaultBeaconNode,
		Step:          stepDefined,
		FeatureSet:    defaultFeatureSet,
	}
}