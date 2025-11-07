// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databricksworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/databricksworkspace/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabricksWorkspaceEnhancedSecurityComplianceOutputReference interface {
	cdktf.ComplexObject
	AutomaticClusterUpdateEnabled() interface{}
	SetAutomaticClusterUpdateEnabled(val interface{})
	AutomaticClusterUpdateEnabledInput() interface{}
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
	ComplianceSecurityProfileEnabled() interface{}
	SetComplianceSecurityProfileEnabled(val interface{})
	ComplianceSecurityProfileEnabledInput() interface{}
	ComplianceSecurityProfileStandards() *[]*string
	SetComplianceSecurityProfileStandards(val *[]*string)
	ComplianceSecurityProfileStandardsInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnhancedSecurityMonitoringEnabled() interface{}
	SetEnhancedSecurityMonitoringEnabled(val interface{})
	EnhancedSecurityMonitoringEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DatabricksWorkspaceEnhancedSecurityCompliance
	SetInternalValue(val *DatabricksWorkspaceEnhancedSecurityCompliance)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetAutomaticClusterUpdateEnabled()
	ResetComplianceSecurityProfileEnabled()
	ResetComplianceSecurityProfileStandards()
	ResetEnhancedSecurityMonitoringEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabricksWorkspaceEnhancedSecurityComplianceOutputReference
type jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) AutomaticClusterUpdateEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticClusterUpdateEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) AutomaticClusterUpdateEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticClusterUpdateEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) ComplianceSecurityProfileEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceSecurityProfileEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) ComplianceSecurityProfileEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceSecurityProfileEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) ComplianceSecurityProfileStandards() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"complianceSecurityProfileStandards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) ComplianceSecurityProfileStandardsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"complianceSecurityProfileStandardsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) EnhancedSecurityMonitoringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedSecurityMonitoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) EnhancedSecurityMonitoringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedSecurityMonitoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) InternalValue() *DatabricksWorkspaceEnhancedSecurityCompliance {
	var returns *DatabricksWorkspaceEnhancedSecurityCompliance
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatabricksWorkspaceEnhancedSecurityComplianceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabricksWorkspaceEnhancedSecurityComplianceOutputReference {
	_init_.Initialize()

	if err := validateNewDatabricksWorkspaceEnhancedSecurityComplianceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspaceEnhancedSecurityComplianceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabricksWorkspaceEnhancedSecurityComplianceOutputReference_Override(d DatabricksWorkspaceEnhancedSecurityComplianceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspaceEnhancedSecurityComplianceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference)SetAutomaticClusterUpdateEnabled(val interface{}) {
	if err := j.validateSetAutomaticClusterUpdateEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticClusterUpdateEnabled",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference)SetComplianceSecurityProfileEnabled(val interface{}) {
	if err := j.validateSetComplianceSecurityProfileEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complianceSecurityProfileEnabled",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference)SetComplianceSecurityProfileStandards(val *[]*string) {
	if err := j.validateSetComplianceSecurityProfileStandardsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complianceSecurityProfileStandards",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference)SetEnhancedSecurityMonitoringEnabled(val interface{}) {
	if err := j.validateSetEnhancedSecurityMonitoringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedSecurityMonitoringEnabled",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference)SetInternalValue(val *DatabricksWorkspaceEnhancedSecurityCompliance) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) ResetAutomaticClusterUpdateEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetAutomaticClusterUpdateEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) ResetComplianceSecurityProfileEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetComplianceSecurityProfileEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) ResetComplianceSecurityProfileStandards() {
	_jsii_.InvokeVoid(
		d,
		"resetComplianceSecurityProfileStandards",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) ResetEnhancedSecurityMonitoringEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetEnhancedSecurityMonitoringEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceEnhancedSecurityComplianceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

