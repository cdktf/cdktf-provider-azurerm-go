// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/kubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference interface {
	cdktf.ComplexObject
	CertChainObjectName() *string
	SetCertChainObjectName(val *string)
	CertChainObjectNameInput() *string
	CertObjectName() *string
	SetCertObjectName(val *string)
	CertObjectNameInput() *string
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
	InternalValue() *KubernetesClusterServiceMeshProfileCertificateAuthority
	SetInternalValue(val *KubernetesClusterServiceMeshProfileCertificateAuthority)
	KeyObjectName() *string
	SetKeyObjectName(val *string)
	KeyObjectNameInput() *string
	KeyVaultId() *string
	SetKeyVaultId(val *string)
	KeyVaultIdInput() *string
	RootCertObjectName() *string
	SetRootCertObjectName(val *string)
	RootCertObjectNameInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference
type jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) CertChainObjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certChainObjectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) CertChainObjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certChainObjectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) CertObjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certObjectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) CertObjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certObjectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) InternalValue() *KubernetesClusterServiceMeshProfileCertificateAuthority {
	var returns *KubernetesClusterServiceMeshProfileCertificateAuthority
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) KeyObjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyObjectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) KeyObjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyObjectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) KeyVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) KeyVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) RootCertObjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootCertObjectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) RootCertObjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootCertObjectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterServiceMeshProfileCertificateAuthorityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference_Override(k KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference)SetCertChainObjectName(val *string) {
	if err := j.validateSetCertChainObjectNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certChainObjectName",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference)SetCertObjectName(val *string) {
	if err := j.validateSetCertObjectNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certObjectName",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference)SetInternalValue(val *KubernetesClusterServiceMeshProfileCertificateAuthority) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference)SetKeyObjectName(val *string) {
	if err := j.validateSetKeyObjectNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyObjectName",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference)SetKeyVaultId(val *string) {
	if err := j.validateSetKeyVaultIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference)SetRootCertObjectName(val *string) {
	if err := j.validateSetRootCertObjectNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootCertObjectName",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KubernetesClusterServiceMeshProfileCertificateAuthorityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

