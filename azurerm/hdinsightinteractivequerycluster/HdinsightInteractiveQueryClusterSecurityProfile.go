// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterSecurityProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/hdinsight_interactive_query_cluster#aadds_resource_id HdinsightInteractiveQueryCluster#aadds_resource_id}.
	AaddsResourceId *string `field:"required" json:"aaddsResourceId" yaml:"aaddsResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/hdinsight_interactive_query_cluster#domain_name HdinsightInteractiveQueryCluster#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/hdinsight_interactive_query_cluster#domain_username HdinsightInteractiveQueryCluster#domain_username}.
	DomainUsername *string `field:"required" json:"domainUsername" yaml:"domainUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/hdinsight_interactive_query_cluster#domain_user_password HdinsightInteractiveQueryCluster#domain_user_password}.
	DomainUserPassword *string `field:"required" json:"domainUserPassword" yaml:"domainUserPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/hdinsight_interactive_query_cluster#ldaps_urls HdinsightInteractiveQueryCluster#ldaps_urls}.
	LdapsUrls *[]*string `field:"required" json:"ldapsUrls" yaml:"ldapsUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/hdinsight_interactive_query_cluster#msi_resource_id HdinsightInteractiveQueryCluster#msi_resource_id}.
	MsiResourceId *string `field:"required" json:"msiResourceId" yaml:"msiResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/hdinsight_interactive_query_cluster#cluster_users_group_dns HdinsightInteractiveQueryCluster#cluster_users_group_dns}.
	ClusterUsersGroupDns *[]*string `field:"optional" json:"clusterUsersGroupDns" yaml:"clusterUsersGroupDns"`
}

