// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabriccluster


type ServiceFabricClusterDiagnosticsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/service_fabric_cluster#blob_endpoint ServiceFabricCluster#blob_endpoint}.
	BlobEndpoint *string `field:"required" json:"blobEndpoint" yaml:"blobEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/service_fabric_cluster#protected_account_key_name ServiceFabricCluster#protected_account_key_name}.
	ProtectedAccountKeyName *string `field:"required" json:"protectedAccountKeyName" yaml:"protectedAccountKeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/service_fabric_cluster#queue_endpoint ServiceFabricCluster#queue_endpoint}.
	QueueEndpoint *string `field:"required" json:"queueEndpoint" yaml:"queueEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/service_fabric_cluster#storage_account_name ServiceFabricCluster#storage_account_name}.
	StorageAccountName *string `field:"required" json:"storageAccountName" yaml:"storageAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/service_fabric_cluster#table_endpoint ServiceFabricCluster#table_endpoint}.
	TableEndpoint *string `field:"required" json:"tableEndpoint" yaml:"tableEndpoint"`
}

