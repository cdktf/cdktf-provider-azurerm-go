// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedapplicationdefinition


type ManagedApplicationDefinitionAuthorization struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/managed_application_definition#role_definition_id ManagedApplicationDefinition#role_definition_id}.
	RoleDefinitionId *string `field:"required" json:"roleDefinitionId" yaml:"roleDefinitionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/managed_application_definition#service_principal_id ManagedApplicationDefinition#service_principal_id}.
	ServicePrincipalId *string `field:"required" json:"servicePrincipalId" yaml:"servicePrincipalId"`
}

