// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagerverifierworkspace


type NetworkManagerVerifierWorkspaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/network_manager_verifier_workspace#create NetworkManagerVerifierWorkspace#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/network_manager_verifier_workspace#delete NetworkManagerVerifierWorkspace#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/network_manager_verifier_workspace#read NetworkManagerVerifierWorkspace#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/network_manager_verifier_workspace#update NetworkManagerVerifierWorkspace#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

