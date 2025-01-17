// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitiveaccountraiblocklist


type CognitiveAccountRaiBlocklistTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/cognitive_account_rai_blocklist#create CognitiveAccountRaiBlocklist#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/cognitive_account_rai_blocklist#delete CognitiveAccountRaiBlocklist#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/cognitive_account_rai_blocklist#read CognitiveAccountRaiBlocklist#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/cognitive_account_rai_blocklist#update CognitiveAccountRaiBlocklist#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

