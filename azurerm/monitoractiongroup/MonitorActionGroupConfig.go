package monitoractiongroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorActionGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#name MonitorActionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#resource_group_name MonitorActionGroup#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#short_name MonitorActionGroup#short_name}.
	ShortName *string `field:"required" json:"shortName" yaml:"shortName"`
	// arm_role_receiver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#arm_role_receiver MonitorActionGroup#arm_role_receiver}
	ArmRoleReceiver interface{} `field:"optional" json:"armRoleReceiver" yaml:"armRoleReceiver"`
	// automation_runbook_receiver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#automation_runbook_receiver MonitorActionGroup#automation_runbook_receiver}
	AutomationRunbookReceiver interface{} `field:"optional" json:"automationRunbookReceiver" yaml:"automationRunbookReceiver"`
	// azure_app_push_receiver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#azure_app_push_receiver MonitorActionGroup#azure_app_push_receiver}
	AzureAppPushReceiver interface{} `field:"optional" json:"azureAppPushReceiver" yaml:"azureAppPushReceiver"`
	// azure_function_receiver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#azure_function_receiver MonitorActionGroup#azure_function_receiver}
	AzureFunctionReceiver interface{} `field:"optional" json:"azureFunctionReceiver" yaml:"azureFunctionReceiver"`
	// email_receiver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#email_receiver MonitorActionGroup#email_receiver}
	EmailReceiver interface{} `field:"optional" json:"emailReceiver" yaml:"emailReceiver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#enabled MonitorActionGroup#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// event_hub_receiver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#event_hub_receiver MonitorActionGroup#event_hub_receiver}
	EventHubReceiver interface{} `field:"optional" json:"eventHubReceiver" yaml:"eventHubReceiver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#id MonitorActionGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// itsm_receiver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#itsm_receiver MonitorActionGroup#itsm_receiver}
	ItsmReceiver interface{} `field:"optional" json:"itsmReceiver" yaml:"itsmReceiver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#location MonitorActionGroup#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// logic_app_receiver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#logic_app_receiver MonitorActionGroup#logic_app_receiver}
	LogicAppReceiver interface{} `field:"optional" json:"logicAppReceiver" yaml:"logicAppReceiver"`
	// sms_receiver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#sms_receiver MonitorActionGroup#sms_receiver}
	SmsReceiver interface{} `field:"optional" json:"smsReceiver" yaml:"smsReceiver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#tags MonitorActionGroup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#timeouts MonitorActionGroup#timeouts}
	Timeouts *MonitorActionGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// voice_receiver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#voice_receiver MonitorActionGroup#voice_receiver}
	VoiceReceiver interface{} `field:"optional" json:"voiceReceiver" yaml:"voiceReceiver"`
	// webhook_receiver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_group#webhook_receiver MonitorActionGroup#webhook_receiver}
	WebhookReceiver interface{} `field:"optional" json:"webhookReceiver" yaml:"webhookReceiver"`
}

