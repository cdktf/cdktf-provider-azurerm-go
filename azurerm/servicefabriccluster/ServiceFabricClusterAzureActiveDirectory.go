// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabriccluster


type ServiceFabricClusterAzureActiveDirectory struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/service_fabric_cluster#client_application_id ServiceFabricCluster#client_application_id}.
	ClientApplicationId *string `field:"required" json:"clientApplicationId" yaml:"clientApplicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/service_fabric_cluster#cluster_application_id ServiceFabricCluster#cluster_application_id}.
	ClusterApplicationId *string `field:"required" json:"clusterApplicationId" yaml:"clusterApplicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/service_fabric_cluster#tenant_id ServiceFabricCluster#tenant_id}.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

