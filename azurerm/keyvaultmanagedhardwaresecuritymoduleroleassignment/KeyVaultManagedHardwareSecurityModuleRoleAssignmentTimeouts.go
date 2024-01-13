// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultmanagedhardwaresecuritymoduleroleassignment


type KeyVaultManagedHardwareSecurityModuleRoleAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/key_vault_managed_hardware_security_module_role_assignment#create KeyVaultManagedHardwareSecurityModuleRoleAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/key_vault_managed_hardware_security_module_role_assignment#delete KeyVaultManagedHardwareSecurityModuleRoleAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/key_vault_managed_hardware_security_module_role_assignment#read KeyVaultManagedHardwareSecurityModuleRoleAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

