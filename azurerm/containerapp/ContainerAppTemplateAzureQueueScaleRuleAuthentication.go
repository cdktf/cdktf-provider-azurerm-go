// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppTemplateAzureQueueScaleRuleAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/container_app#secret_name ContainerApp#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/container_app#trigger_parameter ContainerApp#trigger_parameter}.
	TriggerParameter *string `field:"required" json:"triggerParameter" yaml:"triggerParameter"`
}

