// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package botchannelmsteams

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BotChannelMsTeamsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/bot_channel_ms_teams#bot_name BotChannelMsTeams#bot_name}.
	BotName *string `field:"required" json:"botName" yaml:"botName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/bot_channel_ms_teams#location BotChannelMsTeams#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/bot_channel_ms_teams#resource_group_name BotChannelMsTeams#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/bot_channel_ms_teams#calling_web_hook BotChannelMsTeams#calling_web_hook}.
	CallingWebHook *string `field:"optional" json:"callingWebHook" yaml:"callingWebHook"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/bot_channel_ms_teams#deployment_environment BotChannelMsTeams#deployment_environment}.
	DeploymentEnvironment *string `field:"optional" json:"deploymentEnvironment" yaml:"deploymentEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/bot_channel_ms_teams#enable_calling BotChannelMsTeams#enable_calling}.
	EnableCalling interface{} `field:"optional" json:"enableCalling" yaml:"enableCalling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/bot_channel_ms_teams#id BotChannelMsTeams#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/bot_channel_ms_teams#timeouts BotChannelMsTeams#timeouts}
	Timeouts *BotChannelMsTeamsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

