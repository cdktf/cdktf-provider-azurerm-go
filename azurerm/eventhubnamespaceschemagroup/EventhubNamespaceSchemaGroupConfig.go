// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventhubnamespaceschemagroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventhubNamespaceSchemaGroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/eventhub_namespace_schema_group#name EventhubNamespaceSchemaGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/eventhub_namespace_schema_group#namespace_id EventhubNamespaceSchemaGroup#namespace_id}.
	NamespaceId *string `field:"required" json:"namespaceId" yaml:"namespaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/eventhub_namespace_schema_group#schema_compatibility EventhubNamespaceSchemaGroup#schema_compatibility}.
	SchemaCompatibility *string `field:"required" json:"schemaCompatibility" yaml:"schemaCompatibility"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/eventhub_namespace_schema_group#schema_type EventhubNamespaceSchemaGroup#schema_type}.
	SchemaType *string `field:"required" json:"schemaType" yaml:"schemaType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/eventhub_namespace_schema_group#id EventhubNamespaceSchemaGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/eventhub_namespace_schema_group#timeouts EventhubNamespaceSchemaGroup#timeouts}
	Timeouts *EventhubNamespaceSchemaGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

