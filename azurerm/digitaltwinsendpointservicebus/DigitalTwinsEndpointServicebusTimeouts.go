// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package digitaltwinsendpointservicebus


type DigitalTwinsEndpointServicebusTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/digital_twins_endpoint_servicebus#create DigitalTwinsEndpointServicebus#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/digital_twins_endpoint_servicebus#delete DigitalTwinsEndpointServicebus#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/digital_twins_endpoint_servicebus#read DigitalTwinsEndpointServicebus#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/digital_twins_endpoint_servicebus#update DigitalTwinsEndpointServicebus#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

