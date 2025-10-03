// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsanvolume


type ElasticSanVolumeCreateSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/elastic_san_volume#source_id ElasticSanVolume#source_id}.
	SourceId *string `field:"required" json:"sourceId" yaml:"sourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/elastic_san_volume#source_type ElasticSanVolume#source_type}.
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
}

