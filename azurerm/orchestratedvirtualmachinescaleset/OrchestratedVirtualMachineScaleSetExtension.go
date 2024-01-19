// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetExtension struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/orchestrated_virtual_machine_scale_set#name OrchestratedVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/orchestrated_virtual_machine_scale_set#publisher OrchestratedVirtualMachineScaleSet#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/orchestrated_virtual_machine_scale_set#type OrchestratedVirtualMachineScaleSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/orchestrated_virtual_machine_scale_set#type_handler_version OrchestratedVirtualMachineScaleSet#type_handler_version}.
	TypeHandlerVersion *string `field:"required" json:"typeHandlerVersion" yaml:"typeHandlerVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/orchestrated_virtual_machine_scale_set#auto_upgrade_minor_version_enabled OrchestratedVirtualMachineScaleSet#auto_upgrade_minor_version_enabled}.
	AutoUpgradeMinorVersionEnabled interface{} `field:"optional" json:"autoUpgradeMinorVersionEnabled" yaml:"autoUpgradeMinorVersionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/orchestrated_virtual_machine_scale_set#extensions_to_provision_after_vm_creation OrchestratedVirtualMachineScaleSet#extensions_to_provision_after_vm_creation}.
	ExtensionsToProvisionAfterVmCreation *[]*string `field:"optional" json:"extensionsToProvisionAfterVmCreation" yaml:"extensionsToProvisionAfterVmCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/orchestrated_virtual_machine_scale_set#failure_suppression_enabled OrchestratedVirtualMachineScaleSet#failure_suppression_enabled}.
	FailureSuppressionEnabled interface{} `field:"optional" json:"failureSuppressionEnabled" yaml:"failureSuppressionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/orchestrated_virtual_machine_scale_set#force_extension_execution_on_change OrchestratedVirtualMachineScaleSet#force_extension_execution_on_change}.
	ForceExtensionExecutionOnChange *string `field:"optional" json:"forceExtensionExecutionOnChange" yaml:"forceExtensionExecutionOnChange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/orchestrated_virtual_machine_scale_set#protected_settings OrchestratedVirtualMachineScaleSet#protected_settings}.
	ProtectedSettings *string `field:"optional" json:"protectedSettings" yaml:"protectedSettings"`
	// protected_settings_from_key_vault block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/orchestrated_virtual_machine_scale_set#protected_settings_from_key_vault OrchestratedVirtualMachineScaleSet#protected_settings_from_key_vault}
	ProtectedSettingsFromKeyVault *OrchestratedVirtualMachineScaleSetExtensionProtectedSettingsFromKeyVault `field:"optional" json:"protectedSettingsFromKeyVault" yaml:"protectedSettingsFromKeyVault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/orchestrated_virtual_machine_scale_set#settings OrchestratedVirtualMachineScaleSet#settings}.
	Settings *string `field:"optional" json:"settings" yaml:"settings"`
}

