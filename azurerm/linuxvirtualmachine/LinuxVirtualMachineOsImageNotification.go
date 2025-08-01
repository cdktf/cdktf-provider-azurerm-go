// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachine


type LinuxVirtualMachineOsImageNotification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/linux_virtual_machine#timeout LinuxVirtualMachine#timeout}.
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

