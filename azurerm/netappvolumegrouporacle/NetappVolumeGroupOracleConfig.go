// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolumegrouporacle

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappVolumeGroupOracleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/netapp_volume_group_oracle#account_name NetappVolumeGroupOracle#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/netapp_volume_group_oracle#application_identifier NetappVolumeGroupOracle#application_identifier}.
	ApplicationIdentifier *string `field:"required" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/netapp_volume_group_oracle#group_description NetappVolumeGroupOracle#group_description}.
	GroupDescription *string `field:"required" json:"groupDescription" yaml:"groupDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/netapp_volume_group_oracle#location NetappVolumeGroupOracle#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/netapp_volume_group_oracle#name NetappVolumeGroupOracle#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/netapp_volume_group_oracle#resource_group_name NetappVolumeGroupOracle#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/netapp_volume_group_oracle#volume NetappVolumeGroupOracle#volume}
	Volume interface{} `field:"required" json:"volume" yaml:"volume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/netapp_volume_group_oracle#id NetappVolumeGroupOracle#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/netapp_volume_group_oracle#timeouts NetappVolumeGroupOracle#timeouts}
	Timeouts *NetappVolumeGroupOracleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

