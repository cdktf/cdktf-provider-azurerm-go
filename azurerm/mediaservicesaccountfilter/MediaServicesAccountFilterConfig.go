// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediaservicesaccountfilter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaServicesAccountFilterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/media_services_account_filter#media_services_account_name MediaServicesAccountFilter#media_services_account_name}.
	MediaServicesAccountName *string `field:"required" json:"mediaServicesAccountName" yaml:"mediaServicesAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/media_services_account_filter#name MediaServicesAccountFilter#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/media_services_account_filter#resource_group_name MediaServicesAccountFilter#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/media_services_account_filter#first_quality_bitrate MediaServicesAccountFilter#first_quality_bitrate}.
	FirstQualityBitrate *float64 `field:"optional" json:"firstQualityBitrate" yaml:"firstQualityBitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/media_services_account_filter#id MediaServicesAccountFilter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// presentation_time_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/media_services_account_filter#presentation_time_range MediaServicesAccountFilter#presentation_time_range}
	PresentationTimeRange *MediaServicesAccountFilterPresentationTimeRange `field:"optional" json:"presentationTimeRange" yaml:"presentationTimeRange"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/media_services_account_filter#timeouts MediaServicesAccountFilter#timeouts}
	Timeouts *MediaServicesAccountFilterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// track_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/media_services_account_filter#track_selection MediaServicesAccountFilter#track_selection}
	TrackSelection interface{} `field:"optional" json:"trackSelection" yaml:"trackSelection"`
}

