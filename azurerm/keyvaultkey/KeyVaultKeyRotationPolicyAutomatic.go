// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultkey


type KeyVaultKeyRotationPolicyAutomatic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/key_vault_key#time_after_creation KeyVaultKey#time_after_creation}.
	TimeAfterCreation *string `field:"optional" json:"timeAfterCreation" yaml:"timeAfterCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/key_vault_key#time_before_expiry KeyVaultKey#time_before_expiry}.
	TimeBeforeExpiry *string `field:"optional" json:"timeBeforeExpiry" yaml:"timeBeforeExpiry"`
}

