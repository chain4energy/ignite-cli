package testdata

import (
	"testing"

	"github.com/ignite/cli/ignite/chainconfig"
	"github.com/ignite/cli/ignite/chainconfig/config"
	networkconfigTestData "github.com/ignite/cli/ignite/chainconfig/networkconfig/testdata"
	v0testdata "github.com/ignite/cli/ignite/chainconfig/v0/testdata"
	v1testdata "github.com/ignite/cli/ignite/chainconfig/v1/testdata"
)

var Versions = map[config.Version][]byte{
	0: v0testdata.ConfigYAML,
	1: v1testdata.ConfigYAML,
}

var NetworkConfig = networkconfigTestData.ConfigYAML

func GetLatestConfig(t *testing.T) *chainconfig.Config {
	return v1testdata.GetConfig(t)
}

func GetLatestNetworkConfig(t *testing.T) *chainconfig.Config {
	return networkconfigTestData.GetConfig(t)
}
