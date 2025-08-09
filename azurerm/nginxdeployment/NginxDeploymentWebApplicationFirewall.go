// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nginxdeployment


type NginxDeploymentWebApplicationFirewall struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/nginx_deployment#activation_state_enabled NginxDeployment#activation_state_enabled}.
	ActivationStateEnabled interface{} `field:"required" json:"activationStateEnabled" yaml:"activationStateEnabled"`
}

