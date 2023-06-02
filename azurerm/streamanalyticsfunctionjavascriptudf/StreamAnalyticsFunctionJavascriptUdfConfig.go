package streamanalyticsfunctionjavascriptudf

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamAnalyticsFunctionJavascriptUdfConfig struct {
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
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/stream_analytics_function_javascript_udf#input StreamAnalyticsFunctionJavascriptUdf#input}
	Input interface{} `field:"required" json:"input" yaml:"input"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/stream_analytics_function_javascript_udf#name StreamAnalyticsFunctionJavascriptUdf#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/stream_analytics_function_javascript_udf#output StreamAnalyticsFunctionJavascriptUdf#output}
	Output *StreamAnalyticsFunctionJavascriptUdfOutput `field:"required" json:"output" yaml:"output"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/stream_analytics_function_javascript_udf#resource_group_name StreamAnalyticsFunctionJavascriptUdf#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/stream_analytics_function_javascript_udf#script StreamAnalyticsFunctionJavascriptUdf#script}.
	Script *string `field:"required" json:"script" yaml:"script"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/stream_analytics_function_javascript_udf#stream_analytics_job_name StreamAnalyticsFunctionJavascriptUdf#stream_analytics_job_name}.
	StreamAnalyticsJobName *string `field:"required" json:"streamAnalyticsJobName" yaml:"streamAnalyticsJobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/stream_analytics_function_javascript_udf#id StreamAnalyticsFunctionJavascriptUdf#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/stream_analytics_function_javascript_udf#timeouts StreamAnalyticsFunctionJavascriptUdf#timeouts}
	Timeouts *StreamAnalyticsFunctionJavascriptUdfTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

