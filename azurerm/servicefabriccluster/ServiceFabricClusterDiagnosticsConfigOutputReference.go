// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabriccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/servicefabriccluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceFabricClusterDiagnosticsConfigOutputReference interface {
	cdktf.ComplexObject
	BlobEndpoint() *string
	SetBlobEndpoint(val *string)
	BlobEndpointInput() *string
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
	InternalValue() *ServiceFabricClusterDiagnosticsConfig
	SetInternalValue(val *ServiceFabricClusterDiagnosticsConfig)
	ProtectedAccountKeyName() *string
	SetProtectedAccountKeyName(val *string)
	ProtectedAccountKeyNameInput() *string
	QueueEndpoint() *string
	SetQueueEndpoint(val *string)
	QueueEndpointInput() *string
	StorageAccountName() *string
	SetStorageAccountName(val *string)
	StorageAccountNameInput() *string
	TableEndpoint() *string
	SetTableEndpoint(val *string)
	TableEndpointInput() *string
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

// The jsii proxy struct for ServiceFabricClusterDiagnosticsConfigOutputReference
type jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) BlobEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) BlobEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) InternalValue() *ServiceFabricClusterDiagnosticsConfig {
	var returns *ServiceFabricClusterDiagnosticsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) ProtectedAccountKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectedAccountKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) ProtectedAccountKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectedAccountKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) QueueEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) QueueEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) StorageAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) StorageAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) TableEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) TableEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterDiagnosticsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricClusterDiagnosticsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewServiceFabricClusterDiagnosticsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterDiagnosticsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricClusterDiagnosticsConfigOutputReference_Override(s ServiceFabricClusterDiagnosticsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterDiagnosticsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference)SetBlobEndpoint(val *string) {
	if err := j.validateSetBlobEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blobEndpoint",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference)SetInternalValue(val *ServiceFabricClusterDiagnosticsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference)SetProtectedAccountKeyName(val *string) {
	if err := j.validateSetProtectedAccountKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectedAccountKeyName",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference)SetQueueEndpoint(val *string) {
	if err := j.validateSetQueueEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueEndpoint",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference)SetStorageAccountName(val *string) {
	if err := j.validateSetStorageAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountName",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference)SetTableEndpoint(val *string) {
	if err := j.validateSetTableEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableEndpoint",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

