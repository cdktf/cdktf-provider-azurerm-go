package siterecoveryreplicatedvm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/siterecoveryreplicatedvm/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SiteRecoveryReplicatedVmManagedDiskOutputReference interface {
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
	DiskId() *string
	SetDiskId(val *string)
	DiskIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StagingStorageAccountId() *string
	SetStagingStorageAccountId(val *string)
	StagingStorageAccountIdInput() *string
	TargetDiskEncryption() SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionList
	TargetDiskEncryptionInput() interface{}
	TargetDiskEncryptionSetId() *string
	SetTargetDiskEncryptionSetId(val *string)
	TargetDiskEncryptionSetIdInput() *string
	TargetDiskType() *string
	SetTargetDiskType(val *string)
	TargetDiskTypeInput() *string
	TargetReplicaDiskType() *string
	SetTargetReplicaDiskType(val *string)
	TargetReplicaDiskTypeInput() *string
	TargetResourceGroupId() *string
	SetTargetResourceGroupId(val *string)
	TargetResourceGroupIdInput() *string
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
	PutTargetDiskEncryption(value interface{})
	ResetDiskId()
	ResetStagingStorageAccountId()
	ResetTargetDiskEncryption()
	ResetTargetDiskEncryptionSetId()
	ResetTargetDiskType()
	ResetTargetReplicaDiskType()
	ResetTargetResourceGroupId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SiteRecoveryReplicatedVmManagedDiskOutputReference
type jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) DiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) DiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) StagingStorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingStorageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) StagingStorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingStorageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) TargetDiskEncryption() SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionList {
	var returns SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionList
	_jsii_.Get(
		j,
		"targetDiskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) TargetDiskEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetDiskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) TargetDiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDiskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) TargetDiskEncryptionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDiskEncryptionSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) TargetDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) TargetDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) TargetReplicaDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetReplicaDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) TargetReplicaDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetReplicaDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) TargetResourceGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) TargetResourceGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSiteRecoveryReplicatedVmManagedDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SiteRecoveryReplicatedVmManagedDiskOutputReference {
	_init_.Initialize()

	if err := validateNewSiteRecoveryReplicatedVmManagedDiskOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSiteRecoveryReplicatedVmManagedDiskOutputReference_Override(s SiteRecoveryReplicatedVmManagedDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference)SetDiskId(val *string) {
	if err := j.validateSetDiskIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference)SetStagingStorageAccountId(val *string) {
	if err := j.validateSetStagingStorageAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingStorageAccountId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference)SetTargetDiskEncryptionSetId(val *string) {
	if err := j.validateSetTargetDiskEncryptionSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDiskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference)SetTargetDiskType(val *string) {
	if err := j.validateSetTargetDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDiskType",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference)SetTargetReplicaDiskType(val *string) {
	if err := j.validateSetTargetReplicaDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetReplicaDiskType",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference)SetTargetResourceGroupId(val *string) {
	if err := j.validateSetTargetResourceGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetResourceGroupId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) PutTargetDiskEncryption(value interface{}) {
	if err := s.validatePutTargetDiskEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTargetDiskEncryption",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) ResetDiskId() {
	_jsii_.InvokeVoid(
		s,
		"resetDiskId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) ResetStagingStorageAccountId() {
	_jsii_.InvokeVoid(
		s,
		"resetStagingStorageAccountId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) ResetTargetDiskEncryption() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetDiskEncryption",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) ResetTargetDiskEncryptionSetId() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetDiskEncryptionSetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) ResetTargetDiskType() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetDiskType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) ResetTargetReplicaDiskType() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetReplicaDiskType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) ResetTargetResourceGroupId() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetResourceGroupId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

