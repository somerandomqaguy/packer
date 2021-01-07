package hcl2template

import (
	"crypto/sha256"
	"fmt"
	"runtime"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/packer/packer-plugin-sdk/plugin"
	plugingetter "github.com/hashicorp/packer/packer/plugin-getter"
)

// PluginRequirements returns a sorted list of plugin requirements.
func (cfg *PackerConfig) PluginRequirements() (plugingetter.Requirements, hcl.Diagnostics) {

	var diags hcl.Diagnostics
	var reqs plugingetter.Requirements
	reqPluginsBlocks := cfg.Packer.RequiredPlugins

	// Take all required plugins, make sure there are no conflicting blocks
	// and append them to the list.
	uniq := map[string]*RequiredPlugin{}
	for _, requiredPluginsBlock := range reqPluginsBlocks {
		for name, block := range requiredPluginsBlock.RequiredPlugins {

			if previouslySeenBlock, found := uniq[name]; found {
				diags = append(diags, &hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  fmt.Sprintf("Duplicate required_plugin.%q block", name),
					Detail: fmt.Sprintf("Block previously seen at %s is already named %q.\n", previouslySeenBlock.DeclRange, name) +
						"Names at the left hand side of required_plugins are made available to use in your HCL2 configurations.\n" +
						"To allow to calling to their features correctly two plugins have to have different accessors.",
					Context: &block.DeclRange,
				})
				continue
			}

			reqs = append(reqs, &plugingetter.Requirement{
				Accessor:           name,
				Identifier:         block.Type,
				VersionConstraints: block.Requirement.Required,
			})
			uniq[name] = block
		}

	}

	return reqs, diags
}

func (cfg *PackerConfig) detectPluginBinaries() hcl.Diagnostics {
	opts := plugingetter.ListInstallationsOptions{
		FromFolders: cfg.parser.PluginStore.KnownPluginFolders,
		BinaryInstallationOptions: plugingetter.BinaryInstallationOptions{
			OS:        runtime.GOOS,
			ARCH:      runtime.GOARCH,
			Extension: plugin.FileExtension,
			Checksummers: []plugingetter.Checksummer{
				{Type: "sha256", Hash: sha256.New()},
			},
		},
	}

	pluginReqs, diags := cfg.PluginRequirements()
	if diags.HasErrors() {
		return diags
	}

	for _, pluginRequirement := range pluginReqs {
		installs, err := pluginRequirement.ListInstallations(opts)
		if err != nil {
			panic(err) // fill diag error
		}
		if len(installs) == 0 {
			panic("no plugin installed for " + pluginRequirement.Identifier.ForDisplay())
		}
		install := installs[0]
		err = cfg.parser.PluginStore.DiscoverPlugin(pluginRequirement.Accessor, install.BinaryPath)
		if err != nil {
			panic(err)
		}
	}

	return diags
}
