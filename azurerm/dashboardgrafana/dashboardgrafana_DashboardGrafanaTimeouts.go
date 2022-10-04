package dashboardgrafana


type DashboardGrafanaTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#create DashboardGrafana#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#delete DashboardGrafana#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#read DashboardGrafana#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#update DashboardGrafana#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

