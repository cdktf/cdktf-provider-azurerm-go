// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebappslot


type LinuxWebAppSlotSiteConfigCors struct {
	// Specifies a list of origins that should be allowed to make cross-origin calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/linux_web_app_slot#allowed_origins LinuxWebAppSlot#allowed_origins}
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Are credentials allowed in CORS requests? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/linux_web_app_slot#support_credentials LinuxWebAppSlot#support_credentials}
	SupportCredentials interface{} `field:"optional" json:"supportCredentials" yaml:"supportCredentials"`
}

