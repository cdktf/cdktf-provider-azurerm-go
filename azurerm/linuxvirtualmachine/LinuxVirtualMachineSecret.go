// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachine


type LinuxVirtualMachineSecret struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/linux_virtual_machine#certificate LinuxVirtualMachine#certificate}
	Certificate interface{} `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/linux_virtual_machine#key_vault_id LinuxVirtualMachine#key_vault_id}.
	KeyVaultId *string `field:"required" json:"keyVaultId" yaml:"keyVaultId"`
}

