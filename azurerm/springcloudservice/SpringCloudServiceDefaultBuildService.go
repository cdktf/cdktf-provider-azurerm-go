// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudservice


type SpringCloudServiceDefaultBuildService struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/spring_cloud_service#container_registry_name SpringCloudService#container_registry_name}.
	ContainerRegistryName *string `field:"optional" json:"containerRegistryName" yaml:"containerRegistryName"`
}

