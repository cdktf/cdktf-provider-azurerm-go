package linuxfunctionappslot


type LinuxFunctionAppSlotTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_function_app_slot#create LinuxFunctionAppSlot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_function_app_slot#delete LinuxFunctionAppSlot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_function_app_slot#read LinuxFunctionAppSlot#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_function_app_slot#update LinuxFunctionAppSlot#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

