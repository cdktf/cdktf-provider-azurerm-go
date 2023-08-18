package containerapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAppConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// The ID of the Container App Environment to host this Container App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/container_app#container_app_environment_id ContainerApp#container_app_environment_id}
	ContainerAppEnvironmentId *string `field:"required" json:"containerAppEnvironmentId" yaml:"containerAppEnvironmentId"`
	// The name for this Container App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/container_app#name ContainerApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/container_app#resource_group_name ContainerApp#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/container_app#revision_mode ContainerApp#revision_mode}.
	RevisionMode *string `field:"required" json:"revisionMode" yaml:"revisionMode"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/container_app#template ContainerApp#template}
	Template *ContainerAppTemplate `field:"required" json:"template" yaml:"template"`
	// dapr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/container_app#dapr ContainerApp#dapr}
	Dapr *ContainerAppDapr `field:"optional" json:"dapr" yaml:"dapr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/container_app#id ContainerApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/container_app#identity ContainerApp#identity}
	Identity *ContainerAppIdentity `field:"optional" json:"identity" yaml:"identity"`
	// ingress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/container_app#ingress ContainerApp#ingress}
	Ingress *ContainerAppIngress `field:"optional" json:"ingress" yaml:"ingress"`
	// registry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/container_app#registry ContainerApp#registry}
	Registry interface{} `field:"optional" json:"registry" yaml:"registry"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/container_app#secret ContainerApp#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/container_app#tags ContainerApp#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/container_app#timeouts ContainerApp#timeouts}
	Timeouts *ContainerAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

