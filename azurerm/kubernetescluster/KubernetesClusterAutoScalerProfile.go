// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterAutoScalerProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#balance_similar_node_groups KubernetesCluster#balance_similar_node_groups}.
	BalanceSimilarNodeGroups interface{} `field:"optional" json:"balanceSimilarNodeGroups" yaml:"balanceSimilarNodeGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#empty_bulk_delete_max KubernetesCluster#empty_bulk_delete_max}.
	EmptyBulkDeleteMax *string `field:"optional" json:"emptyBulkDeleteMax" yaml:"emptyBulkDeleteMax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#expander KubernetesCluster#expander}.
	Expander *string `field:"optional" json:"expander" yaml:"expander"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#max_graceful_termination_sec KubernetesCluster#max_graceful_termination_sec}.
	MaxGracefulTerminationSec *string `field:"optional" json:"maxGracefulTerminationSec" yaml:"maxGracefulTerminationSec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#max_node_provisioning_time KubernetesCluster#max_node_provisioning_time}.
	MaxNodeProvisioningTime *string `field:"optional" json:"maxNodeProvisioningTime" yaml:"maxNodeProvisioningTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#max_unready_nodes KubernetesCluster#max_unready_nodes}.
	MaxUnreadyNodes *float64 `field:"optional" json:"maxUnreadyNodes" yaml:"maxUnreadyNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#max_unready_percentage KubernetesCluster#max_unready_percentage}.
	MaxUnreadyPercentage *float64 `field:"optional" json:"maxUnreadyPercentage" yaml:"maxUnreadyPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#new_pod_scale_up_delay KubernetesCluster#new_pod_scale_up_delay}.
	NewPodScaleUpDelay *string `field:"optional" json:"newPodScaleUpDelay" yaml:"newPodScaleUpDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#scale_down_delay_after_add KubernetesCluster#scale_down_delay_after_add}.
	ScaleDownDelayAfterAdd *string `field:"optional" json:"scaleDownDelayAfterAdd" yaml:"scaleDownDelayAfterAdd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#scale_down_delay_after_delete KubernetesCluster#scale_down_delay_after_delete}.
	ScaleDownDelayAfterDelete *string `field:"optional" json:"scaleDownDelayAfterDelete" yaml:"scaleDownDelayAfterDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#scale_down_delay_after_failure KubernetesCluster#scale_down_delay_after_failure}.
	ScaleDownDelayAfterFailure *string `field:"optional" json:"scaleDownDelayAfterFailure" yaml:"scaleDownDelayAfterFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#scale_down_unneeded KubernetesCluster#scale_down_unneeded}.
	ScaleDownUnneeded *string `field:"optional" json:"scaleDownUnneeded" yaml:"scaleDownUnneeded"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#scale_down_unready KubernetesCluster#scale_down_unready}.
	ScaleDownUnready *string `field:"optional" json:"scaleDownUnready" yaml:"scaleDownUnready"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#scale_down_utilization_threshold KubernetesCluster#scale_down_utilization_threshold}.
	ScaleDownUtilizationThreshold *string `field:"optional" json:"scaleDownUtilizationThreshold" yaml:"scaleDownUtilizationThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#scan_interval KubernetesCluster#scan_interval}.
	ScanInterval *string `field:"optional" json:"scanInterval" yaml:"scanInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#skip_nodes_with_local_storage KubernetesCluster#skip_nodes_with_local_storage}.
	SkipNodesWithLocalStorage interface{} `field:"optional" json:"skipNodesWithLocalStorage" yaml:"skipNodesWithLocalStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/kubernetes_cluster#skip_nodes_with_system_pods KubernetesCluster#skip_nodes_with_system_pods}.
	SkipNodesWithSystemPods interface{} `field:"optional" json:"skipNodesWithSystemPods" yaml:"skipNodesWithSystemPods"`
}

