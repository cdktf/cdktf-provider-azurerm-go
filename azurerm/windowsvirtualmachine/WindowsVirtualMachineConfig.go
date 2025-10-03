// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsVirtualMachineConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#location WindowsVirtualMachine#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#name WindowsVirtualMachine#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#network_interface_ids WindowsVirtualMachine#network_interface_ids}.
	NetworkInterfaceIds *[]*string `field:"required" json:"networkInterfaceIds" yaml:"networkInterfaceIds"`
	// os_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#os_disk WindowsVirtualMachine#os_disk}
	OsDisk *WindowsVirtualMachineOsDisk `field:"required" json:"osDisk" yaml:"osDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#resource_group_name WindowsVirtualMachine#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#size WindowsVirtualMachine#size}.
	Size *string `field:"required" json:"size" yaml:"size"`
	// additional_capabilities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#additional_capabilities WindowsVirtualMachine#additional_capabilities}
	AdditionalCapabilities *WindowsVirtualMachineAdditionalCapabilities `field:"optional" json:"additionalCapabilities" yaml:"additionalCapabilities"`
	// additional_unattend_content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#additional_unattend_content WindowsVirtualMachine#additional_unattend_content}
	AdditionalUnattendContent interface{} `field:"optional" json:"additionalUnattendContent" yaml:"additionalUnattendContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#admin_password WindowsVirtualMachine#admin_password}.
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#admin_username WindowsVirtualMachine#admin_username}.
	AdminUsername *string `field:"optional" json:"adminUsername" yaml:"adminUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#allow_extension_operations WindowsVirtualMachine#allow_extension_operations}.
	AllowExtensionOperations interface{} `field:"optional" json:"allowExtensionOperations" yaml:"allowExtensionOperations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#automatic_updates_enabled WindowsVirtualMachine#automatic_updates_enabled}.
	AutomaticUpdatesEnabled interface{} `field:"optional" json:"automaticUpdatesEnabled" yaml:"automaticUpdatesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#availability_set_id WindowsVirtualMachine#availability_set_id}.
	AvailabilitySetId *string `field:"optional" json:"availabilitySetId" yaml:"availabilitySetId"`
	// boot_diagnostics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#boot_diagnostics WindowsVirtualMachine#boot_diagnostics}
	BootDiagnostics *WindowsVirtualMachineBootDiagnostics `field:"optional" json:"bootDiagnostics" yaml:"bootDiagnostics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#bypass_platform_safety_checks_on_user_schedule_enabled WindowsVirtualMachine#bypass_platform_safety_checks_on_user_schedule_enabled}.
	BypassPlatformSafetyChecksOnUserScheduleEnabled interface{} `field:"optional" json:"bypassPlatformSafetyChecksOnUserScheduleEnabled" yaml:"bypassPlatformSafetyChecksOnUserScheduleEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#capacity_reservation_group_id WindowsVirtualMachine#capacity_reservation_group_id}.
	CapacityReservationGroupId *string `field:"optional" json:"capacityReservationGroupId" yaml:"capacityReservationGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#computer_name WindowsVirtualMachine#computer_name}.
	ComputerName *string `field:"optional" json:"computerName" yaml:"computerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#custom_data WindowsVirtualMachine#custom_data}.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#dedicated_host_group_id WindowsVirtualMachine#dedicated_host_group_id}.
	DedicatedHostGroupId *string `field:"optional" json:"dedicatedHostGroupId" yaml:"dedicatedHostGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#dedicated_host_id WindowsVirtualMachine#dedicated_host_id}.
	DedicatedHostId *string `field:"optional" json:"dedicatedHostId" yaml:"dedicatedHostId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#disk_controller_type WindowsVirtualMachine#disk_controller_type}.
	DiskControllerType *string `field:"optional" json:"diskControllerType" yaml:"diskControllerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#edge_zone WindowsVirtualMachine#edge_zone}.
	EdgeZone *string `field:"optional" json:"edgeZone" yaml:"edgeZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#enable_automatic_updates WindowsVirtualMachine#enable_automatic_updates}.
	EnableAutomaticUpdates interface{} `field:"optional" json:"enableAutomaticUpdates" yaml:"enableAutomaticUpdates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#encryption_at_host_enabled WindowsVirtualMachine#encryption_at_host_enabled}.
	EncryptionAtHostEnabled interface{} `field:"optional" json:"encryptionAtHostEnabled" yaml:"encryptionAtHostEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#eviction_policy WindowsVirtualMachine#eviction_policy}.
	EvictionPolicy *string `field:"optional" json:"evictionPolicy" yaml:"evictionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#extensions_time_budget WindowsVirtualMachine#extensions_time_budget}.
	ExtensionsTimeBudget *string `field:"optional" json:"extensionsTimeBudget" yaml:"extensionsTimeBudget"`
	// gallery_application block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#gallery_application WindowsVirtualMachine#gallery_application}
	GalleryApplication interface{} `field:"optional" json:"galleryApplication" yaml:"galleryApplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#hotpatching_enabled WindowsVirtualMachine#hotpatching_enabled}.
	HotpatchingEnabled interface{} `field:"optional" json:"hotpatchingEnabled" yaml:"hotpatchingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#id WindowsVirtualMachine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#identity WindowsVirtualMachine#identity}
	Identity *WindowsVirtualMachineIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#license_type WindowsVirtualMachine#license_type}.
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#max_bid_price WindowsVirtualMachine#max_bid_price}.
	MaxBidPrice *float64 `field:"optional" json:"maxBidPrice" yaml:"maxBidPrice"`
	// os_image_notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#os_image_notification WindowsVirtualMachine#os_image_notification}
	OsImageNotification *WindowsVirtualMachineOsImageNotification `field:"optional" json:"osImageNotification" yaml:"osImageNotification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#os_managed_disk_id WindowsVirtualMachine#os_managed_disk_id}.
	OsManagedDiskId *string `field:"optional" json:"osManagedDiskId" yaml:"osManagedDiskId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#patch_assessment_mode WindowsVirtualMachine#patch_assessment_mode}.
	PatchAssessmentMode *string `field:"optional" json:"patchAssessmentMode" yaml:"patchAssessmentMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#patch_mode WindowsVirtualMachine#patch_mode}.
	PatchMode *string `field:"optional" json:"patchMode" yaml:"patchMode"`
	// plan block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#plan WindowsVirtualMachine#plan}
	Plan *WindowsVirtualMachinePlan `field:"optional" json:"plan" yaml:"plan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#platform_fault_domain WindowsVirtualMachine#platform_fault_domain}.
	PlatformFaultDomain *float64 `field:"optional" json:"platformFaultDomain" yaml:"platformFaultDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#priority WindowsVirtualMachine#priority}.
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#provision_vm_agent WindowsVirtualMachine#provision_vm_agent}.
	ProvisionVmAgent interface{} `field:"optional" json:"provisionVmAgent" yaml:"provisionVmAgent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#proximity_placement_group_id WindowsVirtualMachine#proximity_placement_group_id}.
	ProximityPlacementGroupId *string `field:"optional" json:"proximityPlacementGroupId" yaml:"proximityPlacementGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#reboot_setting WindowsVirtualMachine#reboot_setting}.
	RebootSetting *string `field:"optional" json:"rebootSetting" yaml:"rebootSetting"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#secret WindowsVirtualMachine#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#secure_boot_enabled WindowsVirtualMachine#secure_boot_enabled}.
	SecureBootEnabled interface{} `field:"optional" json:"secureBootEnabled" yaml:"secureBootEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#source_image_id WindowsVirtualMachine#source_image_id}.
	SourceImageId *string `field:"optional" json:"sourceImageId" yaml:"sourceImageId"`
	// source_image_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#source_image_reference WindowsVirtualMachine#source_image_reference}
	SourceImageReference *WindowsVirtualMachineSourceImageReference `field:"optional" json:"sourceImageReference" yaml:"sourceImageReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#tags WindowsVirtualMachine#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// termination_notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#termination_notification WindowsVirtualMachine#termination_notification}
	TerminationNotification *WindowsVirtualMachineTerminationNotification `field:"optional" json:"terminationNotification" yaml:"terminationNotification"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#timeouts WindowsVirtualMachine#timeouts}
	Timeouts *WindowsVirtualMachineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#timezone WindowsVirtualMachine#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#user_data WindowsVirtualMachine#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#virtual_machine_scale_set_id WindowsVirtualMachine#virtual_machine_scale_set_id}.
	VirtualMachineScaleSetId *string `field:"optional" json:"virtualMachineScaleSetId" yaml:"virtualMachineScaleSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#vm_agent_platform_updates_enabled WindowsVirtualMachine#vm_agent_platform_updates_enabled}.
	VmAgentPlatformUpdatesEnabled interface{} `field:"optional" json:"vmAgentPlatformUpdatesEnabled" yaml:"vmAgentPlatformUpdatesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#vtpm_enabled WindowsVirtualMachine#vtpm_enabled}.
	VtpmEnabled interface{} `field:"optional" json:"vtpmEnabled" yaml:"vtpmEnabled"`
	// winrm_listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#winrm_listener WindowsVirtualMachine#winrm_listener}
	WinrmListener interface{} `field:"optional" json:"winrmListener" yaml:"winrmListener"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/windows_virtual_machine#zone WindowsVirtualMachine#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

