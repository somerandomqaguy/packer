// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package ami

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	Filters    map[string]string     `cty:"filters" hcl:"filters"`
	Filter     []config.FlatKeyValue `cty:"filter" hcl:"filter"`
	Owners     []string              `cty:"owners" hcl:"owners"`
	MostRecent *bool                 `mapstructure:"most_recent" cty:"most_recent" hcl:"most_recent"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"filters":     &hcldec.AttrSpec{Name: "filters", Type: cty.Map(cty.String), Required: false},
		"filter":      &hcldec.BlockListSpec{TypeName: "filter", Nested: hcldec.ObjectSpec((*config.FlatKeyValue)(nil).HCL2Spec())},
		"owners":      &hcldec.AttrSpec{Name: "owners", Type: cty.List(cty.String), Required: false},
		"most_recent": &hcldec.AttrSpec{Name: "most_recent", Type: cty.Bool, Required: false},
	}
	return s
}
