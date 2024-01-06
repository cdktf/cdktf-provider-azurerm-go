// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/networksecurityrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/network_security_rule azurerm_network_security_rule}.
type NetworkSecurityRule interface {
	cdktf.TerraformResource
	Access() *string
	SetAccess(val *string)
	AccessInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkSecurityGroupName() *string
	SetNetworkSecurityGroupName(val *string)
	NetworkSecurityGroupNameInput() *string
	// The tree node.
	Node() constructs.Node
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
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
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NetworkSecurityRuleTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *NetworkSecurityRuleTimeouts)
	ResetDescription()
	ResetDestinationAddressPrefix()
	ResetDestinationAddressPrefixes()
	ResetDestinationApplicationSecurityGroupIds()
	ResetDestinationPortRange()
	ResetDestinationPortRanges()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSourceAddressPrefix()
	ResetSourceAddressPrefixes()
	ResetSourceApplicationSecurityGroupIds()
	ResetSourcePortRange()
	ResetSourcePortRanges()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for NetworkSecurityRule
type jsiiProxy_NetworkSecurityRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkSecurityRule) Access() *string {
	var returns *string
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) AccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationAddressPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddressPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationAddressPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddressPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationAddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationApplicationSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationApplicationSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationApplicationSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationApplicationSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationPortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationPortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationPortRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPortRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationPortRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPortRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) NetworkSecurityGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) NetworkSecurityGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourceAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourceAddressPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourceAddressPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourceAddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourceApplicationSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceApplicationSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourceApplicationSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceApplicationSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourcePortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourcePortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourcePortRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePortRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourcePortRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePortRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Timeouts() NetworkSecurityRuleTimeoutsOutputReference {
	var returns NetworkSecurityRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/network_security_rule azurerm_network_security_rule} Resource.
func NewNetworkSecurityRule(scope constructs.Construct, id *string, config *NetworkSecurityRuleConfig) NetworkSecurityRule {
	_init_.Initialize()

	if err := validateNewNetworkSecurityRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkSecurityRule{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkSecurityRule.NetworkSecurityRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/network_security_rule azurerm_network_security_rule} Resource.
func NewNetworkSecurityRule_Override(n NetworkSecurityRule, scope constructs.Construct, id *string, config *NetworkSecurityRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkSecurityRule.NetworkSecurityRule",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetAccess(val *string) {
	if err := j.validateSetAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"access",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetDestinationAddressPrefix(val *string) {
	if err := j.validateSetDestinationAddressPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddressPrefix",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetDestinationAddressPrefixes(val *[]*string) {
	if err := j.validateSetDestinationAddressPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddressPrefixes",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetDestinationApplicationSecurityGroupIds(val *[]*string) {
	if err := j.validateSetDestinationApplicationSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationApplicationSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetDestinationPortRange(val *string) {
	if err := j.validateSetDestinationPortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPortRange",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetDestinationPortRanges(val *[]*string) {
	if err := j.validateSetDestinationPortRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPortRanges",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetNetworkSecurityGroupName(val *string) {
	if err := j.validateSetNetworkSecurityGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkSecurityGroupName",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetSourceAddressPrefix(val *string) {
	if err := j.validateSetSourceAddressPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddressPrefix",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetSourceAddressPrefixes(val *[]*string) {
	if err := j.validateSetSourceAddressPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddressPrefixes",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetSourceApplicationSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSourceApplicationSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceApplicationSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetSourcePortRange(val *string) {
	if err := j.validateSetSourcePortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePortRange",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule)SetSourcePortRanges(val *[]*string) {
	if err := j.validateSetSourcePortRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePortRanges",
		val,
	)
}

// Generates CDKTF code for importing a NetworkSecurityRule resource upon running "cdktf plan <stack-name>".
func NetworkSecurityRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkSecurityRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.networkSecurityRule.NetworkSecurityRule",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func NetworkSecurityRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkSecurityRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.networkSecurityRule.NetworkSecurityRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkSecurityRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkSecurityRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.networkSecurityRule.NetworkSecurityRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkSecurityRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkSecurityRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.networkSecurityRule.NetworkSecurityRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkSecurityRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.networkSecurityRule.NetworkSecurityRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetworkSecurityRule) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetworkSecurityRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkSecurityRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkSecurityRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkSecurityRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkSecurityRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkSecurityRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkSecurityRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkSecurityRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkSecurityRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkSecurityRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetworkSecurityRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkSecurityRule) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkSecurityRule) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkSecurityRule) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkSecurityRule) PutTimeouts(value *NetworkSecurityRuleTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetDestinationAddressPrefix() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationAddressPrefix",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetDestinationAddressPrefixes() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationAddressPrefixes",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetDestinationApplicationSecurityGroupIds() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationApplicationSecurityGroupIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetDestinationPortRange() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationPortRange",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetDestinationPortRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationPortRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetSourceAddressPrefix() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceAddressPrefix",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetSourceAddressPrefixes() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceAddressPrefixes",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetSourceApplicationSecurityGroupIds() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceApplicationSecurityGroupIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetSourcePortRange() {
	_jsii_.InvokeVoid(
		n,
		"resetSourcePortRange",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetSourcePortRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetSourcePortRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

