// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterDiskEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/hdinsight_interactive_query_cluster#encryption_algorithm HdinsightInteractiveQueryCluster#encryption_algorithm}.
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/hdinsight_interactive_query_cluster#encryption_at_host_enabled HdinsightInteractiveQueryCluster#encryption_at_host_enabled}.
	EncryptionAtHostEnabled interface{} `field:"optional" json:"encryptionAtHostEnabled" yaml:"encryptionAtHostEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/hdinsight_interactive_query_cluster#key_vault_key_id HdinsightInteractiveQueryCluster#key_vault_key_id}.
	KeyVaultKeyId *string `field:"optional" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/hdinsight_interactive_query_cluster#key_vault_managed_identity_id HdinsightInteractiveQueryCluster#key_vault_managed_identity_id}.
	KeyVaultManagedIdentityId *string `field:"optional" json:"keyVaultManagedIdentityId" yaml:"keyVaultManagedIdentityId"`
}

