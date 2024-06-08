// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappVolumeConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#account_name NetappVolume#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#location NetappVolume#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#name NetappVolume#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#pool_name NetappVolume#pool_name}.
	PoolName *string `field:"required" json:"poolName" yaml:"poolName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#resource_group_name NetappVolume#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#service_level NetappVolume#service_level}.
	ServiceLevel *string `field:"required" json:"serviceLevel" yaml:"serviceLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#storage_quota_in_gb NetappVolume#storage_quota_in_gb}.
	StorageQuotaInGb *float64 `field:"required" json:"storageQuotaInGb" yaml:"storageQuotaInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#subnet_id NetappVolume#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#volume_path NetappVolume#volume_path}.
	VolumePath *string `field:"required" json:"volumePath" yaml:"volumePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#azure_vmware_data_store_enabled NetappVolume#azure_vmware_data_store_enabled}.
	AzureVmwareDataStoreEnabled interface{} `field:"optional" json:"azureVmwareDataStoreEnabled" yaml:"azureVmwareDataStoreEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#create_from_snapshot_resource_id NetappVolume#create_from_snapshot_resource_id}.
	CreateFromSnapshotResourceId *string `field:"optional" json:"createFromSnapshotResourceId" yaml:"createFromSnapshotResourceId"`
	// data_protection_replication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#data_protection_replication NetappVolume#data_protection_replication}
	DataProtectionReplication *NetappVolumeDataProtectionReplication `field:"optional" json:"dataProtectionReplication" yaml:"dataProtectionReplication"`
	// data_protection_snapshot_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#data_protection_snapshot_policy NetappVolume#data_protection_snapshot_policy}
	DataProtectionSnapshotPolicy *NetappVolumeDataProtectionSnapshotPolicy `field:"optional" json:"dataProtectionSnapshotPolicy" yaml:"dataProtectionSnapshotPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#encryption_key_source NetappVolume#encryption_key_source}.
	EncryptionKeySource *string `field:"optional" json:"encryptionKeySource" yaml:"encryptionKeySource"`
	// export_policy_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#export_policy_rule NetappVolume#export_policy_rule}
	ExportPolicyRule interface{} `field:"optional" json:"exportPolicyRule" yaml:"exportPolicyRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#id NetappVolume#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Enable to allow Kerberos secured volumes.
	//
	// Requires appropriate export rules as well as the parent `azurerm_netapp_account` having a defined AD connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#kerberos_enabled NetappVolume#kerberos_enabled}
	KerberosEnabled interface{} `field:"optional" json:"kerberosEnabled" yaml:"kerberosEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#key_vault_private_endpoint_id NetappVolume#key_vault_private_endpoint_id}.
	KeyVaultPrivateEndpointId *string `field:"optional" json:"keyVaultPrivateEndpointId" yaml:"keyVaultPrivateEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#network_features NetappVolume#network_features}.
	NetworkFeatures *string `field:"optional" json:"networkFeatures" yaml:"networkFeatures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#protocols NetappVolume#protocols}.
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#security_style NetappVolume#security_style}.
	SecurityStyle *string `field:"optional" json:"securityStyle" yaml:"securityStyle"`
	// Enable access based enumeration setting for SMB/Dual Protocol volume.
	//
	// When enabled, users who do not have permission to access a shared folder or file underneath it, do not see that shared resource displayed in their environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#smb_access_based_enumeration_enabled NetappVolume#smb_access_based_enumeration_enabled}
	SmbAccessBasedEnumerationEnabled interface{} `field:"optional" json:"smbAccessBasedEnumerationEnabled" yaml:"smbAccessBasedEnumerationEnabled"`
	// Continuous availability option should be used only for SQL and FSLogix workloads.
	//
	// Using it for any other SMB workloads is not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#smb_continuous_availability_enabled NetappVolume#smb_continuous_availability_enabled}
	SmbContinuousAvailabilityEnabled interface{} `field:"optional" json:"smbContinuousAvailabilityEnabled" yaml:"smbContinuousAvailabilityEnabled"`
	// Enable non browsable share setting for SMB/Dual Protocol volume. When enabled, it restricts windows clients to browse the share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#smb_non_browsable_enabled NetappVolume#smb_non_browsable_enabled}
	SmbNonBrowsableEnabled interface{} `field:"optional" json:"smbNonBrowsableEnabled" yaml:"smbNonBrowsableEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#snapshot_directory_visible NetappVolume#snapshot_directory_visible}.
	SnapshotDirectoryVisible interface{} `field:"optional" json:"snapshotDirectoryVisible" yaml:"snapshotDirectoryVisible"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#tags NetappVolume#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#throughput_in_mibps NetappVolume#throughput_in_mibps}.
	ThroughputInMibps *float64 `field:"optional" json:"throughputInMibps" yaml:"throughputInMibps"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#timeouts NetappVolume#timeouts}
	Timeouts *NetappVolumeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/netapp_volume#zone NetappVolume#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

