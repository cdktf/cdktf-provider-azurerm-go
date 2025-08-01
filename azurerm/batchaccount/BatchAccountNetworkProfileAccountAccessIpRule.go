// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchaccount


type BatchAccountNetworkProfileAccountAccessIpRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/batch_account#ip_range BatchAccount#ip_range}.
	IpRange *string `field:"required" json:"ipRange" yaml:"ipRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/batch_account#action BatchAccount#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
}

