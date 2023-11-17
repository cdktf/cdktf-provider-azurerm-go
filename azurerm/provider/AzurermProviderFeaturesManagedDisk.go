// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeaturesManagedDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs#expand_without_downtime AzurermProvider#expand_without_downtime}.
	ExpandWithoutDowntime interface{} `field:"optional" json:"expandWithoutDowntime" yaml:"expandWithoutDowntime"`
}

