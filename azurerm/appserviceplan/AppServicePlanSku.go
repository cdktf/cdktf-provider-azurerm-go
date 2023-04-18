package appserviceplan


type AppServicePlanSku struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/app_service_plan#size AppServicePlan#size}.
	Size *string `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/app_service_plan#tier AppServicePlan#tier}.
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/app_service_plan#capacity AppServicePlan#capacity}.
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
}

