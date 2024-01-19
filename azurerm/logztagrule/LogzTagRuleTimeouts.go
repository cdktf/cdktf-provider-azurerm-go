// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logztagrule


type LogzTagRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/logz_tag_rule#create LogzTagRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/logz_tag_rule#delete LogzTagRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/logz_tag_rule#read LogzTagRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/logz_tag_rule#update LogzTagRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

