// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelmetadata

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelMetadataConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#content_id SentinelMetadata#content_id}.
	ContentId *string `field:"required" json:"contentId" yaml:"contentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#kind SentinelMetadata#kind}.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#name SentinelMetadata#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#parent_id SentinelMetadata#parent_id}.
	ParentId *string `field:"required" json:"parentId" yaml:"parentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#workspace_id SentinelMetadata#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// author block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#author SentinelMetadata#author}
	Author *SentinelMetadataAuthor `field:"optional" json:"author" yaml:"author"`
	// category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#category SentinelMetadata#category}
	Category *SentinelMetadataCategory `field:"optional" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#content_schema_version SentinelMetadata#content_schema_version}.
	ContentSchemaVersion *string `field:"optional" json:"contentSchemaVersion" yaml:"contentSchemaVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#custom_version SentinelMetadata#custom_version}.
	CustomVersion *string `field:"optional" json:"customVersion" yaml:"customVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#dependency SentinelMetadata#dependency}.
	Dependency *string `field:"optional" json:"dependency" yaml:"dependency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#first_publish_date SentinelMetadata#first_publish_date}.
	FirstPublishDate *string `field:"optional" json:"firstPublishDate" yaml:"firstPublishDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#icon_id SentinelMetadata#icon_id}.
	IconId *string `field:"optional" json:"iconId" yaml:"iconId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#id SentinelMetadata#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#last_publish_date SentinelMetadata#last_publish_date}.
	LastPublishDate *string `field:"optional" json:"lastPublishDate" yaml:"lastPublishDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#preview_images SentinelMetadata#preview_images}.
	PreviewImages *[]*string `field:"optional" json:"previewImages" yaml:"previewImages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#preview_images_dark SentinelMetadata#preview_images_dark}.
	PreviewImagesDark *[]*string `field:"optional" json:"previewImagesDark" yaml:"previewImagesDark"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#providers SentinelMetadata#providers}.
	Providers *[]*string `field:"optional" json:"providers" yaml:"providers"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#source SentinelMetadata#source}
	Source *SentinelMetadataSource `field:"optional" json:"source" yaml:"source"`
	// support block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#support SentinelMetadata#support}
	Support *SentinelMetadataSupport `field:"optional" json:"support" yaml:"support"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#threat_analysis_tactics SentinelMetadata#threat_analysis_tactics}.
	ThreatAnalysisTactics *[]*string `field:"optional" json:"threatAnalysisTactics" yaml:"threatAnalysisTactics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#threat_analysis_techniques SentinelMetadata#threat_analysis_techniques}.
	ThreatAnalysisTechniques *[]*string `field:"optional" json:"threatAnalysisTechniques" yaml:"threatAnalysisTechniques"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#timeouts SentinelMetadata#timeouts}
	Timeouts *SentinelMetadataTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_metadata#version SentinelMetadata#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

