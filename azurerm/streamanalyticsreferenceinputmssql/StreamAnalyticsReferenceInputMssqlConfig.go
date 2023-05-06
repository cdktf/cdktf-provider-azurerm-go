package streamanalyticsreferenceinputmssql

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamAnalyticsReferenceInputMssqlConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/stream_analytics_reference_input_mssql#database StreamAnalyticsReferenceInputMssql#database}.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/stream_analytics_reference_input_mssql#full_snapshot_query StreamAnalyticsReferenceInputMssql#full_snapshot_query}.
	FullSnapshotQuery *string `field:"required" json:"fullSnapshotQuery" yaml:"fullSnapshotQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/stream_analytics_reference_input_mssql#name StreamAnalyticsReferenceInputMssql#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/stream_analytics_reference_input_mssql#password StreamAnalyticsReferenceInputMssql#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/stream_analytics_reference_input_mssql#refresh_type StreamAnalyticsReferenceInputMssql#refresh_type}.
	RefreshType *string `field:"required" json:"refreshType" yaml:"refreshType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/stream_analytics_reference_input_mssql#resource_group_name StreamAnalyticsReferenceInputMssql#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/stream_analytics_reference_input_mssql#server StreamAnalyticsReferenceInputMssql#server}.
	Server *string `field:"required" json:"server" yaml:"server"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/stream_analytics_reference_input_mssql#stream_analytics_job_name StreamAnalyticsReferenceInputMssql#stream_analytics_job_name}.
	StreamAnalyticsJobName *string `field:"required" json:"streamAnalyticsJobName" yaml:"streamAnalyticsJobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/stream_analytics_reference_input_mssql#username StreamAnalyticsReferenceInputMssql#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/stream_analytics_reference_input_mssql#delta_snapshot_query StreamAnalyticsReferenceInputMssql#delta_snapshot_query}.
	DeltaSnapshotQuery *string `field:"optional" json:"deltaSnapshotQuery" yaml:"deltaSnapshotQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/stream_analytics_reference_input_mssql#id StreamAnalyticsReferenceInputMssql#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/stream_analytics_reference_input_mssql#refresh_interval_duration StreamAnalyticsReferenceInputMssql#refresh_interval_duration}.
	RefreshIntervalDuration *string `field:"optional" json:"refreshIntervalDuration" yaml:"refreshIntervalDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/stream_analytics_reference_input_mssql#table StreamAnalyticsReferenceInputMssql#table}.
	Table *string `field:"optional" json:"table" yaml:"table"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/stream_analytics_reference_input_mssql#timeouts StreamAnalyticsReferenceInputMssql#timeouts}
	Timeouts *StreamAnalyticsReferenceInputMssqlTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

