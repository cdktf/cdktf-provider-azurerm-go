// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthbasecluster


type HdinsightHbaseClusterSecurityProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/hdinsight_hbase_cluster#aadds_resource_id HdinsightHbaseCluster#aadds_resource_id}.
	AaddsResourceId *string `field:"required" json:"aaddsResourceId" yaml:"aaddsResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/hdinsight_hbase_cluster#domain_name HdinsightHbaseCluster#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/hdinsight_hbase_cluster#domain_username HdinsightHbaseCluster#domain_username}.
	DomainUsername *string `field:"required" json:"domainUsername" yaml:"domainUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/hdinsight_hbase_cluster#domain_user_password HdinsightHbaseCluster#domain_user_password}.
	DomainUserPassword *string `field:"required" json:"domainUserPassword" yaml:"domainUserPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/hdinsight_hbase_cluster#ldaps_urls HdinsightHbaseCluster#ldaps_urls}.
	LdapsUrls *[]*string `field:"required" json:"ldapsUrls" yaml:"ldapsUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/hdinsight_hbase_cluster#msi_resource_id HdinsightHbaseCluster#msi_resource_id}.
	MsiResourceId *string `field:"required" json:"msiResourceId" yaml:"msiResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/hdinsight_hbase_cluster#cluster_users_group_dns HdinsightHbaseCluster#cluster_users_group_dns}.
	ClusterUsersGroupDns *[]*string `field:"optional" json:"clusterUsersGroupDns" yaml:"clusterUsersGroupDns"`
}

