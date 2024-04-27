// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterDefaultNodePoolLinuxOsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/kubernetes_cluster#swap_file_size_mb KubernetesCluster#swap_file_size_mb}.
	SwapFileSizeMb *float64 `field:"optional" json:"swapFileSizeMb" yaml:"swapFileSizeMb"`
	// sysctl_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/kubernetes_cluster#sysctl_config KubernetesCluster#sysctl_config}
	SysctlConfig *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig `field:"optional" json:"sysctlConfig" yaml:"sysctlConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/kubernetes_cluster#transparent_huge_page_defrag KubernetesCluster#transparent_huge_page_defrag}.
	TransparentHugePageDefrag *string `field:"optional" json:"transparentHugePageDefrag" yaml:"transparentHugePageDefrag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/kubernetes_cluster#transparent_huge_page_enabled KubernetesCluster#transparent_huge_page_enabled}.
	TransparentHugePageEnabled *string `field:"optional" json:"transparentHugePageEnabled" yaml:"transparentHugePageEnabled"`
}

