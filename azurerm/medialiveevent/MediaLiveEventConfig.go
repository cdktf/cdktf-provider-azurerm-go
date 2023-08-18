package medialiveevent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaLiveEventConfig struct {
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
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#input MediaLiveEvent#input}
	Input *MediaLiveEventInput `field:"required" json:"input" yaml:"input"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#location MediaLiveEvent#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#media_services_account_name MediaLiveEvent#media_services_account_name}.
	MediaServicesAccountName *string `field:"required" json:"mediaServicesAccountName" yaml:"mediaServicesAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#name MediaLiveEvent#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#resource_group_name MediaLiveEvent#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#auto_start_enabled MediaLiveEvent#auto_start_enabled}.
	AutoStartEnabled interface{} `field:"optional" json:"autoStartEnabled" yaml:"autoStartEnabled"`
	// cross_site_access_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#cross_site_access_policy MediaLiveEvent#cross_site_access_policy}
	CrossSiteAccessPolicy *MediaLiveEventCrossSiteAccessPolicy `field:"optional" json:"crossSiteAccessPolicy" yaml:"crossSiteAccessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#description MediaLiveEvent#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encoding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#encoding MediaLiveEvent#encoding}
	Encoding *MediaLiveEventEncoding `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#hostname_prefix MediaLiveEvent#hostname_prefix}.
	HostnamePrefix *string `field:"optional" json:"hostnamePrefix" yaml:"hostnamePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#id MediaLiveEvent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// preview block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#preview MediaLiveEvent#preview}
	Preview *MediaLiveEventPreview `field:"optional" json:"preview" yaml:"preview"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#stream_options MediaLiveEvent#stream_options}.
	StreamOptions *[]*string `field:"optional" json:"streamOptions" yaml:"streamOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#tags MediaLiveEvent#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#timeouts MediaLiveEvent#timeouts}
	Timeouts *MediaLiveEventTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#transcription_languages MediaLiveEvent#transcription_languages}.
	TranscriptionLanguages *[]*string `field:"optional" json:"transcriptionLanguages" yaml:"transcriptionLanguages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_live_event#use_static_hostname MediaLiveEvent#use_static_hostname}.
	UseStaticHostname interface{} `field:"optional" json:"useStaticHostname" yaml:"useStaticHostname"`
}

