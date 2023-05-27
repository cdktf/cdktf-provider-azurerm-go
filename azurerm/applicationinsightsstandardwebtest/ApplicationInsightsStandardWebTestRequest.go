package applicationinsightsstandardwebtest


type ApplicationInsightsStandardWebTestRequest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/application_insights_standard_web_test#url ApplicationInsightsStandardWebTest#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/application_insights_standard_web_test#body ApplicationInsightsStandardWebTest#body}.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/application_insights_standard_web_test#follow_redirects_enabled ApplicationInsightsStandardWebTest#follow_redirects_enabled}.
	FollowRedirectsEnabled interface{} `field:"optional" json:"followRedirectsEnabled" yaml:"followRedirectsEnabled"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/application_insights_standard_web_test#header ApplicationInsightsStandardWebTest#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/application_insights_standard_web_test#http_verb ApplicationInsightsStandardWebTest#http_verb}.
	HttpVerb *string `field:"optional" json:"httpVerb" yaml:"httpVerb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/application_insights_standard_web_test#parse_dependent_requests_enabled ApplicationInsightsStandardWebTest#parse_dependent_requests_enabled}.
	ParseDependentRequestsEnabled interface{} `field:"optional" json:"parseDependentRequestsEnabled" yaml:"parseDependentRequestsEnabled"`
}

