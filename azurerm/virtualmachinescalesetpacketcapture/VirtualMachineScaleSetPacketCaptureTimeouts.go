package virtualmachinescalesetpacketcapture


type VirtualMachineScaleSetPacketCaptureTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_scale_set_packet_capture#create VirtualMachineScaleSetPacketCapture#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_scale_set_packet_capture#delete VirtualMachineScaleSetPacketCapture#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_scale_set_packet_capture#read VirtualMachineScaleSetPacketCapture#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
