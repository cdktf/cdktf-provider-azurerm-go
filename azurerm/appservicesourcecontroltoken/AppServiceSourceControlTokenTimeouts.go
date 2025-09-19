// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appservicesourcecontroltoken


type AppServiceSourceControlTokenTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/app_service_source_control_token#create AppServiceSourceControlToken#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/app_service_source_control_token#delete AppServiceSourceControlToken#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/app_service_source_control_token#read AppServiceSourceControlToken#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/app_service_source_control_token#update AppServiceSourceControlToken#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

