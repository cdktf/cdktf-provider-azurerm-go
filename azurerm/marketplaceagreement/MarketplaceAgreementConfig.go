// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package marketplaceagreement

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MarketplaceAgreementConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/marketplace_agreement#offer MarketplaceAgreement#offer}.
	Offer *string `field:"required" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/marketplace_agreement#plan MarketplaceAgreement#plan}.
	Plan *string `field:"required" json:"plan" yaml:"plan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/marketplace_agreement#publisher MarketplaceAgreement#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/marketplace_agreement#id MarketplaceAgreement#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/marketplace_agreement#timeouts MarketplaceAgreement#timeouts}
	Timeouts *MarketplaceAgreementTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

