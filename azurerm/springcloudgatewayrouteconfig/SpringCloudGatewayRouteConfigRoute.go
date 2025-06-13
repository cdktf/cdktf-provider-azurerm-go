// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudgatewayrouteconfig


type SpringCloudGatewayRouteConfigRoute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/spring_cloud_gateway_route_config#order SpringCloudGatewayRouteConfig#order}.
	Order *float64 `field:"required" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/spring_cloud_gateway_route_config#classification_tags SpringCloudGatewayRouteConfig#classification_tags}.
	ClassificationTags *[]*string `field:"optional" json:"classificationTags" yaml:"classificationTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/spring_cloud_gateway_route_config#description SpringCloudGatewayRouteConfig#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/spring_cloud_gateway_route_config#filters SpringCloudGatewayRouteConfig#filters}.
	Filters *[]*string `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/spring_cloud_gateway_route_config#predicates SpringCloudGatewayRouteConfig#predicates}.
	Predicates *[]*string `field:"optional" json:"predicates" yaml:"predicates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/spring_cloud_gateway_route_config#sso_validation_enabled SpringCloudGatewayRouteConfig#sso_validation_enabled}.
	SsoValidationEnabled interface{} `field:"optional" json:"ssoValidationEnabled" yaml:"ssoValidationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/spring_cloud_gateway_route_config#title SpringCloudGatewayRouteConfig#title}.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/spring_cloud_gateway_route_config#token_relay SpringCloudGatewayRouteConfig#token_relay}.
	TokenRelay interface{} `field:"optional" json:"tokenRelay" yaml:"tokenRelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/spring_cloud_gateway_route_config#uri SpringCloudGatewayRouteConfig#uri}.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

