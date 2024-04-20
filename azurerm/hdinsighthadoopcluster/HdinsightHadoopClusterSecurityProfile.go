// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthadoopcluster


type HdinsightHadoopClusterSecurityProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/hdinsight_hadoop_cluster#aadds_resource_id HdinsightHadoopCluster#aadds_resource_id}.
	AaddsResourceId *string `field:"required" json:"aaddsResourceId" yaml:"aaddsResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/hdinsight_hadoop_cluster#domain_name HdinsightHadoopCluster#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/hdinsight_hadoop_cluster#domain_username HdinsightHadoopCluster#domain_username}.
	DomainUsername *string `field:"required" json:"domainUsername" yaml:"domainUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/hdinsight_hadoop_cluster#domain_user_password HdinsightHadoopCluster#domain_user_password}.
	DomainUserPassword *string `field:"required" json:"domainUserPassword" yaml:"domainUserPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/hdinsight_hadoop_cluster#ldaps_urls HdinsightHadoopCluster#ldaps_urls}.
	LdapsUrls *[]*string `field:"required" json:"ldapsUrls" yaml:"ldapsUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/hdinsight_hadoop_cluster#msi_resource_id HdinsightHadoopCluster#msi_resource_id}.
	MsiResourceId *string `field:"required" json:"msiResourceId" yaml:"msiResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/hdinsight_hadoop_cluster#cluster_users_group_dns HdinsightHadoopCluster#cluster_users_group_dns}.
	ClusterUsersGroupDns *[]*string `field:"optional" json:"clusterUsersGroupDns" yaml:"clusterUsersGroupDns"`
}

