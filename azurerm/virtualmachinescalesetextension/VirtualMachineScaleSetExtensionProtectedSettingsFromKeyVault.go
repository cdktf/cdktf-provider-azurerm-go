// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescalesetextension


type VirtualMachineScaleSetExtensionProtectedSettingsFromKeyVault struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/virtual_machine_scale_set_extension#secret_url VirtualMachineScaleSetExtensionA#secret_url}.
	SecretUrl *string `field:"required" json:"secretUrl" yaml:"secretUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/virtual_machine_scale_set_extension#source_vault_id VirtualMachineScaleSetExtensionA#source_vault_id}.
	SourceVaultId *string `field:"required" json:"sourceVaultId" yaml:"sourceVaultId"`
}

