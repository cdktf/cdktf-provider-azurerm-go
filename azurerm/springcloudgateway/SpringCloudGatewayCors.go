// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudgateway


type SpringCloudGatewayCors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/spring_cloud_gateway#allowed_headers SpringCloudGateway#allowed_headers}.
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/spring_cloud_gateway#allowed_methods SpringCloudGateway#allowed_methods}.
	AllowedMethods *[]*string `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/spring_cloud_gateway#allowed_origin_patterns SpringCloudGateway#allowed_origin_patterns}.
	AllowedOriginPatterns *[]*string `field:"optional" json:"allowedOriginPatterns" yaml:"allowedOriginPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/spring_cloud_gateway#allowed_origins SpringCloudGateway#allowed_origins}.
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/spring_cloud_gateway#credentials_allowed SpringCloudGateway#credentials_allowed}.
	CredentialsAllowed interface{} `field:"optional" json:"credentialsAllowed" yaml:"credentialsAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/spring_cloud_gateway#exposed_headers SpringCloudGateway#exposed_headers}.
	ExposedHeaders *[]*string `field:"optional" json:"exposedHeaders" yaml:"exposedHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/spring_cloud_gateway#max_age_seconds SpringCloudGateway#max_age_seconds}.
	MaxAgeSeconds *float64 `field:"optional" json:"maxAgeSeconds" yaml:"maxAgeSeconds"`
}

