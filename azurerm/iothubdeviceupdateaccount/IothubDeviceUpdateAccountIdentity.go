// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iothubdeviceupdateaccount


type IothubDeviceUpdateAccountIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/iothub_device_update_account#type IothubDeviceUpdateAccount#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/iothub_device_update_account#identity_ids IothubDeviceUpdateAccount#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

