// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databoxedgeorder


type DataboxEdgeOrderShipmentAddress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databox_edge_order#address DataboxEdgeOrder#address}.
	Address *[]*string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databox_edge_order#city DataboxEdgeOrder#city}.
	City *string `field:"required" json:"city" yaml:"city"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databox_edge_order#country DataboxEdgeOrder#country}.
	Country *string `field:"required" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databox_edge_order#postal_code DataboxEdgeOrder#postal_code}.
	PostalCode *string `field:"required" json:"postalCode" yaml:"postalCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databox_edge_order#state DataboxEdgeOrder#state}.
	State *string `field:"required" json:"state" yaml:"state"`
}

