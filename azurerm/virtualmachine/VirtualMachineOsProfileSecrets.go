// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine


type VirtualMachineOsProfileSecrets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/virtual_machine#source_vault_id VirtualMachine#source_vault_id}.
	SourceVaultId *string `field:"required" json:"sourceVaultId" yaml:"sourceVaultId"`
	// vault_certificates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/virtual_machine#vault_certificates VirtualMachine#vault_certificates}
	VaultCertificates interface{} `field:"optional" json:"vaultCertificates" yaml:"vaultCertificates"`
}

