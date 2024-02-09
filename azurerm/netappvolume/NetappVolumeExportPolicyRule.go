// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume


type NetappVolumeExportPolicyRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/netapp_volume#allowed_clients NetappVolume#allowed_clients}.
	AllowedClients *[]*string `field:"required" json:"allowedClients" yaml:"allowedClients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/netapp_volume#rule_index NetappVolume#rule_index}.
	RuleIndex *float64 `field:"required" json:"ruleIndex" yaml:"ruleIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/netapp_volume#protocols_enabled NetappVolume#protocols_enabled}.
	ProtocolsEnabled *[]*string `field:"optional" json:"protocolsEnabled" yaml:"protocolsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/netapp_volume#root_access_enabled NetappVolume#root_access_enabled}.
	RootAccessEnabled interface{} `field:"optional" json:"rootAccessEnabled" yaml:"rootAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/netapp_volume#unix_read_only NetappVolume#unix_read_only}.
	UnixReadOnly interface{} `field:"optional" json:"unixReadOnly" yaml:"unixReadOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/netapp_volume#unix_read_write NetappVolume#unix_read_write}.
	UnixReadWrite interface{} `field:"optional" json:"unixReadWrite" yaml:"unixReadWrite"`
}

