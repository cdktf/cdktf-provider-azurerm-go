package appconfigurationfeature


type AppConfigurationFeatureTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_configuration_feature#create AppConfigurationFeature#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_configuration_feature#delete AppConfigurationFeature#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_configuration_feature#read AppConfigurationFeature#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_configuration_feature#update AppConfigurationFeature#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
