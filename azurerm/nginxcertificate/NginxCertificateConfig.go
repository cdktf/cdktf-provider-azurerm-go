// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nginxcertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NginxCertificateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/nginx_certificate#certificate_virtual_path NginxCertificate#certificate_virtual_path}.
	CertificateVirtualPath *string `field:"required" json:"certificateVirtualPath" yaml:"certificateVirtualPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/nginx_certificate#key_vault_secret_id NginxCertificate#key_vault_secret_id}.
	KeyVaultSecretId *string `field:"required" json:"keyVaultSecretId" yaml:"keyVaultSecretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/nginx_certificate#key_virtual_path NginxCertificate#key_virtual_path}.
	KeyVirtualPath *string `field:"required" json:"keyVirtualPath" yaml:"keyVirtualPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/nginx_certificate#name NginxCertificate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/nginx_certificate#nginx_deployment_id NginxCertificate#nginx_deployment_id}.
	NginxDeploymentId *string `field:"required" json:"nginxDeploymentId" yaml:"nginxDeploymentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/nginx_certificate#id NginxCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/nginx_certificate#timeouts NginxCertificate#timeouts}
	Timeouts *NginxCertificateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

