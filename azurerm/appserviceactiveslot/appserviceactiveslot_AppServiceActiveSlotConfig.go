package appserviceactiveslot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppServiceActiveSlotConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_active_slot#app_service_name AppServiceActiveSlot#app_service_name}.
	AppServiceName *string `field:"required" json:"appServiceName" yaml:"appServiceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_active_slot#app_service_slot_name AppServiceActiveSlot#app_service_slot_name}.
	AppServiceSlotName *string `field:"required" json:"appServiceSlotName" yaml:"appServiceSlotName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_active_slot#resource_group_name AppServiceActiveSlot#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_active_slot#id AppServiceActiveSlot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_active_slot#timeouts AppServiceActiveSlot#timeouts}
	Timeouts *AppServiceActiveSlotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}
