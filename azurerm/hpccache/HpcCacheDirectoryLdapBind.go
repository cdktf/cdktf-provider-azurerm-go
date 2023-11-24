// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hpccache


type HpcCacheDirectoryLdapBind struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/hpc_cache#dn HpcCache#dn}.
	Dn *string `field:"required" json:"dn" yaml:"dn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/hpc_cache#password HpcCache#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
}

