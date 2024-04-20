// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package botwebapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BotWebAppConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/bot_web_app#location BotWebApp#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/bot_web_app#microsoft_app_id BotWebApp#microsoft_app_id}.
	MicrosoftAppId *string `field:"required" json:"microsoftAppId" yaml:"microsoftAppId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/bot_web_app#name BotWebApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/bot_web_app#resource_group_name BotWebApp#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/bot_web_app#sku BotWebApp#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/bot_web_app#developer_app_insights_api_key BotWebApp#developer_app_insights_api_key}.
	DeveloperAppInsightsApiKey *string `field:"optional" json:"developerAppInsightsApiKey" yaml:"developerAppInsightsApiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/bot_web_app#developer_app_insights_application_id BotWebApp#developer_app_insights_application_id}.
	DeveloperAppInsightsApplicationId *string `field:"optional" json:"developerAppInsightsApplicationId" yaml:"developerAppInsightsApplicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/bot_web_app#developer_app_insights_key BotWebApp#developer_app_insights_key}.
	DeveloperAppInsightsKey *string `field:"optional" json:"developerAppInsightsKey" yaml:"developerAppInsightsKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/bot_web_app#display_name BotWebApp#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/bot_web_app#endpoint BotWebApp#endpoint}.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/bot_web_app#id BotWebApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/bot_web_app#luis_app_ids BotWebApp#luis_app_ids}.
	LuisAppIds *[]*string `field:"optional" json:"luisAppIds" yaml:"luisAppIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/bot_web_app#luis_key BotWebApp#luis_key}.
	LuisKey *string `field:"optional" json:"luisKey" yaml:"luisKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/bot_web_app#tags BotWebApp#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/bot_web_app#timeouts BotWebApp#timeouts}
	Timeouts *BotWebAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

