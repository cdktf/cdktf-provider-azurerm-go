// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachine


type LinuxVirtualMachineAdditionalCapabilities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_virtual_machine#ultra_ssd_enabled LinuxVirtualMachine#ultra_ssd_enabled}.
	UltraSsdEnabled interface{} `field:"optional" json:"ultraSsdEnabled" yaml:"ultraSsdEnabled"`
}

