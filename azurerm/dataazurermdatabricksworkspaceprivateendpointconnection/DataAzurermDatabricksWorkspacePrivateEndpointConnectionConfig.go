package dataazurermdatabricksworkspaceprivateendpointconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermDatabricksWorkspacePrivateEndpointConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/data-sources/databricks_workspace_private_endpoint_connection#private_endpoint_id DataAzurermDatabricksWorkspacePrivateEndpointConnection#private_endpoint_id}.
	PrivateEndpointId *string `field:"required" json:"privateEndpointId" yaml:"privateEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/data-sources/databricks_workspace_private_endpoint_connection#workspace_id DataAzurermDatabricksWorkspacePrivateEndpointConnection#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/data-sources/databricks_workspace_private_endpoint_connection#id DataAzurermDatabricksWorkspacePrivateEndpointConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/data-sources/databricks_workspace_private_endpoint_connection#timeouts DataAzurermDatabricksWorkspacePrivateEndpointConnection#timeouts}
	Timeouts *DataAzurermDatabricksWorkspacePrivateEndpointConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

