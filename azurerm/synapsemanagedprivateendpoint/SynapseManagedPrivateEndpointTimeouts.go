// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapsemanagedprivateendpoint


type SynapseManagedPrivateEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/synapse_managed_private_endpoint#create SynapseManagedPrivateEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/synapse_managed_private_endpoint#delete SynapseManagedPrivateEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/synapse_managed_private_endpoint#read SynapseManagedPrivateEndpoint#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

