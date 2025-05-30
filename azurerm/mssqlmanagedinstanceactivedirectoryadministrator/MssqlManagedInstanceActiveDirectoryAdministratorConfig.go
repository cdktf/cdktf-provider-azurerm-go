// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlmanagedinstanceactivedirectoryadministrator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlManagedInstanceActiveDirectoryAdministratorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/mssql_managed_instance_active_directory_administrator#login_username MssqlManagedInstanceActiveDirectoryAdministrator#login_username}.
	LoginUsername *string `field:"required" json:"loginUsername" yaml:"loginUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/mssql_managed_instance_active_directory_administrator#managed_instance_id MssqlManagedInstanceActiveDirectoryAdministrator#managed_instance_id}.
	ManagedInstanceId *string `field:"required" json:"managedInstanceId" yaml:"managedInstanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/mssql_managed_instance_active_directory_administrator#object_id MssqlManagedInstanceActiveDirectoryAdministrator#object_id}.
	ObjectId *string `field:"required" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/mssql_managed_instance_active_directory_administrator#tenant_id MssqlManagedInstanceActiveDirectoryAdministrator#tenant_id}.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/mssql_managed_instance_active_directory_administrator#azuread_authentication_only MssqlManagedInstanceActiveDirectoryAdministrator#azuread_authentication_only}.
	AzureadAuthenticationOnly interface{} `field:"optional" json:"azureadAuthenticationOnly" yaml:"azureadAuthenticationOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/mssql_managed_instance_active_directory_administrator#id MssqlManagedInstanceActiveDirectoryAdministrator#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/mssql_managed_instance_active_directory_administrator#timeouts MssqlManagedInstanceActiveDirectoryAdministrator#timeouts}
	Timeouts *MssqlManagedInstanceActiveDirectoryAdministratorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

