// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package botserviceazurebot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BotServiceAzureBotConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#location BotServiceAzureBot#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#microsoft_app_id BotServiceAzureBot#microsoft_app_id}.
	MicrosoftAppId *string `field:"required" json:"microsoftAppId" yaml:"microsoftAppId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#name BotServiceAzureBot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#resource_group_name BotServiceAzureBot#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#sku BotServiceAzureBot#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#cmk_key_vault_key_url BotServiceAzureBot#cmk_key_vault_key_url}.
	CmkKeyVaultKeyUrl *string `field:"optional" json:"cmkKeyVaultKeyUrl" yaml:"cmkKeyVaultKeyUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#developer_app_insights_api_key BotServiceAzureBot#developer_app_insights_api_key}.
	DeveloperAppInsightsApiKey *string `field:"optional" json:"developerAppInsightsApiKey" yaml:"developerAppInsightsApiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#developer_app_insights_application_id BotServiceAzureBot#developer_app_insights_application_id}.
	DeveloperAppInsightsApplicationId *string `field:"optional" json:"developerAppInsightsApplicationId" yaml:"developerAppInsightsApplicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#developer_app_insights_key BotServiceAzureBot#developer_app_insights_key}.
	DeveloperAppInsightsKey *string `field:"optional" json:"developerAppInsightsKey" yaml:"developerAppInsightsKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#display_name BotServiceAzureBot#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#endpoint BotServiceAzureBot#endpoint}.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#icon_url BotServiceAzureBot#icon_url}.
	IconUrl *string `field:"optional" json:"iconUrl" yaml:"iconUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#id BotServiceAzureBot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#local_authentication_enabled BotServiceAzureBot#local_authentication_enabled}.
	LocalAuthenticationEnabled interface{} `field:"optional" json:"localAuthenticationEnabled" yaml:"localAuthenticationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#luis_app_ids BotServiceAzureBot#luis_app_ids}.
	LuisAppIds *[]*string `field:"optional" json:"luisAppIds" yaml:"luisAppIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#luis_key BotServiceAzureBot#luis_key}.
	LuisKey *string `field:"optional" json:"luisKey" yaml:"luisKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#microsoft_app_msi_id BotServiceAzureBot#microsoft_app_msi_id}.
	MicrosoftAppMsiId *string `field:"optional" json:"microsoftAppMsiId" yaml:"microsoftAppMsiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#microsoft_app_tenant_id BotServiceAzureBot#microsoft_app_tenant_id}.
	MicrosoftAppTenantId *string `field:"optional" json:"microsoftAppTenantId" yaml:"microsoftAppTenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#microsoft_app_type BotServiceAzureBot#microsoft_app_type}.
	MicrosoftAppType *string `field:"optional" json:"microsoftAppType" yaml:"microsoftAppType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#public_network_access_enabled BotServiceAzureBot#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#streaming_endpoint_enabled BotServiceAzureBot#streaming_endpoint_enabled}.
	StreamingEndpointEnabled interface{} `field:"optional" json:"streamingEndpointEnabled" yaml:"streamingEndpointEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#tags BotServiceAzureBot#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/bot_service_azure_bot#timeouts BotServiceAzureBot#timeouts}
	Timeouts *BotServiceAzureBotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

