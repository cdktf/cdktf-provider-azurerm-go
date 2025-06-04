// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryreplicationrecoveryplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/siterecoveryreplicationrecoveryplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference interface {
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
	FabricLocation() *string
	SetFabricLocation(val *string)
	FabricLocationInput() *string
	FailOverDirections() *[]*string
	SetFailOverDirections(val *[]*string)
	FailOverDirectionsInput() *[]*string
	FailOverTypes() *[]*string
	SetFailOverTypes(val *[]*string)
	FailOverTypesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManualActionInstruction() *string
	SetManualActionInstruction(val *string)
	ManualActionInstructionInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	RunbookId() *string
	SetRunbookId(val *string)
	RunbookIdInput() *string
	ScriptPath() *string
	SetScriptPath(val *string)
	ScriptPathInput() *string
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
	ResetFabricLocation()
	ResetManualActionInstruction()
	ResetRunbookId()
	ResetScriptPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference
type jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) FabricLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fabricLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) FabricLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fabricLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) FailOverDirections() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"failOverDirections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) FailOverDirectionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"failOverDirectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) FailOverTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"failOverTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) FailOverTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"failOverTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) ManualActionInstruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manualActionInstruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) ManualActionInstructionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manualActionInstructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) RunbookId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runbookId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) RunbookIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runbookIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) ScriptPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) ScriptPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewSiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference {
	_init_.Initialize()

	if err := validateNewSiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.siteRecoveryReplicationRecoveryPlan.SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference_Override(s SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.siteRecoveryReplicationRecoveryPlan.SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference)SetFabricLocation(val *string) {
	if err := j.validateSetFabricLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fabricLocation",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference)SetFailOverDirections(val *[]*string) {
	if err := j.validateSetFailOverDirectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failOverDirections",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference)SetFailOverTypes(val *[]*string) {
	if err := j.validateSetFailOverTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failOverTypes",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference)SetManualActionInstruction(val *string) {
	if err := j.validateSetManualActionInstructionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manualActionInstruction",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference)SetRunbookId(val *string) {
	if err := j.validateSetRunbookIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runbookId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference)SetScriptPath(val *string) {
	if err := j.validateSetScriptPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptPath",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) ResetFabricLocation() {
	_jsii_.InvokeVoid(
		s,
		"resetFabricLocation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) ResetManualActionInstruction() {
	_jsii_.InvokeVoid(
		s,
		"resetManualActionInstruction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) ResetRunbookId() {
	_jsii_.InvokeVoid(
		s,
		"resetRunbookId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) ResetScriptPath() {
	_jsii_.InvokeVoid(
		s,
		"resetScriptPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupPostActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

