// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkservice


type MobileNetworkServiceServiceQosPolicy struct {
	// maximum_bit_rate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/mobile_network_service#maximum_bit_rate MobileNetworkService#maximum_bit_rate}
	MaximumBitRate *MobileNetworkServiceServiceQosPolicyMaximumBitRate `field:"required" json:"maximumBitRate" yaml:"maximumBitRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/mobile_network_service#allocation_and_retention_priority_level MobileNetworkService#allocation_and_retention_priority_level}.
	AllocationAndRetentionPriorityLevel *float64 `field:"optional" json:"allocationAndRetentionPriorityLevel" yaml:"allocationAndRetentionPriorityLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/mobile_network_service#preemption_capability MobileNetworkService#preemption_capability}.
	PreemptionCapability *string `field:"optional" json:"preemptionCapability" yaml:"preemptionCapability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/mobile_network_service#preemption_vulnerability MobileNetworkService#preemption_vulnerability}.
	PreemptionVulnerability *string `field:"optional" json:"preemptionVulnerability" yaml:"preemptionVulnerability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/mobile_network_service#qos_indicator MobileNetworkService#qos_indicator}.
	QosIndicator *float64 `field:"optional" json:"qosIndicator" yaml:"qosIndicator"`
}

