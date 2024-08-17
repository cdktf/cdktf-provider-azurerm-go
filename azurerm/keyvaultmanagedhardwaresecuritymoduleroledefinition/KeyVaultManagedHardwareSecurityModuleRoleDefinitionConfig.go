// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultmanagedhardwaresecuritymoduleroledefinition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyVaultManagedHardwareSecurityModuleRoleDefinitionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/key_vault_managed_hardware_security_module_role_definition#name KeyVaultManagedHardwareSecurityModuleRoleDefinition#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/key_vault_managed_hardware_security_module_role_definition#description KeyVaultManagedHardwareSecurityModuleRoleDefinition#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/key_vault_managed_hardware_security_module_role_definition#id KeyVaultManagedHardwareSecurityModuleRoleDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/key_vault_managed_hardware_security_module_role_definition#managed_hsm_id KeyVaultManagedHardwareSecurityModuleRoleDefinition#managed_hsm_id}.
	ManagedHsmId *string `field:"optional" json:"managedHsmId" yaml:"managedHsmId"`
	// permission block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/key_vault_managed_hardware_security_module_role_definition#permission KeyVaultManagedHardwareSecurityModuleRoleDefinition#permission}
	Permission interface{} `field:"optional" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/key_vault_managed_hardware_security_module_role_definition#role_name KeyVaultManagedHardwareSecurityModuleRoleDefinition#role_name}.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/key_vault_managed_hardware_security_module_role_definition#timeouts KeyVaultManagedHardwareSecurityModuleRoleDefinition#timeouts}
	Timeouts *KeyVaultManagedHardwareSecurityModuleRoleDefinitionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/key_vault_managed_hardware_security_module_role_definition#vault_base_url KeyVaultManagedHardwareSecurityModuleRoleDefinition#vault_base_url}.
	VaultBaseUrl *string `field:"optional" json:"vaultBaseUrl" yaml:"vaultBaseUrl"`
}

