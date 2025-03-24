// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool


type BatchPoolIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/batch_pool#identity_ids BatchPool#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/batch_pool#type BatchPool#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

