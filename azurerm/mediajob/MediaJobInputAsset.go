// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediajob


type MediaJobInputAsset struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/media_job#name MediaJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/media_job#label MediaJob#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
}

