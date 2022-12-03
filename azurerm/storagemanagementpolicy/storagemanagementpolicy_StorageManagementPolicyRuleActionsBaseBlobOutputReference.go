package storagemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/storagemanagementpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageManagementPolicyRuleActionsBaseBlobOutputReference interface {
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
	DeleteAfterDaysSinceCreationGreaterThan() *float64
	SetDeleteAfterDaysSinceCreationGreaterThan(val *float64)
	DeleteAfterDaysSinceCreationGreaterThanInput() *float64
	DeleteAfterDaysSinceLastAccessTimeGreaterThan() *float64
	SetDeleteAfterDaysSinceLastAccessTimeGreaterThan(val *float64)
	DeleteAfterDaysSinceLastAccessTimeGreaterThanInput() *float64
	DeleteAfterDaysSinceModificationGreaterThan() *float64
	SetDeleteAfterDaysSinceModificationGreaterThan(val *float64)
	DeleteAfterDaysSinceModificationGreaterThanInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *StorageManagementPolicyRuleActionsBaseBlob
	SetInternalValue(val *StorageManagementPolicyRuleActionsBaseBlob)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TierToArchiveAfterDaysSinceCreationGreaterThan() *float64
	SetTierToArchiveAfterDaysSinceCreationGreaterThan(val *float64)
	TierToArchiveAfterDaysSinceCreationGreaterThanInput() *float64
	TierToArchiveAfterDaysSinceLastAccessTimeGreaterThan() *float64
	SetTierToArchiveAfterDaysSinceLastAccessTimeGreaterThan(val *float64)
	TierToArchiveAfterDaysSinceLastAccessTimeGreaterThanInput() *float64
	TierToArchiveAfterDaysSinceLastTierChangeGreaterThan() *float64
	SetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan(val *float64)
	TierToArchiveAfterDaysSinceLastTierChangeGreaterThanInput() *float64
	TierToArchiveAfterDaysSinceModificationGreaterThan() *float64
	SetTierToArchiveAfterDaysSinceModificationGreaterThan(val *float64)
	TierToArchiveAfterDaysSinceModificationGreaterThanInput() *float64
	TierToCoolAfterDaysSinceCreationGreaterThan() *float64
	SetTierToCoolAfterDaysSinceCreationGreaterThan(val *float64)
	TierToCoolAfterDaysSinceCreationGreaterThanInput() *float64
	TierToCoolAfterDaysSinceLastAccessTimeGreaterThan() *float64
	SetTierToCoolAfterDaysSinceLastAccessTimeGreaterThan(val *float64)
	TierToCoolAfterDaysSinceLastAccessTimeGreaterThanInput() *float64
	TierToCoolAfterDaysSinceModificationGreaterThan() *float64
	SetTierToCoolAfterDaysSinceModificationGreaterThan(val *float64)
	TierToCoolAfterDaysSinceModificationGreaterThanInput() *float64
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
	ResetDeleteAfterDaysSinceCreationGreaterThan()
	ResetDeleteAfterDaysSinceLastAccessTimeGreaterThan()
	ResetDeleteAfterDaysSinceModificationGreaterThan()
	ResetTierToArchiveAfterDaysSinceCreationGreaterThan()
	ResetTierToArchiveAfterDaysSinceLastAccessTimeGreaterThan()
	ResetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan()
	ResetTierToArchiveAfterDaysSinceModificationGreaterThan()
	ResetTierToCoolAfterDaysSinceCreationGreaterThan()
	ResetTierToCoolAfterDaysSinceLastAccessTimeGreaterThan()
	ResetTierToCoolAfterDaysSinceModificationGreaterThan()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageManagementPolicyRuleActionsBaseBlobOutputReference
type jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) DeleteAfterDaysSinceCreationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceCreationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) DeleteAfterDaysSinceCreationGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceCreationGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) DeleteAfterDaysSinceLastAccessTimeGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceLastAccessTimeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) DeleteAfterDaysSinceLastAccessTimeGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceLastAccessTimeGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) DeleteAfterDaysSinceModificationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceModificationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) DeleteAfterDaysSinceModificationGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceModificationGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) InternalValue() *StorageManagementPolicyRuleActionsBaseBlob {
	var returns *StorageManagementPolicyRuleActionsBaseBlob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceCreationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceCreationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceCreationGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceCreationGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceLastAccessTimeGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceLastAccessTimeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceLastAccessTimeGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceLastAccessTimeGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceLastTierChangeGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceLastTierChangeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceLastTierChangeGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceLastTierChangeGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceModificationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceModificationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceModificationGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceModificationGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToCoolAfterDaysSinceCreationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToCoolAfterDaysSinceCreationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToCoolAfterDaysSinceCreationGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToCoolAfterDaysSinceCreationGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToCoolAfterDaysSinceLastAccessTimeGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToCoolAfterDaysSinceLastAccessTimeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToCoolAfterDaysSinceLastAccessTimeGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToCoolAfterDaysSinceLastAccessTimeGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToCoolAfterDaysSinceModificationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToCoolAfterDaysSinceModificationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToCoolAfterDaysSinceModificationGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToCoolAfterDaysSinceModificationGreaterThanInput",
		&returns,
	)
	return returns
}


