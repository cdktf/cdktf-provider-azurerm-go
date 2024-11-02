// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachine


type WindowsVirtualMachineSecret struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/windows_virtual_machine#certificate WindowsVirtualMachine#certificate}
	Certificate interface{} `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/windows_virtual_machine#key_vault_id WindowsVirtualMachine#key_vault_id}.
	KeyVaultId *string `field:"required" json:"keyVaultId" yaml:"keyVaultId"`
}

