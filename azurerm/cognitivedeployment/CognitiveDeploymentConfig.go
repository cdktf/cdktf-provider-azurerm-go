package cognitivedeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitiveDeploymentConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cognitive_deployment#cognitive_account_id CognitiveDeployment#cognitive_account_id}.
	CognitiveAccountId *string `field:"required" json:"cognitiveAccountId" yaml:"cognitiveAccountId"`
	// model block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cognitive_deployment#model CognitiveDeployment#model}
	Model *CognitiveDeploymentModel `field:"required" json:"model" yaml:"model"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cognitive_deployment#name CognitiveDeployment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// scale block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cognitive_deployment#scale CognitiveDeployment#scale}
	Scale *CognitiveDeploymentScale `field:"required" json:"scale" yaml:"scale"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cognitive_deployment#id CognitiveDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cognitive_deployment#rai_policy_name CognitiveDeployment#rai_policy_name}.
	RaiPolicyName *string `field:"optional" json:"raiPolicyName" yaml:"raiPolicyName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cognitive_deployment#timeouts CognitiveDeployment#timeouts}
	Timeouts *CognitiveDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}
