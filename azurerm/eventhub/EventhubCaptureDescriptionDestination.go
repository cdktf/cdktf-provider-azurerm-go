package eventhub


type EventhubCaptureDescriptionDestination struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventhub#archive_name_format Eventhub#archive_name_format}.
	ArchiveNameFormat *string `field:"required" json:"archiveNameFormat" yaml:"archiveNameFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventhub#blob_container_name Eventhub#blob_container_name}.
	BlobContainerName *string `field:"required" json:"blobContainerName" yaml:"blobContainerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventhub#name Eventhub#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventhub#storage_account_id Eventhub#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
}

