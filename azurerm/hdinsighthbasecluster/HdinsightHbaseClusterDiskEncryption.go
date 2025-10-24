// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthbasecluster


type HdinsightHbaseClusterDiskEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/hdinsight_hbase_cluster#encryption_algorithm HdinsightHbaseCluster#encryption_algorithm}.
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/hdinsight_hbase_cluster#encryption_at_host_enabled HdinsightHbaseCluster#encryption_at_host_enabled}.
	EncryptionAtHostEnabled interface{} `field:"optional" json:"encryptionAtHostEnabled" yaml:"encryptionAtHostEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/hdinsight_hbase_cluster#key_vault_key_id HdinsightHbaseCluster#key_vault_key_id}.
	KeyVaultKeyId *string `field:"optional" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/hdinsight_hbase_cluster#key_vault_managed_identity_id HdinsightHbaseCluster#key_vault_managed_identity_id}.
	KeyVaultManagedIdentityId *string `field:"optional" json:"keyVaultManagedIdentityId" yaml:"keyVaultManagedIdentityId"`
}

