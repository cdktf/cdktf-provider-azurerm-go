package kubernetescluster


type KubernetesClusterWorkloadAutoscalerProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/kubernetes_cluster#keda_enabled KubernetesCluster#keda_enabled}.
	KedaEnabled interface{} `field:"optional" json:"kedaEnabled" yaml:"kedaEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/kubernetes_cluster#vertical_pod_autoscaler_enabled KubernetesCluster#vertical_pod_autoscaler_enabled}.
	VerticalPodAutoscalerEnabled interface{} `field:"optional" json:"verticalPodAutoscalerEnabled" yaml:"verticalPodAutoscalerEnabled"`
}

