package storageobjectreplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/storageobjectreplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageObjectReplicationRulesOutputReference interface {
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
	CopyBlobsCreatedAfter() *string
	SetCopyBlobsCreatedAfter(val *string)
	CopyBlobsCreatedAfterInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DestinationContainerName() *string
	SetDestinationContainerName(val *string)
	DestinationContainerNameInput() *string
	FilterOutBlobsWithPrefix() *[]*string
	SetFilterOutBlobsWithPrefix(val *[]*string)
	FilterOutBlobsWithPrefixInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SourceContainerName() *string
	SetSourceContainerName(val *string)
	SourceContainerNameInput() *string
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
	ResetCopyBlobsCreatedAfter()
	ResetFilterOutBlobsWithPrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageObjectReplicationRulesOutputReference
type jsiiProxy_StorageObjectReplicationRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) CopyBlobsCreatedAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyBlobsCreatedAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) CopyBlobsCreatedAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyBlobsCreatedAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) DestinationContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) DestinationContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) FilterOutBlobsWithPrefix() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filterOutBlobsWithPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) FilterOutBlobsWithPrefixInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filterOutBlobsWithPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) SourceContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) SourceContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageObjectReplicationRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StorageObjectReplicationRulesOutputReference {
	_init_.Initialize()

	if err := validateNewStorageObjectReplicationRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageObjectReplicationRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageObjectReplication.StorageObjectReplicationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStorageObjectReplicationRulesOutputReference_Override(s StorageObjectReplicationRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageObjectReplication.StorageObjectReplicationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference)SetCopyBlobsCreatedAfter(val *string) {
	if err := j.validateSetCopyBlobsCreatedAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyBlobsCreatedAfter",
		val,
	)
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference)SetDestinationContainerName(val *string) {
	if err := j.validateSetDestinationContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationContainerName",
		val,
	)
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference)SetFilterOutBlobsWithPrefix(val *[]*string) {
	if err := j.validateSetFilterOutBlobsWithPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterOutBlobsWithPrefix",
		val,
	)
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference)SetSourceContainerName(val *string) {
	if err := j.validateSetSourceContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceContainerName",
		val,
	)
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageObjectReplicationRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) ResetCopyBlobsCreatedAfter() {
	_jsii_.InvokeVoid(
		s,
		"resetCopyBlobsCreatedAfter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) ResetFilterOutBlobsWithPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetFilterOutBlobsWithPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageObjectReplicationRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

