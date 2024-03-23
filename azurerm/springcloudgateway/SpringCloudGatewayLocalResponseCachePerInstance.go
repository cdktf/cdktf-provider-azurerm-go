// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudgateway


type SpringCloudGatewayLocalResponseCachePerInstance struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/spring_cloud_gateway#size SpringCloudGateway#size}.
	Size *string `field:"optional" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/spring_cloud_gateway#time_to_live SpringCloudGateway#time_to_live}.
	TimeToLive *string `field:"optional" json:"timeToLive" yaml:"timeToLive"`
}

