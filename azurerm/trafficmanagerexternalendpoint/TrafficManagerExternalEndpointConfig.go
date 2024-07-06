// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package trafficmanagerexternalendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TrafficManagerExternalEndpointConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/traffic_manager_external_endpoint#name TrafficManagerExternalEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/traffic_manager_external_endpoint#profile_id TrafficManagerExternalEndpoint#profile_id}.
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/traffic_manager_external_endpoint#target TrafficManagerExternalEndpoint#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/traffic_manager_external_endpoint#always_serve_enabled TrafficManagerExternalEndpoint#always_serve_enabled}.
	AlwaysServeEnabled interface{} `field:"optional" json:"alwaysServeEnabled" yaml:"alwaysServeEnabled"`
	// custom_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/traffic_manager_external_endpoint#custom_header TrafficManagerExternalEndpoint#custom_header}
	CustomHeader interface{} `field:"optional" json:"customHeader" yaml:"customHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/traffic_manager_external_endpoint#enabled TrafficManagerExternalEndpoint#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/traffic_manager_external_endpoint#endpoint_location TrafficManagerExternalEndpoint#endpoint_location}.
	EndpointLocation *string `field:"optional" json:"endpointLocation" yaml:"endpointLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/traffic_manager_external_endpoint#geo_mappings TrafficManagerExternalEndpoint#geo_mappings}.
	GeoMappings *[]*string `field:"optional" json:"geoMappings" yaml:"geoMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/traffic_manager_external_endpoint#id TrafficManagerExternalEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/traffic_manager_external_endpoint#priority TrafficManagerExternalEndpoint#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// subnet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/traffic_manager_external_endpoint#subnet TrafficManagerExternalEndpoint#subnet}
	Subnet interface{} `field:"optional" json:"subnet" yaml:"subnet"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/traffic_manager_external_endpoint#timeouts TrafficManagerExternalEndpoint#timeouts}
	Timeouts *TrafficManagerExternalEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/traffic_manager_external_endpoint#weight TrafficManagerExternalEndpoint#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

