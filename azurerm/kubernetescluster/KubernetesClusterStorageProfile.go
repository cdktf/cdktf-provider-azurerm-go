// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterStorageProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/kubernetes_cluster#blob_driver_enabled KubernetesCluster#blob_driver_enabled}.
	BlobDriverEnabled interface{} `field:"optional" json:"blobDriverEnabled" yaml:"blobDriverEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/kubernetes_cluster#disk_driver_enabled KubernetesCluster#disk_driver_enabled}.
	DiskDriverEnabled interface{} `field:"optional" json:"diskDriverEnabled" yaml:"diskDriverEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/kubernetes_cluster#file_driver_enabled KubernetesCluster#file_driver_enabled}.
	FileDriverEnabled interface{} `field:"optional" json:"fileDriverEnabled" yaml:"fileDriverEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/kubernetes_cluster#snapshot_controller_enabled KubernetesCluster#snapshot_controller_enabled}.
	SnapshotControllerEnabled interface{} `field:"optional" json:"snapshotControllerEnabled" yaml:"snapshotControllerEnabled"`
}

