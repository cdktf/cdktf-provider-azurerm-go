// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package graphservicesaccount


type GraphServicesAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/graph_services_account#create GraphServicesAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/graph_services_account#delete GraphServicesAccount#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/graph_services_account#read GraphServicesAccount#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/graph_services_account#update GraphServicesAccount#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

