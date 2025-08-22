// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboardgrafana


type DashboardGrafanaSmtp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/dashboard_grafana#from_address DashboardGrafana#from_address}.
	FromAddress *string `field:"required" json:"fromAddress" yaml:"fromAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/dashboard_grafana#host DashboardGrafana#host}.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/dashboard_grafana#password DashboardGrafana#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/dashboard_grafana#start_tls_policy DashboardGrafana#start_tls_policy}.
	StartTlsPolicy *string `field:"required" json:"startTlsPolicy" yaml:"startTlsPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/dashboard_grafana#user DashboardGrafana#user}.
	User *string `field:"required" json:"user" yaml:"user"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/dashboard_grafana#enabled DashboardGrafana#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/dashboard_grafana#from_name DashboardGrafana#from_name}.
	FromName *string `field:"optional" json:"fromName" yaml:"fromName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/dashboard_grafana#verification_skip_enabled DashboardGrafana#verification_skip_enabled}.
	VerificationSkipEnabled interface{} `field:"optional" json:"verificationSkipEnabled" yaml:"verificationSkipEnabled"`
}

