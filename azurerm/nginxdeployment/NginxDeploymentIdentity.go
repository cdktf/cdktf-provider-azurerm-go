// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nginxdeployment


type NginxDeploymentIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/nginx_deployment#type NginxDeployment#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/nginx_deployment#identity_ids NginxDeployment#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

