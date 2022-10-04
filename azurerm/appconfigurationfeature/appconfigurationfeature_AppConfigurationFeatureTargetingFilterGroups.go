package appconfigurationfeature


type AppConfigurationFeatureTargetingFilterGroups struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_configuration_feature#name AppConfigurationFeature#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_configuration_feature#rollout_percentage AppConfigurationFeature#rollout_percentage}.
	RolloutPercentage *float64 `field:"required" json:"rolloutPercentage" yaml:"rolloutPercentage"`
}

