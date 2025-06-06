// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/kubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterServiceMeshProfileOutputReference interface {
	cdktf.ComplexObject
	CertificateAuthority() KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference
	CertificateAuthorityInput() *KubernetesClusterServiceMeshProfileCertificateAuthority
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
	ExternalIngressGatewayEnabled() interface{}
	SetExternalIngressGatewayEnabled(val interface{})
	ExternalIngressGatewayEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalIngressGatewayEnabled() interface{}
	SetInternalIngressGatewayEnabled(val interface{})
	InternalIngressGatewayEnabledInput() interface{}
	InternalValue() *KubernetesClusterServiceMeshProfile
	SetInternalValue(val *KubernetesClusterServiceMeshProfile)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Revisions() *[]*string
	SetRevisions(val *[]*string)
	RevisionsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutCertificateAuthority(value *KubernetesClusterServiceMeshProfileCertificateAuthority)
	ResetCertificateAuthority()
	ResetExternalIngressGatewayEnabled()
	ResetInternalIngressGatewayEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterServiceMeshProfileOutputReference
type jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) CertificateAuthority() KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference {
	var returns KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference
	_jsii_.Get(
		j,
		"certificateAuthority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) CertificateAuthorityInput() *KubernetesClusterServiceMeshProfileCertificateAuthority {
	var returns *KubernetesClusterServiceMeshProfileCertificateAuthority
	_jsii_.Get(
		j,
		"certificateAuthorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) ExternalIngressGatewayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalIngressGatewayEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) ExternalIngressGatewayEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalIngressGatewayEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) InternalIngressGatewayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalIngressGatewayEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) InternalIngressGatewayEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalIngressGatewayEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) InternalValue() *KubernetesClusterServiceMeshProfile {
	var returns *KubernetesClusterServiceMeshProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) Revisions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"revisions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) RevisionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"revisionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterServiceMeshProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterServiceMeshProfileOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterServiceMeshProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterServiceMeshProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterServiceMeshProfileOutputReference_Override(k KubernetesClusterServiceMeshProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterServiceMeshProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference)SetExternalIngressGatewayEnabled(val interface{}) {
	if err := j.validateSetExternalIngressGatewayEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalIngressGatewayEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference)SetInternalIngressGatewayEnabled(val interface{}) {
	if err := j.validateSetInternalIngressGatewayEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalIngressGatewayEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference)SetInternalValue(val *KubernetesClusterServiceMeshProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference)SetRevisions(val *[]*string) {
	if err := j.validateSetRevisionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revisions",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) PutCertificateAuthority(value *KubernetesClusterServiceMeshProfileCertificateAuthority) {
	if err := k.validatePutCertificateAuthorityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCertificateAuthority",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) ResetCertificateAuthority() {
	_jsii_.InvokeVoid(
		k,
		"resetCertificateAuthority",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) ResetExternalIngressGatewayEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetExternalIngressGatewayEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) ResetInternalIngressGatewayEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetInternalIngressGatewayEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

