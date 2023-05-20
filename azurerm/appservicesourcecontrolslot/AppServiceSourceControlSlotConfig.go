package appservicesourcecontrolslot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppServiceSourceControlSlotConfig struct {
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
	// The ID of the Linux or Windows Web App Slot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/app_service_source_control_slot#slot_id AppServiceSourceControlSlot#slot_id}
	SlotId *string `field:"required" json:"slotId" yaml:"slotId"`
	// The URL for the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/app_service_source_control_slot#branch AppServiceSourceControlSlot#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// github_action_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/app_service_source_control_slot#github_action_configuration AppServiceSourceControlSlot#github_action_configuration}
	GithubActionConfiguration *AppServiceSourceControlSlotGithubActionConfiguration `field:"optional" json:"githubActionConfiguration" yaml:"githubActionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/app_service_source_control_slot#id AppServiceSourceControlSlot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The branch name to use for deployments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/app_service_source_control_slot#repo_url AppServiceSourceControlSlot#repo_url}
	RepoUrl *string `field:"optional" json:"repoUrl" yaml:"repoUrl"`
	// Should the Deployment Rollback be enabled? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/app_service_source_control_slot#rollback_enabled AppServiceSourceControlSlot#rollback_enabled}
	RollbackEnabled interface{} `field:"optional" json:"rollbackEnabled" yaml:"rollbackEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/app_service_source_control_slot#timeouts AppServiceSourceControlSlot#timeouts}
	Timeouts *AppServiceSourceControlSlotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Should the Slot use local Git configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/app_service_source_control_slot#use_local_git AppServiceSourceControlSlot#use_local_git}
	UseLocalGit interface{} `field:"optional" json:"useLocalGit" yaml:"useLocalGit"`
	// Should code be deployed manually.
	//
	// Set to `true` to disable continuous integration, such as webhooks into online repos such as GitHub. Defaults to `false`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/app_service_source_control_slot#use_manual_integration AppServiceSourceControlSlot#use_manual_integration}
	UseManualIntegration interface{} `field:"optional" json:"useManualIntegration" yaml:"useManualIntegration"`
	// The repository specified is Mercurial. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/app_service_source_control_slot#use_mercurial AppServiceSourceControlSlot#use_mercurial}
	UseMercurial interface{} `field:"optional" json:"useMercurial" yaml:"useMercurial"`
}

