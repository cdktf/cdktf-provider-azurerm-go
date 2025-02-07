// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine


type VirtualMachineOsProfileWindowsConfigWinrm struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/virtual_machine#protocol VirtualMachine#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/virtual_machine#certificate_url VirtualMachine#certificate_url}.
	CertificateUrl *string `field:"optional" json:"certificateUrl" yaml:"certificateUrl"`
}

