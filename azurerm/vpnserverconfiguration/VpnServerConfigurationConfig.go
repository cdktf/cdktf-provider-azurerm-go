// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnserverconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnServerConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/vpn_server_configuration#location VpnServerConfiguration#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/vpn_server_configuration#name VpnServerConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/vpn_server_configuration#resource_group_name VpnServerConfiguration#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/vpn_server_configuration#vpn_authentication_types VpnServerConfiguration#vpn_authentication_types}.
	VpnAuthenticationTypes *[]*string `field:"required" json:"vpnAuthenticationTypes" yaml:"vpnAuthenticationTypes"`
	// azure_active_directory_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/vpn_server_configuration#azure_active_directory_authentication VpnServerConfiguration#azure_active_directory_authentication}
	AzureActiveDirectoryAuthentication interface{} `field:"optional" json:"azureActiveDirectoryAuthentication" yaml:"azureActiveDirectoryAuthentication"`
	// client_revoked_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/vpn_server_configuration#client_revoked_certificate VpnServerConfiguration#client_revoked_certificate}
	ClientRevokedCertificate interface{} `field:"optional" json:"clientRevokedCertificate" yaml:"clientRevokedCertificate"`
	// client_root_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/vpn_server_configuration#client_root_certificate VpnServerConfiguration#client_root_certificate}
	ClientRootCertificate interface{} `field:"optional" json:"clientRootCertificate" yaml:"clientRootCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/vpn_server_configuration#id VpnServerConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ipsec_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/vpn_server_configuration#ipsec_policy VpnServerConfiguration#ipsec_policy}
	IpsecPolicy *VpnServerConfigurationIpsecPolicy `field:"optional" json:"ipsecPolicy" yaml:"ipsecPolicy"`
	// radius block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/vpn_server_configuration#radius VpnServerConfiguration#radius}
	Radius *VpnServerConfigurationRadius `field:"optional" json:"radius" yaml:"radius"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/vpn_server_configuration#tags VpnServerConfiguration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/vpn_server_configuration#timeouts VpnServerConfiguration#timeouts}
	Timeouts *VpnServerConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/vpn_server_configuration#vpn_protocols VpnServerConfiguration#vpn_protocols}.
	VpnProtocols *[]*string `field:"optional" json:"vpnProtocols" yaml:"vpnProtocols"`
}

