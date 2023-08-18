package kubernetescluster


type KubernetesClusterOmsAgent struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/kubernetes_cluster#log_analytics_workspace_id KubernetesCluster#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/kubernetes_cluster#msi_auth_for_monitoring_enabled KubernetesCluster#msi_auth_for_monitoring_enabled}.
	MsiAuthForMonitoringEnabled interface{} `field:"optional" json:"msiAuthForMonitoringEnabled" yaml:"msiAuthForMonitoringEnabled"`
}

