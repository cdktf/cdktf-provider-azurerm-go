// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesfleetupdaterun

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/kubernetesfleetupdaterun/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesFleetUpdateRunManagedClusterUpdateOutputReference interface {
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
	InternalValue() *KubernetesFleetUpdateRunManagedClusterUpdate
	SetInternalValue(val *KubernetesFleetUpdateRunManagedClusterUpdate)
	NodeImageSelection() KubernetesFleetUpdateRunManagedClusterUpdateNodeImageSelectionOutputReference
	NodeImageSelectionInput() *KubernetesFleetUpdateRunManagedClusterUpdateNodeImageSelection
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Upgrade() KubernetesFleetUpdateRunManagedClusterUpdateUpgradeOutputReference
	UpgradeInput() *KubernetesFleetUpdateRunManagedClusterUpdateUpgrade
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
	PutNodeImageSelection(value *KubernetesFleetUpdateRunManagedClusterUpdateNodeImageSelection)
	PutUpgrade(value *KubernetesFleetUpdateRunManagedClusterUpdateUpgrade)
	ResetNodeImageSelection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesFleetUpdateRunManagedClusterUpdateOutputReference
type jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) InternalValue() *KubernetesFleetUpdateRunManagedClusterUpdate {
	var returns *KubernetesFleetUpdateRunManagedClusterUpdate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) NodeImageSelection() KubernetesFleetUpdateRunManagedClusterUpdateNodeImageSelectionOutputReference {
	var returns KubernetesFleetUpdateRunManagedClusterUpdateNodeImageSelectionOutputReference
	_jsii_.Get(
		j,
		"nodeImageSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) NodeImageSelectionInput() *KubernetesFleetUpdateRunManagedClusterUpdateNodeImageSelection {
	var returns *KubernetesFleetUpdateRunManagedClusterUpdateNodeImageSelection
	_jsii_.Get(
		j,
		"nodeImageSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) Upgrade() KubernetesFleetUpdateRunManagedClusterUpdateUpgradeOutputReference {
	var returns KubernetesFleetUpdateRunManagedClusterUpdateUpgradeOutputReference
	_jsii_.Get(
		j,
		"upgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) UpgradeInput() *KubernetesFleetUpdateRunManagedClusterUpdateUpgrade {
	var returns *KubernetesFleetUpdateRunManagedClusterUpdateUpgrade
	_jsii_.Get(
		j,
		"upgradeInput",
		&returns,
	)
	return returns
}


func NewKubernetesFleetUpdateRunManagedClusterUpdateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesFleetUpdateRunManagedClusterUpdateOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesFleetUpdateRunManagedClusterUpdateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesFleetUpdateRun.KubernetesFleetUpdateRunManagedClusterUpdateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesFleetUpdateRunManagedClusterUpdateOutputReference_Override(k KubernetesFleetUpdateRunManagedClusterUpdateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesFleetUpdateRun.KubernetesFleetUpdateRunManagedClusterUpdateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference)SetInternalValue(val *KubernetesFleetUpdateRunManagedClusterUpdate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) PutNodeImageSelection(value *KubernetesFleetUpdateRunManagedClusterUpdateNodeImageSelection) {
	if err := k.validatePutNodeImageSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putNodeImageSelection",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) PutUpgrade(value *KubernetesFleetUpdateRunManagedClusterUpdateUpgrade) {
	if err := k.validatePutUpgradeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putUpgrade",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) ResetNodeImageSelection() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeImageSelection",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFleetUpdateRunManagedClusterUpdateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

