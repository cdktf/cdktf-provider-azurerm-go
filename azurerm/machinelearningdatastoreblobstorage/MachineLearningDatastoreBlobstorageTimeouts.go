package machinelearningdatastoreblobstorage


type MachineLearningDatastoreBlobstorageTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/machine_learning_datastore_blobstorage#create MachineLearningDatastoreBlobstorage#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/machine_learning_datastore_blobstorage#delete MachineLearningDatastoreBlobstorage#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/machine_learning_datastore_blobstorage#read MachineLearningDatastoreBlobstorage#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/machine_learning_datastore_blobstorage#update MachineLearningDatastoreBlobstorage#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
