// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedservicesnowflake


type DataFactoryLinkedServiceSnowflakeTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/data_factory_linked_service_snowflake#create DataFactoryLinkedServiceSnowflake#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/data_factory_linked_service_snowflake#delete DataFactoryLinkedServiceSnowflake#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/data_factory_linked_service_snowflake#read DataFactoryLinkedServiceSnowflake#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/data_factory_linked_service_snowflake#update DataFactoryLinkedServiceSnowflake#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

