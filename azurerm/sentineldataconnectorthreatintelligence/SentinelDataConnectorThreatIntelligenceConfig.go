// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentineldataconnectorthreatintelligence

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelDataConnectorThreatIntelligenceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/sentinel_data_connector_threat_intelligence#log_analytics_workspace_id SentinelDataConnectorThreatIntelligence#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/sentinel_data_connector_threat_intelligence#name SentinelDataConnectorThreatIntelligence#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/sentinel_data_connector_threat_intelligence#id SentinelDataConnectorThreatIntelligence#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/sentinel_data_connector_threat_intelligence#lookback_date SentinelDataConnectorThreatIntelligence#lookback_date}.
	LookbackDate *string `field:"optional" json:"lookbackDate" yaml:"lookbackDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/sentinel_data_connector_threat_intelligence#tenant_id SentinelDataConnectorThreatIntelligence#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/sentinel_data_connector_threat_intelligence#timeouts SentinelDataConnectorThreatIntelligence#timeouts}
	Timeouts *SentinelDataConnectorThreatIntelligenceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

