package config

import (
	"time"

	"github.com/spf13/viper"
)

type (
	// Config provides access to app configs.
	Config interface {
		MerkleEndpoint() string
		MerkleContract() string
		MerkleFilePath() string
		DaemonEndpoint() string
		DaemonContract() string
		DaemonInterval() time.Duration
	}
	// config stores all application configs.
	config struct {
		Merkle struct {
			Endpoint string
			Contract string
			FilePath string
		}
		Daemon struct {
			Endpoint string
			Contract string
			Interval time.Duration
		}
	}
)

// New returns a config provider.
func New(configFilePath string) (Config, error) {
	viper.SetConfigFile(configFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg config
	err = viper.UnmarshalExact(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// MerkleEndpoint returns api endpoint URL for MerkleProof problem
func (cfg config) MerkleEndpoint() string {
	return cfg.Merkle.Endpoint
}

// MerkleContract returns smart contract hex for MerkleProof problem.
func (cfg config) MerkleContract() string {
	return cfg.Merkle.Contract
}

// MerkleFilePath returns a path to the file with source data.
func (cfg config) MerkleFilePath() string {
	return cfg.Merkle.FilePath
}

// DaemonEndpoint returns api endpoint URL for Daemon problem.
func (cfg config) DaemonEndpoint() string {
	return cfg.Daemon.Endpoint
}

// DaemonContract returns smart contract hex for Daemon problem.
func (cfg config) DaemonContract() string {
	return cfg.Daemon.Contract
}

// DaemonContract returns interval for daemon ticker.
func (cfg config) DaemonInterval() time.Duration {
	return cfg.Daemon.Interval
}