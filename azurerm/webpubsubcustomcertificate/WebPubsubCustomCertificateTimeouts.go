package webpubsubcustomcertificate


type WebPubsubCustomCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_pubsub_custom_certificate#create WebPubsubCustomCertificate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_pubsub_custom_certificate#delete WebPubsubCustomCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_pubsub_custom_certificate#read WebPubsubCustomCertificate#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

