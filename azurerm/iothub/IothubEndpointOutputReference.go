package iothub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/iothub/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IothubEndpointOutputReference interface {
	cdktf.ComplexObject
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
	BatchFrequencyInSeconds() *float64
	SetBatchFrequencyInSeconds(val *float64)
	BatchFrequencyInSecondsInput() *float64
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
	ConnectionString() *string
	SetConnectionString(val *string)
	ConnectionStringInput() *string
	ContainerName() *string
	SetContainerName(val *string)
	ContainerNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	EndpointUri() *string
	SetEndpointUri(val *string)
	EndpointUriInput() *string
	EntityPath() *string
	SetEntityPath(val *string)
	EntityPathInput() *string
	FileNameFormat() *string
	SetFileNameFormat(val *string)
	FileNameFormatInput() *string
	// Experimental.
	Fqn() *string
	IdentityId() *string
	SetIdentityId(val *string)
	IdentityIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxChunkSizeInBytes() *float64
	SetMaxChunkSizeInBytes(val *float64)
	MaxChunkSizeInBytesInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
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
	ResetAuthenticationType()
	ResetBatchFrequencyInSeconds()
	ResetConnectionString()
	ResetContainerName()
	ResetEncoding()
	ResetEndpointUri()
	ResetEntityPath()
	ResetFileNameFormat()
	ResetIdentityId()
	ResetMaxChunkSizeInBytes()
	ResetName()
	ResetResourceGroupName()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IothubEndpointOutputReference
type jsiiProxy_IothubEndpointOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IothubEndpointOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) BatchFrequencyInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchFrequencyInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) BatchFrequencyInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchFrequencyInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) ConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) ConnectionStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) ContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) EndpointUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) EndpointUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) EntityPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) EntityPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) FileNameFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileNameFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) FileNameFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileNameFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) IdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) IdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) MaxChunkSizeInBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxChunkSizeInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) MaxChunkSizeInBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxChunkSizeInBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewIothubEndpointOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IothubEndpointOutputReference {
	_init_.Initialize()

	if err := validateNewIothubEndpointOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IothubEndpointOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.iothub.IothubEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIothubEndpointOutputReference_Override(i IothubEndpointOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.iothub.IothubEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetBatchFrequencyInSeconds(val *float64) {
	if err := j.validateSetBatchFrequencyInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchFrequencyInSeconds",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetConnectionString(val *string) {
	if err := j.validateSetConnectionStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionString",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetContainerName(val *string) {
	if err := j.validateSetContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerName",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetEndpointUri(val *string) {
	if err := j.validateSetEndpointUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointUri",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetEntityPath(val *string) {
	if err := j.validateSetEntityPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityPath",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetFileNameFormat(val *string) {
	if err := j.validateSetFileNameFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileNameFormat",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetIdentityId(val *string) {
	if err := j.validateSetIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityId",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetMaxChunkSizeInBytes(val *float64) {
	if err := j.validateSetMaxChunkSizeInBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxChunkSizeInBytes",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (i *jsiiProxy_IothubEndpointOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointOutputReference) ResetAuthenticationType() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthenticationType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointOutputReference) ResetBatchFrequencyInSeconds() {
	_jsii_.InvokeVoid(
		i,
		"resetBatchFrequencyInSeconds",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointOutputReference) ResetConnectionString() {
	_jsii_.InvokeVoid(
		i,
		"resetConnectionString",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointOutputReference) ResetContainerName() {
	_jsii_.InvokeVoid(
		i,
		"resetContainerName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointOutputReference) ResetEncoding() {
	_jsii_.InvokeVoid(
		i,
		"resetEncoding",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointOutputReference) ResetEndpointUri() {
	_jsii_.InvokeVoid(
		i,
		"resetEndpointUri",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointOutputReference) ResetEntityPath() {
	_jsii_.InvokeVoid(
		i,
		"resetEntityPath",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointOutputReference) ResetFileNameFormat() {
	_jsii_.InvokeVoid(
		i,
		"resetFileNameFormat",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointOutputReference) ResetIdentityId() {
	_jsii_.InvokeVoid(
		i,
		"resetIdentityId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointOutputReference) ResetMaxChunkSizeInBytes() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxChunkSizeInBytes",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		i,
		"resetName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointOutputReference) ResetResourceGroupName() {
	_jsii_.InvokeVoid(
		i,
		"resetResourceGroupName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		i,
		"resetType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

