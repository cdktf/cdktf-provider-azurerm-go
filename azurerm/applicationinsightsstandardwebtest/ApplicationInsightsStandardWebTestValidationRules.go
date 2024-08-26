// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsstandardwebtest


type ApplicationInsightsStandardWebTestValidationRules struct {
	// content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/application_insights_standard_web_test#content ApplicationInsightsStandardWebTest#content}
	Content *ApplicationInsightsStandardWebTestValidationRulesContent `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/application_insights_standard_web_test#expected_status_code ApplicationInsightsStandardWebTest#expected_status_code}.
	ExpectedStatusCode *float64 `field:"optional" json:"expectedStatusCode" yaml:"expectedStatusCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/application_insights_standard_web_test#ssl_cert_remaining_lifetime ApplicationInsightsStandardWebTest#ssl_cert_remaining_lifetime}.
	SslCertRemainingLifetime *float64 `field:"optional" json:"sslCertRemainingLifetime" yaml:"sslCertRemainingLifetime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/application_insights_standard_web_test#ssl_check_enabled ApplicationInsightsStandardWebTest#ssl_check_enabled}.
	SslCheckEnabled interface{} `field:"optional" json:"sslCheckEnabled" yaml:"sslCheckEnabled"`
}

