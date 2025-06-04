// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkpacketcorecontrolplane

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/mobilenetworkpacketcorecontrolplane/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MobileNetworkPacketCoreControlPlanePlatformOutputReference interface {
	cdktf.ComplexObject
	ArcKubernetesClusterId() *string
	SetArcKubernetesClusterId(val *string)
	ArcKubernetesClusterIdInput() *string
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
	CustomLocationId() *string
	SetCustomLocationId(val *string)
	CustomLocationIdInput() *string
	EdgeDeviceId() *string
	SetEdgeDeviceId(val *string)
	EdgeDeviceIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MobileNetworkPacketCoreControlPlanePlatform
	SetInternalValue(val *MobileNetworkPacketCoreControlPlanePlatform)
	StackHciClusterId() *string
	SetStackHciClusterId(val *string)
	StackHciClusterIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetArcKubernetesClusterId()
	ResetCustomLocationId()
	ResetEdgeDeviceId()
	ResetStackHciClusterId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MobileNetworkPacketCoreControlPlanePlatformOutputReference
type jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) ArcKubernetesClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arcKubernetesClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) ArcKubernetesClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arcKubernetesClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) CustomLocationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLocationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) CustomLocationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLocationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) EdgeDeviceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeDeviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) EdgeDeviceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeDeviceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) InternalValue() *MobileNetworkPacketCoreControlPlanePlatform {
	var returns *MobileNetworkPacketCoreControlPlanePlatform
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) StackHciClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackHciClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) StackHciClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackHciClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewMobileNetworkPacketCoreControlPlanePlatformOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MobileNetworkPacketCoreControlPlanePlatformOutputReference {
	_init_.Initialize()

	if err := validateNewMobileNetworkPacketCoreControlPlanePlatformOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreControlPlane.MobileNetworkPacketCoreControlPlanePlatformOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMobileNetworkPacketCoreControlPlanePlatformOutputReference_Override(m MobileNetworkPacketCoreControlPlanePlatformOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreControlPlane.MobileNetworkPacketCoreControlPlanePlatformOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference)SetArcKubernetesClusterId(val *string) {
	if err := j.validateSetArcKubernetesClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arcKubernetesClusterId",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference)SetCustomLocationId(val *string) {
	if err := j.validateSetCustomLocationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLocationId",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference)SetEdgeDeviceId(val *string) {
	if err := j.validateSetEdgeDeviceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeDeviceId",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference)SetInternalValue(val *MobileNetworkPacketCoreControlPlanePlatform) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference)SetStackHciClusterId(val *string) {
	if err := j.validateSetStackHciClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackHciClusterId",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) ResetArcKubernetesClusterId() {
	_jsii_.InvokeVoid(
		m,
		"resetArcKubernetesClusterId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) ResetCustomLocationId() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomLocationId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) ResetEdgeDeviceId() {
	_jsii_.InvokeVoid(
		m,
		"resetEdgeDeviceId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) ResetStackHciClusterId() {
	_jsii_.InvokeVoid(
		m,
		"resetStackHciClusterId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlanePlatformOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

