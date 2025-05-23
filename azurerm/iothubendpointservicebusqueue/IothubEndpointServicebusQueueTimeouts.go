// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iothubendpointservicebusqueue


type IothubEndpointServicebusQueueTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/iothub_endpoint_servicebus_queue#create IothubEndpointServicebusQueue#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/iothub_endpoint_servicebus_queue#delete IothubEndpointServicebusQueue#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/iothub_endpoint_servicebus_queue#read IothubEndpointServicebusQueue#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/iothub_endpoint_servicebus_queue#update IothubEndpointServicebusQueue#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

