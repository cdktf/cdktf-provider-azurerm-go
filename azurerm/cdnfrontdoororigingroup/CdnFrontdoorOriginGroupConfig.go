// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoororigingroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnFrontdoorOriginGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_origin_group#cdn_frontdoor_profile_id CdnFrontdoorOriginGroup#cdn_frontdoor_profile_id}.
	CdnFrontdoorProfileId *string `field:"required" json:"cdnFrontdoorProfileId" yaml:"cdnFrontdoorProfileId"`
	// load_balancing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_origin_group#load_balancing CdnFrontdoorOriginGroup#load_balancing}
	LoadBalancing *CdnFrontdoorOriginGroupLoadBalancing `field:"required" json:"loadBalancing" yaml:"loadBalancing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_origin_group#name CdnFrontdoorOriginGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// health_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_origin_group#health_probe CdnFrontdoorOriginGroup#health_probe}
	HealthProbe *CdnFrontdoorOriginGroupHealthProbe `field:"optional" json:"healthProbe" yaml:"healthProbe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_origin_group#id CdnFrontdoorOriginGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_origin_group#restore_traffic_time_to_healed_or_new_endpoint_in_minutes CdnFrontdoorOriginGroup#restore_traffic_time_to_healed_or_new_endpoint_in_minutes}.
	RestoreTrafficTimeToHealedOrNewEndpointInMinutes *float64 `field:"optional" json:"restoreTrafficTimeToHealedOrNewEndpointInMinutes" yaml:"restoreTrafficTimeToHealedOrNewEndpointInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_origin_group#session_affinity_enabled CdnFrontdoorOriginGroup#session_affinity_enabled}.
	SessionAffinityEnabled interface{} `field:"optional" json:"sessionAffinityEnabled" yaml:"sessionAffinityEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_origin_group#timeouts CdnFrontdoorOriginGroup#timeouts}
	Timeouts *CdnFrontdoorOriginGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

