// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob


type ContainerAppJobEventTriggerConfigScaleRulesAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/container_app_job#secret_name ContainerAppJob#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/container_app_job#trigger_parameter ContainerAppJob#trigger_parameter}.
	TriggerParameter *string `field:"required" json:"triggerParameter" yaml:"triggerParameter"`
}

