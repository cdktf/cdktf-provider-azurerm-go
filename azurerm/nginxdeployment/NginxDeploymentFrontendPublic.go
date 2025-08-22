// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nginxdeployment


type NginxDeploymentFrontendPublic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/nginx_deployment#ip_address NginxDeployment#ip_address}.
	IpAddress *[]*string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}

