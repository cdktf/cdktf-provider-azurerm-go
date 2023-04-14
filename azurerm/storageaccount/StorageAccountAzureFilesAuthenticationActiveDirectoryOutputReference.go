package storageaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/storageaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference interface {
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
	DomainGuid() *string
	SetDomainGuid(val *string)
	DomainGuidInput() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	DomainSid() *string
	SetDomainSid(val *string)
	DomainSidInput() *string
	ForestName() *string
	SetForestName(val *string)
	ForestNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *StorageAccountAzureFilesAuthenticationActiveDirectory
	SetInternalValue(val *StorageAccountAzureFilesAuthenticationActiveDirectory)
	NetbiosDomainName() *string
	SetNetbiosDomainName(val *string)
	NetbiosDomainNameInput() *string
	StorageSid() *string
	SetStorageSid(val *string)
	StorageSidInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference
type jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) DomainGuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainGuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) DomainGuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainGuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) DomainSid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) DomainSidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) ForestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) ForestNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forestNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) InternalValue() *StorageAccountAzureFilesAuthenticationActiveDirectory {
	var returns *StorageAccountAzureFilesAuthenticationActiveDirectory
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) NetbiosDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netbiosDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) NetbiosDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netbiosDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) StorageSid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) StorageSidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference {
	_init_.Initialize()

	if err := validateNewStorageAccountAzureFilesAuthenticationActiveDirectoryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference_Override(s StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference)SetDomainGuid(val *string) {
	if err := j.validateSetDomainGuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainGuid",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference)SetDomainSid(val *string) {
	if err := j.validateSetDomainSidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainSid",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference)SetForestName(val *string) {
	if err := j.validateSetForestNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forestName",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference)SetInternalValue(val *StorageAccountAzureFilesAuthenticationActiveDirectory) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference)SetNetbiosDomainName(val *string) {
	if err := j.validateSetNetbiosDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netbiosDomainName",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference)SetStorageSid(val *string) {
	if err := j.validateSetStorageSidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSid",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

