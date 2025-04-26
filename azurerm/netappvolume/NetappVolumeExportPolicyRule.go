// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume


type NetappVolumeExportPolicyRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/netapp_volume#allowed_clients NetappVolume#allowed_clients}.
	AllowedClients *[]*string `field:"required" json:"allowedClients" yaml:"allowedClients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/netapp_volume#rule_index NetappVolume#rule_index}.
	RuleIndex *float64 `field:"required" json:"ruleIndex" yaml:"ruleIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/netapp_volume#kerberos_5i_read_only_enabled NetappVolume#kerberos_5i_read_only_enabled}.
	Kerberos5IReadOnlyEnabled interface{} `field:"optional" json:"kerberos5IReadOnlyEnabled" yaml:"kerberos5IReadOnlyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/netapp_volume#kerberos_5i_read_write_enabled NetappVolume#kerberos_5i_read_write_enabled}.
	Kerberos5IReadWriteEnabled interface{} `field:"optional" json:"kerberos5IReadWriteEnabled" yaml:"kerberos5IReadWriteEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/netapp_volume#kerberos_5p_read_only_enabled NetappVolume#kerberos_5p_read_only_enabled}.
	Kerberos5PReadOnlyEnabled interface{} `field:"optional" json:"kerberos5PReadOnlyEnabled" yaml:"kerberos5PReadOnlyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/netapp_volume#kerberos_5p_read_write_enabled NetappVolume#kerberos_5p_read_write_enabled}.
	Kerberos5PReadWriteEnabled interface{} `field:"optional" json:"kerberos5PReadWriteEnabled" yaml:"kerberos5PReadWriteEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/netapp_volume#kerberos_5_read_only_enabled NetappVolume#kerberos_5_read_only_enabled}.
	Kerberos5ReadOnlyEnabled interface{} `field:"optional" json:"kerberos5ReadOnlyEnabled" yaml:"kerberos5ReadOnlyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/netapp_volume#kerberos_5_read_write_enabled NetappVolume#kerberos_5_read_write_enabled}.
	Kerberos5ReadWriteEnabled interface{} `field:"optional" json:"kerberos5ReadWriteEnabled" yaml:"kerberos5ReadWriteEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/netapp_volume#protocols_enabled NetappVolume#protocols_enabled}.
	ProtocolsEnabled *[]*string `field:"optional" json:"protocolsEnabled" yaml:"protocolsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/netapp_volume#root_access_enabled NetappVolume#root_access_enabled}.
	RootAccessEnabled interface{} `field:"optional" json:"rootAccessEnabled" yaml:"rootAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/netapp_volume#unix_read_only NetappVolume#unix_read_only}.
	UnixReadOnly interface{} `field:"optional" json:"unixReadOnly" yaml:"unixReadOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/netapp_volume#unix_read_write NetappVolume#unix_read_write}.
	UnixReadWrite interface{} `field:"optional" json:"unixReadWrite" yaml:"unixReadWrite"`
}

