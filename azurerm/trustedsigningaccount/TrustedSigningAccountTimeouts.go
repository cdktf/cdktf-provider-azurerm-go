// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package trustedsigningaccount


type TrustedSigningAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/trusted_signing_account#create TrustedSigningAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/trusted_signing_account#delete TrustedSigningAccount#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/trusted_signing_account#read TrustedSigningAccount#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/trusted_signing_account#update TrustedSigningAccount#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

