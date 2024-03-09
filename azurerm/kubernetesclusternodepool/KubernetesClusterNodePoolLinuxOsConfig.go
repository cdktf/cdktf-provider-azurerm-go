// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesclusternodepool


type KubernetesClusterNodePoolLinuxOsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/kubernetes_cluster_node_pool#swap_file_size_mb KubernetesClusterNodePool#swap_file_size_mb}.
	SwapFileSizeMb *float64 `field:"optional" json:"swapFileSizeMb" yaml:"swapFileSizeMb"`
	// sysctl_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/kubernetes_cluster_node_pool#sysctl_config KubernetesClusterNodePool#sysctl_config}
	SysctlConfig *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig `field:"optional" json:"sysctlConfig" yaml:"sysctlConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/kubernetes_cluster_node_pool#transparent_huge_page_defrag KubernetesClusterNodePool#transparent_huge_page_defrag}.
	TransparentHugePageDefrag *string `field:"optional" json:"transparentHugePageDefrag" yaml:"transparentHugePageDefrag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/kubernetes_cluster_node_pool#transparent_huge_page_enabled KubernetesClusterNodePool#transparent_huge_page_enabled}.
	TransparentHugePageEnabled *string `field:"optional" json:"transparentHugePageEnabled" yaml:"transparentHugePageEnabled"`
}

