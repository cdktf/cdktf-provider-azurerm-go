package webappactiveslot


type WebAppActiveSlotTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_app_active_slot#create WebAppActiveSlot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_app_active_slot#delete WebAppActiveSlot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_app_active_slot#read WebAppActiveSlot#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_app_active_slot#update WebAppActiveSlot#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
