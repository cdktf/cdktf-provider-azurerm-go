// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/kubernetes_cluster#admin_group_object_ids KubernetesCluster#admin_group_object_ids}.
	AdminGroupObjectIds *[]*string `field:"optional" json:"adminGroupObjectIds" yaml:"adminGroupObjectIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/kubernetes_cluster#azure_rbac_enabled KubernetesCluster#azure_rbac_enabled}.
	AzureRbacEnabled interface{} `field:"optional" json:"azureRbacEnabled" yaml:"azureRbacEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/kubernetes_cluster#client_app_id KubernetesCluster#client_app_id}.
	ClientAppId *string `field:"optional" json:"clientAppId" yaml:"clientAppId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/kubernetes_cluster#managed KubernetesCluster#managed}.
	Managed interface{} `field:"optional" json:"managed" yaml:"managed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/kubernetes_cluster#server_app_id KubernetesCluster#server_app_id}.
	ServerAppId *string `field:"optional" json:"serverAppId" yaml:"serverAppId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/kubernetes_cluster#server_app_secret KubernetesCluster#server_app_secret}.
	ServerAppSecret *string `field:"optional" json:"serverAppSecret" yaml:"serverAppSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/kubernetes_cluster#tenant_id KubernetesCluster#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

