// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package trafficmanagerprofile


type TrafficManagerProfileDnsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/traffic_manager_profile#relative_name TrafficManagerProfile#relative_name}.
	RelativeName *string `field:"required" json:"relativeName" yaml:"relativeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/traffic_manager_profile#ttl TrafficManagerProfile#ttl}.
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
}

