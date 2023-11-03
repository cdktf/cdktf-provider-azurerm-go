// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactory


type DataFactoryVstsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_factory#account_name DataFactory#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_factory#branch_name DataFactory#branch_name}.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_factory#project_name DataFactory#project_name}.
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_factory#repository_name DataFactory#repository_name}.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_factory#root_folder DataFactory#root_folder}.
	RootFolder *string `field:"required" json:"rootFolder" yaml:"rootFolder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_factory#tenant_id DataFactory#tenant_id}.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_factory#publishing_enabled DataFactory#publishing_enabled}.
	PublishingEnabled interface{} `field:"optional" json:"publishingEnabled" yaml:"publishingEnabled"`
}

