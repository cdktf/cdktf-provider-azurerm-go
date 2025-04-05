// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageobjectreplication


type StorageObjectReplicationRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/storage_object_replication#destination_container_name StorageObjectReplication#destination_container_name}.
	DestinationContainerName *string `field:"required" json:"destinationContainerName" yaml:"destinationContainerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/storage_object_replication#source_container_name StorageObjectReplication#source_container_name}.
	SourceContainerName *string `field:"required" json:"sourceContainerName" yaml:"sourceContainerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/storage_object_replication#copy_blobs_created_after StorageObjectReplication#copy_blobs_created_after}.
	CopyBlobsCreatedAfter *string `field:"optional" json:"copyBlobsCreatedAfter" yaml:"copyBlobsCreatedAfter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/storage_object_replication#filter_out_blobs_with_prefix StorageObjectReplication#filter_out_blobs_with_prefix}.
	FilterOutBlobsWithPrefix *[]*string `field:"optional" json:"filterOutBlobsWithPrefix" yaml:"filterOutBlobsWithPrefix"`
}

