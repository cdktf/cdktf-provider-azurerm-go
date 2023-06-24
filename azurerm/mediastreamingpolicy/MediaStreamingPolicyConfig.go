package mediastreamingpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaStreamingPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_streaming_policy#media_services_account_name MediaStreamingPolicy#media_services_account_name}.
	MediaServicesAccountName *string `field:"required" json:"mediaServicesAccountName" yaml:"mediaServicesAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_streaming_policy#name MediaStreamingPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_streaming_policy#resource_group_name MediaStreamingPolicy#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// common_encryption_cbcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_streaming_policy#common_encryption_cbcs MediaStreamingPolicy#common_encryption_cbcs}
	CommonEncryptionCbcs *MediaStreamingPolicyCommonEncryptionCbcs `field:"optional" json:"commonEncryptionCbcs" yaml:"commonEncryptionCbcs"`
	// common_encryption_cenc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_streaming_policy#common_encryption_cenc MediaStreamingPolicy#common_encryption_cenc}
	CommonEncryptionCenc *MediaStreamingPolicyCommonEncryptionCenc `field:"optional" json:"commonEncryptionCenc" yaml:"commonEncryptionCenc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_streaming_policy#default_content_key_policy_name MediaStreamingPolicy#default_content_key_policy_name}.
	DefaultContentKeyPolicyName *string `field:"optional" json:"defaultContentKeyPolicyName" yaml:"defaultContentKeyPolicyName"`
	// envelope_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_streaming_policy#envelope_encryption MediaStreamingPolicy#envelope_encryption}
	EnvelopeEncryption *MediaStreamingPolicyEnvelopeEncryption `field:"optional" json:"envelopeEncryption" yaml:"envelopeEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_streaming_policy#id MediaStreamingPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// no_encryption_enabled_protocols block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_streaming_policy#no_encryption_enabled_protocols MediaStreamingPolicy#no_encryption_enabled_protocols}
	NoEncryptionEnabledProtocols *MediaStreamingPolicyNoEncryptionEnabledProtocols `field:"optional" json:"noEncryptionEnabledProtocols" yaml:"noEncryptionEnabledProtocols"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_streaming_policy#timeouts MediaStreamingPolicy#timeouts}
	Timeouts *MediaStreamingPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

