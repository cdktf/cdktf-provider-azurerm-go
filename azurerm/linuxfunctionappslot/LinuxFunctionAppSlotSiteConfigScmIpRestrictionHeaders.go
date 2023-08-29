// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionappslot


type LinuxFunctionAppSlotSiteConfigScmIpRestrictionHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_function_app_slot#x_azure_fdid LinuxFunctionAppSlot#x_azure_fdid}.
	XAzureFdid *[]*string `field:"optional" json:"xAzureFdid" yaml:"xAzureFdid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_function_app_slot#x_fd_health_probe LinuxFunctionAppSlot#x_fd_health_probe}.
	XFdHealthProbe *[]*string `field:"optional" json:"xFdHealthProbe" yaml:"xFdHealthProbe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_function_app_slot#x_forwarded_for LinuxFunctionAppSlot#x_forwarded_for}.
	XForwardedFor *[]*string `field:"optional" json:"xForwardedFor" yaml:"xForwardedFor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_function_app_slot#x_forwarded_host LinuxFunctionAppSlot#x_forwarded_host}.
	XForwardedHost *[]*string `field:"optional" json:"xForwardedHost" yaml:"xForwardedHost"`
}

