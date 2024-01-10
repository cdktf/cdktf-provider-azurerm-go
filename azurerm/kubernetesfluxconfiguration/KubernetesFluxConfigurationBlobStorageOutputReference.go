// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesfluxconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/kubernetesfluxconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesFluxConfigurationBlobStorageOutputReference interface {
	cdktf.ComplexObject
	AccountKey() *string
	SetAccountKey(val *string)
	AccountKeyInput() *string
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
	ContainerId() *string
	SetContainerId(val *string)
	ContainerIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesFluxConfigurationBlobStorage
	SetInternalValue(val *KubernetesFluxConfigurationBlobStorage)
	LocalAuthReference() *string
	SetLocalAuthReference(val *string)
	LocalAuthReferenceInput() *string
	ManagedIdentity() KubernetesFluxConfigurationBlobStorageManagedIdentityOutputReference
	ManagedIdentityInput() *KubernetesFluxConfigurationBlobStorageManagedIdentity
	SasToken() *string
	SetSasToken(val *string)
	SasTokenInput() *string
	ServicePrincipal() KubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference
	ServicePrincipalInput() *KubernetesFluxConfigurationBlobStorageServicePrincipal
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
	PutManagedIdentity(value *KubernetesFluxConfigurationBlobStorageManagedIdentity)
	PutServicePrincipal(value *KubernetesFluxConfigurationBlobStorageServicePrincipal)
	ResetAccountKey()
	ResetLocalAuthReference()
	ResetManagedIdentity()
	ResetSasToken()
	ResetServicePrincipal()
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

// The jsii proxy struct for KubernetesFluxConfigurationBlobStorageOutputReference
type jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) AccountKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) AccountKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ContainerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ContainerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) InternalValue() *KubernetesFluxConfigurationBlobStorage {
	var returns *KubernetesFluxConfigurationBlobStorage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) LocalAuthReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAuthReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) LocalAuthReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAuthReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ManagedIdentity() KubernetesFluxConfigurationBlobStorageManagedIdentityOutputReference {
	var returns KubernetesFluxConfigurationBlobStorageManagedIdentityOutputReference
	_jsii_.Get(
		j,
		"managedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ManagedIdentityInput() *KubernetesFluxConfigurationBlobStorageManagedIdentity {
	var returns *KubernetesFluxConfigurationBlobStorageManagedIdentity
	_jsii_.Get(
		j,
		"managedIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) SasToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) SasTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ServicePrincipal() KubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference {
	var returns KubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference
	_jsii_.Get(
		j,
		"servicePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ServicePrincipalInput() *KubernetesFluxConfigurationBlobStorageServicePrincipal {
	var returns *KubernetesFluxConfigurationBlobStorageServicePrincipal
	_jsii_.Get(
		j,
		"servicePrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) SyncIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) SyncIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}


func NewKubernetesFluxConfigurationBlobStorageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesFluxConfigurationBlobStorageOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesFluxConfigurationBlobStorageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesFluxConfiguration.KubernetesFluxConfigurationBlobStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesFluxConfigurationBlobStorageOutputReference_Override(k KubernetesFluxConfigurationBlobStorageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesFluxConfiguration.KubernetesFluxConfigurationBlobStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference)SetAccountKey(val *string) {
	if err := j.validateSetAccountKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountKey",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference)SetContainerId(val *string) {
	if err := j.validateSetContainerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerId",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference)SetInternalValue(val *KubernetesFluxConfigurationBlobStorage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference)SetLocalAuthReference(val *string) {
	if err := j.validateSetLocalAuthReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAuthReference",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference)SetSasToken(val *string) {
	if err := j.validateSetSasTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sasToken",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference)SetSyncIntervalInSeconds(val *float64) {
	if err := j.validateSetSyncIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference)SetTimeoutInSeconds(val *float64) {
	if err := j.validateSetTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) PutManagedIdentity(value *KubernetesFluxConfigurationBlobStorageManagedIdentity) {
	if err := k.validatePutManagedIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putManagedIdentity",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) PutServicePrincipal(value *KubernetesFluxConfigurationBlobStorageServicePrincipal) {
	if err := k.validatePutServicePrincipalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putServicePrincipal",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ResetAccountKey() {
	_jsii_.InvokeVoid(
		k,
		"resetAccountKey",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ResetLocalAuthReference() {
	_jsii_.InvokeVoid(
		k,
		"resetLocalAuthReference",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ResetManagedIdentity() {
	_jsii_.InvokeVoid(
		k,
		"resetManagedIdentity",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ResetSasToken() {
	_jsii_.InvokeVoid(
		k,
		"resetSasToken",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ResetServicePrincipal() {
	_jsii_.InvokeVoid(
		k,
		"resetServicePrincipal",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ResetSyncIntervalInSeconds() {
	_jsii_.InvokeVoid(
		k,
		"resetSyncIntervalInSeconds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KubernetesFluxConfigurationBlobStorageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

