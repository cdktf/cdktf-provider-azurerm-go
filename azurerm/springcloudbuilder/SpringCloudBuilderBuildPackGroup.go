// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudbuilder


type SpringCloudBuilderBuildPackGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/spring_cloud_builder#name SpringCloudBuilder#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/spring_cloud_builder#build_pack_ids SpringCloudBuilder#build_pack_ids}.
	BuildPackIds *[]*string `field:"optional" json:"buildPackIds" yaml:"buildPackIds"`
}

