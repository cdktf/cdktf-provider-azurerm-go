// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume


type NetappVolumeDataProtectionReplication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/netapp_volume#remote_volume_location NetappVolume#remote_volume_location}.
	RemoteVolumeLocation *string `field:"required" json:"remoteVolumeLocation" yaml:"remoteVolumeLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/netapp_volume#remote_volume_resource_id NetappVolume#remote_volume_resource_id}.
	RemoteVolumeResourceId *string `field:"required" json:"remoteVolumeResourceId" yaml:"remoteVolumeResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/netapp_volume#replication_frequency NetappVolume#replication_frequency}.
	ReplicationFrequency *string `field:"required" json:"replicationFrequency" yaml:"replicationFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/netapp_volume#endpoint_type NetappVolume#endpoint_type}.
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
}