func NewStorageManagementPolicyRuleActionsBaseBlobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageManagementPolicyRuleActionsBaseBlobOutputReference {
	_init_.Initialize()

	if err := validateNewStorageManagementPolicyRuleActionsBaseBlobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsBaseBlobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageManagementPolicyRuleActionsBaseBlobOutputReference_Override(s StorageManagementPolicyRuleActionsBaseBlobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsBaseBlobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference)SetDeleteAfterDaysSinceCreationGreaterThan(val *float64) {
	if err := j.validateSetDeleteAfterDaysSinceCreationGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAfterDaysSinceCreationGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference)SetDeleteAfterDaysSinceLastAccessTimeGreaterThan(val *float64) {
	if err := j.validateSetDeleteAfterDaysSinceLastAccessTimeGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAfterDaysSinceLastAccessTimeGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference)SetDeleteAfterDaysSinceModificationGreaterThan(val *float64) {
	if err := j.validateSetDeleteAfterDaysSinceModificationGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAfterDaysSinceModificationGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference)SetInternalValue(val *StorageManagementPolicyRuleActionsBaseBlob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference)SetTierToArchiveAfterDaysSinceCreationGreaterThan(val *float64) {
	if err := j.validateSetTierToArchiveAfterDaysSinceCreationGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tierToArchiveAfterDaysSinceCreationGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference)SetTierToArchiveAfterDaysSinceLastAccessTimeGreaterThan(val *float64) {
	if err := j.validateSetTierToArchiveAfterDaysSinceLastAccessTimeGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tierToArchiveAfterDaysSinceLastAccessTimeGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference)SetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan(val *float64) {
	if err := j.validateSetTierToArchiveAfterDaysSinceLastTierChangeGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tierToArchiveAfterDaysSinceLastTierChangeGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference)SetTierToArchiveAfterDaysSinceModificationGreaterThan(val *float64) {
	if err := j.validateSetTierToArchiveAfterDaysSinceModificationGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tierToArchiveAfterDaysSinceModificationGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference)SetTierToCoolAfterDaysSinceCreationGreaterThan(val *float64) {
	if err := j.validateSetTierToCoolAfterDaysSinceCreationGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tierToCoolAfterDaysSinceCreationGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference)SetTierToCoolAfterDaysSinceLastAccessTimeGreaterThan(val *float64) {
	if err := j.validateSetTierToCoolAfterDaysSinceLastAccessTimeGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tierToCoolAfterDaysSinceLastAccessTimeGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference)SetTierToCoolAfterDaysSinceModificationGreaterThan(val *float64) {
	if err := j.validateSetTierToCoolAfterDaysSinceModificationGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tierToCoolAfterDaysSinceModificationGreaterThan",
		val,
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetDeleteAfterDaysSinceCreationGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteAfterDaysSinceCreationGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetDeleteAfterDaysSinceLastAccessTimeGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteAfterDaysSinceLastAccessTimeGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetDeleteAfterDaysSinceModificationGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteAfterDaysSinceModificationGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetTierToArchiveAfterDaysSinceCreationGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetTierToArchiveAfterDaysSinceCreationGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetTierToArchiveAfterDaysSinceLastAccessTimeGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetTierToArchiveAfterDaysSinceLastAccessTimeGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetTierToArchiveAfterDaysSinceModificationGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetTierToArchiveAfterDaysSinceModificationGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetTierToCoolAfterDaysSinceCreationGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetTierToCoolAfterDaysSinceCreationGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetTierToCoolAfterDaysSinceLastAccessTimeGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetTierToCoolAfterDaysSinceLastAccessTimeGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetTierToCoolAfterDaysSinceModificationGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetTierToCoolAfterDaysSinceModificationGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

