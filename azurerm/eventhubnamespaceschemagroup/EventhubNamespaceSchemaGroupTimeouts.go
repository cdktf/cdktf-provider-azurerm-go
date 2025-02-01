// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventhubnamespaceschemagroup


type EventhubNamespaceSchemaGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/eventhub_namespace_schema_group#create EventhubNamespaceSchemaGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/eventhub_namespace_schema_group#delete EventhubNamespaceSchemaGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/eventhub_namespace_schema_group#read EventhubNamespaceSchemaGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

