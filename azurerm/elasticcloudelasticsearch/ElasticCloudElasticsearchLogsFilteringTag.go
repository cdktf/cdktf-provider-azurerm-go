// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticcloudelasticsearch


type ElasticCloudElasticsearchLogsFilteringTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/elastic_cloud_elasticsearch#action ElasticCloudElasticsearch#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/elastic_cloud_elasticsearch#name ElasticCloudElasticsearch#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/elastic_cloud_elasticsearch#value ElasticCloudElasticsearch#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

