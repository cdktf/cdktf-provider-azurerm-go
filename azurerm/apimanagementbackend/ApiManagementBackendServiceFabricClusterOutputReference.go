// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/apimanagementbackend/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementBackendServiceFabricClusterOutputReference interface {
	cdktf.ComplexObject
	ClientCertificateId() *string
	SetClientCertificateId(val *string)
	ClientCertificateIdInput() *string
	ClientCertificateThumbprint() *string
	SetClientCertificateThumbprint(val *string)
	ClientCertificateThumbprintInput() *string
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
	InternalValue() *ApiManagementBackendServiceFabricCluster
	SetInternalValue(val *ApiManagementBackendServiceFabricCluster)
	ManagementEndpoints() *[]*string
	SetManagementEndpoints(val *[]*string)
	ManagementEndpointsInput() *[]*string
	MaxPartitionResolutionRetries() *float64
	SetMaxPartitionResolutionRetries(val *float64)
	MaxPartitionResolutionRetriesInput() *float64
	ServerCertificateThumbprints() *[]*string
	SetServerCertificateThumbprints(val *[]*string)
	ServerCertificateThumbprintsInput() *[]*string
	ServerX509Name() ApiManagementBackendServiceFabricClusterServerX509NameList
	ServerX509NameInput() interface{}
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
	PutServerX509Name(value interface{})
	ResetClientCertificateId()
	ResetClientCertificateThumbprint()
	ResetServerCertificateThumbprints()
	ResetServerX509Name()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementBackendServiceFabricClusterOutputReference
type jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ClientCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ClientCertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ClientCertificateThumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateThumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ClientCertificateThumbprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateThumbprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) InternalValue() *ApiManagementBackendServiceFabricCluster {
	var returns *ApiManagementBackendServiceFabricCluster
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ManagementEndpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"managementEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ManagementEndpointsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"managementEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) MaxPartitionResolutionRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPartitionResolutionRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) MaxPartitionResolutionRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPartitionResolutionRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ServerCertificateThumbprints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverCertificateThumbprints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ServerCertificateThumbprintsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverCertificateThumbprintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ServerX509Name() ApiManagementBackendServiceFabricClusterServerX509NameList {
	var returns ApiManagementBackendServiceFabricClusterServerX509NameList
	_jsii_.Get(
		j,
		"serverX509Name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ServerX509NameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverX509NameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiManagementBackendServiceFabricClusterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementBackendServiceFabricClusterOutputReference {
	_init_.Initialize()

	if err := validateNewApiManagementBackendServiceFabricClusterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementBackend.ApiManagementBackendServiceFabricClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementBackendServiceFabricClusterOutputReference_Override(a ApiManagementBackendServiceFabricClusterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementBackend.ApiManagementBackendServiceFabricClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference)SetClientCertificateId(val *string) {
	if err := j.validateSetClientCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference)SetClientCertificateThumbprint(val *string) {
	if err := j.validateSetClientCertificateThumbprintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateThumbprint",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference)SetInternalValue(val *ApiManagementBackendServiceFabricCluster) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference)SetManagementEndpoints(val *[]*string) {
	if err := j.validateSetManagementEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managementEndpoints",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference)SetMaxPartitionResolutionRetries(val *float64) {
	if err := j.validateSetMaxPartitionResolutionRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPartitionResolutionRetries",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference)SetServerCertificateThumbprints(val *[]*string) {
	if err := j.validateSetServerCertificateThumbprintsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCertificateThumbprints",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) PutServerX509Name(value interface{}) {
	if err := a.validatePutServerX509NameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServerX509Name",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ResetClientCertificateId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificateId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ResetClientCertificateThumbprint() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificateThumbprint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ResetServerCertificateThumbprints() {
	_jsii_.InvokeVoid(
		a,
		"resetServerCertificateThumbprints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ResetServerX509Name() {
	_jsii_.InvokeVoid(
		a,
		"resetServerX509Name",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendServiceFabricClusterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

