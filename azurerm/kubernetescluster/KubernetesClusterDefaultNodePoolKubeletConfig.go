// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterDefaultNodePoolKubeletConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/kubernetes_cluster#allowed_unsafe_sysctls KubernetesCluster#allowed_unsafe_sysctls}.
	AllowedUnsafeSysctls *[]*string `field:"optional" json:"allowedUnsafeSysctls" yaml:"allowedUnsafeSysctls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/kubernetes_cluster#container_log_max_line KubernetesCluster#container_log_max_line}.
	ContainerLogMaxLine *float64 `field:"optional" json:"containerLogMaxLine" yaml:"containerLogMaxLine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/kubernetes_cluster#container_log_max_size_mb KubernetesCluster#container_log_max_size_mb}.
	ContainerLogMaxSizeMb *float64 `field:"optional" json:"containerLogMaxSizeMb" yaml:"containerLogMaxSizeMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/kubernetes_cluster#cpu_cfs_quota_enabled KubernetesCluster#cpu_cfs_quota_enabled}.
	CpuCfsQuotaEnabled interface{} `field:"optional" json:"cpuCfsQuotaEnabled" yaml:"cpuCfsQuotaEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/kubernetes_cluster#cpu_cfs_quota_period KubernetesCluster#cpu_cfs_quota_period}.
	CpuCfsQuotaPeriod *string `field:"optional" json:"cpuCfsQuotaPeriod" yaml:"cpuCfsQuotaPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/kubernetes_cluster#cpu_manager_policy KubernetesCluster#cpu_manager_policy}.
	CpuManagerPolicy *string `field:"optional" json:"cpuManagerPolicy" yaml:"cpuManagerPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/kubernetes_cluster#image_gc_high_threshold KubernetesCluster#image_gc_high_threshold}.
	ImageGcHighThreshold *float64 `field:"optional" json:"imageGcHighThreshold" yaml:"imageGcHighThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/kubernetes_cluster#image_gc_low_threshold KubernetesCluster#image_gc_low_threshold}.
	ImageGcLowThreshold *float64 `field:"optional" json:"imageGcLowThreshold" yaml:"imageGcLowThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/kubernetes_cluster#pod_max_pid KubernetesCluster#pod_max_pid}.
	PodMaxPid *float64 `field:"optional" json:"podMaxPid" yaml:"podMaxPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/kubernetes_cluster#topology_manager_policy KubernetesCluster#topology_manager_policy}.
	TopologyManagerPolicy *string `field:"optional" json:"topologyManagerPolicy" yaml:"topologyManagerPolicy"`
}

