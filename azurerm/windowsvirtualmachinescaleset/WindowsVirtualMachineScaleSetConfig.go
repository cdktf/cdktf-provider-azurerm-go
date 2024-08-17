// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachinescaleset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsVirtualMachineScaleSetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#admin_password WindowsVirtualMachineScaleSet#admin_password}.
	AdminPassword *string `field:"required" json:"adminPassword" yaml:"adminPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#admin_username WindowsVirtualMachineScaleSet#admin_username}.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#instances WindowsVirtualMachineScaleSet#instances}.
	Instances *float64 `field:"required" json:"instances" yaml:"instances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#location WindowsVirtualMachineScaleSet#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#name WindowsVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#network_interface WindowsVirtualMachineScaleSet#network_interface}
	NetworkInterface interface{} `field:"required" json:"networkInterface" yaml:"networkInterface"`
	// os_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#os_disk WindowsVirtualMachineScaleSet#os_disk}
	OsDisk *WindowsVirtualMachineScaleSetOsDisk `field:"required" json:"osDisk" yaml:"osDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#resource_group_name WindowsVirtualMachineScaleSet#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#sku WindowsVirtualMachineScaleSet#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// additional_capabilities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#additional_capabilities WindowsVirtualMachineScaleSet#additional_capabilities}
	AdditionalCapabilities *WindowsVirtualMachineScaleSetAdditionalCapabilities `field:"optional" json:"additionalCapabilities" yaml:"additionalCapabilities"`
	// additional_unattend_content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#additional_unattend_content WindowsVirtualMachineScaleSet#additional_unattend_content}
	AdditionalUnattendContent interface{} `field:"optional" json:"additionalUnattendContent" yaml:"additionalUnattendContent"`
	// automatic_instance_repair block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#automatic_instance_repair WindowsVirtualMachineScaleSet#automatic_instance_repair}
	AutomaticInstanceRepair *WindowsVirtualMachineScaleSetAutomaticInstanceRepair `field:"optional" json:"automaticInstanceRepair" yaml:"automaticInstanceRepair"`
	// automatic_os_upgrade_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#automatic_os_upgrade_policy WindowsVirtualMachineScaleSet#automatic_os_upgrade_policy}
	AutomaticOsUpgradePolicy *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy `field:"optional" json:"automaticOsUpgradePolicy" yaml:"automaticOsUpgradePolicy"`
	// boot_diagnostics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#boot_diagnostics WindowsVirtualMachineScaleSet#boot_diagnostics}
	BootDiagnostics *WindowsVirtualMachineScaleSetBootDiagnostics `field:"optional" json:"bootDiagnostics" yaml:"bootDiagnostics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#capacity_reservation_group_id WindowsVirtualMachineScaleSet#capacity_reservation_group_id}.
	CapacityReservationGroupId *string `field:"optional" json:"capacityReservationGroupId" yaml:"capacityReservationGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#computer_name_prefix WindowsVirtualMachineScaleSet#computer_name_prefix}.
	ComputerNamePrefix *string `field:"optional" json:"computerNamePrefix" yaml:"computerNamePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#custom_data WindowsVirtualMachineScaleSet#custom_data}.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
	// data_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#data_disk WindowsVirtualMachineScaleSet#data_disk}
	DataDisk interface{} `field:"optional" json:"dataDisk" yaml:"dataDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#do_not_run_extensions_on_overprovisioned_machines WindowsVirtualMachineScaleSet#do_not_run_extensions_on_overprovisioned_machines}.
	DoNotRunExtensionsOnOverprovisionedMachines interface{} `field:"optional" json:"doNotRunExtensionsOnOverprovisionedMachines" yaml:"doNotRunExtensionsOnOverprovisionedMachines"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#edge_zone WindowsVirtualMachineScaleSet#edge_zone}.
	EdgeZone *string `field:"optional" json:"edgeZone" yaml:"edgeZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#enable_automatic_updates WindowsVirtualMachineScaleSet#enable_automatic_updates}.
	EnableAutomaticUpdates interface{} `field:"optional" json:"enableAutomaticUpdates" yaml:"enableAutomaticUpdates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#encryption_at_host_enabled WindowsVirtualMachineScaleSet#encryption_at_host_enabled}.
	EncryptionAtHostEnabled interface{} `field:"optional" json:"encryptionAtHostEnabled" yaml:"encryptionAtHostEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#eviction_policy WindowsVirtualMachineScaleSet#eviction_policy}.
	EvictionPolicy *string `field:"optional" json:"evictionPolicy" yaml:"evictionPolicy"`
	// extension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#extension WindowsVirtualMachineScaleSet#extension}
	Extension interface{} `field:"optional" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#extension_operations_enabled WindowsVirtualMachineScaleSet#extension_operations_enabled}.
	ExtensionOperationsEnabled interface{} `field:"optional" json:"extensionOperationsEnabled" yaml:"extensionOperationsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#extensions_time_budget WindowsVirtualMachineScaleSet#extensions_time_budget}.
	ExtensionsTimeBudget *string `field:"optional" json:"extensionsTimeBudget" yaml:"extensionsTimeBudget"`
	// gallery_application block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#gallery_application WindowsVirtualMachineScaleSet#gallery_application}
	GalleryApplication interface{} `field:"optional" json:"galleryApplication" yaml:"galleryApplication"`
	// gallery_applications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#gallery_applications WindowsVirtualMachineScaleSet#gallery_applications}
	GalleryApplications interface{} `field:"optional" json:"galleryApplications" yaml:"galleryApplications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#health_probe_id WindowsVirtualMachineScaleSet#health_probe_id}.
	HealthProbeId *string `field:"optional" json:"healthProbeId" yaml:"healthProbeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#host_group_id WindowsVirtualMachineScaleSet#host_group_id}.
	HostGroupId *string `field:"optional" json:"hostGroupId" yaml:"hostGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#id WindowsVirtualMachineScaleSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#identity WindowsVirtualMachineScaleSet#identity}
	Identity *WindowsVirtualMachineScaleSetIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#license_type WindowsVirtualMachineScaleSet#license_type}.
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#max_bid_price WindowsVirtualMachineScaleSet#max_bid_price}.
	MaxBidPrice *float64 `field:"optional" json:"maxBidPrice" yaml:"maxBidPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#overprovision WindowsVirtualMachineScaleSet#overprovision}.
	Overprovision interface{} `field:"optional" json:"overprovision" yaml:"overprovision"`
	// plan block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#plan WindowsVirtualMachineScaleSet#plan}
	Plan *WindowsVirtualMachineScaleSetPlan `field:"optional" json:"plan" yaml:"plan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#platform_fault_domain_count WindowsVirtualMachineScaleSet#platform_fault_domain_count}.
	PlatformFaultDomainCount *float64 `field:"optional" json:"platformFaultDomainCount" yaml:"platformFaultDomainCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#priority WindowsVirtualMachineScaleSet#priority}.
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#provision_vm_agent WindowsVirtualMachineScaleSet#provision_vm_agent}.
	ProvisionVmAgent interface{} `field:"optional" json:"provisionVmAgent" yaml:"provisionVmAgent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#proximity_placement_group_id WindowsVirtualMachineScaleSet#proximity_placement_group_id}.
	ProximityPlacementGroupId *string `field:"optional" json:"proximityPlacementGroupId" yaml:"proximityPlacementGroupId"`
	// rolling_upgrade_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#rolling_upgrade_policy WindowsVirtualMachineScaleSet#rolling_upgrade_policy}
	RollingUpgradePolicy *WindowsVirtualMachineScaleSetRollingUpgradePolicy `field:"optional" json:"rollingUpgradePolicy" yaml:"rollingUpgradePolicy"`
	// scale_in block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#scale_in WindowsVirtualMachineScaleSet#scale_in}
	ScaleIn *WindowsVirtualMachineScaleSetScaleIn `field:"optional" json:"scaleIn" yaml:"scaleIn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#scale_in_policy WindowsVirtualMachineScaleSet#scale_in_policy}.
	ScaleInPolicy *string `field:"optional" json:"scaleInPolicy" yaml:"scaleInPolicy"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#secret WindowsVirtualMachineScaleSet#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#secure_boot_enabled WindowsVirtualMachineScaleSet#secure_boot_enabled}.
	SecureBootEnabled interface{} `field:"optional" json:"secureBootEnabled" yaml:"secureBootEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#single_placement_group WindowsVirtualMachineScaleSet#single_placement_group}.
	SinglePlacementGroup interface{} `field:"optional" json:"singlePlacementGroup" yaml:"singlePlacementGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#source_image_id WindowsVirtualMachineScaleSet#source_image_id}.
	SourceImageId *string `field:"optional" json:"sourceImageId" yaml:"sourceImageId"`
	// source_image_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#source_image_reference WindowsVirtualMachineScaleSet#source_image_reference}
	SourceImageReference *WindowsVirtualMachineScaleSetSourceImageReference `field:"optional" json:"sourceImageReference" yaml:"sourceImageReference"`
	// spot_restore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#spot_restore WindowsVirtualMachineScaleSet#spot_restore}
	SpotRestore *WindowsVirtualMachineScaleSetSpotRestore `field:"optional" json:"spotRestore" yaml:"spotRestore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#tags WindowsVirtualMachineScaleSet#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// terminate_notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#terminate_notification WindowsVirtualMachineScaleSet#terminate_notification}
	TerminateNotification *WindowsVirtualMachineScaleSetTerminateNotification `field:"optional" json:"terminateNotification" yaml:"terminateNotification"`
	// termination_notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#termination_notification WindowsVirtualMachineScaleSet#termination_notification}
	TerminationNotification *WindowsVirtualMachineScaleSetTerminationNotification `field:"optional" json:"terminationNotification" yaml:"terminationNotification"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#timeouts WindowsVirtualMachineScaleSet#timeouts}
	Timeouts *WindowsVirtualMachineScaleSetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#timezone WindowsVirtualMachineScaleSet#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#upgrade_mode WindowsVirtualMachineScaleSet#upgrade_mode}.
	UpgradeMode *string `field:"optional" json:"upgradeMode" yaml:"upgradeMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#user_data WindowsVirtualMachineScaleSet#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#vtpm_enabled WindowsVirtualMachineScaleSet#vtpm_enabled}.
	VtpmEnabled interface{} `field:"optional" json:"vtpmEnabled" yaml:"vtpmEnabled"`
	// winrm_listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#winrm_listener WindowsVirtualMachineScaleSet#winrm_listener}
	WinrmListener interface{} `field:"optional" json:"winrmListener" yaml:"winrmListener"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#zone_balance WindowsVirtualMachineScaleSet#zone_balance}.
	ZoneBalance interface{} `field:"optional" json:"zoneBalance" yaml:"zoneBalance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/windows_virtual_machine_scale_set#zones WindowsVirtualMachineScaleSet#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

