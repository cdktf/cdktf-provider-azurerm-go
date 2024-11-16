// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetExtension struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_virtual_machine_scale_set#name WindowsVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_virtual_machine_scale_set#publisher WindowsVirtualMachineScaleSet#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_virtual_machine_scale_set#type WindowsVirtualMachineScaleSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_virtual_machine_scale_set#type_handler_version WindowsVirtualMachineScaleSet#type_handler_version}.
	TypeHandlerVersion *string `field:"required" json:"typeHandlerVersion" yaml:"typeHandlerVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_virtual_machine_scale_set#automatic_upgrade_enabled WindowsVirtualMachineScaleSet#automatic_upgrade_enabled}.
	AutomaticUpgradeEnabled interface{} `field:"optional" json:"automaticUpgradeEnabled" yaml:"automaticUpgradeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_virtual_machine_scale_set#auto_upgrade_minor_version WindowsVirtualMachineScaleSet#auto_upgrade_minor_version}.
	AutoUpgradeMinorVersion interface{} `field:"optional" json:"autoUpgradeMinorVersion" yaml:"autoUpgradeMinorVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_virtual_machine_scale_set#force_update_tag WindowsVirtualMachineScaleSet#force_update_tag}.
	ForceUpdateTag *string `field:"optional" json:"forceUpdateTag" yaml:"forceUpdateTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_virtual_machine_scale_set#protected_settings WindowsVirtualMachineScaleSet#protected_settings}.
	ProtectedSettings *string `field:"optional" json:"protectedSettings" yaml:"protectedSettings"`
	// protected_settings_from_key_vault block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_virtual_machine_scale_set#protected_settings_from_key_vault WindowsVirtualMachineScaleSet#protected_settings_from_key_vault}
	ProtectedSettingsFromKeyVault *WindowsVirtualMachineScaleSetExtensionProtectedSettingsFromKeyVault `field:"optional" json:"protectedSettingsFromKeyVault" yaml:"protectedSettingsFromKeyVault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_virtual_machine_scale_set#provision_after_extensions WindowsVirtualMachineScaleSet#provision_after_extensions}.
	ProvisionAfterExtensions *[]*string `field:"optional" json:"provisionAfterExtensions" yaml:"provisionAfterExtensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_virtual_machine_scale_set#settings WindowsVirtualMachineScaleSet#settings}.
	Settings *string `field:"optional" json:"settings" yaml:"settings"`
}

