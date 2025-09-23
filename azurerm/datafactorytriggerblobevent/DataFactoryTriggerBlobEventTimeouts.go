// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorytriggerblobevent


type DataFactoryTriggerBlobEventTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_trigger_blob_event#create DataFactoryTriggerBlobEvent#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_trigger_blob_event#delete DataFactoryTriggerBlobEvent#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_trigger_blob_event#read DataFactoryTriggerBlobEvent#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_trigger_blob_event#update DataFactoryTriggerBlobEvent#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

