// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgriddomain


type EventgridDomainInputMappingFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/eventgrid_domain#data_version EventgridDomain#data_version}.
	DataVersion *string `field:"optional" json:"dataVersion" yaml:"dataVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/eventgrid_domain#event_time EventgridDomain#event_time}.
	EventTime *string `field:"optional" json:"eventTime" yaml:"eventTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/eventgrid_domain#event_type EventgridDomain#event_type}.
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/eventgrid_domain#id EventgridDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/eventgrid_domain#subject EventgridDomain#subject}.
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/eventgrid_domain#topic EventgridDomain#topic}.
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

