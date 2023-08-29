// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nginxdeployment


type NginxDeploymentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/nginx_deployment#create NginxDeployment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/nginx_deployment#delete NginxDeployment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/nginx_deployment#read NginxDeployment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/nginx_deployment#update NginxDeployment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

