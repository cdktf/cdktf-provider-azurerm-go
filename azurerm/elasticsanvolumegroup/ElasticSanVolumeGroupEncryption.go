// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsanvolumegroup


type ElasticSanVolumeGroupEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/elastic_san_volume_group#key_vault_key_id ElasticSanVolumeGroup#key_vault_key_id}.
	KeyVaultKeyId *string `field:"required" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/elastic_san_volume_group#user_assigned_identity_id ElasticSanVolumeGroup#user_assigned_identity_id}.
	UserAssignedIdentityId *string `field:"optional" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
}

