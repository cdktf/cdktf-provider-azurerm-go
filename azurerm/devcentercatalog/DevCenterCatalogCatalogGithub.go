// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devcentercatalog


type DevCenterCatalogCatalogGithub struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/dev_center_catalog#branch DevCenterCatalog#branch}.
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/dev_center_catalog#key_vault_key_url DevCenterCatalog#key_vault_key_url}.
	KeyVaultKeyUrl *string `field:"required" json:"keyVaultKeyUrl" yaml:"keyVaultKeyUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/dev_center_catalog#path DevCenterCatalog#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/dev_center_catalog#uri DevCenterCatalog#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

