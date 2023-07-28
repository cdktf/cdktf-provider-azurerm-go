package applicationinsightswebtest


type ApplicationInsightsWebTestTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/application_insights_web_test#create ApplicationInsightsWebTest#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/application_insights_web_test#delete ApplicationInsightsWebTest#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/application_insights_web_test#read ApplicationInsightsWebTest#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/application_insights_web_test#update ApplicationInsightsWebTest#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

