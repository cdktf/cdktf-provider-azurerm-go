package applicationinsightsstandardwebtest


type ApplicationInsightsStandardWebTestValidationRulesContent struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_insights_standard_web_test#content_match ApplicationInsightsStandardWebTest#content_match}.
	ContentMatch *string `field:"required" json:"contentMatch" yaml:"contentMatch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_insights_standard_web_test#ignore_case ApplicationInsightsStandardWebTest#ignore_case}.
	IgnoreCase interface{} `field:"optional" json:"ignoreCase" yaml:"ignoreCase"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_insights_standard_web_test#pass_if_text_found ApplicationInsightsStandardWebTest#pass_if_text_found}.
	PassIfTextFound interface{} `field:"optional" json:"passIfTextFound" yaml:"passIfTextFound"`
}
