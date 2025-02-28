// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightkafkacluster


type HdinsightKafkaClusterDiskEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#encryption_algorithm HdinsightKafkaCluster#encryption_algorithm}.
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#encryption_at_host_enabled HdinsightKafkaCluster#encryption_at_host_enabled}.
	EncryptionAtHostEnabled interface{} `field:"optional" json:"encryptionAtHostEnabled" yaml:"encryptionAtHostEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#key_vault_key_id HdinsightKafkaCluster#key_vault_key_id}.
	KeyVaultKeyId *string `field:"optional" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#key_vault_managed_identity_id HdinsightKafkaCluster#key_vault_managed_identity_id}.
	KeyVaultManagedIdentityId *string `field:"optional" json:"keyVaultManagedIdentityId" yaml:"keyVaultManagedIdentityId"`
}

