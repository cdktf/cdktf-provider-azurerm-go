// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinepacketcapture


type VirtualMachinePacketCaptureTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/virtual_machine_packet_capture#create VirtualMachinePacketCapture#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/virtual_machine_packet_capture#delete VirtualMachinePacketCapture#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/virtual_machine_packet_capture#read VirtualMachinePacketCapture#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

