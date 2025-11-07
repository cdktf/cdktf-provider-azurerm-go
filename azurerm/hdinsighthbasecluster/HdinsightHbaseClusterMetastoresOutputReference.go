// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthbasecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/hdinsighthbasecluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HdinsightHbaseClusterMetastoresOutputReference interface {
	cdktf.ComplexObject
	Ambari() HdinsightHbaseClusterMetastoresAmbariOutputReference
	AmbariInput() *HdinsightHbaseClusterMetastoresAmbari
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
	Hive() HdinsightHbaseClusterMetastoresHiveOutputReference
	HiveInput() *HdinsightHbaseClusterMetastoresHive
	InternalValue() *HdinsightHbaseClusterMetastores
	SetInternalValue(val *HdinsightHbaseClusterMetastores)
	Oozie() HdinsightHbaseClusterMetastoresOozieOutputReference
	OozieInput() *HdinsightHbaseClusterMetastoresOozie
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutAmbari(value *HdinsightHbaseClusterMetastoresAmbari)
	PutHive(value *HdinsightHbaseClusterMetastoresHive)
	PutOozie(value *HdinsightHbaseClusterMetastoresOozie)
	ResetAmbari()
	ResetHive()
	ResetOozie()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHbaseClusterMetastoresOutputReference
type jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) Ambari() HdinsightHbaseClusterMetastoresAmbariOutputReference {
	var returns HdinsightHbaseClusterMetastoresAmbariOutputReference
	_jsii_.Get(
		j,
		"ambari",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) AmbariInput() *HdinsightHbaseClusterMetastoresAmbari {
	var returns *HdinsightHbaseClusterMetastoresAmbari
	_jsii_.Get(
		j,
		"ambariInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) Hive() HdinsightHbaseClusterMetastoresHiveOutputReference {
	var returns HdinsightHbaseClusterMetastoresHiveOutputReference
	_jsii_.Get(
		j,
		"hive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) HiveInput() *HdinsightHbaseClusterMetastoresHive {
	var returns *HdinsightHbaseClusterMetastoresHive
	_jsii_.Get(
		j,
		"hiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) InternalValue() *HdinsightHbaseClusterMetastores {
	var returns *HdinsightHbaseClusterMetastores
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) Oozie() HdinsightHbaseClusterMetastoresOozieOutputReference {
	var returns HdinsightHbaseClusterMetastoresOozieOutputReference
	_jsii_.Get(
		j,
		"oozie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) OozieInput() *HdinsightHbaseClusterMetastoresOozie {
	var returns *HdinsightHbaseClusterMetastoresOozie
	_jsii_.Get(
		j,
		"oozieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightHbaseClusterMetastoresOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHbaseClusterMetastoresOutputReference {
	_init_.Initialize()

	if err := validateNewHdinsightHbaseClusterMetastoresOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHbaseCluster.HdinsightHbaseClusterMetastoresOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHbaseClusterMetastoresOutputReference_Override(h HdinsightHbaseClusterMetastoresOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHbaseCluster.HdinsightHbaseClusterMetastoresOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference)SetInternalValue(val *HdinsightHbaseClusterMetastores) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) PutAmbari(value *HdinsightHbaseClusterMetastoresAmbari) {
	if err := h.validatePutAmbariParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putAmbari",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) PutHive(value *HdinsightHbaseClusterMetastoresHive) {
	if err := h.validatePutHiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putHive",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) PutOozie(value *HdinsightHbaseClusterMetastoresOozie) {
	if err := h.validatePutOozieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putOozie",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) ResetAmbari() {
	_jsii_.InvokeVoid(
		h,
		"resetAmbari",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) ResetHive() {
	_jsii_.InvokeVoid(
		h,
		"resetHive",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) ResetOozie() {
	_jsii_.InvokeVoid(
		h,
		"resetOozie",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterMetastoresOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

