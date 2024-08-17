// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebapp


type LinuxWebAppSiteConfigCors struct {
	// Specifies a list of origins that should be allowed to make cross-origin calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/linux_web_app#allowed_origins LinuxWebApp#allowed_origins}
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Are credentials allowed in CORS requests? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/linux_web_app#support_credentials LinuxWebApp#support_credentials}
	SupportCredentials interface{} `field:"optional" json:"supportCredentials" yaml:"supportCredentials"`
}

