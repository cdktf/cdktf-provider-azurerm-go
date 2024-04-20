// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudservice


type SpringCloudServiceContainerRegistry struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/spring_cloud_service#name SpringCloudService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/spring_cloud_service#password SpringCloudService#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/spring_cloud_service#server SpringCloudService#server}.
	Server *string `field:"required" json:"server" yaml:"server"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/spring_cloud_service#username SpringCloudService#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

