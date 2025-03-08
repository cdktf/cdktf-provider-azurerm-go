// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementapioperation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementApiOperationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/api_management_api_operation#api_management_name ApiManagementApiOperation#api_management_name}.
	ApiManagementName *string `field:"required" json:"apiManagementName" yaml:"apiManagementName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/api_management_api_operation#api_name ApiManagementApiOperation#api_name}.
	ApiName *string `field:"required" json:"apiName" yaml:"apiName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/api_management_api_operation#display_name ApiManagementApiOperation#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/api_management_api_operation#method ApiManagementApiOperation#method}.
	Method *string `field:"required" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/api_management_api_operation#operation_id ApiManagementApiOperation#operation_id}.
	OperationId *string `field:"required" json:"operationId" yaml:"operationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/api_management_api_operation#resource_group_name ApiManagementApiOperation#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/api_management_api_operation#url_template ApiManagementApiOperation#url_template}.
	UrlTemplate *string `field:"required" json:"urlTemplate" yaml:"urlTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/api_management_api_operation#description ApiManagementApiOperation#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/api_management_api_operation#id ApiManagementApiOperation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/api_management_api_operation#request ApiManagementApiOperation#request}
	Request *ApiManagementApiOperationRequest `field:"optional" json:"request" yaml:"request"`
	// response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/api_management_api_operation#response ApiManagementApiOperation#response}
	Response interface{} `field:"optional" json:"response" yaml:"response"`
	// template_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/api_management_api_operation#template_parameter ApiManagementApiOperation#template_parameter}
	TemplateParameter interface{} `field:"optional" json:"templateParameter" yaml:"templateParameter"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/api_management_api_operation#timeouts ApiManagementApiOperation#timeouts}
	Timeouts *ApiManagementApiOperationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

