package iotsecuritysolution


type IotSecuritySolutionAdditionalWorkspace struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#data_types IotSecuritySolution#data_types}.
	DataTypes *[]*string `field:"required" json:"dataTypes" yaml:"dataTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#workspace_id IotSecuritySolution#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

