// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentineldataconnectorawss3


type SentinelDataConnectorAwsS3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/sentinel_data_connector_aws_s3#create SentinelDataConnectorAwsS3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/sentinel_data_connector_aws_s3#delete SentinelDataConnectorAwsS3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/sentinel_data_connector_aws_s3#read SentinelDataConnectorAwsS3#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/sentinel_data_connector_aws_s3#update SentinelDataConnectorAwsS3#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

