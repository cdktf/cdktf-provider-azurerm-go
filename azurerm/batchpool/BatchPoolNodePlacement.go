// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool


type BatchPoolNodePlacement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/batch_pool#policy BatchPool#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

