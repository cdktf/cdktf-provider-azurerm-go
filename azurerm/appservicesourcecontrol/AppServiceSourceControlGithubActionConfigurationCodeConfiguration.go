// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appservicesourcecontrol


type AppServiceSourceControlGithubActionConfigurationCodeConfiguration struct {
	// The value to use for the Runtime Stack in the workflow file content for code base apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/app_service_source_control#runtime_stack AppServiceSourceControlA#runtime_stack}
	RuntimeStack *string `field:"required" json:"runtimeStack" yaml:"runtimeStack"`
	// The value to use for the Runtime Version in the workflow file content for code base apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/app_service_source_control#runtime_version AppServiceSourceControlA#runtime_version}
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
}

