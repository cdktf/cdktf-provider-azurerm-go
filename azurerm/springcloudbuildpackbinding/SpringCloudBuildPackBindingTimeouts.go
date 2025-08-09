// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudbuildpackbinding


type SpringCloudBuildPackBindingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/spring_cloud_build_pack_binding#create SpringCloudBuildPackBinding#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/spring_cloud_build_pack_binding#delete SpringCloudBuildPackBinding#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/spring_cloud_build_pack_binding#read SpringCloudBuildPackBinding#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/spring_cloud_build_pack_binding#update SpringCloudBuildPackBinding#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

