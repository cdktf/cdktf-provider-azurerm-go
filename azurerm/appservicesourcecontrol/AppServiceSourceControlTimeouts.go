package appservicesourcecontrol


type AppServiceSourceControlTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#create AppServiceSourceControlA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#delete AppServiceSourceControlA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#read AppServiceSourceControlA#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
