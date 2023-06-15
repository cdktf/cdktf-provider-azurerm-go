package storageblobinventorypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/storageblobinventorypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageBlobInventoryPolicyRulesFilterOutputReference interface {
	cdktf.ComplexObject
	BlobTypes() *[]*string
	SetBlobTypes(val *[]*string)
	BlobTypesInput() *[]*string
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
	ExcludePrefixes() *[]*string
	SetExcludePrefixes(val *[]*string)
	ExcludePrefixesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeBlobVersions() interface{}
	SetIncludeBlobVersions(val interface{})
	IncludeBlobVersionsInput() interface{}
	IncludeDeleted() interface{}
	SetIncludeDeleted(val interface{})
	IncludeDeletedInput() interface{}
	IncludeSnapshots() interface{}
	SetIncludeSnapshots(val interface{})
	IncludeSnapshotsInput() interface{}
	InternalValue() *StorageBlobInventoryPolicyRulesFilter
	SetInternalValue(val *StorageBlobInventoryPolicyRulesFilter)
	PrefixMatch() *[]*string
	SetPrefixMatch(val *[]*string)
	PrefixMatchInput() *[]*string
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
	ResetExcludePrefixes()
	ResetIncludeBlobVersions()
	ResetIncludeDeleted()
	ResetIncludeSnapshots()
	ResetPrefixMatch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageBlobInventoryPolicyRulesFilterOutputReference
type jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) BlobTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blobTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) BlobTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blobTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) ExcludePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) ExcludePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) IncludeBlobVersions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeBlobVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) IncludeBlobVersionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeBlobVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) IncludeDeleted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeDeleted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) IncludeDeletedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeDeletedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) IncludeSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) IncludeSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) InternalValue() *StorageBlobInventoryPolicyRulesFilter {
	var returns *StorageBlobInventoryPolicyRulesFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) PrefixMatch() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prefixMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) PrefixMatchInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prefixMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageBlobInventoryPolicyRulesFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageBlobInventoryPolicyRulesFilterOutputReference {
	_init_.Initialize()

	if err := validateNewStorageBlobInventoryPolicyRulesFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageBlobInventoryPolicy.StorageBlobInventoryPolicyRulesFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageBlobInventoryPolicyRulesFilterOutputReference_Override(s StorageBlobInventoryPolicyRulesFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageBlobInventoryPolicy.StorageBlobInventoryPolicyRulesFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference)SetBlobTypes(val *[]*string) {
	if err := j.validateSetBlobTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blobTypes",
		val,
	)
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference)SetExcludePrefixes(val *[]*string) {
	if err := j.validateSetExcludePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludePrefixes",
		val,
	)
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference)SetIncludeBlobVersions(val interface{}) {
	if err := j.validateSetIncludeBlobVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeBlobVersions",
		val,
	)
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference)SetIncludeDeleted(val interface{}) {
	if err := j.validateSetIncludeDeletedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeDeleted",
		val,
	)
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference)SetIncludeSnapshots(val interface{}) {
	if err := j.validateSetIncludeSnapshotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSnapshots",
		val,
	)
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference)SetInternalValue(val *StorageBlobInventoryPolicyRulesFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference)SetPrefixMatch(val *[]*string) {
	if err := j.validateSetPrefixMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixMatch",
		val,
	)
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) ResetExcludePrefixes() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludePrefixes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) ResetIncludeBlobVersions() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludeBlobVersions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) ResetIncludeDeleted() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludeDeleted",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) ResetIncludeSnapshots() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludeSnapshots",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) ResetPrefixMatch() {
	_jsii_.InvokeVoid(
		s,
		"resetPrefixMatch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageBlobInventoryPolicyRulesFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

