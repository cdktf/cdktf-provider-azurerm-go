// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolumegrouporacle

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/netappvolumegrouporacle/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference interface {
	cdktf.ComplexObject
	AllowedClients() *string
	SetAllowedClients(val *string)
	AllowedClientsInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Nfsv3Enabled() interface{}
	SetNfsv3Enabled(val interface{})
	Nfsv3EnabledInput() interface{}
	Nfsv41Enabled() interface{}
	SetNfsv41Enabled(val interface{})
	Nfsv41EnabledInput() interface{}
	RootAccessEnabled() interface{}
	SetRootAccessEnabled(val interface{})
	RootAccessEnabledInput() interface{}
	RuleIndex() *float64
	SetRuleIndex(val *float64)
	RuleIndexInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnixReadOnly() interface{}
	SetUnixReadOnly(val interface{})
	UnixReadOnlyInput() interface{}
	UnixReadWrite() interface{}
	SetUnixReadWrite(val interface{})
	UnixReadWriteInput() interface{}
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
	ResetRootAccessEnabled()
	ResetUnixReadOnly()
	ResetUnixReadWrite()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference
type jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) AllowedClients() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) AllowedClientsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) Nfsv3Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsv3Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) Nfsv3EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsv3EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) Nfsv41Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsv41Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) Nfsv41EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsv41EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) RootAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) RootAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) RuleIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) RuleIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) UnixReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unixReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) UnixReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unixReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) UnixReadWrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unixReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) UnixReadWriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unixReadWriteInput",
		&returns,
	)
	return returns
}


func NewNetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference {
	_init_.Initialize()

	if err := validateNewNetappVolumeGroupOracleVolumeExportPolicyRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolumeGroupOracle.NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference_Override(n NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolumeGroupOracle.NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference)SetAllowedClients(val *string) {
	if err := j.validateSetAllowedClientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedClients",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference)SetNfsv3Enabled(val interface{}) {
	if err := j.validateSetNfsv3EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsv3Enabled",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference)SetNfsv41Enabled(val interface{}) {
	if err := j.validateSetNfsv41EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsv41Enabled",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference)SetRootAccessEnabled(val interface{}) {
	if err := j.validateSetRootAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference)SetRuleIndex(val *float64) {
	if err := j.validateSetRuleIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference)SetUnixReadOnly(val interface{}) {
	if err := j.validateSetUnixReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unixReadOnly",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference)SetUnixReadWrite(val interface{}) {
	if err := j.validateSetUnixReadWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unixReadWrite",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) ResetRootAccessEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetRootAccessEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) ResetUnixReadOnly() {
	_jsii_.InvokeVoid(
		n,
		"resetUnixReadOnly",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) ResetUnixReadWrite() {
	_jsii_.InvokeVoid(
		n,
		"resetUnixReadWrite",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeExportPolicyRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

