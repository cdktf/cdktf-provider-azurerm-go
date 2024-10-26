// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarefhirservice


type HealthcareFhirServiceAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/healthcare_fhir_service#audience HealthcareFhirService#audience}.
	Audience *string `field:"required" json:"audience" yaml:"audience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/healthcare_fhir_service#authority HealthcareFhirService#authority}.
	Authority *string `field:"required" json:"authority" yaml:"authority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/healthcare_fhir_service#smart_proxy_enabled HealthcareFhirService#smart_proxy_enabled}.
	SmartProxyEnabled interface{} `field:"optional" json:"smartProxyEnabled" yaml:"smartProxyEnabled"`
}

