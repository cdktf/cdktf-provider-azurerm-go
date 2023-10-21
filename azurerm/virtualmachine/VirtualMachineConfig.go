// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#location VirtualMachine#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#name VirtualMachine#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#network_interface_ids VirtualMachine#network_interface_ids}.
	NetworkInterfaceIds *[]*string `field:"required" json:"networkInterfaceIds" yaml:"networkInterfaceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#resource_group_name VirtualMachine#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// storage_os_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#storage_os_disk VirtualMachine#storage_os_disk}
	StorageOsDisk *VirtualMachineStorageOsDisk `field:"required" json:"storageOsDisk" yaml:"storageOsDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#vm_size VirtualMachine#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// additional_capabilities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#additional_capabilities VirtualMachine#additional_capabilities}
	AdditionalCapabilities *VirtualMachineAdditionalCapabilities `field:"optional" json:"additionalCapabilities" yaml:"additionalCapabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#availability_set_id VirtualMachine#availability_set_id}.
	AvailabilitySetId *string `field:"optional" json:"availabilitySetId" yaml:"availabilitySetId"`
	// boot_diagnostics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#boot_diagnostics VirtualMachine#boot_diagnostics}
	BootDiagnostics *VirtualMachineBootDiagnostics `field:"optional" json:"bootDiagnostics" yaml:"bootDiagnostics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#delete_data_disks_on_termination VirtualMachine#delete_data_disks_on_termination}.
	DeleteDataDisksOnTermination interface{} `field:"optional" json:"deleteDataDisksOnTermination" yaml:"deleteDataDisksOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#delete_os_disk_on_termination VirtualMachine#delete_os_disk_on_termination}.
	DeleteOsDiskOnTermination interface{} `field:"optional" json:"deleteOsDiskOnTermination" yaml:"deleteOsDiskOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#id VirtualMachine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#identity VirtualMachine#identity}
	Identity *VirtualMachineIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#license_type VirtualMachine#license_type}.
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// os_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#os_profile VirtualMachine#os_profile}
	OsProfile *VirtualMachineOsProfile `field:"optional" json:"osProfile" yaml:"osProfile"`
	// os_profile_linux_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#os_profile_linux_config VirtualMachine#os_profile_linux_config}
	OsProfileLinuxConfig *VirtualMachineOsProfileLinuxConfig `field:"optional" json:"osProfileLinuxConfig" yaml:"osProfileLinuxConfig"`
	// os_profile_secrets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#os_profile_secrets VirtualMachine#os_profile_secrets}
	OsProfileSecrets interface{} `field:"optional" json:"osProfileSecrets" yaml:"osProfileSecrets"`
	// os_profile_windows_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#os_profile_windows_config VirtualMachine#os_profile_windows_config}
	OsProfileWindowsConfig *VirtualMachineOsProfileWindowsConfig `field:"optional" json:"osProfileWindowsConfig" yaml:"osProfileWindowsConfig"`
	// plan block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#plan VirtualMachine#plan}
	Plan *VirtualMachinePlan `field:"optional" json:"plan" yaml:"plan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#primary_network_interface_id VirtualMachine#primary_network_interface_id}.
	PrimaryNetworkInterfaceId *string `field:"optional" json:"primaryNetworkInterfaceId" yaml:"primaryNetworkInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#proximity_placement_group_id VirtualMachine#proximity_placement_group_id}.
	ProximityPlacementGroupId *string `field:"optional" json:"proximityPlacementGroupId" yaml:"proximityPlacementGroupId"`
	// storage_data_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#storage_data_disk VirtualMachine#storage_data_disk}
	StorageDataDisk interface{} `field:"optional" json:"storageDataDisk" yaml:"storageDataDisk"`
	// storage_image_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#storage_image_reference VirtualMachine#storage_image_reference}
	StorageImageReference *VirtualMachineStorageImageReference `field:"optional" json:"storageImageReference" yaml:"storageImageReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#tags VirtualMachine#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#timeouts VirtualMachine#timeouts}
	Timeouts *VirtualMachineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/virtual_machine#zones VirtualMachine#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

