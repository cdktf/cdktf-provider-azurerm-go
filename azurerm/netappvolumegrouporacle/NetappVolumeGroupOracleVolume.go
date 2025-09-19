// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolumegrouporacle


type NetappVolumeGroupOracleVolume struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#capacity_pool_id NetappVolumeGroupOracle#capacity_pool_id}.
	CapacityPoolId *string `field:"required" json:"capacityPoolId" yaml:"capacityPoolId"`
	// export_policy_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#export_policy_rule NetappVolumeGroupOracle#export_policy_rule}
	ExportPolicyRule interface{} `field:"required" json:"exportPolicyRule" yaml:"exportPolicyRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#name NetappVolumeGroupOracle#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#protocols NetappVolumeGroupOracle#protocols}.
	Protocols *[]*string `field:"required" json:"protocols" yaml:"protocols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#security_style NetappVolumeGroupOracle#security_style}.
	SecurityStyle *string `field:"required" json:"securityStyle" yaml:"securityStyle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#service_level NetappVolumeGroupOracle#service_level}.
	ServiceLevel *string `field:"required" json:"serviceLevel" yaml:"serviceLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#snapshot_directory_visible NetappVolumeGroupOracle#snapshot_directory_visible}.
	SnapshotDirectoryVisible interface{} `field:"required" json:"snapshotDirectoryVisible" yaml:"snapshotDirectoryVisible"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#storage_quota_in_gb NetappVolumeGroupOracle#storage_quota_in_gb}.
	StorageQuotaInGb *float64 `field:"required" json:"storageQuotaInGb" yaml:"storageQuotaInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#subnet_id NetappVolumeGroupOracle#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#throughput_in_mibps NetappVolumeGroupOracle#throughput_in_mibps}.
	ThroughputInMibps *float64 `field:"required" json:"throughputInMibps" yaml:"throughputInMibps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#volume_path NetappVolumeGroupOracle#volume_path}.
	VolumePath *string `field:"required" json:"volumePath" yaml:"volumePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#volume_spec_name NetappVolumeGroupOracle#volume_spec_name}.
	VolumeSpecName *string `field:"required" json:"volumeSpecName" yaml:"volumeSpecName"`
	// data_protection_replication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#data_protection_replication NetappVolumeGroupOracle#data_protection_replication}
	DataProtectionReplication *NetappVolumeGroupOracleVolumeDataProtectionReplication `field:"optional" json:"dataProtectionReplication" yaml:"dataProtectionReplication"`
	// data_protection_snapshot_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#data_protection_snapshot_policy NetappVolumeGroupOracle#data_protection_snapshot_policy}
	DataProtectionSnapshotPolicy *NetappVolumeGroupOracleVolumeDataProtectionSnapshotPolicy `field:"optional" json:"dataProtectionSnapshotPolicy" yaml:"dataProtectionSnapshotPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#encryption_key_source NetappVolumeGroupOracle#encryption_key_source}.
	EncryptionKeySource *string `field:"optional" json:"encryptionKeySource" yaml:"encryptionKeySource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#key_vault_private_endpoint_id NetappVolumeGroupOracle#key_vault_private_endpoint_id}.
	KeyVaultPrivateEndpointId *string `field:"optional" json:"keyVaultPrivateEndpointId" yaml:"keyVaultPrivateEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#network_features NetappVolumeGroupOracle#network_features}.
	NetworkFeatures *string `field:"optional" json:"networkFeatures" yaml:"networkFeatures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#proximity_placement_group_id NetappVolumeGroupOracle#proximity_placement_group_id}.
	ProximityPlacementGroupId *string `field:"optional" json:"proximityPlacementGroupId" yaml:"proximityPlacementGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#tags NetappVolumeGroupOracle#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/netapp_volume_group_oracle#zone NetappVolumeGroupOracle#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

