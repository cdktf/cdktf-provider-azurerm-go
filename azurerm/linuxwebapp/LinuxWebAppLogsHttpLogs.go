package linuxwebapp


type LinuxWebAppLogsHttpLogs struct {
	// azure_blob_storage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#azure_blob_storage LinuxWebApp#azure_blob_storage}
	AzureBlobStorage *LinuxWebAppLogsHttpLogsAzureBlobStorage `field:"optional" json:"azureBlobStorage" yaml:"azureBlobStorage"`
	// file_system block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#file_system LinuxWebApp#file_system}
	FileSystem *LinuxWebAppLogsHttpLogsFileSystem `field:"optional" json:"fileSystem" yaml:"fileSystem"`
}
