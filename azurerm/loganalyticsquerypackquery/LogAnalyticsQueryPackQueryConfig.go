// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loganalyticsquerypackquery

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogAnalyticsQueryPackQueryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/log_analytics_query_pack_query#body LogAnalyticsQueryPackQuery#body}.
	Body *string `field:"required" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/log_analytics_query_pack_query#display_name LogAnalyticsQueryPackQuery#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/log_analytics_query_pack_query#query_pack_id LogAnalyticsQueryPackQuery#query_pack_id}.
	QueryPackId *string `field:"required" json:"queryPackId" yaml:"queryPackId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/log_analytics_query_pack_query#additional_settings_json LogAnalyticsQueryPackQuery#additional_settings_json}.
	AdditionalSettingsJson *string `field:"optional" json:"additionalSettingsJson" yaml:"additionalSettingsJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/log_analytics_query_pack_query#categories LogAnalyticsQueryPackQuery#categories}.
	Categories *[]*string `field:"optional" json:"categories" yaml:"categories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/log_analytics_query_pack_query#description LogAnalyticsQueryPackQuery#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/log_analytics_query_pack_query#id LogAnalyticsQueryPackQuery#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/log_analytics_query_pack_query#name LogAnalyticsQueryPackQuery#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/log_analytics_query_pack_query#resource_types LogAnalyticsQueryPackQuery#resource_types}.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/log_analytics_query_pack_query#solutions LogAnalyticsQueryPackQuery#solutions}.
	Solutions *[]*string `field:"optional" json:"solutions" yaml:"solutions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/log_analytics_query_pack_query#tags LogAnalyticsQueryPackQuery#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/log_analytics_query_pack_query#timeouts LogAnalyticsQueryPackQuery#timeouts}
	Timeouts *LogAnalyticsQueryPackQueryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

