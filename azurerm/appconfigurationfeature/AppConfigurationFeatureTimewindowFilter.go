package appconfigurationfeature


type AppConfigurationFeatureTimewindowFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_configuration_feature#end AppConfigurationFeature#end}.
	End *string `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_configuration_feature#start AppConfigurationFeature#start}.
	Start *string `field:"optional" json:"start" yaml:"start"`
}

