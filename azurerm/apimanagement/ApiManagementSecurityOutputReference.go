// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagement

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/apimanagement/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementSecurityOutputReference interface {
	cdktf.ComplexObject
	BackendSsl30Enabled() interface{}
	SetBackendSsl30Enabled(val interface{})
	BackendSsl30EnabledInput() interface{}
	BackendTls10Enabled() interface{}
	SetBackendTls10Enabled(val interface{})
	BackendTls10EnabledInput() interface{}
	BackendTls11Enabled() interface{}
	SetBackendTls11Enabled(val interface{})
	BackendTls11EnabledInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableBackendSsl30() interface{}
	SetEnableBackendSsl30(val interface{})
	EnableBackendSsl30Input() interface{}
	EnableBackendTls10() interface{}
	SetEnableBackendTls10(val interface{})
	EnableBackendTls10Input() interface{}
	EnableBackendTls11() interface{}
	SetEnableBackendTls11(val interface{})
	EnableBackendTls11Input() interface{}
	EnableFrontendSsl30() interface{}
	SetEnableFrontendSsl30(val interface{})
	EnableFrontendSsl30Input() interface{}
	EnableFrontendTls10() interface{}
	SetEnableFrontendTls10(val interface{})
	EnableFrontendTls10Input() interface{}
	EnableFrontendTls11() interface{}
	SetEnableFrontendTls11(val interface{})
	EnableFrontendTls11Input() interface{}
	// Experimental.
	Fqn() *string
	FrontendSsl30Enabled() interface{}
	SetFrontendSsl30Enabled(val interface{})
	FrontendSsl30EnabledInput() interface{}
	FrontendTls10Enabled() interface{}
	SetFrontendTls10Enabled(val interface{})
	FrontendTls10EnabledInput() interface{}
	FrontendTls11Enabled() interface{}
	SetFrontendTls11Enabled(val interface{})
	FrontendTls11EnabledInput() interface{}
	InternalValue() *ApiManagementSecurity
	SetInternalValue(val *ApiManagementSecurity)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsEcdheEcdsaWithAes128CbcShaCiphersEnabled() interface{}
	SetTlsEcdheEcdsaWithAes128CbcShaCiphersEnabled(val interface{})
	TlsEcdheEcdsaWithAes128CbcShaCiphersEnabledInput() interface{}
	TlsEcdheEcdsaWithAes256CbcShaCiphersEnabled() interface{}
	SetTlsEcdheEcdsaWithAes256CbcShaCiphersEnabled(val interface{})
	TlsEcdheEcdsaWithAes256CbcShaCiphersEnabledInput() interface{}
	TlsEcdheRsaWithAes128CbcShaCiphersEnabled() interface{}
	SetTlsEcdheRsaWithAes128CbcShaCiphersEnabled(val interface{})
	TlsEcdheRsaWithAes128CbcShaCiphersEnabledInput() interface{}
	TlsEcdheRsaWithAes256CbcShaCiphersEnabled() interface{}
	SetTlsEcdheRsaWithAes256CbcShaCiphersEnabled(val interface{})
	TlsEcdheRsaWithAes256CbcShaCiphersEnabledInput() interface{}
	TlsRsaWithAes128CbcSha256CiphersEnabled() interface{}
	SetTlsRsaWithAes128CbcSha256CiphersEnabled(val interface{})
	TlsRsaWithAes128CbcSha256CiphersEnabledInput() interface{}
	TlsRsaWithAes128CbcShaCiphersEnabled() interface{}
	SetTlsRsaWithAes128CbcShaCiphersEnabled(val interface{})
	TlsRsaWithAes128CbcShaCiphersEnabledInput() interface{}
	TlsRsaWithAes128GcmSha256CiphersEnabled() interface{}
	SetTlsRsaWithAes128GcmSha256CiphersEnabled(val interface{})
	TlsRsaWithAes128GcmSha256CiphersEnabledInput() interface{}
	TlsRsaWithAes256CbcSha256CiphersEnabled() interface{}
	SetTlsRsaWithAes256CbcSha256CiphersEnabled(val interface{})
	TlsRsaWithAes256CbcSha256CiphersEnabledInput() interface{}
	TlsRsaWithAes256CbcShaCiphersEnabled() interface{}
	SetTlsRsaWithAes256CbcShaCiphersEnabled(val interface{})
	TlsRsaWithAes256CbcShaCiphersEnabledInput() interface{}
	TlsRsaWithAes256GcmSha384CiphersEnabled() interface{}
	SetTlsRsaWithAes256GcmSha384CiphersEnabled(val interface{})
	TlsRsaWithAes256GcmSha384CiphersEnabledInput() interface{}
	TripleDesCiphersEnabled() interface{}
	SetTripleDesCiphersEnabled(val interface{})
	TripleDesCiphersEnabledInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetBackendSsl30Enabled()
	ResetBackendTls10Enabled()
	ResetBackendTls11Enabled()
	ResetEnableBackendSsl30()
	ResetEnableBackendTls10()
	ResetEnableBackendTls11()
	ResetEnableFrontendSsl30()
	ResetEnableFrontendTls10()
	ResetEnableFrontendTls11()
	ResetFrontendSsl30Enabled()
	ResetFrontendTls10Enabled()
	ResetFrontendTls11Enabled()
	ResetTlsEcdheEcdsaWithAes128CbcShaCiphersEnabled()
	ResetTlsEcdheEcdsaWithAes256CbcShaCiphersEnabled()
	ResetTlsEcdheRsaWithAes128CbcShaCiphersEnabled()
	ResetTlsEcdheRsaWithAes256CbcShaCiphersEnabled()
	ResetTlsRsaWithAes128CbcSha256CiphersEnabled()
	ResetTlsRsaWithAes128CbcShaCiphersEnabled()
	ResetTlsRsaWithAes128GcmSha256CiphersEnabled()
	ResetTlsRsaWithAes256CbcSha256CiphersEnabled()
	ResetTlsRsaWithAes256CbcShaCiphersEnabled()
	ResetTlsRsaWithAes256GcmSha384CiphersEnabled()
	ResetTripleDesCiphersEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementSecurityOutputReference
type jsiiProxy_ApiManagementSecurityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) BackendSsl30Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendSsl30Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) BackendSsl30EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendSsl30EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) BackendTls10Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendTls10Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) BackendTls10EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendTls10EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) BackendTls11Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendTls11Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) BackendTls11EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendTls11EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableBackendSsl30() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBackendSsl30",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableBackendSsl30Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBackendSsl30Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableBackendTls10() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBackendTls10",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableBackendTls10Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBackendTls10Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableBackendTls11() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBackendTls11",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableBackendTls11Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBackendTls11Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableFrontendSsl30() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFrontendSsl30",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableFrontendSsl30Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFrontendSsl30Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableFrontendTls10() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFrontendTls10",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableFrontendTls10Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFrontendTls10Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableFrontendTls11() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFrontendTls11",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableFrontendTls11Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFrontendTls11Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) FrontendSsl30Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frontendSsl30Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) FrontendSsl30EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frontendSsl30EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) FrontendTls10Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frontendTls10Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) FrontendTls10EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frontendTls10EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) FrontendTls11Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frontendTls11Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) FrontendTls11EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frontendTls11EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) InternalValue() *ApiManagementSecurity {
	var returns *ApiManagementSecurity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheEcdsaWithAes128CbcShaCiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheEcdsaWithAes128CbcShaCiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheEcdsaWithAes256CbcShaCiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheEcdsaWithAes256CbcShaCiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheRsaWithAes128CbcShaCiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheRsaWithAes128CbcShaCiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheRsaWithAes128CbcShaCiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheRsaWithAes128CbcShaCiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheRsaWithAes256CbcShaCiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheRsaWithAes256CbcShaCiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheRsaWithAes256CbcShaCiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheRsaWithAes256CbcShaCiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes128CbcSha256CiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes128CbcSha256CiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes128CbcSha256CiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes128CbcSha256CiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes128CbcShaCiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes128CbcShaCiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes128CbcShaCiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes128CbcShaCiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes128GcmSha256CiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes128GcmSha256CiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes128GcmSha256CiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes128GcmSha256CiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes256CbcSha256CiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes256CbcSha256CiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes256CbcSha256CiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes256CbcSha256CiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes256CbcShaCiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes256CbcShaCiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes256CbcShaCiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes256CbcShaCiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes256GcmSha384CiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes256GcmSha384CiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes256GcmSha384CiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes256GcmSha384CiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TripleDesCiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tripleDesCiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TripleDesCiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tripleDesCiphersEnabledInput",
		&returns,
	)
	return returns
}


func NewApiManagementSecurityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementSecurityOutputReference {
	_init_.Initialize()

	if err := validateNewApiManagementSecurityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagementSecurityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementSecurityOutputReference_Override(a ApiManagementSecurityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetBackendSsl30Enabled(val interface{}) {
	if err := j.validateSetBackendSsl30EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendSsl30Enabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetBackendTls10Enabled(val interface{}) {
	if err := j.validateSetBackendTls10EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendTls10Enabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetBackendTls11Enabled(val interface{}) {
	if err := j.validateSetBackendTls11EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendTls11Enabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetEnableBackendSsl30(val interface{}) {
	if err := j.validateSetEnableBackendSsl30Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBackendSsl30",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetEnableBackendTls10(val interface{}) {
	if err := j.validateSetEnableBackendTls10Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBackendTls10",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetEnableBackendTls11(val interface{}) {
	if err := j.validateSetEnableBackendTls11Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBackendTls11",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetEnableFrontendSsl30(val interface{}) {
	if err := j.validateSetEnableFrontendSsl30Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableFrontendSsl30",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetEnableFrontendTls10(val interface{}) {
	if err := j.validateSetEnableFrontendTls10Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableFrontendTls10",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetEnableFrontendTls11(val interface{}) {
	if err := j.validateSetEnableFrontendTls11Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableFrontendTls11",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetFrontendSsl30Enabled(val interface{}) {
	if err := j.validateSetFrontendSsl30EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frontendSsl30Enabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetFrontendTls10Enabled(val interface{}) {
	if err := j.validateSetFrontendTls10EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frontendTls10Enabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetFrontendTls11Enabled(val interface{}) {
	if err := j.validateSetFrontendTls11EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frontendTls11Enabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetInternalValue(val *ApiManagementSecurity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetTlsEcdheEcdsaWithAes128CbcShaCiphersEnabled(val interface{}) {
	if err := j.validateSetTlsEcdheEcdsaWithAes128CbcShaCiphersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetTlsEcdheEcdsaWithAes256CbcShaCiphersEnabled(val interface{}) {
	if err := j.validateSetTlsEcdheEcdsaWithAes256CbcShaCiphersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetTlsEcdheRsaWithAes128CbcShaCiphersEnabled(val interface{}) {
	if err := j.validateSetTlsEcdheRsaWithAes128CbcShaCiphersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsEcdheRsaWithAes128CbcShaCiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetTlsEcdheRsaWithAes256CbcShaCiphersEnabled(val interface{}) {
	if err := j.validateSetTlsEcdheRsaWithAes256CbcShaCiphersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsEcdheRsaWithAes256CbcShaCiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetTlsRsaWithAes128CbcSha256CiphersEnabled(val interface{}) {
	if err := j.validateSetTlsRsaWithAes128CbcSha256CiphersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsRsaWithAes128CbcSha256CiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetTlsRsaWithAes128CbcShaCiphersEnabled(val interface{}) {
	if err := j.validateSetTlsRsaWithAes128CbcShaCiphersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsRsaWithAes128CbcShaCiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetTlsRsaWithAes128GcmSha256CiphersEnabled(val interface{}) {
	if err := j.validateSetTlsRsaWithAes128GcmSha256CiphersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsRsaWithAes128GcmSha256CiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetTlsRsaWithAes256CbcSha256CiphersEnabled(val interface{}) {
	if err := j.validateSetTlsRsaWithAes256CbcSha256CiphersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsRsaWithAes256CbcSha256CiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetTlsRsaWithAes256CbcShaCiphersEnabled(val interface{}) {
	if err := j.validateSetTlsRsaWithAes256CbcShaCiphersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsRsaWithAes256CbcShaCiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetTlsRsaWithAes256GcmSha384CiphersEnabled(val interface{}) {
	if err := j.validateSetTlsRsaWithAes256GcmSha384CiphersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsRsaWithAes256GcmSha384CiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference)SetTripleDesCiphersEnabled(val interface{}) {
	if err := j.validateSetTripleDesCiphersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tripleDesCiphersEnabled",
		val,
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetBackendSsl30Enabled() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendSsl30Enabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetBackendTls10Enabled() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendTls10Enabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetBackendTls11Enabled() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendTls11Enabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetEnableBackendSsl30() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBackendSsl30",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetEnableBackendTls10() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBackendTls10",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetEnableBackendTls11() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBackendTls11",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetEnableFrontendSsl30() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableFrontendSsl30",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetEnableFrontendTls10() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableFrontendTls10",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetEnableFrontendTls11() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableFrontendTls11",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetFrontendSsl30Enabled() {
	_jsii_.InvokeVoid(
		a,
		"resetFrontendSsl30Enabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetFrontendTls10Enabled() {
	_jsii_.InvokeVoid(
		a,
		"resetFrontendTls10Enabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetFrontendTls11Enabled() {
	_jsii_.InvokeVoid(
		a,
		"resetFrontendTls11Enabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsEcdheEcdsaWithAes128CbcShaCiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsEcdheEcdsaWithAes128CbcShaCiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsEcdheEcdsaWithAes256CbcShaCiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsEcdheEcdsaWithAes256CbcShaCiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsEcdheRsaWithAes128CbcShaCiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsEcdheRsaWithAes128CbcShaCiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsEcdheRsaWithAes256CbcShaCiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsEcdheRsaWithAes256CbcShaCiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsRsaWithAes128CbcSha256CiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsRsaWithAes128CbcSha256CiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsRsaWithAes128CbcShaCiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsRsaWithAes128CbcShaCiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsRsaWithAes128GcmSha256CiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsRsaWithAes128GcmSha256CiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsRsaWithAes256CbcSha256CiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsRsaWithAes256CbcSha256CiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsRsaWithAes256CbcShaCiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsRsaWithAes256CbcShaCiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsRsaWithAes256GcmSha384CiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsRsaWithAes256GcmSha384CiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTripleDesCiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTripleDesCiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

