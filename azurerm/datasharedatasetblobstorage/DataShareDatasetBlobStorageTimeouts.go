package datasharedatasetblobstorage


type DataShareDatasetBlobStorageTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_share_dataset_blob_storage#create DataShareDatasetBlobStorage#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_share_dataset_blob_storage#delete DataShareDatasetBlobStorage#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_share_dataset_blob_storage#read DataShareDatasetBlobStorage#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

