package nginxcertificate


type NginxCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/nginx_certificate#create NginxCertificate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/nginx_certificate#delete NginxCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/nginx_certificate#read NginxCertificate#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

