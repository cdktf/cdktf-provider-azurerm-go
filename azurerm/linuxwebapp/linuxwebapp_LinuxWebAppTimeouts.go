package linuxwebapp


type LinuxWebAppTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#create LinuxWebApp#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#delete LinuxWebApp#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#read LinuxWebApp#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#update LinuxWebApp#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
