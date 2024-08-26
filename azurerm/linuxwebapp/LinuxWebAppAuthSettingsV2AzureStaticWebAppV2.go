// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebapp


type LinuxWebAppAuthSettingsV2AzureStaticWebAppV2 struct {
	// The ID of the Client to use to authenticate with Azure Static Web App Authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/linux_web_app#client_id LinuxWebApp#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
}

