// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package machinelearningworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/machinelearningworkspace/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MachineLearningWorkspaceFeatureStoreOutputReference interface {
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
	ComputerSparkRuntimeVersion() *string
	SetComputerSparkRuntimeVersion(val *string)
	ComputerSparkRuntimeVersionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MachineLearningWorkspaceFeatureStore
	SetInternalValue(val *MachineLearningWorkspaceFeatureStore)
	OfflineConnectionName() *string
	SetOfflineConnectionName(val *string)
	OfflineConnectionNameInput() *string
	OnlineConnectionName() *string
	SetOnlineConnectionName(val *string)
	OnlineConnectionNameInput() *string
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
	ResetComputerSparkRuntimeVersion()
	ResetOfflineConnectionName()
	ResetOnlineConnectionName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MachineLearningWorkspaceFeatureStoreOutputReference
type jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) ComputerSparkRuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerSparkRuntimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) ComputerSparkRuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerSparkRuntimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) InternalValue() *MachineLearningWorkspaceFeatureStore {
	var returns *MachineLearningWorkspaceFeatureStore
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) OfflineConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offlineConnectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) OfflineConnectionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offlineConnectionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) OnlineConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onlineConnectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) OnlineConnectionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onlineConnectionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMachineLearningWorkspaceFeatureStoreOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MachineLearningWorkspaceFeatureStoreOutputReference {
	_init_.Initialize()

	if err := validateNewMachineLearningWorkspaceFeatureStoreOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.machineLearningWorkspace.MachineLearningWorkspaceFeatureStoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMachineLearningWorkspaceFeatureStoreOutputReference_Override(m MachineLearningWorkspaceFeatureStoreOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.machineLearningWorkspace.MachineLearningWorkspaceFeatureStoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference)SetComputerSparkRuntimeVersion(val *string) {
	if err := j.validateSetComputerSparkRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computerSparkRuntimeVersion",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference)SetInternalValue(val *MachineLearningWorkspaceFeatureStore) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference)SetOfflineConnectionName(val *string) {
	if err := j.validateSetOfflineConnectionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offlineConnectionName",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference)SetOnlineConnectionName(val *string) {
	if err := j.validateSetOnlineConnectionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onlineConnectionName",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) ResetComputerSparkRuntimeVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetComputerSparkRuntimeVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) ResetOfflineConnectionName() {
	_jsii_.InvokeVoid(
		m,
		"resetOfflineConnectionName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) ResetOnlineConnectionName() {
	_jsii_.InvokeVoid(
		m,
		"resetOnlineConnectionName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceFeatureStoreOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

