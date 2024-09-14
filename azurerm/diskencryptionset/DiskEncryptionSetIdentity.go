// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package diskencryptionset


type DiskEncryptionSetIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/disk_encryption_set#type DiskEncryptionSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/disk_encryption_set#identity_ids DiskEncryptionSet#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

