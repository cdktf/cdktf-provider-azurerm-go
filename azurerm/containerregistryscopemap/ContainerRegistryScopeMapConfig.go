package containerregistryscopemap

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerRegistryScopeMapConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_scope_map#actions ContainerRegistryScopeMap#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_scope_map#container_registry_name ContainerRegistryScopeMap#container_registry_name}.
	ContainerRegistryName *string `field:"required" json:"containerRegistryName" yaml:"containerRegistryName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_scope_map#name ContainerRegistryScopeMap#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_scope_map#resource_group_name ContainerRegistryScopeMap#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_scope_map#description ContainerRegistryScopeMap#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_scope_map#id ContainerRegistryScopeMap#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_scope_map#timeouts ContainerRegistryScopeMap#timeouts}
	Timeouts *ContainerRegistryScopeMapTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}
