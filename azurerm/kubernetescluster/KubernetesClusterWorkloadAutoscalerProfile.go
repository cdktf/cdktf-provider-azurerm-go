package kubernetescluster


type KubernetesClusterWorkloadAutoscalerProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#keda_enabled KubernetesCluster#keda_enabled}.
	KedaEnabled interface{} `field:"optional" json:"kedaEnabled" yaml:"kedaEnabled"`
}

