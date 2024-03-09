// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databoxedgeorder


type DataboxEdgeOrderContact struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/databox_edge_order#company_name DataboxEdgeOrder#company_name}.
	CompanyName *string `field:"required" json:"companyName" yaml:"companyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/databox_edge_order#emails DataboxEdgeOrder#emails}.
	Emails *[]*string `field:"required" json:"emails" yaml:"emails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/databox_edge_order#name DataboxEdgeOrder#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/databox_edge_order#phone_number DataboxEdgeOrder#phone_number}.
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
}

