// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationloadbalancerfrontend


type ApplicationLoadBalancerFrontendTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/application_load_balancer_frontend#create ApplicationLoadBalancerFrontend#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/application_load_balancer_frontend#delete ApplicationLoadBalancerFrontend#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/application_load_balancer_frontend#read ApplicationLoadBalancerFrontend#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/application_load_balancer_frontend#update ApplicationLoadBalancerFrontend#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

