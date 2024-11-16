// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package digitaltwinsendpointeventgrid


type DigitalTwinsEndpointEventgridTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/digital_twins_endpoint_eventgrid#create DigitalTwinsEndpointEventgrid#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/digital_twins_endpoint_eventgrid#delete DigitalTwinsEndpointEventgrid#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/digital_twins_endpoint_eventgrid#read DigitalTwinsEndpointEventgrid#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/digital_twins_endpoint_eventgrid#update DigitalTwinsEndpointEventgrid#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

