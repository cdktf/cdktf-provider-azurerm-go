// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#location CdnEndpoint#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#name CdnEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// origin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#origin CdnEndpoint#origin}
	Origin interface{} `field:"required" json:"origin" yaml:"origin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#profile_name CdnEndpoint#profile_name}.
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#resource_group_name CdnEndpoint#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#content_types_to_compress CdnEndpoint#content_types_to_compress}.
	ContentTypesToCompress *[]*string `field:"optional" json:"contentTypesToCompress" yaml:"contentTypesToCompress"`
	// delivery_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#delivery_rule CdnEndpoint#delivery_rule}
	DeliveryRule interface{} `field:"optional" json:"deliveryRule" yaml:"deliveryRule"`
	// geo_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#geo_filter CdnEndpoint#geo_filter}
	GeoFilter interface{} `field:"optional" json:"geoFilter" yaml:"geoFilter"`
	// global_delivery_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#global_delivery_rule CdnEndpoint#global_delivery_rule}
	GlobalDeliveryRule *CdnEndpointGlobalDeliveryRule `field:"optional" json:"globalDeliveryRule" yaml:"globalDeliveryRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#id CdnEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#is_compression_enabled CdnEndpoint#is_compression_enabled}.
	IsCompressionEnabled interface{} `field:"optional" json:"isCompressionEnabled" yaml:"isCompressionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#is_http_allowed CdnEndpoint#is_http_allowed}.
	IsHttpAllowed interface{} `field:"optional" json:"isHttpAllowed" yaml:"isHttpAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#is_https_allowed CdnEndpoint#is_https_allowed}.
	IsHttpsAllowed interface{} `field:"optional" json:"isHttpsAllowed" yaml:"isHttpsAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#optimization_type CdnEndpoint#optimization_type}.
	OptimizationType *string `field:"optional" json:"optimizationType" yaml:"optimizationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#origin_host_header CdnEndpoint#origin_host_header}.
	OriginHostHeader *string `field:"optional" json:"originHostHeader" yaml:"originHostHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#origin_path CdnEndpoint#origin_path}.
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#probe_path CdnEndpoint#probe_path}.
	ProbePath *string `field:"optional" json:"probePath" yaml:"probePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#querystring_caching_behaviour CdnEndpoint#querystring_caching_behaviour}.
	QuerystringCachingBehaviour *string `field:"optional" json:"querystringCachingBehaviour" yaml:"querystringCachingBehaviour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#tags CdnEndpoint#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cdn_endpoint#timeouts CdnEndpoint#timeouts}
	Timeouts *CdnEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

