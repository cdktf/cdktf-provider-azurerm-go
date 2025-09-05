// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolumegrouporacle


type NetappVolumeGroupOracleVolumeDataProtectionReplication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/netapp_volume_group_oracle#remote_volume_location NetappVolumeGroupOracle#remote_volume_location}.
	RemoteVolumeLocation *string `field:"required" json:"remoteVolumeLocation" yaml:"remoteVolumeLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/netapp_volume_group_oracle#remote_volume_resource_id NetappVolumeGroupOracle#remote_volume_resource_id}.
	RemoteVolumeResourceId *string `field:"required" json:"remoteVolumeResourceId" yaml:"remoteVolumeResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/netapp_volume_group_oracle#replication_frequency NetappVolumeGroupOracle#replication_frequency}.
	ReplicationFrequency *string `field:"required" json:"replicationFrequency" yaml:"replicationFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/netapp_volume_group_oracle#endpoint_type NetappVolumeGroupOracle#endpoint_type}.
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
}

