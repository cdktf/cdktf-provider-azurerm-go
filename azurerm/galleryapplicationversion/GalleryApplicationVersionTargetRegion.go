package galleryapplicationversion


type GalleryApplicationVersionTargetRegion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#name GalleryApplicationVersion#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#regional_replica_count GalleryApplicationVersion#regional_replica_count}.
	RegionalReplicaCount *float64 `field:"required" json:"regionalReplicaCount" yaml:"regionalReplicaCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#storage_account_type GalleryApplicationVersion#storage_account_type}.
	StorageAccountType *string `field:"optional" json:"storageAccountType" yaml:"storageAccountType"`
}
