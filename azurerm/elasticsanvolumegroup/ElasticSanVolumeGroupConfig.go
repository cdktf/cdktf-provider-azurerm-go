// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsanvolumegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticSanVolumeGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/elastic_san_volume_group#elastic_san_id ElasticSanVolumeGroup#elastic_san_id}.
	ElasticSanId *string `field:"required" json:"elasticSanId" yaml:"elasticSanId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/elastic_san_volume_group#name ElasticSanVolumeGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/elastic_san_volume_group#encryption ElasticSanVolumeGroup#encryption}
	Encryption *ElasticSanVolumeGroupEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/elastic_san_volume_group#encryption_type ElasticSanVolumeGroup#encryption_type}.
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/elastic_san_volume_group#id ElasticSanVolumeGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/elastic_san_volume_group#identity ElasticSanVolumeGroup#identity}
	Identity *ElasticSanVolumeGroupIdentity `field:"optional" json:"identity" yaml:"identity"`
	// network_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/elastic_san_volume_group#network_rule ElasticSanVolumeGroup#network_rule}
	NetworkRule interface{} `field:"optional" json:"networkRule" yaml:"networkRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/elastic_san_volume_group#protocol_type ElasticSanVolumeGroup#protocol_type}.
	ProtocolType *string `field:"optional" json:"protocolType" yaml:"protocolType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/elastic_san_volume_group#timeouts ElasticSanVolumeGroup#timeouts}
	Timeouts *ElasticSanVolumeGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

