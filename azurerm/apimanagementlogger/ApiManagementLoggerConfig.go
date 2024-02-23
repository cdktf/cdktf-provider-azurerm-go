// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementlogger

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementLoggerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/api_management_logger#api_management_name ApiManagementLogger#api_management_name}.
	ApiManagementName *string `field:"required" json:"apiManagementName" yaml:"apiManagementName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/api_management_logger#name ApiManagementLogger#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/api_management_logger#resource_group_name ApiManagementLogger#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// application_insights block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/api_management_logger#application_insights ApiManagementLogger#application_insights}
	ApplicationInsights *ApiManagementLoggerApplicationInsights `field:"optional" json:"applicationInsights" yaml:"applicationInsights"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/api_management_logger#buffered ApiManagementLogger#buffered}.
	Buffered interface{} `field:"optional" json:"buffered" yaml:"buffered"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/api_management_logger#description ApiManagementLogger#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// eventhub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/api_management_logger#eventhub ApiManagementLogger#eventhub}
	Eventhub *ApiManagementLoggerEventhub `field:"optional" json:"eventhub" yaml:"eventhub"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/api_management_logger#id ApiManagementLogger#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/api_management_logger#resource_id ApiManagementLogger#resource_id}.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/api_management_logger#timeouts ApiManagementLogger#timeouts}
	Timeouts *ApiManagementLoggerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

