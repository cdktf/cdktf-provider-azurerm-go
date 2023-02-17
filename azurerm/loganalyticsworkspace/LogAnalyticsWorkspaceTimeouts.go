package loganalyticsworkspace


type LogAnalyticsWorkspaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_workspace#create LogAnalyticsWorkspace#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_workspace#delete LogAnalyticsWorkspace#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_workspace#read LogAnalyticsWorkspace#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_workspace#update LogAnalyticsWorkspace#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

