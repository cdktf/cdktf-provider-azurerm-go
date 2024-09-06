// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventhubnamespacecustomermanagedkey


type EventhubNamespaceCustomerManagedKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/eventhub_namespace_customer_managed_key#create EventhubNamespaceCustomerManagedKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/eventhub_namespace_customer_managed_key#delete EventhubNamespaceCustomerManagedKey#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/eventhub_namespace_customer_managed_key#read EventhubNamespaceCustomerManagedKey#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/eventhub_namespace_customer_managed_key#update EventhubNamespaceCustomerManagedKey#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

