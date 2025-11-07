// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/workloadssapthreetiervirtualinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference interface {
	cdktf.ComplexObject
	ApplicationServerConfiguration() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationApplicationServerConfigurationOutputReference
	ApplicationServerConfigurationInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationApplicationServerConfiguration
	AppResourceGroupName() *string
	SetAppResourceGroupName(val *string)
	AppResourceGroupNameInput() *string
	CentralServerConfiguration() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationCentralServerConfigurationOutputReference
	CentralServerConfigurationInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationCentralServerConfiguration
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
	DatabaseServerConfiguration() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference
	DatabaseServerConfigurationInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfiguration
	// Experimental.
	Fqn() *string
	HighAvailabilityType() *string
	SetHighAvailabilityType(val *string)
	HighAvailabilityTypeInput() *string
	InternalValue() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfiguration
	SetInternalValue(val *WorkloadsSapThreeTierVirtualInstanceThreeTierConfiguration)
	ResourceNames() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference
	ResourceNamesInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNames
	SecondaryIpEnabled() interface{}
	SetSecondaryIpEnabled(val interface{})
	SecondaryIpEnabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransportCreateAndMount() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationTransportCreateAndMountOutputReference
	TransportCreateAndMountInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationTransportCreateAndMount
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
	PutApplicationServerConfiguration(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationApplicationServerConfiguration)
	PutCentralServerConfiguration(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationCentralServerConfiguration)
	PutDatabaseServerConfiguration(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfiguration)
	PutResourceNames(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNames)
	PutTransportCreateAndMount(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationTransportCreateAndMount)
	ResetHighAvailabilityType()
	ResetResourceNames()
	ResetSecondaryIpEnabled()
	ResetTransportCreateAndMount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference
type jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) ApplicationServerConfiguration() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationApplicationServerConfigurationOutputReference {
	var returns WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationApplicationServerConfigurationOutputReference
	_jsii_.Get(
		j,
		"applicationServerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) ApplicationServerConfigurationInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationApplicationServerConfiguration {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationApplicationServerConfiguration
	_jsii_.Get(
		j,
		"applicationServerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) AppResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appResourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) AppResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appResourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) CentralServerConfiguration() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationCentralServerConfigurationOutputReference {
	var returns WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationCentralServerConfigurationOutputReference
	_jsii_.Get(
		j,
		"centralServerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) CentralServerConfigurationInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationCentralServerConfiguration {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationCentralServerConfiguration
	_jsii_.Get(
		j,
		"centralServerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) DatabaseServerConfiguration() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference {
	var returns WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference
	_jsii_.Get(
		j,
		"databaseServerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) DatabaseServerConfigurationInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfiguration {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfiguration
	_jsii_.Get(
		j,
		"databaseServerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) HighAvailabilityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"highAvailabilityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) HighAvailabilityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"highAvailabilityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) InternalValue() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfiguration {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) ResourceNames() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference {
	var returns WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference
	_jsii_.Get(
		j,
		"resourceNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) ResourceNamesInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNames {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNames
	_jsii_.Get(
		j,
		"resourceNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) SecondaryIpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryIpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) SecondaryIpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryIpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) TransportCreateAndMount() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationTransportCreateAndMountOutputReference {
	var returns WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationTransportCreateAndMountOutputReference
	_jsii_.Get(
		j,
		"transportCreateAndMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) TransportCreateAndMountInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationTransportCreateAndMount {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationTransportCreateAndMount
	_jsii_.Get(
		j,
		"transportCreateAndMountInput",
		&returns,
	)
	return returns
}


func NewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapThreeTierVirtualInstance.WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference_Override(w WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapThreeTierVirtualInstance.WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference)SetAppResourceGroupName(val *string) {
	if err := j.validateSetAppResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appResourceGroupName",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference)SetHighAvailabilityType(val *string) {
	if err := j.validateSetHighAvailabilityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"highAvailabilityType",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference)SetInternalValue(val *WorkloadsSapThreeTierVirtualInstanceThreeTierConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference)SetSecondaryIpEnabled(val interface{}) {
	if err := j.validateSetSecondaryIpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryIpEnabled",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) PutApplicationServerConfiguration(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationApplicationServerConfiguration) {
	if err := w.validatePutApplicationServerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putApplicationServerConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) PutCentralServerConfiguration(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationCentralServerConfiguration) {
	if err := w.validatePutCentralServerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCentralServerConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) PutDatabaseServerConfiguration(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfiguration) {
	if err := w.validatePutDatabaseServerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putDatabaseServerConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) PutResourceNames(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNames) {
	if err := w.validatePutResourceNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putResourceNames",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) PutTransportCreateAndMount(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationTransportCreateAndMount) {
	if err := w.validatePutTransportCreateAndMountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTransportCreateAndMount",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) ResetHighAvailabilityType() {
	_jsii_.InvokeVoid(
		w,
		"resetHighAvailabilityType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) ResetResourceNames() {
	_jsii_.InvokeVoid(
		w,
		"resetResourceNames",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) ResetSecondaryIpEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetSecondaryIpEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) ResetTransportCreateAndMount() {
	_jsii_.InvokeVoid(
		w,
		"resetTransportCreateAndMount",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

