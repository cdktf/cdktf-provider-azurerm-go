// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppTemplateCustomScaleRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/container_app#custom_rule_type ContainerApp#custom_rule_type}.
	CustomRuleType *string `field:"required" json:"customRuleType" yaml:"customRuleType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/container_app#metadata ContainerApp#metadata}.
	Metadata *map[string]*string `field:"required" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/container_app#name ContainerApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/container_app#authentication ContainerApp#authentication}
	Authentication interface{} `field:"optional" json:"authentication" yaml:"authentication"`
}

