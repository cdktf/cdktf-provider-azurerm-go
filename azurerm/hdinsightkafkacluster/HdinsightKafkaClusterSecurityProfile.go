// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightkafkacluster


type HdinsightKafkaClusterSecurityProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#aadds_resource_id HdinsightKafkaCluster#aadds_resource_id}.
	AaddsResourceId *string `field:"required" json:"aaddsResourceId" yaml:"aaddsResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#domain_name HdinsightKafkaCluster#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#domain_username HdinsightKafkaCluster#domain_username}.
	DomainUsername *string `field:"required" json:"domainUsername" yaml:"domainUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#domain_user_password HdinsightKafkaCluster#domain_user_password}.
	DomainUserPassword *string `field:"required" json:"domainUserPassword" yaml:"domainUserPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#ldaps_urls HdinsightKafkaCluster#ldaps_urls}.
	LdapsUrls *[]*string `field:"required" json:"ldapsUrls" yaml:"ldapsUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#msi_resource_id HdinsightKafkaCluster#msi_resource_id}.
	MsiResourceId *string `field:"required" json:"msiResourceId" yaml:"msiResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#cluster_users_group_dns HdinsightKafkaCluster#cluster_users_group_dns}.
	ClusterUsersGroupDns *[]*string `field:"optional" json:"clusterUsersGroupDns" yaml:"clusterUsersGroupDns"`
}

