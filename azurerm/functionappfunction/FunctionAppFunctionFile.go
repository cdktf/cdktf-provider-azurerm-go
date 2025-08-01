// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappfunction


type FunctionAppFunctionFile struct {
	// The content of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/function_app_function#content FunctionAppFunction#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The filename of the file to be uploaded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/function_app_function#name FunctionAppFunction#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

