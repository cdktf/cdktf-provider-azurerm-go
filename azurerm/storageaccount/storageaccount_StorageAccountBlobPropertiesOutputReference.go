package storageaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/storageaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageAccountBlobPropertiesOutputReference interface {
	cdktf.ComplexObject
	ChangeFeedEnabled() interface{}
	SetChangeFeedEnabled(val interface{})
	ChangeFeedEnabledInput() interface{}
	ChangeFeedRetentionInDays() *float64
	SetChangeFeedRetentionInDays(val *float64)
	ChangeFeedRetentionInDaysInput() *float64
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
	ContainerDeleteRetentionPolicy() StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference
	ContainerDeleteRetentionPolicyInput() *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy
	CorsRule() StorageAccountBlobPropertiesCorsRuleList
	CorsRuleInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultServiceVersion() *string
	SetDefaultServiceVersion(val *string)
	DefaultServiceVersionInput() *string
	DeleteRetentionPolicy() StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference
	DeleteRetentionPolicyInput() *StorageAccountBlobPropertiesDeleteRetentionPolicy
	// Experimental.
	Fqn() *string
	InternalValue() *StorageAccountBlobProperties
	SetInternalValue(val *StorageAccountBlobProperties)
	LastAccessTimeEnabled() interface{}
	SetLastAccessTimeEnabled(val interface{})
	LastAccessTimeEnabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VersioningEnabled() interface{}
	SetVersioningEnabled(val interface{})
	VersioningEnabledInput() interface{}
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
	PutContainerDeleteRetentionPolicy(value *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy)
	PutCorsRule(value interface{})
	PutDeleteRetentionPolicy(value *StorageAccountBlobPropertiesDeleteRetentionPolicy)
	ResetChangeFeedEnabled()
	ResetChangeFeedRetentionInDays()
	ResetContainerDeleteRetentionPolicy()
	ResetCorsRule()
	ResetDefaultServiceVersion()
	ResetDeleteRetentionPolicy()
	ResetLastAccessTimeEnabled()
	ResetVersioningEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountBlobPropertiesOutputReference
type jsiiProxy_StorageAccountBlobPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ChangeFeedEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"changeFeedEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ChangeFeedEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"changeFeedEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ChangeFeedRetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeFeedRetentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ChangeFeedRetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeFeedRetentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ContainerDeleteRetentionPolicy() StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference {
	var returns StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference
	_jsii_.Get(
		j,
		"containerDeleteRetentionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ContainerDeleteRetentionPolicyInput() *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy {
	var returns *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy
	_jsii_.Get(
		j,
		"containerDeleteRetentionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) CorsRule() StorageAccountBlobPropertiesCorsRuleList {
	var returns StorageAccountBlobPropertiesCorsRuleList
	_jsii_.Get(
		j,
		"corsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) CorsRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) DefaultServiceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultServiceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) DefaultServiceVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultServiceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) DeleteRetentionPolicy() StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference {
	var returns StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference
	_jsii_.Get(
		j,
		"deleteRetentionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) DeleteRetentionPolicyInput() *StorageAccountBlobPropertiesDeleteRetentionPolicy {
	var returns *StorageAccountBlobPropertiesDeleteRetentionPolicy
	_jsii_.Get(
		j,
		"deleteRetentionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) InternalValue() *StorageAccountBlobProperties {
	var returns *StorageAccountBlobProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) LastAccessTimeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastAccessTimeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) LastAccessTimeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastAccessTimeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) VersioningEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versioningEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) VersioningEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versioningEnabledInput",
		&returns,
	)
	return returns
}


func NewStorageAccountBlobPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountBlobPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewStorageAccountBlobPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageAccountBlobPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountBlobPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountBlobPropertiesOutputReference_Override(s StorageAccountBlobPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountBlobPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference)SetChangeFeedEnabled(val interface{}) {
	if err := j.validateSetChangeFeedEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changeFeedEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference)SetChangeFeedRetentionInDays(val *float64) {
	if err := j.validateSetChangeFeedRetentionInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changeFeedRetentionInDays",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference)SetDefaultServiceVersion(val *string) {
	if err := j.validateSetDefaultServiceVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultServiceVersion",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference)SetInternalValue(val *StorageAccountBlobProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference)SetLastAccessTimeEnabled(val interface{}) {
	if err := j.validateSetLastAccessTimeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastAccessTimeEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference)SetVersioningEnabled(val interface{}) {
	if err := j.validateSetVersioningEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versioningEnabled",
		val,
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) PutContainerDeleteRetentionPolicy(value *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy) {
	if err := s.validatePutContainerDeleteRetentionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putContainerDeleteRetentionPolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) PutCorsRule(value interface{}) {
	if err := s.validatePutCorsRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCorsRule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) PutDeleteRetentionPolicy(value *StorageAccountBlobPropertiesDeleteRetentionPolicy) {
	if err := s.validatePutDeleteRetentionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDeleteRetentionPolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetChangeFeedEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetChangeFeedEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetChangeFeedRetentionInDays() {
	_jsii_.InvokeVoid(
		s,
		"resetChangeFeedRetentionInDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetContainerDeleteRetentionPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetContainerDeleteRetentionPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetCorsRule() {
	_jsii_.InvokeVoid(
		s,
		"resetCorsRule",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetDefaultServiceVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultServiceVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetDeleteRetentionPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteRetentionPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetLastAccessTimeEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetLastAccessTimeEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetVersioningEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetVersioningEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

