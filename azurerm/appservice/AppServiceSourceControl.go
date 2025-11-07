// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appservice


type AppServiceSourceControl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/app_service#branch AppService#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/app_service#manual_integration AppService#manual_integration}.
	ManualIntegration interface{} `field:"optional" json:"manualIntegration" yaml:"manualIntegration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/app_service#repo_url AppService#repo_url}.
	RepoUrl *string `field:"optional" json:"repoUrl" yaml:"repoUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/app_service#rollback_enabled AppService#rollback_enabled}.
	RollbackEnabled interface{} `field:"optional" json:"rollbackEnabled" yaml:"rollbackEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/app_service#use_mercurial AppService#use_mercurial}.
	UseMercurial interface{} `field:"optional" json:"useMercurial" yaml:"useMercurial"`
}

