// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/hdinsight_interactive_query_cluster#create HdinsightInteractiveQueryCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/hdinsight_interactive_query_cluster#delete HdinsightInteractiveQueryCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/hdinsight_interactive_query_cluster#read HdinsightInteractiveQueryCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/hdinsight_interactive_query_cluster#update HdinsightInteractiveQueryCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

