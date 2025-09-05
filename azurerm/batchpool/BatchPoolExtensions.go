// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool


type BatchPoolExtensions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/batch_pool#name BatchPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/batch_pool#publisher BatchPool#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/batch_pool#type BatchPool#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/batch_pool#automatic_upgrade_enabled BatchPool#automatic_upgrade_enabled}.
	AutomaticUpgradeEnabled interface{} `field:"optional" json:"automaticUpgradeEnabled" yaml:"automaticUpgradeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/batch_pool#auto_upgrade_minor_version BatchPool#auto_upgrade_minor_version}.
	AutoUpgradeMinorVersion interface{} `field:"optional" json:"autoUpgradeMinorVersion" yaml:"autoUpgradeMinorVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/batch_pool#protected_settings BatchPool#protected_settings}.
	ProtectedSettings *string `field:"optional" json:"protectedSettings" yaml:"protectedSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/batch_pool#provision_after_extensions BatchPool#provision_after_extensions}.
	ProvisionAfterExtensions *[]*string `field:"optional" json:"provisionAfterExtensions" yaml:"provisionAfterExtensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/batch_pool#settings_json BatchPool#settings_json}.
	SettingsJson *string `field:"optional" json:"settingsJson" yaml:"settingsJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/batch_pool#type_handler_version BatchPool#type_handler_version}.
	TypeHandlerVersion *string `field:"optional" json:"typeHandlerVersion" yaml:"typeHandlerVersion"`
}

