// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/workloadssapthreetiervirtualinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference interface {
	cdktf.ComplexObject
	ApplicationServer() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerOutputReference
	ApplicationServerInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServer
	CentralServer() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerOutputReference
	CentralServerInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServer
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
	DatabaseServer() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference
	DatabaseServerInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServer
	// Experimental.
	Fqn() *string
	InternalValue() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNames
	SetInternalValue(val *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNames)
	SharedStorage() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesSharedStorageOutputReference
	SharedStorageInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesSharedStorage
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
	PutApplicationServer(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServer)
	PutCentralServer(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServer)
	PutDatabaseServer(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServer)
	PutSharedStorage(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesSharedStorage)
	ResetApplicationServer()
	ResetCentralServer()
	ResetDatabaseServer()
	ResetSharedStorage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference
type jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) ApplicationServer() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerOutputReference {
	var returns WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerOutputReference
	_jsii_.Get(
		j,
		"applicationServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) ApplicationServerInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServer {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServer
	_jsii_.Get(
		j,
		"applicationServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) CentralServer() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerOutputReference {
	var returns WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerOutputReference
	_jsii_.Get(
		j,
		"centralServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) CentralServerInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServer {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServer
	_jsii_.Get(
		j,
		"centralServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) DatabaseServer() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference {
	var returns WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference
	_jsii_.Get(
		j,
		"databaseServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) DatabaseServerInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServer {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServer
	_jsii_.Get(
		j,
		"databaseServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) InternalValue() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNames {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNames
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) SharedStorage() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesSharedStorageOutputReference {
	var returns WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesSharedStorageOutputReference
	_jsii_.Get(
		j,
		"sharedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) SharedStorageInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesSharedStorage {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesSharedStorage
	_jsii_.Get(
		j,
		"sharedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference {
	_init_.Initialize()

	if err := validateNewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapThreeTierVirtualInstance.WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference_Override(w WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapThreeTierVirtualInstance.WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference)SetInternalValue(val *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNames) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) PutApplicationServer(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServer) {
	if err := w.validatePutApplicationServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putApplicationServer",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) PutCentralServer(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServer) {
	if err := w.validatePutCentralServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCentralServer",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) PutDatabaseServer(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServer) {
	if err := w.validatePutDatabaseServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putDatabaseServer",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) PutSharedStorage(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesSharedStorage) {
	if err := w.validatePutSharedStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSharedStorage",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) ResetApplicationServer() {
	_jsii_.InvokeVoid(
		w,
		"resetApplicationServer",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) ResetCentralServer() {
	_jsii_.InvokeVoid(
		w,
		"resetCentralServer",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) ResetDatabaseServer() {
	_jsii_.InvokeVoid(
		w,
		"resetDatabaseServer",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) ResetSharedStorage() {
	_jsii_.InvokeVoid(
		w,
		"resetSharedStorage",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

