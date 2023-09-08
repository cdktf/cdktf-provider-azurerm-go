// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworksimpolicy


type MobileNetworkSimPolicySliceDataNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mobile_network_sim_policy#allowed_services_ids MobileNetworkSimPolicy#allowed_services_ids}.
	AllowedServicesIds *[]*string `field:"required" json:"allowedServicesIds" yaml:"allowedServicesIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mobile_network_sim_policy#data_network_id MobileNetworkSimPolicy#data_network_id}.
	DataNetworkId *string `field:"required" json:"dataNetworkId" yaml:"dataNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mobile_network_sim_policy#qos_indicator MobileNetworkSimPolicy#qos_indicator}.
	QosIndicator *float64 `field:"required" json:"qosIndicator" yaml:"qosIndicator"`
	// session_aggregate_maximum_bit_rate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mobile_network_sim_policy#session_aggregate_maximum_bit_rate MobileNetworkSimPolicy#session_aggregate_maximum_bit_rate}
	SessionAggregateMaximumBitRate *MobileNetworkSimPolicySliceDataNetworkSessionAggregateMaximumBitRate `field:"required" json:"sessionAggregateMaximumBitRate" yaml:"sessionAggregateMaximumBitRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mobile_network_sim_policy#additional_allowed_session_types MobileNetworkSimPolicy#additional_allowed_session_types}.
	AdditionalAllowedSessionTypes *[]*string `field:"optional" json:"additionalAllowedSessionTypes" yaml:"additionalAllowedSessionTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mobile_network_sim_policy#allocation_and_retention_priority_level MobileNetworkSimPolicy#allocation_and_retention_priority_level}.
	AllocationAndRetentionPriorityLevel *float64 `field:"optional" json:"allocationAndRetentionPriorityLevel" yaml:"allocationAndRetentionPriorityLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mobile_network_sim_policy#default_session_type MobileNetworkSimPolicy#default_session_type}.
	DefaultSessionType *string `field:"optional" json:"defaultSessionType" yaml:"defaultSessionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mobile_network_sim_policy#max_buffered_packets MobileNetworkSimPolicy#max_buffered_packets}.
	MaxBufferedPackets *float64 `field:"optional" json:"maxBufferedPackets" yaml:"maxBufferedPackets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mobile_network_sim_policy#preemption_capability MobileNetworkSimPolicy#preemption_capability}.
	PreemptionCapability *string `field:"optional" json:"preemptionCapability" yaml:"preemptionCapability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mobile_network_sim_policy#preemption_vulnerability MobileNetworkSimPolicy#preemption_vulnerability}.
	PreemptionVulnerability *string `field:"optional" json:"preemptionVulnerability" yaml:"preemptionVulnerability"`
}

