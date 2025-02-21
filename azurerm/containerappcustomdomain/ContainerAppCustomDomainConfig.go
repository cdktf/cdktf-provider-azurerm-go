// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappcustomdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAppCustomDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/container_app_custom_domain#container_app_id ContainerAppCustomDomain#container_app_id}.
	ContainerAppId *string `field:"required" json:"containerAppId" yaml:"containerAppId"`
	// The hostname of the Certificate. Must be the CN or a named SAN in the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/container_app_custom_domain#name ContainerAppCustomDomain#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Binding type. Possible values include `Disabled` and `SniEnabled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/container_app_custom_domain#certificate_binding_type ContainerAppCustomDomain#certificate_binding_type}
	CertificateBindingType *string `field:"optional" json:"certificateBindingType" yaml:"certificateBindingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/container_app_custom_domain#container_app_environment_certificate_id ContainerAppCustomDomain#container_app_environment_certificate_id}.
	ContainerAppEnvironmentCertificateId *string `field:"optional" json:"containerAppEnvironmentCertificateId" yaml:"containerAppEnvironmentCertificateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/container_app_custom_domain#id ContainerAppCustomDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/container_app_custom_domain#timeouts ContainerAppCustomDomain#timeouts}
	Timeouts *ContainerAppCustomDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

