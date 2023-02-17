package kubernetescluster


type KubernetesClusterOmsAgent struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#log_analytics_workspace_id KubernetesCluster#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
}

