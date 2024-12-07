// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool


type BatchPoolUserAccountsLinuxUserConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/batch_pool#gid BatchPool#gid}.
	Gid *float64 `field:"optional" json:"gid" yaml:"gid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/batch_pool#ssh_private_key BatchPool#ssh_private_key}.
	SshPrivateKey *string `field:"optional" json:"sshPrivateKey" yaml:"sshPrivateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/batch_pool#uid BatchPool#uid}.
	Uid *float64 `field:"optional" json:"uid" yaml:"uid"`
}

