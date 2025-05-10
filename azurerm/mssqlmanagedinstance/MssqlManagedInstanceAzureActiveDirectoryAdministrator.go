// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlmanagedinstance


type MssqlManagedInstanceAzureActiveDirectoryAdministrator struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/mssql_managed_instance#login_username MssqlManagedInstance#login_username}.
	LoginUsername *string `field:"required" json:"loginUsername" yaml:"loginUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/mssql_managed_instance#object_id MssqlManagedInstance#object_id}.
	ObjectId *string `field:"required" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/mssql_managed_instance#principal_type MssqlManagedInstance#principal_type}.
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/mssql_managed_instance#azuread_authentication_only_enabled MssqlManagedInstance#azuread_authentication_only_enabled}.
	AzureadAuthenticationOnlyEnabled interface{} `field:"optional" json:"azureadAuthenticationOnlyEnabled" yaml:"azureadAuthenticationOnlyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/mssql_managed_instance#tenant_id MssqlManagedInstance#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

