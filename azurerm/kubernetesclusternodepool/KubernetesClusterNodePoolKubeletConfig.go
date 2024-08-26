// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesclusternodepool


type KubernetesClusterNodePoolKubeletConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/kubernetes_cluster_node_pool#allowed_unsafe_sysctls KubernetesClusterNodePool#allowed_unsafe_sysctls}.
	AllowedUnsafeSysctls *[]*string `field:"optional" json:"allowedUnsafeSysctls" yaml:"allowedUnsafeSysctls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/kubernetes_cluster_node_pool#container_log_max_line KubernetesClusterNodePool#container_log_max_line}.
	ContainerLogMaxLine *float64 `field:"optional" json:"containerLogMaxLine" yaml:"containerLogMaxLine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/kubernetes_cluster_node_pool#container_log_max_size_mb KubernetesClusterNodePool#container_log_max_size_mb}.
	ContainerLogMaxSizeMb *float64 `field:"optional" json:"containerLogMaxSizeMb" yaml:"containerLogMaxSizeMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/kubernetes_cluster_node_pool#cpu_cfs_quota_enabled KubernetesClusterNodePool#cpu_cfs_quota_enabled}.
	CpuCfsQuotaEnabled interface{} `field:"optional" json:"cpuCfsQuotaEnabled" yaml:"cpuCfsQuotaEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/kubernetes_cluster_node_pool#cpu_cfs_quota_period KubernetesClusterNodePool#cpu_cfs_quota_period}.
	CpuCfsQuotaPeriod *string `field:"optional" json:"cpuCfsQuotaPeriod" yaml:"cpuCfsQuotaPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/kubernetes_cluster_node_pool#cpu_manager_policy KubernetesClusterNodePool#cpu_manager_policy}.
	CpuManagerPolicy *string `field:"optional" json:"cpuManagerPolicy" yaml:"cpuManagerPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/kubernetes_cluster_node_pool#image_gc_high_threshold KubernetesClusterNodePool#image_gc_high_threshold}.
	ImageGcHighThreshold *float64 `field:"optional" json:"imageGcHighThreshold" yaml:"imageGcHighThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/kubernetes_cluster_node_pool#image_gc_low_threshold KubernetesClusterNodePool#image_gc_low_threshold}.
	ImageGcLowThreshold *float64 `field:"optional" json:"imageGcLowThreshold" yaml:"imageGcLowThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/kubernetes_cluster_node_pool#pod_max_pid KubernetesClusterNodePool#pod_max_pid}.
	PodMaxPid *float64 `field:"optional" json:"podMaxPid" yaml:"podMaxPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/kubernetes_cluster_node_pool#topology_manager_policy KubernetesClusterNodePool#topology_manager_policy}.
	TopologyManagerPolicy *string `field:"optional" json:"topologyManagerPolicy" yaml:"topologyManagerPolicy"`
}

