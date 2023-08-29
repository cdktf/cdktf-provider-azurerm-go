// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotsecuritysolution

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotSecuritySolutionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#display_name IotSecuritySolution#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#iothub_ids IotSecuritySolution#iothub_ids}.
	IothubIds *[]*string `field:"required" json:"iothubIds" yaml:"iothubIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#location IotSecuritySolution#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#name IotSecuritySolution#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#resource_group_name IotSecuritySolution#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// additional_workspace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#additional_workspace IotSecuritySolution#additional_workspace}
	AdditionalWorkspace interface{} `field:"optional" json:"additionalWorkspace" yaml:"additionalWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#disabled_data_sources IotSecuritySolution#disabled_data_sources}.
	DisabledDataSources *[]*string `field:"optional" json:"disabledDataSources" yaml:"disabledDataSources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#enabled IotSecuritySolution#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#events_to_export IotSecuritySolution#events_to_export}.
	EventsToExport *[]*string `field:"optional" json:"eventsToExport" yaml:"eventsToExport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#id IotSecuritySolution#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#log_analytics_workspace_id IotSecuritySolution#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"optional" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#log_unmasked_ips_enabled IotSecuritySolution#log_unmasked_ips_enabled}.
	LogUnmaskedIpsEnabled interface{} `field:"optional" json:"logUnmaskedIpsEnabled" yaml:"logUnmaskedIpsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#query_for_resources IotSecuritySolution#query_for_resources}.
	QueryForResources *string `field:"optional" json:"queryForResources" yaml:"queryForResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#query_subscription_ids IotSecuritySolution#query_subscription_ids}.
	QuerySubscriptionIds *[]*string `field:"optional" json:"querySubscriptionIds" yaml:"querySubscriptionIds"`
	// recommendations_enabled block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#recommendations_enabled IotSecuritySolution#recommendations_enabled}
	RecommendationsEnabled *IotSecuritySolutionRecommendationsEnabled `field:"optional" json:"recommendationsEnabled" yaml:"recommendationsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#tags IotSecuritySolution#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/iot_security_solution#timeouts IotSecuritySolution#timeouts}
	Timeouts *IotSecuritySolutionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

