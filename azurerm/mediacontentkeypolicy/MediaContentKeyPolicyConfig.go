package mediacontentkeypolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaContentKeyPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/media_content_key_policy#media_services_account_name MediaContentKeyPolicy#media_services_account_name}.
	MediaServicesAccountName *string `field:"required" json:"mediaServicesAccountName" yaml:"mediaServicesAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/media_content_key_policy#name MediaContentKeyPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// policy_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/media_content_key_policy#policy_option MediaContentKeyPolicy#policy_option}
	PolicyOption interface{} `field:"required" json:"policyOption" yaml:"policyOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/media_content_key_policy#resource_group_name MediaContentKeyPolicy#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/media_content_key_policy#description MediaContentKeyPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/media_content_key_policy#id MediaContentKeyPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/media_content_key_policy#timeouts MediaContentKeyPolicy#timeouts}
	Timeouts *MediaContentKeyPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

