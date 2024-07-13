// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesclusternodepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterNodePoolConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#kubernetes_cluster_id KubernetesClusterNodePool#kubernetes_cluster_id}.
	KubernetesClusterId *string `field:"required" json:"kubernetesClusterId" yaml:"kubernetesClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#name KubernetesClusterNodePool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#vm_size KubernetesClusterNodePool#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#capacity_reservation_group_id KubernetesClusterNodePool#capacity_reservation_group_id}.
	CapacityReservationGroupId *string `field:"optional" json:"capacityReservationGroupId" yaml:"capacityReservationGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#custom_ca_trust_enabled KubernetesClusterNodePool#custom_ca_trust_enabled}.
	CustomCaTrustEnabled interface{} `field:"optional" json:"customCaTrustEnabled" yaml:"customCaTrustEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#enable_auto_scaling KubernetesClusterNodePool#enable_auto_scaling}.
	EnableAutoScaling interface{} `field:"optional" json:"enableAutoScaling" yaml:"enableAutoScaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#enable_host_encryption KubernetesClusterNodePool#enable_host_encryption}.
	EnableHostEncryption interface{} `field:"optional" json:"enableHostEncryption" yaml:"enableHostEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#enable_node_public_ip KubernetesClusterNodePool#enable_node_public_ip}.
	EnableNodePublicIp interface{} `field:"optional" json:"enableNodePublicIp" yaml:"enableNodePublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#eviction_policy KubernetesClusterNodePool#eviction_policy}.
	EvictionPolicy *string `field:"optional" json:"evictionPolicy" yaml:"evictionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#fips_enabled KubernetesClusterNodePool#fips_enabled}.
	FipsEnabled interface{} `field:"optional" json:"fipsEnabled" yaml:"fipsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#gpu_instance KubernetesClusterNodePool#gpu_instance}.
	GpuInstance *string `field:"optional" json:"gpuInstance" yaml:"gpuInstance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#host_group_id KubernetesClusterNodePool#host_group_id}.
	HostGroupId *string `field:"optional" json:"hostGroupId" yaml:"hostGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#id KubernetesClusterNodePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kubelet_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#kubelet_config KubernetesClusterNodePool#kubelet_config}
	KubeletConfig *KubernetesClusterNodePoolKubeletConfig `field:"optional" json:"kubeletConfig" yaml:"kubeletConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#kubelet_disk_type KubernetesClusterNodePool#kubelet_disk_type}.
	KubeletDiskType *string `field:"optional" json:"kubeletDiskType" yaml:"kubeletDiskType"`
	// linux_os_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#linux_os_config KubernetesClusterNodePool#linux_os_config}
	LinuxOsConfig *KubernetesClusterNodePoolLinuxOsConfig `field:"optional" json:"linuxOsConfig" yaml:"linuxOsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#max_count KubernetesClusterNodePool#max_count}.
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#max_pods KubernetesClusterNodePool#max_pods}.
	MaxPods *float64 `field:"optional" json:"maxPods" yaml:"maxPods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#message_of_the_day KubernetesClusterNodePool#message_of_the_day}.
	MessageOfTheDay *string `field:"optional" json:"messageOfTheDay" yaml:"messageOfTheDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#min_count KubernetesClusterNodePool#min_count}.
	MinCount *float64 `field:"optional" json:"minCount" yaml:"minCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#mode KubernetesClusterNodePool#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#node_count KubernetesClusterNodePool#node_count}.
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#node_labels KubernetesClusterNodePool#node_labels}.
	NodeLabels *map[string]*string `field:"optional" json:"nodeLabels" yaml:"nodeLabels"`
	// node_network_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#node_network_profile KubernetesClusterNodePool#node_network_profile}
	NodeNetworkProfile *KubernetesClusterNodePoolNodeNetworkProfile `field:"optional" json:"nodeNetworkProfile" yaml:"nodeNetworkProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#node_public_ip_prefix_id KubernetesClusterNodePool#node_public_ip_prefix_id}.
	NodePublicIpPrefixId *string `field:"optional" json:"nodePublicIpPrefixId" yaml:"nodePublicIpPrefixId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#node_taints KubernetesClusterNodePool#node_taints}.
	NodeTaints *[]*string `field:"optional" json:"nodeTaints" yaml:"nodeTaints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#orchestrator_version KubernetesClusterNodePool#orchestrator_version}.
	OrchestratorVersion *string `field:"optional" json:"orchestratorVersion" yaml:"orchestratorVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#os_disk_size_gb KubernetesClusterNodePool#os_disk_size_gb}.
	OsDiskSizeGb *float64 `field:"optional" json:"osDiskSizeGb" yaml:"osDiskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#os_disk_type KubernetesClusterNodePool#os_disk_type}.
	OsDiskType *string `field:"optional" json:"osDiskType" yaml:"osDiskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#os_sku KubernetesClusterNodePool#os_sku}.
	OsSku *string `field:"optional" json:"osSku" yaml:"osSku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#os_type KubernetesClusterNodePool#os_type}.
	OsType *string `field:"optional" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#pod_subnet_id KubernetesClusterNodePool#pod_subnet_id}.
	PodSubnetId *string `field:"optional" json:"podSubnetId" yaml:"podSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#priority KubernetesClusterNodePool#priority}.
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#proximity_placement_group_id KubernetesClusterNodePool#proximity_placement_group_id}.
	ProximityPlacementGroupId *string `field:"optional" json:"proximityPlacementGroupId" yaml:"proximityPlacementGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#scale_down_mode KubernetesClusterNodePool#scale_down_mode}.
	ScaleDownMode *string `field:"optional" json:"scaleDownMode" yaml:"scaleDownMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#snapshot_id KubernetesClusterNodePool#snapshot_id}.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#spot_max_price KubernetesClusterNodePool#spot_max_price}.
	SpotMaxPrice *float64 `field:"optional" json:"spotMaxPrice" yaml:"spotMaxPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#tags KubernetesClusterNodePool#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#timeouts KubernetesClusterNodePool#timeouts}
	Timeouts *KubernetesClusterNodePoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#ultra_ssd_enabled KubernetesClusterNodePool#ultra_ssd_enabled}.
	UltraSsdEnabled interface{} `field:"optional" json:"ultraSsdEnabled" yaml:"ultraSsdEnabled"`
	// upgrade_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#upgrade_settings KubernetesClusterNodePool#upgrade_settings}
	UpgradeSettings *KubernetesClusterNodePoolUpgradeSettings `field:"optional" json:"upgradeSettings" yaml:"upgradeSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#vnet_subnet_id KubernetesClusterNodePool#vnet_subnet_id}.
	VnetSubnetId *string `field:"optional" json:"vnetSubnetId" yaml:"vnetSubnetId"`
	// windows_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#windows_profile KubernetesClusterNodePool#windows_profile}
	WindowsProfile *KubernetesClusterNodePoolWindowsProfile `field:"optional" json:"windowsProfile" yaml:"windowsProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#workload_runtime KubernetesClusterNodePool#workload_runtime}.
	WorkloadRuntime *string `field:"optional" json:"workloadRuntime" yaml:"workloadRuntime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/kubernetes_cluster_node_pool#zones KubernetesClusterNodePool#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

