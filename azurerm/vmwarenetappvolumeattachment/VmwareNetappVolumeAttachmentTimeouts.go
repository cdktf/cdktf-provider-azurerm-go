package vmwarenetappvolumeattachment


type VmwareNetappVolumeAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vmware_netapp_volume_attachment#create VmwareNetappVolumeAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vmware_netapp_volume_attachment#delete VmwareNetappVolumeAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vmware_netapp_volume_attachment#read VmwareNetappVolumeAttachment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
