// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappslot


type FunctionAppSlotSiteConfigScmIpRestrictionHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/function_app_slot#x_azure_fdid FunctionAppSlot#x_azure_fdid}.
	XAzureFdid *[]*string `field:"optional" json:"xAzureFdid" yaml:"xAzureFdid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/function_app_slot#x_fd_health_probe FunctionAppSlot#x_fd_health_probe}.
	XFdHealthProbe *[]*string `field:"optional" json:"xFdHealthProbe" yaml:"xFdHealthProbe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/function_app_slot#x_forwarded_for FunctionAppSlot#x_forwarded_for}.
	XForwardedFor *[]*string `field:"optional" json:"xForwardedFor" yaml:"xForwardedFor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/function_app_slot#x_forwarded_host FunctionAppSlot#x_forwarded_host}.
	XForwardedHost *[]*string `field:"optional" json:"xForwardedHost" yaml:"xForwardedHost"`
}

