// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package videoanalyzer


type VideoAnalyzerStorageAccount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/video_analyzer#id VideoAnalyzer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/video_analyzer#user_assigned_identity_id VideoAnalyzer#user_assigned_identity_id}.
	UserAssignedIdentityId *string `field:"required" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
}

