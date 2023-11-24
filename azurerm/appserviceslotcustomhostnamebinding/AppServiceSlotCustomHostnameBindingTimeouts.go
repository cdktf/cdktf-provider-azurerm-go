// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appserviceslotcustomhostnamebinding


type AppServiceSlotCustomHostnameBindingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/app_service_slot_custom_hostname_binding#create AppServiceSlotCustomHostnameBinding#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/app_service_slot_custom_hostname_binding#delete AppServiceSlotCustomHostnameBinding#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/app_service_slot_custom_hostname_binding#read AppServiceSlotCustomHostnameBinding#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

