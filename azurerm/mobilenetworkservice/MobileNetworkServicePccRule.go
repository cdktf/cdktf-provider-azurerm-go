// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkservice


type MobileNetworkServicePccRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/mobile_network_service#name MobileNetworkService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/mobile_network_service#precedence MobileNetworkService#precedence}.
	Precedence *float64 `field:"required" json:"precedence" yaml:"precedence"`
	// service_data_flow_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/mobile_network_service#service_data_flow_template MobileNetworkService#service_data_flow_template}
	ServiceDataFlowTemplate interface{} `field:"required" json:"serviceDataFlowTemplate" yaml:"serviceDataFlowTemplate"`
	// qos_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/mobile_network_service#qos_policy MobileNetworkService#qos_policy}
	QosPolicy *MobileNetworkServicePccRuleQosPolicy `field:"optional" json:"qosPolicy" yaml:"qosPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/mobile_network_service#traffic_control_enabled MobileNetworkService#traffic_control_enabled}.
	TrafficControlEnabled interface{} `field:"optional" json:"trafficControlEnabled" yaml:"trafficControlEnabled"`
}

