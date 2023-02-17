package firewallpolicyrulecollectiongroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/firewallpolicyrulecollectiongroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DestinationAddresses() *[]*string
	SetDestinationAddresses(val *[]*string)
	DestinationAddressesInput() *[]*string
	DestinationFqdns() *[]*string
	SetDestinationFqdns(val *[]*string)
	DestinationFqdnsInput() *[]*string
	DestinationFqdnTags() *[]*string
	SetDestinationFqdnTags(val *[]*string)
	DestinationFqdnTagsInput() *[]*string
	DestinationUrls() *[]*string
	SetDestinationUrls(val *[]*string)
	DestinationUrlsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Protocols() FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList
	ProtocolsInput() interface{}
	SourceAddresses() *[]*string
	SetSourceAddresses(val *[]*string)
	SourceAddressesInput() *[]*string
	SourceIpGroups() *[]*string
	SetSourceIpGroups(val *[]*string)
	SourceIpGroupsInput() *[]*string
	TerminateTls() interface{}
	SetTerminateTls(val interface{})
	TerminateTlsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebCategories() *[]*string
	SetWebCategories(val *[]*string)
	WebCategoriesInput() *[]*string
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
	PutProtocols(value interface{})
	ResetDescription()
	ResetDestinationAddresses()
	ResetDestinationFqdns()
	ResetDestinationFqdnTags()
	ResetDestinationUrls()
	ResetProtocols()
	ResetSourceAddresses()
	ResetSourceIpGroups()
	ResetTerminateTls()
	ResetWebCategories()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference
type jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationFqdnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationFqdnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationFqdnTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationFqdnTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationFqdnTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationFqdnTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) Protocols() FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList {
	var returns FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ProtocolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SourceAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SourceAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SourceIpGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIpGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SourceIpGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIpGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) TerminateTls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) TerminateTlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) WebCategories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"webCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) WebCategoriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"webCategoriesInput",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference {
	_init_.Initialize()

	if err := validateNewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference_Override(f FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference)SetDestinationAddresses(val *[]*string) {
	if err := j.validateSetDestinationAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddresses",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference)SetDestinationFqdns(val *[]*string) {
	if err := j.validateSetDestinationFqdnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationFqdns",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference)SetDestinationFqdnTags(val *[]*string) {
	if err := j.validateSetDestinationFqdnTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationFqdnTags",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference)SetDestinationUrls(val *[]*string) {
	if err := j.validateSetDestinationUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationUrls",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference)SetSourceAddresses(val *[]*string) {
	if err := j.validateSetSourceAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddresses",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference)SetSourceIpGroups(val *[]*string) {
	if err := j.validateSetSourceIpGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIpGroups",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference)SetTerminateTls(val interface{}) {
	if err := j.validateSetTerminateTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminateTls",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference)SetWebCategories(val *[]*string) {
	if err := j.validateSetWebCategoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webCategories",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) PutProtocols(value interface{}) {
	if err := f.validatePutProtocolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putProtocols",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		f,
		"resetDescription",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetDestinationAddresses() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationAddresses",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetDestinationFqdns() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationFqdns",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetDestinationFqdnTags() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationFqdnTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetDestinationUrls() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationUrls",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetProtocols() {
	_jsii_.InvokeVoid(
		f,
		"resetProtocols",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetSourceAddresses() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceAddresses",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetSourceIpGroups() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceIpGroups",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetTerminateTls() {
	_jsii_.InvokeVoid(
		f,
		"resetTerminateTls",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetWebCategories() {
	_jsii_.InvokeVoid(
		f,
		"resetWebCategories",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

