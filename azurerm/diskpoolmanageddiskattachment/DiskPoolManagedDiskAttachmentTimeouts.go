package diskpoolmanageddiskattachment


type DiskPoolManagedDiskAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/disk_pool_managed_disk_attachment#create DiskPoolManagedDiskAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/disk_pool_managed_disk_attachment#delete DiskPoolManagedDiskAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/disk_pool_managed_disk_attachment#read DiskPoolManagedDiskAttachment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
