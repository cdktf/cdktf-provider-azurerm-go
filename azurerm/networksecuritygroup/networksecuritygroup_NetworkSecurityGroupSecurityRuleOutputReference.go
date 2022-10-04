package networksecuritygroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/networksecuritygroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecurityGroupSecurityRuleOutputReference interface {
	cdktf.ComplexObject
	Access() *string
	SetAccess(val *string)
	AccessInput() *string
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
	DestinationAddressPrefix() *string
	SetDestinationAddressPrefix(val *string)
	DestinationAddressPrefixes() *[]*string
	SetDestinationAddressPrefixes(val *[]*string)
	DestinationAddressPrefixesInput() *[]*string
	DestinationAddressPrefixInput() *string
	DestinationApplicationSecurityGroupIds() *[]*string
	SetDestinationApplicationSecurityGroupIds(val *[]*string)
	DestinationApplicationSecurityGroupIdsInput() *[]*string
	DestinationPortRange() *string
	SetDestinationPortRange(val *string)
	DestinationPortRangeInput() *string
	DestinationPortRanges() *[]*string
	SetDestinationPortRanges(val *[]*string)
	DestinationPortRangesInput() *[]*string
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	SourceAddressPrefix() *string
	SetSourceAddressPrefix(val *string)
	SourceAddressPrefixes() *[]*string
	SetSourceAddressPrefixes(val *[]*string)
	SourceAddressPrefixesInput() *[]*string
	SourceAddressPrefixInput() *string
	SourceApplicationSecurityGroupIds() *[]*string
	SetSourceApplicationSecurityGroupIds(val *[]*string)
	SourceApplicationSecurityGroupIdsInput() *[]*string
	SourcePortRange() *string
	SetSourcePortRange(val *string)
	SourcePortRangeInput() *string
	SourcePortRanges() *[]*string
	SetSourcePortRanges(val *[]*string)
	SourcePortRangesInput() *[]*string
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
	ResetAccess()
	ResetDescription()
	ResetDestinationAddressPrefix()
	ResetDestinationAddressPrefixes()
	ResetDestinationApplicationSecurityGroupIds()
	ResetDestinationPortRange()
	ResetDestinationPortRanges()
	ResetDirection()
	ResetName()
	ResetPriority()
	ResetProtocol()
	ResetSourceAddressPrefix()
	ResetSourceAddressPrefixes()
	ResetSourceApplicationSecurityGroupIds()
	ResetSourcePortRange()
	ResetSourcePortRanges()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkSecurityGroupSecurityRuleOutputReference
type jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) Access() *string {
	var returns *string
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) AccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) DestinationAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) DestinationAddressPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddressPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) DestinationAddressPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddressPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) DestinationAddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) DestinationApplicationSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationApplicationSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) DestinationApplicationSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationApplicationSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) DestinationPortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) DestinationPortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) DestinationPortRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPortRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) DestinationPortRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPortRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) SourceAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) SourceAddressPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) SourceAddressPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) SourceAddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) SourceApplicationSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceApplicationSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) SourceApplicationSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceApplicationSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) SourcePortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) SourcePortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) SourcePortRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePortRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) SourcePortRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePortRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkSecurityGroupSecurityRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetworkSecurityGroupSecurityRuleOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkSecurityGroupSecurityRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkSecurityGroup.NetworkSecurityGroupSecurityRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetworkSecurityGroupSecurityRuleOutputReference_Override(n NetworkSecurityGroupSecurityRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkSecurityGroup.NetworkSecurityGroupSecurityRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetAccess(val *string) {
	if err := j.validateSetAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"access",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetDestinationAddressPrefix(val *string) {
	if err := j.validateSetDestinationAddressPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddressPrefix",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetDestinationAddressPrefixes(val *[]*string) {
	if err := j.validateSetDestinationAddressPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddressPrefixes",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetDestinationApplicationSecurityGroupIds(val *[]*string) {
	if err := j.validateSetDestinationApplicationSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationApplicationSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetDestinationPortRange(val *string) {
	if err := j.validateSetDestinationPortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPortRange",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetDestinationPortRanges(val *[]*string) {
	if err := j.validateSetDestinationPortRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPortRanges",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetSourceAddressPrefix(val *string) {
	if err := j.validateSetSourceAddressPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddressPrefix",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetSourceAddressPrefixes(val *[]*string) {
	if err := j.validateSetSourceAddressPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddressPrefixes",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetSourceApplicationSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSourceApplicationSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceApplicationSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetSourcePortRange(val *string) {
	if err := j.validateSetSourcePortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePortRange",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetSourcePortRanges(val *[]*string) {
	if err := j.validateSetSourcePortRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePortRanges",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetAccess() {
	_jsii_.InvokeVoid(
		n,
		"resetAccess",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetDestinationAddressPrefix() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationAddressPrefix",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetDestinationAddressPrefixes() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationAddressPrefixes",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetDestinationApplicationSecurityGroupIds() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationApplicationSecurityGroupIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetDestinationPortRange() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationPortRange",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetDestinationPortRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationPortRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetDirection() {
	_jsii_.InvokeVoid(
		n,
		"resetDirection",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		n,
		"resetName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		n,
		"resetPriority",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		n,
		"resetProtocol",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetSourceAddressPrefix() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceAddressPrefix",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetSourceAddressPrefixes() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceAddressPrefixes",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetSourceApplicationSecurityGroupIds() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceApplicationSecurityGroupIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetSourcePortRange() {
	_jsii_.InvokeVoid(
		n,
		"resetSourcePortRange",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ResetSourcePortRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetSourcePortRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

