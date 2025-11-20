// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhciextension

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StackHciExtensionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_extension#arc_setting_id StackHciExtension#arc_setting_id}.
	ArcSettingId *string `field:"required" json:"arcSettingId" yaml:"arcSettingId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_extension#name StackHciExtension#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_extension#publisher StackHciExtension#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_extension#type StackHciExtension#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_extension#automatic_upgrade_enabled StackHciExtension#automatic_upgrade_enabled}.
	AutomaticUpgradeEnabled interface{} `field:"optional" json:"automaticUpgradeEnabled" yaml:"automaticUpgradeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_extension#auto_upgrade_minor_version_enabled StackHciExtension#auto_upgrade_minor_version_enabled}.
	AutoUpgradeMinorVersionEnabled interface{} `field:"optional" json:"autoUpgradeMinorVersionEnabled" yaml:"autoUpgradeMinorVersionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_extension#id StackHciExtension#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_extension#protected_settings StackHciExtension#protected_settings}.
	ProtectedSettings *string `field:"optional" json:"protectedSettings" yaml:"protectedSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_extension#settings StackHciExtension#settings}.
	Settings *string `field:"optional" json:"settings" yaml:"settings"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_extension#timeouts StackHciExtension#timeouts}
	Timeouts *StackHciExtensionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/stack_hci_extension#type_handler_version StackHciExtension#type_handler_version}.
	TypeHandlerVersion *string `field:"optional" json:"typeHandlerVersion" yaml:"typeHandlerVersion"`
}

