// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedservicekeyvault


type DataFactoryLinkedServiceKeyVaultTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/data_factory_linked_service_key_vault#create DataFactoryLinkedServiceKeyVault#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/data_factory_linked_service_key_vault#delete DataFactoryLinkedServiceKeyVault#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/data_factory_linked_service_key_vault#read DataFactoryLinkedServiceKeyVault#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/data_factory_linked_service_key_vault#update DataFactoryLinkedServiceKeyVault#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

