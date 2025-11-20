// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcideploymentsetting


type StackHciDeploymentSettingScaleUnit struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#active_directory_organizational_unit_path StackHciDeploymentSetting#active_directory_organizational_unit_path}.
	ActiveDirectoryOrganizationalUnitPath *string `field:"required" json:"activeDirectoryOrganizationalUnitPath" yaml:"activeDirectoryOrganizationalUnitPath"`
	// cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#cluster StackHciDeploymentSetting#cluster}
	Cluster *StackHciDeploymentSettingScaleUnitCluster `field:"required" json:"cluster" yaml:"cluster"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#domain_fqdn StackHciDeploymentSetting#domain_fqdn}.
	DomainFqdn *string `field:"required" json:"domainFqdn" yaml:"domainFqdn"`
	// host_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#host_network StackHciDeploymentSetting#host_network}
	HostNetwork *StackHciDeploymentSettingScaleUnitHostNetwork `field:"required" json:"hostNetwork" yaml:"hostNetwork"`
	// infrastructure_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#infrastructure_network StackHciDeploymentSetting#infrastructure_network}
	InfrastructureNetwork interface{} `field:"required" json:"infrastructureNetwork" yaml:"infrastructureNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#name_prefix StackHciDeploymentSetting#name_prefix}.
	NamePrefix *string `field:"required" json:"namePrefix" yaml:"namePrefix"`
	// optional_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#optional_service StackHciDeploymentSetting#optional_service}
	OptionalService *StackHciDeploymentSettingScaleUnitOptionalService `field:"required" json:"optionalService" yaml:"optionalService"`
	// physical_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#physical_node StackHciDeploymentSetting#physical_node}
	PhysicalNode interface{} `field:"required" json:"physicalNode" yaml:"physicalNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#secrets_location StackHciDeploymentSetting#secrets_location}.
	SecretsLocation *string `field:"required" json:"secretsLocation" yaml:"secretsLocation"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#storage StackHciDeploymentSetting#storage}
	Storage *StackHciDeploymentSettingScaleUnitStorage `field:"required" json:"storage" yaml:"storage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#bitlocker_boot_volume_enabled StackHciDeploymentSetting#bitlocker_boot_volume_enabled}.
	BitlockerBootVolumeEnabled interface{} `field:"optional" json:"bitlockerBootVolumeEnabled" yaml:"bitlockerBootVolumeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#bitlocker_data_volume_enabled StackHciDeploymentSetting#bitlocker_data_volume_enabled}.
	BitlockerDataVolumeEnabled interface{} `field:"optional" json:"bitlockerDataVolumeEnabled" yaml:"bitlockerDataVolumeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#credential_guard_enabled StackHciDeploymentSetting#credential_guard_enabled}.
	CredentialGuardEnabled interface{} `field:"optional" json:"credentialGuardEnabled" yaml:"credentialGuardEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#drift_control_enabled StackHciDeploymentSetting#drift_control_enabled}.
	DriftControlEnabled interface{} `field:"optional" json:"driftControlEnabled" yaml:"driftControlEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#drtm_protection_enabled StackHciDeploymentSetting#drtm_protection_enabled}.
	DrtmProtectionEnabled interface{} `field:"optional" json:"drtmProtectionEnabled" yaml:"drtmProtectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#episodic_data_upload_enabled StackHciDeploymentSetting#episodic_data_upload_enabled}.
	EpisodicDataUploadEnabled interface{} `field:"optional" json:"episodicDataUploadEnabled" yaml:"episodicDataUploadEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#eu_location_enabled StackHciDeploymentSetting#eu_location_enabled}.
	EuLocationEnabled interface{} `field:"optional" json:"euLocationEnabled" yaml:"euLocationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#hvci_protection_enabled StackHciDeploymentSetting#hvci_protection_enabled}.
	HvciProtectionEnabled interface{} `field:"optional" json:"hvciProtectionEnabled" yaml:"hvciProtectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#side_channel_mitigation_enabled StackHciDeploymentSetting#side_channel_mitigation_enabled}.
	SideChannelMitigationEnabled interface{} `field:"optional" json:"sideChannelMitigationEnabled" yaml:"sideChannelMitigationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#smb_cluster_encryption_enabled StackHciDeploymentSetting#smb_cluster_encryption_enabled}.
	SmbClusterEncryptionEnabled interface{} `field:"optional" json:"smbClusterEncryptionEnabled" yaml:"smbClusterEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#smb_signing_enabled StackHciDeploymentSetting#smb_signing_enabled}.
	SmbSigningEnabled interface{} `field:"optional" json:"smbSigningEnabled" yaml:"smbSigningEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#streaming_data_client_enabled StackHciDeploymentSetting#streaming_data_client_enabled}.
	StreamingDataClientEnabled interface{} `field:"optional" json:"streamingDataClientEnabled" yaml:"streamingDataClientEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_deployment_setting#wdac_enabled StackHciDeploymentSetting#wdac_enabled}.
	WdacEnabled interface{} `field:"optional" json:"wdacEnabled" yaml:"wdacEnabled"`
}

