// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package federatedidentitycredential


type FederatedIdentityCredentialTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/federated_identity_credential#create FederatedIdentityCredential#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/federated_identity_credential#delete FederatedIdentityCredential#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/federated_identity_credential#read FederatedIdentityCredential#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/federated_identity_credential#update FederatedIdentityCredential#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

