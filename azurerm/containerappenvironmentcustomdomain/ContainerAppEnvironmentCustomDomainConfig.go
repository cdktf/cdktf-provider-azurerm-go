// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappenvironmentcustomdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAppEnvironmentCustomDomainConfig struct {
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
	// The Custom Domain Certificate Private Key as a base64 encoded PFX or PEM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/container_app_environment_custom_domain#certificate_blob_base64 ContainerAppEnvironmentCustomDomain#certificate_blob_base64}
	CertificateBlobBase64 *string `field:"required" json:"certificateBlobBase64" yaml:"certificateBlobBase64"`
	// The Custom Domain Certificate password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/container_app_environment_custom_domain#certificate_password ContainerAppEnvironmentCustomDomain#certificate_password}
	CertificatePassword *string `field:"required" json:"certificatePassword" yaml:"certificatePassword"`
	// The Container App Managed Environment ID to configure this Custom Domain on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/container_app_environment_custom_domain#container_app_environment_id ContainerAppEnvironmentCustomDomain#container_app_environment_id}
	ContainerAppEnvironmentId *string `field:"required" json:"containerAppEnvironmentId" yaml:"containerAppEnvironmentId"`
	// The Custom Domain DNS suffix for this Container App Environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/container_app_environment_custom_domain#dns_suffix ContainerAppEnvironmentCustomDomain#dns_suffix}
	DnsSuffix *string `field:"required" json:"dnsSuffix" yaml:"dnsSuffix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/container_app_environment_custom_domain#id ContainerAppEnvironmentCustomDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/container_app_environment_custom_domain#timeouts ContainerAppEnvironmentCustomDomain#timeouts}
	Timeouts *ContainerAppEnvironmentCustomDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

