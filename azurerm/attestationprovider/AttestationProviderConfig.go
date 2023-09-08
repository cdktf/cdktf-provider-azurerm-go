// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package attestationprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AttestationProviderConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/attestation_provider#location AttestationProvider#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/attestation_provider#name AttestationProvider#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/attestation_provider#resource_group_name AttestationProvider#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/attestation_provider#id AttestationProvider#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/attestation_provider#open_enclave_policy_base64 AttestationProvider#open_enclave_policy_base64}.
	OpenEnclavePolicyBase64 *string `field:"optional" json:"openEnclavePolicyBase64" yaml:"openEnclavePolicyBase64"`
	// policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/attestation_provider#policy AttestationProvider#policy}
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/attestation_provider#policy_signing_certificate_data AttestationProvider#policy_signing_certificate_data}.
	PolicySigningCertificateData *string `field:"optional" json:"policySigningCertificateData" yaml:"policySigningCertificateData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/attestation_provider#sev_snp_policy_base64 AttestationProvider#sev_snp_policy_base64}.
	SevSnpPolicyBase64 *string `field:"optional" json:"sevSnpPolicyBase64" yaml:"sevSnpPolicyBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/attestation_provider#sgx_enclave_policy_base64 AttestationProvider#sgx_enclave_policy_base64}.
	SgxEnclavePolicyBase64 *string `field:"optional" json:"sgxEnclavePolicyBase64" yaml:"sgxEnclavePolicyBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/attestation_provider#tags AttestationProvider#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/attestation_provider#timeouts AttestationProvider#timeouts}
	Timeouts *AttestationProviderTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/attestation_provider#tpm_policy_base64 AttestationProvider#tpm_policy_base64}.
	TpmPolicyBase64 *string `field:"optional" json:"tpmPolicyBase64" yaml:"tpmPolicyBase64"`
}

