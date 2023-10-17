// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package trafficmanagerprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/trafficmanagerprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TrafficManagerProfileMonitorConfigOutputReference interface {
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
	CustomHeader() TrafficManagerProfileMonitorConfigCustomHeaderList
	CustomHeaderInput() interface{}
	ExpectedStatusCodeRanges() *[]*string
	SetExpectedStatusCodeRanges(val *[]*string)
	ExpectedStatusCodeRangesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *TrafficManagerProfileMonitorConfig
	SetInternalValue(val *TrafficManagerProfileMonitorConfig)
	IntervalInSeconds() *float64
	SetIntervalInSeconds(val *float64)
	IntervalInSecondsInput() *float64
	Path() *string
	SetPath(val *string)
	PathInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	ToleratedNumberOfFailures() *float64
	SetToleratedNumberOfFailures(val *float64)
	ToleratedNumberOfFailuresInput() *float64
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
	PutCustomHeader(value interface{})
	ResetCustomHeader()
	ResetExpectedStatusCodeRanges()
	ResetIntervalInSeconds()
	ResetPath()
	ResetTimeoutInSeconds()
	ResetToleratedNumberOfFailures()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TrafficManagerProfileMonitorConfigOutputReference
type jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) CustomHeader() TrafficManagerProfileMonitorConfigCustomHeaderList {
	var returns TrafficManagerProfileMonitorConfigCustomHeaderList
	_jsii_.Get(
		j,
		"customHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) CustomHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) ExpectedStatusCodeRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"expectedStatusCodeRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) ExpectedStatusCodeRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"expectedStatusCodeRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) InternalValue() *TrafficManagerProfileMonitorConfig {
	var returns *TrafficManagerProfileMonitorConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) IntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) IntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) ToleratedNumberOfFailures() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toleratedNumberOfFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) ToleratedNumberOfFailuresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toleratedNumberOfFailuresInput",
		&returns,
	)
	return returns
}


func NewTrafficManagerProfileMonitorConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TrafficManagerProfileMonitorConfigOutputReference {
	_init_.Initialize()

	if err := validateNewTrafficManagerProfileMonitorConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.trafficManagerProfile.TrafficManagerProfileMonitorConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTrafficManagerProfileMonitorConfigOutputReference_Override(t TrafficManagerProfileMonitorConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.trafficManagerProfile.TrafficManagerProfileMonitorConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference)SetExpectedStatusCodeRanges(val *[]*string) {
	if err := j.validateSetExpectedStatusCodeRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedStatusCodeRanges",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference)SetInternalValue(val *TrafficManagerProfileMonitorConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference)SetIntervalInSeconds(val *float64) {
	if err := j.validateSetIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference)SetTimeoutInSeconds(val *float64) {
	if err := j.validateSetTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference)SetToleratedNumberOfFailures(val *float64) {
	if err := j.validateSetToleratedNumberOfFailuresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toleratedNumberOfFailures",
		val,
	)
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) PutCustomHeader(value interface{}) {
	if err := t.validatePutCustomHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) ResetCustomHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) ResetExpectedStatusCodeRanges() {
	_jsii_.InvokeVoid(
		t,
		"resetExpectedStatusCodeRanges",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) ResetIntervalInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetIntervalInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		t,
		"resetPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) ResetToleratedNumberOfFailures() {
	_jsii_.InvokeVoid(
		t,
		"resetToleratedNumberOfFailures",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerProfileMonitorConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

