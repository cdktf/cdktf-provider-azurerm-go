package virtualmachinedatadiskattachment


type VirtualMachineDataDiskAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_data_disk_attachment#create VirtualMachineDataDiskAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_data_disk_attachment#delete VirtualMachineDataDiskAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_data_disk_attachment#read VirtualMachineDataDiskAttachment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_data_disk_attachment#update VirtualMachineDataDiskAttachment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

