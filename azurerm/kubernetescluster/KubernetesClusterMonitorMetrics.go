// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterMonitorMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/kubernetes_cluster#annotations_allowed KubernetesCluster#annotations_allowed}.
	AnnotationsAllowed *string `field:"optional" json:"annotationsAllowed" yaml:"annotationsAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/kubernetes_cluster#labels_allowed KubernetesCluster#labels_allowed}.
	LabelsAllowed *string `field:"optional" json:"labelsAllowed" yaml:"labelsAllowed"`
}

