package roledefinition


type RoleDefinitionPermissions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/role_definition#actions RoleDefinition#actions}.
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/role_definition#data_actions RoleDefinition#data_actions}.
	DataActions *[]*string `field:"optional" json:"dataActions" yaml:"dataActions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/role_definition#not_actions RoleDefinition#not_actions}.
	NotActions *[]*string `field:"optional" json:"notActions" yaml:"notActions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/role_definition#not_data_actions RoleDefinition#not_data_actions}.
	NotDataActions *[]*string `field:"optional" json:"notDataActions" yaml:"notDataActions"`
}
