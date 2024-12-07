// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetwork


type VirtualNetworkEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/virtual_network#enforcement VirtualNetwork#enforcement}.
	Enforcement *string `field:"required" json:"enforcement" yaml:"enforcement"`
}

