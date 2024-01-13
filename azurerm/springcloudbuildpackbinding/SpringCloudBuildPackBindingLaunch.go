// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudbuildpackbinding


type SpringCloudBuildPackBindingLaunch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/spring_cloud_build_pack_binding#properties SpringCloudBuildPackBinding#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/spring_cloud_build_pack_binding#secrets SpringCloudBuildPackBinding#secrets}.
	Secrets *map[string]*string `field:"optional" json:"secrets" yaml:"secrets"`
}

