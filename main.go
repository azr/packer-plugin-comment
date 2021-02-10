package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/hashicorp/packer-plugin-sdk/version"
)

var (
	// The main version number that is being run at the moment.
	Version = "2.0.0"

	// A pre-release marker for the Version. If this is "" (empty string)
	// then it means that it is a final release. Otherwise, this is a pre-release
	// such as "dev" (in development), "beta", "rc1", etc.
	VersionPrerelease = ""

	PluginVersion = version.InitializePluginVersion(Version, VersionPrerelease)
)

func main() {
	pps := plugin.NewSet()

	pps.RegisterProvisioner(plugin.DEFAULT_NAME, new(CommentProvisioner))
	pps.SetVersion(PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
