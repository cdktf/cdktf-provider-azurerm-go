// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcareservice


type HealthcareServiceAuthenticationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/healthcare_service#audience HealthcareService#audience}.
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/healthcare_service#authority HealthcareService#authority}.
	Authority *string `field:"optional" json:"authority" yaml:"authority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/healthcare_service#smart_proxy_enabled HealthcareService#smart_proxy_enabled}.
	SmartProxyEnabled interface{} `field:"optional" json:"smartProxyEnabled" yaml:"smartProxyEnabled"`
}

