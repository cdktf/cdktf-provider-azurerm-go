package linuxfunctionapp


type LinuxFunctionAppTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_function_app#create LinuxFunctionApp#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_function_app#delete LinuxFunctionApp#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_function_app#read LinuxFunctionApp#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_function_app#update LinuxFunctionApp#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
