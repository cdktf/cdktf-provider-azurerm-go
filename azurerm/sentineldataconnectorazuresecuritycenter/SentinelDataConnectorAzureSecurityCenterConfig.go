// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentineldataconnectorazuresecuritycenter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelDataConnectorAzureSecurityCenterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/sentinel_data_connector_azure_security_center#log_analytics_workspace_id SentinelDataConnectorAzureSecurityCenter#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/sentinel_data_connector_azure_security_center#name SentinelDataConnectorAzureSecurityCenter#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/sentinel_data_connector_azure_security_center#id SentinelDataConnectorAzureSecurityCenter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/sentinel_data_connector_azure_security_center#subscription_id SentinelDataConnectorAzureSecurityCenter#subscription_id}.
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/sentinel_data_connector_azure_security_center#timeouts SentinelDataConnectorAzureSecurityCenter#timeouts}
	Timeouts *SentinelDataConnectorAzureSecurityCenterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

