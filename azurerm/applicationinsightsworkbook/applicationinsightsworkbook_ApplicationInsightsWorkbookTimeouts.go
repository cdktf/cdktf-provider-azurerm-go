package applicationinsightsworkbook


type ApplicationInsightsWorkbookTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_insights_workbook#create ApplicationInsightsWorkbook#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_insights_workbook#delete ApplicationInsightsWorkbook#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_insights_workbook#read ApplicationInsightsWorkbook#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_insights_workbook#update ApplicationInsightsWorkbook#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
