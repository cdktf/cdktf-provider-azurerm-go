// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudapplicationliveview


type SpringCloudApplicationLiveViewTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/spring_cloud_application_live_view#create SpringCloudApplicationLiveView#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/spring_cloud_application_live_view#delete SpringCloudApplicationLiveView#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/spring_cloud_application_live_view#read SpringCloudApplicationLiveView#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

