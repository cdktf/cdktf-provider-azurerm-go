// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapseworkspace

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SynapseWorkspaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#location SynapseWorkspace#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#name SynapseWorkspace#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#resource_group_name SynapseWorkspace#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#storage_data_lake_gen2_filesystem_id SynapseWorkspace#storage_data_lake_gen2_filesystem_id}.
	StorageDataLakeGen2FilesystemId *string `field:"required" json:"storageDataLakeGen2FilesystemId" yaml:"storageDataLakeGen2FilesystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#azuread_authentication_only SynapseWorkspace#azuread_authentication_only}.
	AzureadAuthenticationOnly interface{} `field:"optional" json:"azureadAuthenticationOnly" yaml:"azureadAuthenticationOnly"`
	// azure_devops_repo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#azure_devops_repo SynapseWorkspace#azure_devops_repo}
	AzureDevopsRepo *SynapseWorkspaceAzureDevopsRepo `field:"optional" json:"azureDevopsRepo" yaml:"azureDevopsRepo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#compute_subnet_id SynapseWorkspace#compute_subnet_id}.
	ComputeSubnetId *string `field:"optional" json:"computeSubnetId" yaml:"computeSubnetId"`
	// customer_managed_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#customer_managed_key SynapseWorkspace#customer_managed_key}
	CustomerManagedKey *SynapseWorkspaceCustomerManagedKey `field:"optional" json:"customerManagedKey" yaml:"customerManagedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#data_exfiltration_protection_enabled SynapseWorkspace#data_exfiltration_protection_enabled}.
	DataExfiltrationProtectionEnabled interface{} `field:"optional" json:"dataExfiltrationProtectionEnabled" yaml:"dataExfiltrationProtectionEnabled"`
	// github_repo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#github_repo SynapseWorkspace#github_repo}
	GithubRepo *SynapseWorkspaceGithubRepo `field:"optional" json:"githubRepo" yaml:"githubRepo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#id SynapseWorkspace#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#identity SynapseWorkspace#identity}
	Identity *SynapseWorkspaceIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#linking_allowed_for_aad_tenant_ids SynapseWorkspace#linking_allowed_for_aad_tenant_ids}.
	LinkingAllowedForAadTenantIds *[]*string `field:"optional" json:"linkingAllowedForAadTenantIds" yaml:"linkingAllowedForAadTenantIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#managed_resource_group_name SynapseWorkspace#managed_resource_group_name}.
	ManagedResourceGroupName *string `field:"optional" json:"managedResourceGroupName" yaml:"managedResourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#managed_virtual_network_enabled SynapseWorkspace#managed_virtual_network_enabled}.
	ManagedVirtualNetworkEnabled interface{} `field:"optional" json:"managedVirtualNetworkEnabled" yaml:"managedVirtualNetworkEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#public_network_access_enabled SynapseWorkspace#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#purview_id SynapseWorkspace#purview_id}.
	PurviewId *string `field:"optional" json:"purviewId" yaml:"purviewId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#sql_administrator_login SynapseWorkspace#sql_administrator_login}.
	SqlAdministratorLogin *string `field:"optional" json:"sqlAdministratorLogin" yaml:"sqlAdministratorLogin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#sql_administrator_login_password SynapseWorkspace#sql_administrator_login_password}.
	SqlAdministratorLoginPassword *string `field:"optional" json:"sqlAdministratorLoginPassword" yaml:"sqlAdministratorLoginPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#sql_identity_control_enabled SynapseWorkspace#sql_identity_control_enabled}.
	SqlIdentityControlEnabled interface{} `field:"optional" json:"sqlIdentityControlEnabled" yaml:"sqlIdentityControlEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#tags SynapseWorkspace#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/synapse_workspace#timeouts SynapseWorkspace#timeouts}
	Timeouts *SynapseWorkspaceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

