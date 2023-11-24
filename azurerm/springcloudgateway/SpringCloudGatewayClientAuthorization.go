// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudgateway


type SpringCloudGatewayClientAuthorization struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/spring_cloud_gateway#certificate_ids SpringCloudGateway#certificate_ids}.
	CertificateIds *[]*string `field:"optional" json:"certificateIds" yaml:"certificateIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/spring_cloud_gateway#verification_enabled SpringCloudGateway#verification_enabled}.
	VerificationEnabled interface{} `field:"optional" json:"verificationEnabled" yaml:"verificationEnabled"`
}

