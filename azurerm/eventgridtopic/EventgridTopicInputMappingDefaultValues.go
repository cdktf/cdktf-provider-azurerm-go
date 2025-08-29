// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgridtopic


type EventgridTopicInputMappingDefaultValues struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/eventgrid_topic#data_version EventgridTopic#data_version}.
	DataVersion *string `field:"optional" json:"dataVersion" yaml:"dataVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/eventgrid_topic#event_type EventgridTopic#event_type}.
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/eventgrid_topic#subject EventgridTopic#subject}.
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

