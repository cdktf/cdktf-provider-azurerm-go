package labserviceplan


type LabServicePlanSupport struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lab_service_plan#email LabServicePlan#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lab_service_plan#instructions LabServicePlan#instructions}.
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lab_service_plan#phone LabServicePlan#phone}.
	Phone *string `field:"optional" json:"phone" yaml:"phone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lab_service_plan#url LabServicePlan#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}
