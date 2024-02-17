// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logzsubaccount


type LogzSubAccountUser struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/logz_sub_account#email LogzSubAccount#email}.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/logz_sub_account#first_name LogzSubAccount#first_name}.
	FirstName *string `field:"required" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/logz_sub_account#last_name LogzSubAccount#last_name}.
	LastName *string `field:"required" json:"lastName" yaml:"lastName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/logz_sub_account#phone_number LogzSubAccount#phone_number}.
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
}

