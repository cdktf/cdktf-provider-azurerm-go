// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorycustomermanagedkey


type DataFactoryCustomerManagedKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/data_factory_customer_managed_key#create DataFactoryCustomerManagedKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/data_factory_customer_managed_key#delete DataFactoryCustomerManagedKey#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/data_factory_customer_managed_key#read DataFactoryCustomerManagedKey#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/data_factory_customer_managed_key#update DataFactoryCustomerManagedKey#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

