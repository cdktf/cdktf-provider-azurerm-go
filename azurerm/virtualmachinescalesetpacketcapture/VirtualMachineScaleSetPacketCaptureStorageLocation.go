package virtualmachinescalesetpacketcapture


type VirtualMachineScaleSetPacketCaptureStorageLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_scale_set_packet_capture#file_path VirtualMachineScaleSetPacketCapture#file_path}.
	FilePath *string `field:"optional" json:"filePath" yaml:"filePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_scale_set_packet_capture#storage_account_id VirtualMachineScaleSetPacketCapture#storage_account_id}.
	StorageAccountId *string `field:"optional" json:"storageAccountId" yaml:"storageAccountId"`
}
