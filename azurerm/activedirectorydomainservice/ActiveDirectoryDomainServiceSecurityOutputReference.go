// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package activedirectorydomainservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/activedirectorydomainservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActiveDirectoryDomainServiceSecurityOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *ActiveDirectoryDomainServiceSecurity
	SetInternalValue(val *ActiveDirectoryDomainServiceSecurity)
	KerberosArmoringEnabled() interface{}
	SetKerberosArmoringEnabled(val interface{})
	KerberosArmoringEnabledInput() interface{}
	KerberosRc4EncryptionEnabled() interface{}
	SetKerberosRc4EncryptionEnabled(val interface{})
	KerberosRc4EncryptionEnabledInput() interface{}
	NtlmV1Enabled() interface{}
	SetNtlmV1Enabled(val interface{})
	NtlmV1EnabledInput() interface{}
	SyncKerberosPasswords() interface{}
	SetSyncKerberosPasswords(val interface{})
	SyncKerberosPasswordsInput() interface{}
	SyncNtlmPasswords() interface{}
	SetSyncNtlmPasswords(val interface{})
	SyncNtlmPasswordsInput() interface{}
	SyncOnPremPasswords() interface{}
	SetSyncOnPremPasswords(val interface{})
	SyncOnPremPasswordsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsV1Enabled() interface{}
	SetTlsV1Enabled(val interface{})
	TlsV1EnabledInput() interface{}
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
	ResetKerberosArmoringEnabled()
	ResetKerberosRc4EncryptionEnabled()
	ResetNtlmV1Enabled()
	ResetSyncKerberosPasswords()
	ResetSyncNtlmPasswords()
	ResetSyncOnPremPasswords()
	ResetTlsV1Enabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ActiveDirectoryDomainServiceSecurityOutputReference
type jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) InternalValue() *ActiveDirectoryDomainServiceSecurity {
	var returns *ActiveDirectoryDomainServiceSecurity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) KerberosArmoringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberosArmoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) KerberosArmoringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberosArmoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) KerberosRc4EncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberosRc4EncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) KerberosRc4EncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberosRc4EncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) NtlmV1Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ntlmV1Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) NtlmV1EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ntlmV1EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SyncKerberosPasswords() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncKerberosPasswords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SyncKerberosPasswordsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncKerberosPasswordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SyncNtlmPasswords() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncNtlmPasswords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SyncNtlmPasswordsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncNtlmPasswordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SyncOnPremPasswords() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncOnPremPasswords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SyncOnPremPasswordsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncOnPremPasswordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) TlsV1Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsV1Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) TlsV1EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsV1EnabledInput",
		&returns,
	)
	return returns
}


func NewActiveDirectoryDomainServiceSecurityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ActiveDirectoryDomainServiceSecurityOutputReference {
	_init_.Initialize()

	if err := validateNewActiveDirectoryDomainServiceSecurityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainServiceSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewActiveDirectoryDomainServiceSecurityOutputReference_Override(a ActiveDirectoryDomainServiceSecurityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainServiceSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference)SetInternalValue(val *ActiveDirectoryDomainServiceSecurity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference)SetKerberosArmoringEnabled(val interface{}) {
	if err := j.validateSetKerberosArmoringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosArmoringEnabled",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference)SetKerberosRc4EncryptionEnabled(val interface{}) {
	if err := j.validateSetKerberosRc4EncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosRc4EncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference)SetNtlmV1Enabled(val interface{}) {
	if err := j.validateSetNtlmV1EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ntlmV1Enabled",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference)SetSyncKerberosPasswords(val interface{}) {
	if err := j.validateSetSyncKerberosPasswordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncKerberosPasswords",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference)SetSyncNtlmPasswords(val interface{}) {
	if err := j.validateSetSyncNtlmPasswordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncNtlmPasswords",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference)SetSyncOnPremPasswords(val interface{}) {
	if err := j.validateSetSyncOnPremPasswordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncOnPremPasswords",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference)SetTlsV1Enabled(val interface{}) {
	if err := j.validateSetTlsV1EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsV1Enabled",
		val,
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ResetKerberosArmoringEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetKerberosArmoringEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ResetKerberosRc4EncryptionEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetKerberosRc4EncryptionEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ResetNtlmV1Enabled() {
	_jsii_.InvokeVoid(
		a,
		"resetNtlmV1Enabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ResetSyncKerberosPasswords() {
	_jsii_.InvokeVoid(
		a,
		"resetSyncKerberosPasswords",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ResetSyncNtlmPasswords() {
	_jsii_.InvokeVoid(
		a,
		"resetSyncNtlmPasswords",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ResetSyncOnPremPasswords() {
	_jsii_.InvokeVoid(
		a,
		"resetSyncOnPremPasswords",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ResetTlsV1Enabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsV1Enabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

