package storagemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/storagemanagementpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageManagementPolicyRuleActionsSnapshotOutputReference interface {
	cdktf.ComplexObject
	ChangeTierToArchiveAfterDaysSinceCreation() *float64
	SetChangeTierToArchiveAfterDaysSinceCreation(val *float64)
	ChangeTierToArchiveAfterDaysSinceCreationInput() *float64
	ChangeTierToCoolAfterDaysSinceCreation() *float64
	SetChangeTierToCoolAfterDaysSinceCreation(val *float64)
	ChangeTierToCoolAfterDaysSinceCreationInput() *float64
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
	DeleteAfterDaysSinceCreationGreaterThan() *float64
	SetDeleteAfterDaysSinceCreationGreaterThan(val *float64)
	DeleteAfterDaysSinceCreationGreaterThanInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *StorageManagementPolicyRuleActionsSnapshot
	SetInternalValue(val *StorageManagementPolicyRuleActionsSnapshot)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TierToArchiveAfterDaysSinceLastTierChangeGreaterThan() *float64
	SetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan(val *float64)
	TierToArchiveAfterDaysSinceLastTierChangeGreaterThanInput() *float64
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
	ResetChangeTierToArchiveAfterDaysSinceCreation()
	ResetChangeTierToCoolAfterDaysSinceCreation()
	ResetDeleteAfterDaysSinceCreationGreaterThan()
	ResetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageManagementPolicyRuleActionsSnapshotOutputReference
type jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ChangeTierToArchiveAfterDaysSinceCreation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeTierToArchiveAfterDaysSinceCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ChangeTierToArchiveAfterDaysSinceCreationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeTierToArchiveAfterDaysSinceCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ChangeTierToCoolAfterDaysSinceCreation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeTierToCoolAfterDaysSinceCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ChangeTierToCoolAfterDaysSinceCreationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeTierToCoolAfterDaysSinceCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) DeleteAfterDaysSinceCreationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceCreationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) DeleteAfterDaysSinceCreationGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceCreationGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) InternalValue() *StorageManagementPolicyRuleActionsSnapshot {
	var returns *StorageManagementPolicyRuleActionsSnapshot
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) TierToArchiveAfterDaysSinceLastTierChangeGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceLastTierChangeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) TierToArchiveAfterDaysSinceLastTierChangeGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceLastTierChangeGreaterThanInput",
		&returns,
	)
	return returns
}


func NewStorageManagementPolicyRuleActionsSnapshotOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageManagementPolicyRuleActionsSnapshotOutputReference {
	_init_.Initialize()

	if err := validateNewStorageManagementPolicyRuleActionsSnapshotOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsSnapshotOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageManagementPolicyRuleActionsSnapshotOutputReference_Override(s StorageManagementPolicyRuleActionsSnapshotOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsSnapshotOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference)SetChangeTierToArchiveAfterDaysSinceCreation(val *float64) {
	if err := j.validateSetChangeTierToArchiveAfterDaysSinceCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changeTierToArchiveAfterDaysSinceCreation",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference)SetChangeTierToCoolAfterDaysSinceCreation(val *float64) {
	if err := j.validateSetChangeTierToCoolAfterDaysSinceCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changeTierToCoolAfterDaysSinceCreation",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference)SetDeleteAfterDaysSinceCreationGreaterThan(val *float64) {
	if err := j.validateSetDeleteAfterDaysSinceCreationGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAfterDaysSinceCreationGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference)SetInternalValue(val *StorageManagementPolicyRuleActionsSnapshot) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference)SetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan(val *float64) {
	if err := j.validateSetTierToArchiveAfterDaysSinceLastTierChangeGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tierToArchiveAfterDaysSinceLastTierChangeGreaterThan",
		val,
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ResetChangeTierToArchiveAfterDaysSinceCreation() {
	_jsii_.InvokeVoid(
		s,
		"resetChangeTierToArchiveAfterDaysSinceCreation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ResetChangeTierToCoolAfterDaysSinceCreation() {
	_jsii_.InvokeVoid(
		s,
		"resetChangeTierToCoolAfterDaysSinceCreation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ResetDeleteAfterDaysSinceCreationGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteAfterDaysSinceCreationGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ResetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
