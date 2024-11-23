// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultmanagedhardwaresecuritymoduleroledefinition


type KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermission struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/key_vault_managed_hardware_security_module_role_definition#actions KeyVaultManagedHardwareSecurityModuleRoleDefinition#actions}.
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/key_vault_managed_hardware_security_module_role_definition#data_actions KeyVaultManagedHardwareSecurityModuleRoleDefinition#data_actions}.
	DataActions *[]*string `field:"optional" json:"dataActions" yaml:"dataActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/key_vault_managed_hardware_security_module_role_definition#not_actions KeyVaultManagedHardwareSecurityModuleRoleDefinition#not_actions}.
	NotActions *[]*string `field:"optional" json:"notActions" yaml:"notActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/key_vault_managed_hardware_security_module_role_definition#not_data_actions KeyVaultManagedHardwareSecurityModuleRoleDefinition#not_data_actions}.
	NotDataActions *[]*string `field:"optional" json:"notDataActions" yaml:"notDataActions"`
}

