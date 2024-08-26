// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesfluxconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/kubernetesfluxconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesFluxConfigurationBucketOutputReference interface {
	cdktf.ComplexObject
	AccessKey() *string
	SetAccessKey(val *string)
	AccessKeyInput() *string
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
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
	InternalValue() *KubernetesFluxConfigurationBucket
	SetInternalValue(val *KubernetesFluxConfigurationBucket)
	LocalAuthReference() *string
	SetLocalAuthReference(val *string)
	LocalAuthReferenceInput() *string
	SecretKeyBase64() *string
	SetSecretKeyBase64(val *string)
	SecretKeyBase64Input() *string
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
	TlsEnabled() interface{}
	SetTlsEnabled(val interface{})
	TlsEnabledInput() interface{}
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
	ResetAccessKey()
	ResetLocalAuthReference()
	ResetSecretKeyBase64()
	ResetSyncIntervalInSeconds()
	ResetTimeoutInSeconds()
	ResetTlsEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesFluxConfigurationBucketOutputReference
type jsiiProxy_KubernetesFluxConfigurationBucketOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) InternalValue() *KubernetesFluxConfigurationBucket {
	var returns *KubernetesFluxConfigurationBucket
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) LocalAuthReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAuthReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) LocalAuthReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAuthReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) SecretKeyBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) SecretKeyBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) SyncIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) SyncIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) TlsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) TlsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewKubernetesFluxConfigurationBucketOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesFluxConfigurationBucketOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesFluxConfigurationBucketOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesFluxConfigurationBucketOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesFluxConfiguration.KubernetesFluxConfigurationBucketOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesFluxConfigurationBucketOutputReference_Override(k KubernetesFluxConfigurationBucketOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesFluxConfiguration.KubernetesFluxConfigurationBucketOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference)SetAccessKey(val *string) {
	if err := j.validateSetAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference)SetInternalValue(val *KubernetesFluxConfigurationBucket) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference)SetLocalAuthReference(val *string) {
	if err := j.validateSetLocalAuthReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAuthReference",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference)SetSecretKeyBase64(val *string) {
	if err := j.validateSetSecretKeyBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretKeyBase64",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference)SetSyncIntervalInSeconds(val *float64) {
	if err := j.validateSetSyncIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference)SetTimeoutInSeconds(val *float64) {
	if err := j.validateSetTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference)SetTlsEnabled(val interface{}) {
	if err := j.validateSetTlsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) ResetAccessKey() {
	_jsii_.InvokeVoid(
		k,
		"resetAccessKey",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) ResetLocalAuthReference() {
	_jsii_.InvokeVoid(
		k,
		"resetLocalAuthReference",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) ResetSecretKeyBase64() {
	_jsii_.InvokeVoid(
		k,
		"resetSecretKeyBase64",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) ResetSyncIntervalInSeconds() {
	_jsii_.InvokeVoid(
		k,
		"resetSyncIntervalInSeconds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) ResetTlsEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetTlsEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationBucketOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

