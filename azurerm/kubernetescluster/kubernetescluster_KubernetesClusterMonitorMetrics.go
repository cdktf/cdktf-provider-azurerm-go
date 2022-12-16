package kubernetescluster


type KubernetesClusterMonitorMetrics struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#annotations_allowed KubernetesCluster#annotations_allowed}.
	AnnotationsAllowed *string `field:"optional" json:"annotationsAllowed" yaml:"annotationsAllowed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#labels_allowed KubernetesCluster#labels_allowed}.
	LabelsAllowed *string `field:"optional" json:"labelsAllowed" yaml:"labelsAllowed"`
}

