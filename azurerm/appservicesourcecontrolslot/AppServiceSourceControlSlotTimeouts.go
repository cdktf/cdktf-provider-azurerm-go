package appservicesourcecontrolslot


type AppServiceSourceControlSlotTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#create AppServiceSourceControlSlot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#delete AppServiceSourceControlSlot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#read AppServiceSourceControlSlot#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
