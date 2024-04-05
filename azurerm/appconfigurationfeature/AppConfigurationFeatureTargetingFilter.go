// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appconfigurationfeature


type AppConfigurationFeatureTargetingFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/app_configuration_feature#default_rollout_percentage AppConfigurationFeature#default_rollout_percentage}.
	DefaultRolloutPercentage *float64 `field:"required" json:"defaultRolloutPercentage" yaml:"defaultRolloutPercentage"`
	// groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/app_configuration_feature#groups AppConfigurationFeature#groups}
	Groups interface{} `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/app_configuration_feature#users AppConfigurationFeature#users}.
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}

