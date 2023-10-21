// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionapp


type FunctionAppSourceControl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/function_app#branch FunctionApp#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/function_app#manual_integration FunctionApp#manual_integration}.
	ManualIntegration interface{} `field:"optional" json:"manualIntegration" yaml:"manualIntegration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/function_app#repo_url FunctionApp#repo_url}.
	RepoUrl *string `field:"optional" json:"repoUrl" yaml:"repoUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/function_app#rollback_enabled FunctionApp#rollback_enabled}.
	RollbackEnabled interface{} `field:"optional" json:"rollbackEnabled" yaml:"rollbackEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/function_app#use_mercurial FunctionApp#use_mercurial}.
	UseMercurial interface{} `field:"optional" json:"useMercurial" yaml:"useMercurial"`
}

