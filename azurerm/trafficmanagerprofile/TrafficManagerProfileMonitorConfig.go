// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package trafficmanagerprofile


type TrafficManagerProfileMonitorConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/traffic_manager_profile#port TrafficManagerProfile#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/traffic_manager_profile#protocol TrafficManagerProfile#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// custom_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/traffic_manager_profile#custom_header TrafficManagerProfile#custom_header}
	CustomHeader interface{} `field:"optional" json:"customHeader" yaml:"customHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/traffic_manager_profile#expected_status_code_ranges TrafficManagerProfile#expected_status_code_ranges}.
	ExpectedStatusCodeRanges *[]*string `field:"optional" json:"expectedStatusCodeRanges" yaml:"expectedStatusCodeRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/traffic_manager_profile#interval_in_seconds TrafficManagerProfile#interval_in_seconds}.
	IntervalInSeconds *float64 `field:"optional" json:"intervalInSeconds" yaml:"intervalInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/traffic_manager_profile#path TrafficManagerProfile#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/traffic_manager_profile#timeout_in_seconds TrafficManagerProfile#timeout_in_seconds}.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/traffic_manager_profile#tolerated_number_of_failures TrafficManagerProfile#tolerated_number_of_failures}.
	ToleratedNumberOfFailures *float64 `field:"optional" json:"toleratedNumberOfFailures" yaml:"toleratedNumberOfFailures"`
}

