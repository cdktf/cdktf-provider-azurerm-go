// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachineextension

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineExtensionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/virtual_machine_extension#name VirtualMachineExtension#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/virtual_machine_extension#publisher VirtualMachineExtension#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/virtual_machine_extension#type VirtualMachineExtension#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/virtual_machine_extension#type_handler_version VirtualMachineExtension#type_handler_version}.
	TypeHandlerVersion *string `field:"required" json:"typeHandlerVersion" yaml:"typeHandlerVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/virtual_machine_extension#virtual_machine_id VirtualMachineExtension#virtual_machine_id}.
	VirtualMachineId *string `field:"required" json:"virtualMachineId" yaml:"virtualMachineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/virtual_machine_extension#automatic_upgrade_enabled VirtualMachineExtension#automatic_upgrade_enabled}.
	AutomaticUpgradeEnabled interface{} `field:"optional" json:"automaticUpgradeEnabled" yaml:"automaticUpgradeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/virtual_machine_extension#auto_upgrade_minor_version VirtualMachineExtension#auto_upgrade_minor_version}.
	AutoUpgradeMinorVersion interface{} `field:"optional" json:"autoUpgradeMinorVersion" yaml:"autoUpgradeMinorVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/virtual_machine_extension#failure_suppression_enabled VirtualMachineExtension#failure_suppression_enabled}.
	FailureSuppressionEnabled interface{} `field:"optional" json:"failureSuppressionEnabled" yaml:"failureSuppressionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/virtual_machine_extension#id VirtualMachineExtension#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/virtual_machine_extension#protected_settings VirtualMachineExtension#protected_settings}.
	ProtectedSettings *string `field:"optional" json:"protectedSettings" yaml:"protectedSettings"`
	// protected_settings_from_key_vault block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/virtual_machine_extension#protected_settings_from_key_vault VirtualMachineExtension#protected_settings_from_key_vault}
	ProtectedSettingsFromKeyVault *VirtualMachineExtensionProtectedSettingsFromKeyVault `field:"optional" json:"protectedSettingsFromKeyVault" yaml:"protectedSettingsFromKeyVault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/virtual_machine_extension#provision_after_extensions VirtualMachineExtension#provision_after_extensions}.
	ProvisionAfterExtensions *[]*string `field:"optional" json:"provisionAfterExtensions" yaml:"provisionAfterExtensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/virtual_machine_extension#settings VirtualMachineExtension#settings}.
	Settings *string `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/virtual_machine_extension#tags VirtualMachineExtension#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/virtual_machine_extension#timeouts VirtualMachineExtension#timeouts}
	Timeouts *VirtualMachineExtensionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

