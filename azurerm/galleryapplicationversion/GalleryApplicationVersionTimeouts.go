package galleryapplicationversion


type GalleryApplicationVersionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#create GalleryApplicationVersion#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#delete GalleryApplicationVersion#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#read GalleryApplicationVersion#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#update GalleryApplicationVersion#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
