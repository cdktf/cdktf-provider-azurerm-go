// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorytriggertumblingwindow

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryTriggerTumblingWindowConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#data_factory_id DataFactoryTriggerTumblingWindow#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#frequency DataFactoryTriggerTumblingWindow#frequency}.
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#interval DataFactoryTriggerTumblingWindow#interval}.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#name DataFactoryTriggerTumblingWindow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// pipeline block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#pipeline DataFactoryTriggerTumblingWindow#pipeline}
	Pipeline *DataFactoryTriggerTumblingWindowPipeline `field:"required" json:"pipeline" yaml:"pipeline"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#start_time DataFactoryTriggerTumblingWindow#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#activated DataFactoryTriggerTumblingWindow#activated}.
	Activated interface{} `field:"optional" json:"activated" yaml:"activated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#additional_properties DataFactoryTriggerTumblingWindow#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#annotations DataFactoryTriggerTumblingWindow#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#delay DataFactoryTriggerTumblingWindow#delay}.
	Delay *string `field:"optional" json:"delay" yaml:"delay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#description DataFactoryTriggerTumblingWindow#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#end_time DataFactoryTriggerTumblingWindow#end_time}.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#id DataFactoryTriggerTumblingWindow#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#max_concurrency DataFactoryTriggerTumblingWindow#max_concurrency}.
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// retry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#retry DataFactoryTriggerTumblingWindow#retry}
	Retry *DataFactoryTriggerTumblingWindowRetry `field:"optional" json:"retry" yaml:"retry"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#timeouts DataFactoryTriggerTumblingWindow#timeouts}
	Timeouts *DataFactoryTriggerTumblingWindowTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// trigger_dependency block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/data_factory_trigger_tumbling_window#trigger_dependency DataFactoryTriggerTumblingWindow#trigger_dependency}
	TriggerDependency interface{} `field:"optional" json:"triggerDependency" yaml:"triggerDependency"`
}

