package mysqlflexibleserveractivedirectoryadministrator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MysqlFlexibleServerActiveDirectoryAdministratorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mysql_flexible_server_active_directory_administrator#identity_id MysqlFlexibleServerActiveDirectoryAdministrator#identity_id}.
	IdentityId *string `field:"required" json:"identityId" yaml:"identityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mysql_flexible_server_active_directory_administrator#login MysqlFlexibleServerActiveDirectoryAdministrator#login}.
	Login *string `field:"required" json:"login" yaml:"login"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mysql_flexible_server_active_directory_administrator#object_id MysqlFlexibleServerActiveDirectoryAdministrator#object_id}.
	ObjectId *string `field:"required" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mysql_flexible_server_active_directory_administrator#server_id MysqlFlexibleServerActiveDirectoryAdministrator#server_id}.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mysql_flexible_server_active_directory_administrator#tenant_id MysqlFlexibleServerActiveDirectoryAdministrator#tenant_id}.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mysql_flexible_server_active_directory_administrator#id MysqlFlexibleServerActiveDirectoryAdministrator#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mysql_flexible_server_active_directory_administrator#timeouts MysqlFlexibleServerActiveDirectoryAdministrator#timeouts}
	Timeouts *MysqlFlexibleServerActiveDirectoryAdministratorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

