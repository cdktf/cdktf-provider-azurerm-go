package costanomalyalert

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CostAnomalyAlertConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cost_anomaly_alert#display_name CostAnomalyAlert#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cost_anomaly_alert#email_addresses CostAnomalyAlert#email_addresses}.
	EmailAddresses *[]*string `field:"required" json:"emailAddresses" yaml:"emailAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cost_anomaly_alert#email_subject CostAnomalyAlert#email_subject}.
	EmailSubject *string `field:"required" json:"emailSubject" yaml:"emailSubject"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cost_anomaly_alert#name CostAnomalyAlert#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cost_anomaly_alert#id CostAnomalyAlert#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cost_anomaly_alert#message CostAnomalyAlert#message}.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cost_anomaly_alert#timeouts CostAnomalyAlert#timeouts}
	Timeouts *CostAnomalyAlertTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}
