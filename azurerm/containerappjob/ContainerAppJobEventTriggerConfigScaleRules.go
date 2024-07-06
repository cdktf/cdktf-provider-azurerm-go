// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob


type ContainerAppJobEventTriggerConfigScaleRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/container_app_job#custom_rule_type ContainerAppJob#custom_rule_type}.
	CustomRuleType *string `field:"required" json:"customRuleType" yaml:"customRuleType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/container_app_job#metadata ContainerAppJob#metadata}.
	Metadata *map[string]*string `field:"required" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/container_app_job#name ContainerAppJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/container_app_job#authentication ContainerAppJob#authentication}
	Authentication interface{} `field:"optional" json:"authentication" yaml:"authentication"`
}

