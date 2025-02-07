// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nginxconfiguration


type NginxConfigurationProtectedFile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/nginx_configuration#content NginxConfiguration#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/nginx_configuration#virtual_path NginxConfiguration#virtual_path}.
	VirtualPath *string `field:"required" json:"virtualPath" yaml:"virtualPath"`
}

