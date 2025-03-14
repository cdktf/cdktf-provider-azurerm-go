// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appservicesourcecontrol

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppServiceSourceControlAConfig struct {
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
	// The ID of the Windows or Linux Web App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/app_service_source_control#app_id AppServiceSourceControlA#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The branch name to use for deployments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/app_service_source_control#branch AppServiceSourceControlA#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// github_action_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/app_service_source_control#github_action_configuration AppServiceSourceControlA#github_action_configuration}
	GithubActionConfiguration *AppServiceSourceControlGithubActionConfiguration `field:"optional" json:"githubActionConfiguration" yaml:"githubActionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/app_service_source_control#id AppServiceSourceControlA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The URL for the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/app_service_source_control#repo_url AppServiceSourceControlA#repo_url}
	RepoUrl *string `field:"optional" json:"repoUrl" yaml:"repoUrl"`
	// Should the Deployment Rollback be enabled? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/app_service_source_control#rollback_enabled AppServiceSourceControlA#rollback_enabled}
	RollbackEnabled interface{} `field:"optional" json:"rollbackEnabled" yaml:"rollbackEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/app_service_source_control#timeouts AppServiceSourceControlA#timeouts}
	Timeouts *AppServiceSourceControlTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Should the App use local Git configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/app_service_source_control#use_local_git AppServiceSourceControlA#use_local_git}
	UseLocalGit interface{} `field:"optional" json:"useLocalGit" yaml:"useLocalGit"`
	// Should code be deployed manually.
	//
	// Set to `false` to enable continuous integration, such as webhooks into online repos such as GitHub. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/app_service_source_control#use_manual_integration AppServiceSourceControlA#use_manual_integration}
	UseManualIntegration interface{} `field:"optional" json:"useManualIntegration" yaml:"useManualIntegration"`
	// The repository specified is Mercurial. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/app_service_source_control#use_mercurial AppServiceSourceControlA#use_mercurial}
	UseMercurial interface{} `field:"optional" json:"useMercurial" yaml:"useMercurial"`
}

