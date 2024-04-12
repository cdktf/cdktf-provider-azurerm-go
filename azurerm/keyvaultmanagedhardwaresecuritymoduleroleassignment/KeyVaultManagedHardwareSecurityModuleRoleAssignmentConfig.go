// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultmanagedhardwaresecuritymoduleroleassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyVaultManagedHardwareSecurityModuleRoleAssignmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/key_vault_managed_hardware_security_module_role_assignment#name KeyVaultManagedHardwareSecurityModuleRoleAssignment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/key_vault_managed_hardware_security_module_role_assignment#principal_id KeyVaultManagedHardwareSecurityModuleRoleAssignment#principal_id}.
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/key_vault_managed_hardware_security_module_role_assignment#role_definition_id KeyVaultManagedHardwareSecurityModuleRoleAssignment#role_definition_id}.
	RoleDefinitionId *string `field:"required" json:"roleDefinitionId" yaml:"roleDefinitionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/key_vault_managed_hardware_security_module_role_assignment#scope KeyVaultManagedHardwareSecurityModuleRoleAssignment#scope}.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/key_vault_managed_hardware_security_module_role_assignment#vault_base_url KeyVaultManagedHardwareSecurityModuleRoleAssignment#vault_base_url}.
	VaultBaseUrl *string `field:"required" json:"vaultBaseUrl" yaml:"vaultBaseUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/key_vault_managed_hardware_security_module_role_assignment#id KeyVaultManagedHardwareSecurityModuleRoleAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/key_vault_managed_hardware_security_module_role_assignment#timeouts KeyVaultManagedHardwareSecurityModuleRoleAssignment#timeouts}
	Timeouts *KeyVaultManagedHardwareSecurityModuleRoleAssignmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

