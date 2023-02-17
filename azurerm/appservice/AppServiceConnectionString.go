package appservice


type AppServiceConnectionString struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#name AppService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#type AppService#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#value AppService#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

