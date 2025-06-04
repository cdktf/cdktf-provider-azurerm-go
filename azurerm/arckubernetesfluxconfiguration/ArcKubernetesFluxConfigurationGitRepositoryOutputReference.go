// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arckubernetesfluxconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/arckubernetesfluxconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArcKubernetesFluxConfigurationGitRepositoryOutputReference interface {
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
	HttpsCaCertBase64() *string
	SetHttpsCaCertBase64(val *string)
	HttpsCaCertBase64Input() *string
	HttpsKeyBase64() *string
	SetHttpsKeyBase64(val *string)
	HttpsKeyBase64Input() *string
	HttpsUser() *string
	SetHttpsUser(val *string)
	HttpsUserInput() *string
	InternalValue() *ArcKubernetesFluxConfigurationGitRepository
	SetInternalValue(val *ArcKubernetesFluxConfigurationGitRepository)
	LocalAuthReference() *string
	SetLocalAuthReference(val *string)
	LocalAuthReferenceInput() *string
	ReferenceType() *string
	SetReferenceType(val *string)
	ReferenceTypeInput() *string
	ReferenceValue() *string
	SetReferenceValue(val *string)
	ReferenceValueInput() *string
	SshKnownHostsBase64() *string
	SetSshKnownHostsBase64(val *string)
	SshKnownHostsBase64Input() *string
	SshPrivateKeyBase64() *string
	SetSshPrivateKeyBase64(val *string)
	SshPrivateKeyBase64Input() *string
	SyncIntervalInSeconds() *float64
	SetSyncIntervalInSeconds(val *float64)
	SyncIntervalInSecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutInSeconds() *float64
	SetTimeoutInSeconds(val *float64)
	TimeoutInSecondsInput() *float64
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	ResetHttpsCaCertBase64()
	ResetHttpsKeyBase64()
	ResetHttpsUser()
	ResetLocalAuthReference()
	ResetSshKnownHostsBase64()
	ResetSshPrivateKeyBase64()
	ResetSyncIntervalInSeconds()
	ResetTimeoutInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ArcKubernetesFluxConfigurationGitRepositoryOutputReference
type jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) HttpsCaCertBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsCaCertBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) HttpsCaCertBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsCaCertBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) HttpsKeyBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsKeyBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) HttpsKeyBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsKeyBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) HttpsUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) HttpsUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) InternalValue() *ArcKubernetesFluxConfigurationGitRepository {
	var returns *ArcKubernetesFluxConfigurationGitRepository
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) LocalAuthReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAuthReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) LocalAuthReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAuthReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ReferenceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ReferenceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ReferenceValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ReferenceValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) SshKnownHostsBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshKnownHostsBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) SshKnownHostsBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshKnownHostsBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) SshPrivateKeyBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPrivateKeyBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) SshPrivateKeyBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPrivateKeyBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) SyncIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) SyncIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewArcKubernetesFluxConfigurationGitRepositoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ArcKubernetesFluxConfigurationGitRepositoryOutputReference {
	_init_.Initialize()

	if err := validateNewArcKubernetesFluxConfigurationGitRepositoryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.arcKubernetesFluxConfiguration.ArcKubernetesFluxConfigurationGitRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewArcKubernetesFluxConfigurationGitRepositoryOutputReference_Override(a ArcKubernetesFluxConfigurationGitRepositoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.arcKubernetesFluxConfiguration.ArcKubernetesFluxConfigurationGitRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetHttpsCaCertBase64(val *string) {
	if err := j.validateSetHttpsCaCertBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsCaCertBase64",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetHttpsKeyBase64(val *string) {
	if err := j.validateSetHttpsKeyBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsKeyBase64",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetHttpsUser(val *string) {
	if err := j.validateSetHttpsUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsUser",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetInternalValue(val *ArcKubernetesFluxConfigurationGitRepository) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetLocalAuthReference(val *string) {
	if err := j.validateSetLocalAuthReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAuthReference",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetReferenceType(val *string) {
	if err := j.validateSetReferenceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceType",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetReferenceValue(val *string) {
	if err := j.validateSetReferenceValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceValue",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetSshKnownHostsBase64(val *string) {
	if err := j.validateSetSshKnownHostsBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKnownHostsBase64",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetSshPrivateKeyBase64(val *string) {
	if err := j.validateSetSshPrivateKeyBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPrivateKeyBase64",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetSyncIntervalInSeconds(val *float64) {
	if err := j.validateSetSyncIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetTimeoutInSeconds(val *float64) {
	if err := j.validateSetTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ResetHttpsCaCertBase64() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpsCaCertBase64",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ResetHttpsKeyBase64() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpsKeyBase64",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ResetHttpsUser() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpsUser",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ResetLocalAuthReference() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalAuthReference",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ResetSshKnownHostsBase64() {
	_jsii_.InvokeVoid(
		a,
		"resetSshKnownHostsBase64",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ResetSshPrivateKeyBase64() {
	_jsii_.InvokeVoid(
		a,
		"resetSshPrivateKeyBase64",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ResetSyncIntervalInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetSyncIntervalInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationGitRepositoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

