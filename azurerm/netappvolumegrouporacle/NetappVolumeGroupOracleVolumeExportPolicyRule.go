// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolumegrouporacle


type NetappVolumeGroupOracleVolumeExportPolicyRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/netapp_volume_group_oracle#allowed_clients NetappVolumeGroupOracle#allowed_clients}.
	AllowedClients *string `field:"required" json:"allowedClients" yaml:"allowedClients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/netapp_volume_group_oracle#nfsv3_enabled NetappVolumeGroupOracle#nfsv3_enabled}.
	Nfsv3Enabled interface{} `field:"required" json:"nfsv3Enabled" yaml:"nfsv3Enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/netapp_volume_group_oracle#nfsv41_enabled NetappVolumeGroupOracle#nfsv41_enabled}.
	Nfsv41Enabled interface{} `field:"required" json:"nfsv41Enabled" yaml:"nfsv41Enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/netapp_volume_group_oracle#rule_index NetappVolumeGroupOracle#rule_index}.
	RuleIndex *float64 `field:"required" json:"ruleIndex" yaml:"ruleIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/netapp_volume_group_oracle#root_access_enabled NetappVolumeGroupOracle#root_access_enabled}.
	RootAccessEnabled interface{} `field:"optional" json:"rootAccessEnabled" yaml:"rootAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/netapp_volume_group_oracle#unix_read_only NetappVolumeGroupOracle#unix_read_only}.
	UnixReadOnly interface{} `field:"optional" json:"unixReadOnly" yaml:"unixReadOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/netapp_volume_group_oracle#unix_read_write NetappVolumeGroupOracle#unix_read_write}.
	UnixReadWrite interface{} `field:"optional" json:"unixReadWrite" yaml:"unixReadWrite"`
}

