package applicationinsightsstandardwebtest


type ApplicationInsightsStandardWebTestRequestHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/application_insights_standard_web_test#name ApplicationInsightsStandardWebTest#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/application_insights_standard_web_test#value ApplicationInsightsStandardWebTest#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

