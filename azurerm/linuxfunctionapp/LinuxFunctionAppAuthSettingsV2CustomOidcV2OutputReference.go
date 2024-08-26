// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/linuxfunctionapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference interface {
	cdktf.ComplexObject
	AuthorisationEndpoint() *string
	CertificationUri() *string
	ClientCredentialMethod() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecretSettingName() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IssuerEndpoint() *string
	Name() *string
	SetName(val *string)
	NameClaimType() *string
	SetNameClaimType(val *string)
	NameClaimTypeInput() *string
	NameInput() *string
	OpenidConfigurationEndpoint() *string
	SetOpenidConfigurationEndpoint(val *string)
	OpenidConfigurationEndpointInput() *string
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenEndpoint() *string
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
	ResetNameClaimType()
	ResetScopes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference
type jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) AuthorisationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorisationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) CertificationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) ClientCredentialMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCredentialMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) ClientSecretSettingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretSettingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) IssuerEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) NameClaimType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameClaimType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) NameClaimTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameClaimTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) OpenidConfigurationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openidConfigurationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) OpenidConfigurationEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openidConfigurationEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) TokenEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpoint",
		&returns,
	)
	return returns
}


func NewLinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference {
	_init_.Initialize()

	if err := validateNewLinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxFunctionApp.LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference_Override(l LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxFunctionApp.LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference)SetNameClaimType(val *string) {
	if err := j.validateSetNameClaimTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameClaimType",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference)SetOpenidConfigurationEndpoint(val *string) {
	if err := j.validateSetOpenidConfigurationEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openidConfigurationEndpoint",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) ResetNameClaimType() {
	_jsii_.InvokeVoid(
		l,
		"resetNameClaimType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) ResetScopes() {
	_jsii_.InvokeVoid(
		l,
		"resetScopes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2CustomOidcV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

