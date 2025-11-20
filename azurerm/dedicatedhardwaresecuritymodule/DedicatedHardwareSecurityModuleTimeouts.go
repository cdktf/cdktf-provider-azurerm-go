// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dedicatedhardwaresecuritymodule


type DedicatedHardwareSecurityModuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/dedicated_hardware_security_module#create DedicatedHardwareSecurityModule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/dedicated_hardware_security_module#delete DedicatedHardwareSecurityModule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/dedicated_hardware_security_module#read DedicatedHardwareSecurityModule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/dedicated_hardware_security_module#update DedicatedHardwareSecurityModule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

