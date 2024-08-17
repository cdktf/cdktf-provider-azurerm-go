// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewayHttpListener struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/application_gateway#frontend_ip_configuration_name ApplicationGateway#frontend_ip_configuration_name}.
	FrontendIpConfigurationName *string `field:"required" json:"frontendIpConfigurationName" yaml:"frontendIpConfigurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/application_gateway#frontend_port_name ApplicationGateway#frontend_port_name}.
	FrontendPortName *string `field:"required" json:"frontendPortName" yaml:"frontendPortName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/application_gateway#protocol ApplicationGateway#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// custom_error_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/application_gateway#custom_error_configuration ApplicationGateway#custom_error_configuration}
	CustomErrorConfiguration interface{} `field:"optional" json:"customErrorConfiguration" yaml:"customErrorConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/application_gateway#firewall_policy_id ApplicationGateway#firewall_policy_id}.
	FirewallPolicyId *string `field:"optional" json:"firewallPolicyId" yaml:"firewallPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/application_gateway#host_name ApplicationGateway#host_name}.
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/application_gateway#host_names ApplicationGateway#host_names}.
	HostNames *[]*string `field:"optional" json:"hostNames" yaml:"hostNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/application_gateway#require_sni ApplicationGateway#require_sni}.
	RequireSni interface{} `field:"optional" json:"requireSni" yaml:"requireSni"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/application_gateway#ssl_certificate_name ApplicationGateway#ssl_certificate_name}.
	SslCertificateName *string `field:"optional" json:"sslCertificateName" yaml:"sslCertificateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/application_gateway#ssl_profile_name ApplicationGateway#ssl_profile_name}.
	SslProfileName *string `field:"optional" json:"sslProfileName" yaml:"sslProfileName"`
}

