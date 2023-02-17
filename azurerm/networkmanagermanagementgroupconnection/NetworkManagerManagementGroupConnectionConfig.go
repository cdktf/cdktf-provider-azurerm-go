package networkmanagermanagementgroupconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkManagerManagementGroupConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_management_group_connection#management_group_id NetworkManagerManagementGroupConnection#management_group_id}.
	ManagementGroupId *string `field:"required" json:"managementGroupId" yaml:"managementGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_management_group_connection#name NetworkManagerManagementGroupConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_management_group_connection#network_manager_id NetworkManagerManagementGroupConnection#network_manager_id}.
	NetworkManagerId *string `field:"required" json:"networkManagerId" yaml:"networkManagerId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_management_group_connection#description NetworkManagerManagementGroupConnection#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_management_group_connection#id NetworkManagerManagementGroupConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_management_group_connection#timeouts NetworkManagerManagementGroupConnection#timeouts}
	Timeouts *NetworkManagerManagementGroupConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

