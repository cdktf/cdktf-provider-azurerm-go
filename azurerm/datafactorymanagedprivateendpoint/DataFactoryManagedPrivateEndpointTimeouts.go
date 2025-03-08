// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorymanagedprivateendpoint


type DataFactoryManagedPrivateEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/data_factory_managed_private_endpoint#create DataFactoryManagedPrivateEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/data_factory_managed_private_endpoint#delete DataFactoryManagedPrivateEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/data_factory_managed_private_endpoint#read DataFactoryManagedPrivateEndpoint#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

