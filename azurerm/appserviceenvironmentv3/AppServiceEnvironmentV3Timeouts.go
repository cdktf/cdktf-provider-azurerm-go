// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appserviceenvironmentv3


type AppServiceEnvironmentV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/app_service_environment_v3#create AppServiceEnvironmentV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/app_service_environment_v3#delete AppServiceEnvironmentV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/app_service_environment_v3#read AppServiceEnvironmentV3#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/app_service_environment_v3#update AppServiceEnvironmentV3#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

