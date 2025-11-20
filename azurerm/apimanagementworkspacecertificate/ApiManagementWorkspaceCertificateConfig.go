// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementworkspacecertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementWorkspaceCertificateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/api_management_workspace_certificate#api_management_workspace_id ApiManagementWorkspaceCertificate#api_management_workspace_id}.
	ApiManagementWorkspaceId *string `field:"required" json:"apiManagementWorkspaceId" yaml:"apiManagementWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/api_management_workspace_certificate#name ApiManagementWorkspaceCertificate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/api_management_workspace_certificate#certificate_data_base64 ApiManagementWorkspaceCertificate#certificate_data_base64}.
	CertificateDataBase64 *string `field:"optional" json:"certificateDataBase64" yaml:"certificateDataBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/api_management_workspace_certificate#id ApiManagementWorkspaceCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/api_management_workspace_certificate#key_vault_secret_id ApiManagementWorkspaceCertificate#key_vault_secret_id}.
	KeyVaultSecretId *string `field:"optional" json:"keyVaultSecretId" yaml:"keyVaultSecretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/api_management_workspace_certificate#password ApiManagementWorkspaceCertificate#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/api_management_workspace_certificate#timeouts ApiManagementWorkspaceCertificate#timeouts}
	Timeouts *ApiManagementWorkspaceCertificateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/api_management_workspace_certificate#user_assigned_identity_client_id ApiManagementWorkspaceCertificate#user_assigned_identity_client_id}.
	UserAssignedIdentityClientId *string `field:"optional" json:"userAssignedIdentityClientId" yaml:"userAssignedIdentityClientId"`
}

