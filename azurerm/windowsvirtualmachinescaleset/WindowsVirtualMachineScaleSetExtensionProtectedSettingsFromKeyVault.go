// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetExtensionProtectedSettingsFromKeyVault struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/windows_virtual_machine_scale_set#secret_url WindowsVirtualMachineScaleSet#secret_url}.
	SecretUrl *string `field:"required" json:"secretUrl" yaml:"secretUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/windows_virtual_machine_scale_set#source_vault_id WindowsVirtualMachineScaleSet#source_vault_id}.
	SourceVaultId *string `field:"required" json:"sourceVaultId" yaml:"sourceVaultId"`
}

