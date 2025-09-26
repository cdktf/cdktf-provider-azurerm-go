// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqljobstep

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlJobStepConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/mssql_job_step#job_id MssqlJobStep#job_id}.
	JobId *string `field:"required" json:"jobId" yaml:"jobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/mssql_job_step#job_step_index MssqlJobStep#job_step_index}.
	JobStepIndex *float64 `field:"required" json:"jobStepIndex" yaml:"jobStepIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/mssql_job_step#job_target_group_id MssqlJobStep#job_target_group_id}.
	JobTargetGroupId *string `field:"required" json:"jobTargetGroupId" yaml:"jobTargetGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/mssql_job_step#name MssqlJobStep#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/mssql_job_step#sql_script MssqlJobStep#sql_script}.
	SqlScript *string `field:"required" json:"sqlScript" yaml:"sqlScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/mssql_job_step#id MssqlJobStep#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/mssql_job_step#initial_retry_interval_seconds MssqlJobStep#initial_retry_interval_seconds}.
	InitialRetryIntervalSeconds *float64 `field:"optional" json:"initialRetryIntervalSeconds" yaml:"initialRetryIntervalSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/mssql_job_step#job_credential_id MssqlJobStep#job_credential_id}.
	JobCredentialId *string `field:"optional" json:"jobCredentialId" yaml:"jobCredentialId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/mssql_job_step#maximum_retry_interval_seconds MssqlJobStep#maximum_retry_interval_seconds}.
	MaximumRetryIntervalSeconds *float64 `field:"optional" json:"maximumRetryIntervalSeconds" yaml:"maximumRetryIntervalSeconds"`
	// output_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/mssql_job_step#output_target MssqlJobStep#output_target}
	OutputTarget *MssqlJobStepOutputTarget `field:"optional" json:"outputTarget" yaml:"outputTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/mssql_job_step#retry_attempts MssqlJobStep#retry_attempts}.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/mssql_job_step#retry_interval_backoff_multiplier MssqlJobStep#retry_interval_backoff_multiplier}.
	RetryIntervalBackoffMultiplier *float64 `field:"optional" json:"retryIntervalBackoffMultiplier" yaml:"retryIntervalBackoffMultiplier"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/mssql_job_step#timeouts MssqlJobStep#timeouts}
	Timeouts *MssqlJobStepTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/mssql_job_step#timeout_seconds MssqlJobStep#timeout_seconds}.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

