package virtualmachinescalesetpacketcapture


type VirtualMachineScaleSetPacketCaptureMachineScope struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_scale_set_packet_capture#exclude_instance_ids VirtualMachineScaleSetPacketCapture#exclude_instance_ids}.
	ExcludeInstanceIds *[]*string `field:"optional" json:"excludeInstanceIds" yaml:"excludeInstanceIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_scale_set_packet_capture#include_instance_ids VirtualMachineScaleSetPacketCapture#include_instance_ids}.
	IncludeInstanceIds *[]*string `field:"optional" json:"includeInstanceIds" yaml:"includeInstanceIds"`
}
