// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddisk

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDiskConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#create_option ManagedDisk#create_option}.
	CreateOption *string `field:"required" json:"createOption" yaml:"createOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#location ManagedDisk#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#name ManagedDisk#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#resource_group_name ManagedDisk#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#storage_account_type ManagedDisk#storage_account_type}.
	StorageAccountType *string `field:"required" json:"storageAccountType" yaml:"storageAccountType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#disk_access_id ManagedDisk#disk_access_id}.
	DiskAccessId *string `field:"optional" json:"diskAccessId" yaml:"diskAccessId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#disk_encryption_set_id ManagedDisk#disk_encryption_set_id}.
	DiskEncryptionSetId *string `field:"optional" json:"diskEncryptionSetId" yaml:"diskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#disk_iops_read_only ManagedDisk#disk_iops_read_only}.
	DiskIopsReadOnly *float64 `field:"optional" json:"diskIopsReadOnly" yaml:"diskIopsReadOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#disk_iops_read_write ManagedDisk#disk_iops_read_write}.
	DiskIopsReadWrite *float64 `field:"optional" json:"diskIopsReadWrite" yaml:"diskIopsReadWrite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#disk_mbps_read_only ManagedDisk#disk_mbps_read_only}.
	DiskMbpsReadOnly *float64 `field:"optional" json:"diskMbpsReadOnly" yaml:"diskMbpsReadOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#disk_mbps_read_write ManagedDisk#disk_mbps_read_write}.
	DiskMbpsReadWrite *float64 `field:"optional" json:"diskMbpsReadWrite" yaml:"diskMbpsReadWrite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#disk_size_gb ManagedDisk#disk_size_gb}.
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#edge_zone ManagedDisk#edge_zone}.
	EdgeZone *string `field:"optional" json:"edgeZone" yaml:"edgeZone"`
	// encryption_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#encryption_settings ManagedDisk#encryption_settings}
	EncryptionSettings *ManagedDiskEncryptionSettings `field:"optional" json:"encryptionSettings" yaml:"encryptionSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#gallery_image_reference_id ManagedDisk#gallery_image_reference_id}.
	GalleryImageReferenceId *string `field:"optional" json:"galleryImageReferenceId" yaml:"galleryImageReferenceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#hyper_v_generation ManagedDisk#hyper_v_generation}.
	HyperVGeneration *string `field:"optional" json:"hyperVGeneration" yaml:"hyperVGeneration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#id ManagedDisk#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#image_reference_id ManagedDisk#image_reference_id}.
	ImageReferenceId *string `field:"optional" json:"imageReferenceId" yaml:"imageReferenceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#logical_sector_size ManagedDisk#logical_sector_size}.
	LogicalSectorSize *float64 `field:"optional" json:"logicalSectorSize" yaml:"logicalSectorSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#max_shares ManagedDisk#max_shares}.
	MaxShares *float64 `field:"optional" json:"maxShares" yaml:"maxShares"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#network_access_policy ManagedDisk#network_access_policy}.
	NetworkAccessPolicy *string `field:"optional" json:"networkAccessPolicy" yaml:"networkAccessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#on_demand_bursting_enabled ManagedDisk#on_demand_bursting_enabled}.
	OnDemandBurstingEnabled interface{} `field:"optional" json:"onDemandBurstingEnabled" yaml:"onDemandBurstingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#optimized_frequent_attach_enabled ManagedDisk#optimized_frequent_attach_enabled}.
	OptimizedFrequentAttachEnabled interface{} `field:"optional" json:"optimizedFrequentAttachEnabled" yaml:"optimizedFrequentAttachEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#os_type ManagedDisk#os_type}.
	OsType *string `field:"optional" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#performance_plus_enabled ManagedDisk#performance_plus_enabled}.
	PerformancePlusEnabled interface{} `field:"optional" json:"performancePlusEnabled" yaml:"performancePlusEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#public_network_access_enabled ManagedDisk#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#secure_vm_disk_encryption_set_id ManagedDisk#secure_vm_disk_encryption_set_id}.
	SecureVmDiskEncryptionSetId *string `field:"optional" json:"secureVmDiskEncryptionSetId" yaml:"secureVmDiskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#security_type ManagedDisk#security_type}.
	SecurityType *string `field:"optional" json:"securityType" yaml:"securityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#source_resource_id ManagedDisk#source_resource_id}.
	SourceResourceId *string `field:"optional" json:"sourceResourceId" yaml:"sourceResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#source_uri ManagedDisk#source_uri}.
	SourceUri *string `field:"optional" json:"sourceUri" yaml:"sourceUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#storage_account_id ManagedDisk#storage_account_id}.
	StorageAccountId *string `field:"optional" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#tags ManagedDisk#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#tier ManagedDisk#tier}.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#timeouts ManagedDisk#timeouts}
	Timeouts *ManagedDiskTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#trusted_launch_enabled ManagedDisk#trusted_launch_enabled}.
	TrustedLaunchEnabled interface{} `field:"optional" json:"trustedLaunchEnabled" yaml:"trustedLaunchEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#upload_size_bytes ManagedDisk#upload_size_bytes}.
	UploadSizeBytes *float64 `field:"optional" json:"uploadSizeBytes" yaml:"uploadSizeBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/managed_disk#zone ManagedDisk#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

