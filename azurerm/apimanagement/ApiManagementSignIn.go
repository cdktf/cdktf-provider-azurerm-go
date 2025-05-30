// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagement


type ApiManagementSignIn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/api_management#enabled ApiManagement#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

