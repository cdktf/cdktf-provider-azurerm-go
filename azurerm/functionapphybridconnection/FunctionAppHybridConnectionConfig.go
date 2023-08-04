package functionapphybridconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionAppHybridConnectionConfig struct {
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
	// The ID of the Function App for this Hybrid Connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/function_app_hybrid_connection#function_app_id FunctionAppHybridConnection#function_app_id}
	FunctionAppId *string `field:"required" json:"functionAppId" yaml:"functionAppId"`
	// The hostname of the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/function_app_hybrid_connection#hostname FunctionAppHybridConnection#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// The port to use for the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/function_app_hybrid_connection#port FunctionAppHybridConnection#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The ID of the Relay Hybrid Connection to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/function_app_hybrid_connection#relay_id FunctionAppHybridConnection#relay_id}
	RelayId *string `field:"required" json:"relayId" yaml:"relayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/function_app_hybrid_connection#id FunctionAppHybridConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the Relay key with `Send` permission to use. Defaults to `RootManageSharedAccessKey`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/function_app_hybrid_connection#send_key_name FunctionAppHybridConnection#send_key_name}
	SendKeyName *string `field:"optional" json:"sendKeyName" yaml:"sendKeyName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/function_app_hybrid_connection#timeouts FunctionAppHybridConnection#timeouts}
	Timeouts *FunctionAppHybridConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

