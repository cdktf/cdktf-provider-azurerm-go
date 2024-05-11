// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterDefaultNodePool struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#name KubernetesCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#vm_size KubernetesCluster#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#capacity_reservation_group_id KubernetesCluster#capacity_reservation_group_id}.
	CapacityReservationGroupId *string `field:"optional" json:"capacityReservationGroupId" yaml:"capacityReservationGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#custom_ca_trust_enabled KubernetesCluster#custom_ca_trust_enabled}.
	CustomCaTrustEnabled interface{} `field:"optional" json:"customCaTrustEnabled" yaml:"customCaTrustEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#enable_auto_scaling KubernetesCluster#enable_auto_scaling}.
	EnableAutoScaling interface{} `field:"optional" json:"enableAutoScaling" yaml:"enableAutoScaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#enable_host_encryption KubernetesCluster#enable_host_encryption}.
	EnableHostEncryption interface{} `field:"optional" json:"enableHostEncryption" yaml:"enableHostEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#enable_node_public_ip KubernetesCluster#enable_node_public_ip}.
	EnableNodePublicIp interface{} `field:"optional" json:"enableNodePublicIp" yaml:"enableNodePublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#fips_enabled KubernetesCluster#fips_enabled}.
	FipsEnabled interface{} `field:"optional" json:"fipsEnabled" yaml:"fipsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#gpu_instance KubernetesCluster#gpu_instance}.
	GpuInstance *string `field:"optional" json:"gpuInstance" yaml:"gpuInstance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#host_group_id KubernetesCluster#host_group_id}.
	HostGroupId *string `field:"optional" json:"hostGroupId" yaml:"hostGroupId"`
	// kubelet_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#kubelet_config KubernetesCluster#kubelet_config}
	KubeletConfig *KubernetesClusterDefaultNodePoolKubeletConfig `field:"optional" json:"kubeletConfig" yaml:"kubeletConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#kubelet_disk_type KubernetesCluster#kubelet_disk_type}.
	KubeletDiskType *string `field:"optional" json:"kubeletDiskType" yaml:"kubeletDiskType"`
	// linux_os_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#linux_os_config KubernetesCluster#linux_os_config}
	LinuxOsConfig *KubernetesClusterDefaultNodePoolLinuxOsConfig `field:"optional" json:"linuxOsConfig" yaml:"linuxOsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#max_count KubernetesCluster#max_count}.
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#max_pods KubernetesCluster#max_pods}.
	MaxPods *float64 `field:"optional" json:"maxPods" yaml:"maxPods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#message_of_the_day KubernetesCluster#message_of_the_day}.
	MessageOfTheDay *string `field:"optional" json:"messageOfTheDay" yaml:"messageOfTheDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#min_count KubernetesCluster#min_count}.
	MinCount *float64 `field:"optional" json:"minCount" yaml:"minCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#node_count KubernetesCluster#node_count}.
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#node_labels KubernetesCluster#node_labels}.
	NodeLabels *map[string]*string `field:"optional" json:"nodeLabels" yaml:"nodeLabels"`
	// node_network_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#node_network_profile KubernetesCluster#node_network_profile}
	NodeNetworkProfile *KubernetesClusterDefaultNodePoolNodeNetworkProfile `field:"optional" json:"nodeNetworkProfile" yaml:"nodeNetworkProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#node_public_ip_prefix_id KubernetesCluster#node_public_ip_prefix_id}.
	NodePublicIpPrefixId *string `field:"optional" json:"nodePublicIpPrefixId" yaml:"nodePublicIpPrefixId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#node_taints KubernetesCluster#node_taints}.
	NodeTaints *[]*string `field:"optional" json:"nodeTaints" yaml:"nodeTaints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#only_critical_addons_enabled KubernetesCluster#only_critical_addons_enabled}.
	OnlyCriticalAddonsEnabled interface{} `field:"optional" json:"onlyCriticalAddonsEnabled" yaml:"onlyCriticalAddonsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#orchestrator_version KubernetesCluster#orchestrator_version}.
	OrchestratorVersion *string `field:"optional" json:"orchestratorVersion" yaml:"orchestratorVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#os_disk_size_gb KubernetesCluster#os_disk_size_gb}.
	OsDiskSizeGb *float64 `field:"optional" json:"osDiskSizeGb" yaml:"osDiskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#os_disk_type KubernetesCluster#os_disk_type}.
	OsDiskType *string `field:"optional" json:"osDiskType" yaml:"osDiskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#os_sku KubernetesCluster#os_sku}.
	OsSku *string `field:"optional" json:"osSku" yaml:"osSku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#pod_subnet_id KubernetesCluster#pod_subnet_id}.
	PodSubnetId *string `field:"optional" json:"podSubnetId" yaml:"podSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#proximity_placement_group_id KubernetesCluster#proximity_placement_group_id}.
	ProximityPlacementGroupId *string `field:"optional" json:"proximityPlacementGroupId" yaml:"proximityPlacementGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#scale_down_mode KubernetesCluster#scale_down_mode}.
	ScaleDownMode *string `field:"optional" json:"scaleDownMode" yaml:"scaleDownMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#snapshot_id KubernetesCluster#snapshot_id}.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#tags KubernetesCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#temporary_name_for_rotation KubernetesCluster#temporary_name_for_rotation}.
	TemporaryNameForRotation *string `field:"optional" json:"temporaryNameForRotation" yaml:"temporaryNameForRotation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#type KubernetesCluster#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#ultra_ssd_enabled KubernetesCluster#ultra_ssd_enabled}.
	UltraSsdEnabled interface{} `field:"optional" json:"ultraSsdEnabled" yaml:"ultraSsdEnabled"`
	// upgrade_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#upgrade_settings KubernetesCluster#upgrade_settings}
	UpgradeSettings *KubernetesClusterDefaultNodePoolUpgradeSettings `field:"optional" json:"upgradeSettings" yaml:"upgradeSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#vnet_subnet_id KubernetesCluster#vnet_subnet_id}.
	VnetSubnetId *string `field:"optional" json:"vnetSubnetId" yaml:"vnetSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#workload_runtime KubernetesCluster#workload_runtime}.
	WorkloadRuntime *string `field:"optional" json:"workloadRuntime" yaml:"workloadRuntime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/kubernetes_cluster#zones KubernetesCluster#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

