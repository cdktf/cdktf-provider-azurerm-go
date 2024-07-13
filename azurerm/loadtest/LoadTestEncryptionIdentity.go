// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadtest


type LoadTestEncryptionIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/load_test#identity_id LoadTest#identity_id}.
	IdentityId *string `field:"required" json:"identityId" yaml:"identityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/load_test#type LoadTest#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

