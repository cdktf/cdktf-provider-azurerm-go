// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appservicesourcecontrolslot


type AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration struct {
	// The value to use for the Runtime Stack in the workflow file content for code base apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.113.0/docs/resources/app_service_source_control_slot#runtime_stack AppServiceSourceControlSlot#runtime_stack}
	RuntimeStack *string `field:"required" json:"runtimeStack" yaml:"runtimeStack"`
	// The value to use for the Runtime Version in the workflow file content for code base apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.113.0/docs/resources/app_service_source_control_slot#runtime_version AppServiceSourceControlSlot#runtime_version}
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
}

