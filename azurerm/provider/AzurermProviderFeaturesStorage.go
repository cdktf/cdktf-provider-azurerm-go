// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeaturesStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs#data_plane_available AzurermProvider#data_plane_available}.
	DataPlaneAvailable interface{} `field:"optional" json:"dataPlaneAvailable" yaml:"dataPlaneAvailable"`
}

