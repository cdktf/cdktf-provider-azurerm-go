package logicappactioncustom


type LogicAppActionCustomTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_action_custom#create LogicAppActionCustom#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_action_custom#delete LogicAppActionCustom#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_action_custom#read LogicAppActionCustom#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_action_custom#update LogicAppActionCustom#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
