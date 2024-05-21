// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppIngress struct {
	// The target port on the container for the Ingress traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/container_app#target_port ContainerApp#target_port}
	TargetPort *float64 `field:"required" json:"targetPort" yaml:"targetPort"`
	// traffic_weight block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/container_app#traffic_weight ContainerApp#traffic_weight}
	TrafficWeight interface{} `field:"required" json:"trafficWeight" yaml:"trafficWeight"`
	// Should this ingress allow insecure connections?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/container_app#allow_insecure_connections ContainerApp#allow_insecure_connections}
	AllowInsecureConnections interface{} `field:"optional" json:"allowInsecureConnections" yaml:"allowInsecureConnections"`
	// custom_domain block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/container_app#custom_domain ContainerApp#custom_domain}
	CustomDomain *ContainerAppIngressCustomDomain `field:"optional" json:"customDomain" yaml:"customDomain"`
	// The exposed port on the container for the Ingress traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/container_app#exposed_port ContainerApp#exposed_port}
	ExposedPort *float64 `field:"optional" json:"exposedPort" yaml:"exposedPort"`
	// Is this an external Ingress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/container_app#external_enabled ContainerApp#external_enabled}
	ExternalEnabled interface{} `field:"optional" json:"externalEnabled" yaml:"externalEnabled"`
	// ip_security_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/container_app#ip_security_restriction ContainerApp#ip_security_restriction}
	IpSecurityRestriction interface{} `field:"optional" json:"ipSecurityRestriction" yaml:"ipSecurityRestriction"`
	// The transport method for the Ingress. Possible values include `auto`, `http`, and `http2`, `tcp`. Defaults to `auto`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/container_app#transport ContainerApp#transport}
	Transport *string `field:"optional" json:"transport" yaml:"transport"`
}

