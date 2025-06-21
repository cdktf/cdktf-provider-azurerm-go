// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitiveaccountraipolicy


type CognitiveAccountRaiPolicyContentFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/cognitive_account_rai_policy#block_enabled CognitiveAccountRaiPolicy#block_enabled}.
	BlockEnabled interface{} `field:"required" json:"blockEnabled" yaml:"blockEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/cognitive_account_rai_policy#filter_enabled CognitiveAccountRaiPolicy#filter_enabled}.
	FilterEnabled interface{} `field:"required" json:"filterEnabled" yaml:"filterEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/cognitive_account_rai_policy#name CognitiveAccountRaiPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/cognitive_account_rai_policy#severity_threshold CognitiveAccountRaiPolicy#severity_threshold}.
	SeverityThreshold *string `field:"required" json:"severityThreshold" yaml:"severityThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/cognitive_account_rai_policy#source CognitiveAccountRaiPolicy#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
}

