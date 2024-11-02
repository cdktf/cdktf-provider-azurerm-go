// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappfunction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionAppFunctionConfig struct {
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
	// The config for this Function in JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/function_app_function#config_json FunctionAppFunction#config_json}
	ConfigJson *string `field:"required" json:"configJson" yaml:"configJson"`
	// The ID of the Function App in which this function should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/function_app_function#function_app_id FunctionAppFunction#function_app_id}
	FunctionAppId *string `field:"required" json:"functionAppId" yaml:"functionAppId"`
	// The name of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/function_app_function#name FunctionAppFunction#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Should this function be enabled. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/function_app_function#enabled FunctionAppFunction#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/function_app_function#file FunctionAppFunction#file}
	File interface{} `field:"optional" json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/function_app_function#id FunctionAppFunction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The language the Function is written in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/function_app_function#language FunctionAppFunction#language}
	Language *string `field:"optional" json:"language" yaml:"language"`
	// The test data for the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/function_app_function#test_data FunctionAppFunction#test_data}
	TestData *string `field:"optional" json:"testData" yaml:"testData"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/function_app_function#timeouts FunctionAppFunction#timeouts}
	Timeouts *FunctionAppFunctionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

