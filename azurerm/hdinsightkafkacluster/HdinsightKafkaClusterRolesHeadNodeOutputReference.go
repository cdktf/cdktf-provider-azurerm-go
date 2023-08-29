// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightkafkacluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/hdinsightkafkacluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HdinsightKafkaClusterRolesHeadNodeOutputReference interface {
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
	InternalValue() *HdinsightKafkaClusterRolesHeadNode
	SetInternalValue(val *HdinsightKafkaClusterRolesHeadNode)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	ScriptActions() HdinsightKafkaClusterRolesHeadNodeScriptActionsList
	ScriptActionsInput() interface{}
	SshKeys() *[]*string
	SetSshKeys(val *[]*string)
	SshKeysInput() *[]*string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	VirtualNetworkId() *string
	SetVirtualNetworkId(val *string)
	VirtualNetworkIdInput() *string
	VmSize() *string
	SetVmSize(val *string)
	VmSizeInput() *string
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
	PutScriptActions(value interface{})
	ResetPassword()
	ResetScriptActions()
	ResetSshKeys()
	ResetSubnetId()
	ResetVirtualNetworkId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightKafkaClusterRolesHeadNodeOutputReference
type jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) InternalValue() *HdinsightKafkaClusterRolesHeadNode {
	var returns *HdinsightKafkaClusterRolesHeadNode
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) ScriptActions() HdinsightKafkaClusterRolesHeadNodeScriptActionsList {
	var returns HdinsightKafkaClusterRolesHeadNodeScriptActionsList
	_jsii_.Get(
		j,
		"scriptActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) ScriptActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scriptActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) SshKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) SshKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}


func NewHdinsightKafkaClusterRolesHeadNodeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightKafkaClusterRolesHeadNodeOutputReference {
	_init_.Initialize()

	if err := validateNewHdinsightKafkaClusterRolesHeadNodeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightKafkaCluster.HdinsightKafkaClusterRolesHeadNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightKafkaClusterRolesHeadNodeOutputReference_Override(h HdinsightKafkaClusterRolesHeadNodeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightKafkaCluster.HdinsightKafkaClusterRolesHeadNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference)SetInternalValue(val *HdinsightKafkaClusterRolesHeadNode) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference)SetSshKeys(val *[]*string) {
	if err := j.validateSetSshKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKeys",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference)SetVirtualNetworkId(val *string) {
	if err := j.validateSetVirtualNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference)SetVmSize(val *string) {
	if err := j.validateSetVmSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) PutScriptActions(value interface{}) {
	if err := h.validatePutScriptActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putScriptActions",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		h,
		"resetPassword",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) ResetScriptActions() {
	_jsii_.InvokeVoid(
		h,
		"resetScriptActions",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) ResetSshKeys() {
	_jsii_.InvokeVoid(
		h,
		"resetSshKeys",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		h,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) ResetVirtualNetworkId() {
	_jsii_.InvokeVoid(
		h,
		"resetVirtualNetworkId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesHeadNodeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

