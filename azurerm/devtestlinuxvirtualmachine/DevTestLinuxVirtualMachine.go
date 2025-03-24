// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devtestlinuxvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/devtestlinuxvirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/dev_test_linux_virtual_machine azurerm_dev_test_linux_virtual_machine}.
type DevTestLinuxVirtualMachine interface {
	cdktf.TerraformResource
	AllowClaim() interface{}
	SetAllowClaim(val interface{})
	AllowClaimInput() interface{}
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
	DisallowPublicIpAddress() interface{}
	SetDisallowPublicIpAddress(val interface{})
	DisallowPublicIpAddressInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Fqdn() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GalleryImageReference() DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference
	GalleryImageReferenceInput() *DevTestLinuxVirtualMachineGalleryImageReference
	Id() *string
	SetId(val *string)
	IdInput() *string
	InboundNatRule() DevTestLinuxVirtualMachineInboundNatRuleList
	InboundNatRuleInput() interface{}
	LabName() *string
	SetLabName(val *string)
	LabNameInput() *string
	LabSubnetName() *string
	SetLabSubnetName(val *string)
	LabSubnetNameInput() *string
	LabVirtualNetworkId() *string
	SetLabVirtualNetworkId(val *string)
	LabVirtualNetworkIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Notes() *string
	SetNotes(val *string)
	NotesInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
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
	Size() *string
	SetSize(val *string)
	SizeInput() *string
	SshKey() *string
	SetSshKey(val *string)
	SshKeyInput() *string
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DevTestLinuxVirtualMachineTimeoutsOutputReference
	TimeoutsInput() interface{}
	UniqueIdentifier() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	PutGalleryImageReference(value *DevTestLinuxVirtualMachineGalleryImageReference)
	PutInboundNatRule(value interface{})
	PutTimeouts(value *DevTestLinuxVirtualMachineTimeouts)
	ResetAllowClaim()
	ResetDisallowPublicIpAddress()
	ResetId()
	ResetInboundNatRule()
	ResetNotes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetSshKey()
	ResetTags()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DevTestLinuxVirtualMachine
type jsiiProxy_DevTestLinuxVirtualMachine struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) AllowClaim() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) AllowClaimInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) DisallowPublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disallowPublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) DisallowPublicIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disallowPublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) GalleryImageReference() DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference {
	var returns DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference
	_jsii_.Get(
		j,
		"galleryImageReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) GalleryImageReferenceInput() *DevTestLinuxVirtualMachineGalleryImageReference {
	var returns *DevTestLinuxVirtualMachineGalleryImageReference
	_jsii_.Get(
		j,
		"galleryImageReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) InboundNatRule() DevTestLinuxVirtualMachineInboundNatRuleList {
	var returns DevTestLinuxVirtualMachineInboundNatRuleList
	_jsii_.Get(
		j,
		"inboundNatRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) InboundNatRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inboundNatRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) LabName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) LabNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) LabSubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labSubnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) LabSubnetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labSubnetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) LabVirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labVirtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) LabVirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labVirtualNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Notes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) NotesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Size() *string {
	var returns *string
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SshKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SshKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Timeouts() DevTestLinuxVirtualMachineTimeoutsOutputReference {
	var returns DevTestLinuxVirtualMachineTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) UniqueIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/dev_test_linux_virtual_machine azurerm_dev_test_linux_virtual_machine} Resource.
func NewDevTestLinuxVirtualMachine(scope constructs.Construct, id *string, config *DevTestLinuxVirtualMachineConfig) DevTestLinuxVirtualMachine {
	_init_.Initialize()

	if err := validateNewDevTestLinuxVirtualMachineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DevTestLinuxVirtualMachine{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachine",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/dev_test_linux_virtual_machine azurerm_dev_test_linux_virtual_machine} Resource.
func NewDevTestLinuxVirtualMachine_Override(d DevTestLinuxVirtualMachine, scope constructs.Construct, id *string, config *DevTestLinuxVirtualMachineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachine",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetAllowClaim(val interface{}) {
	if err := j.validateSetAllowClaimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowClaim",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetDisallowPublicIpAddress(val interface{}) {
	if err := j.validateSetDisallowPublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disallowPublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetLabName(val *string) {
	if err := j.validateSetLabNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labName",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetLabSubnetName(val *string) {
	if err := j.validateSetLabSubnetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labSubnetName",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetLabVirtualNetworkId(val *string) {
	if err := j.validateSetLabVirtualNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labVirtualNetworkId",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetNotes(val *string) {
	if err := j.validateSetNotesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notes",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetSize(val *string) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetSshKey(val *string) {
	if err := j.validateSetSshKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKey",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Generates CDKTF code for importing a DevTestLinuxVirtualMachine resource upon running "cdktf plan <stack-name>".
func DevTestLinuxVirtualMachine_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDevTestLinuxVirtualMachine_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachine",
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
func DevTestLinuxVirtualMachine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDevTestLinuxVirtualMachine_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DevTestLinuxVirtualMachine_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDevTestLinuxVirtualMachine_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachine",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DevTestLinuxVirtualMachine_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDevTestLinuxVirtualMachine_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachine",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DevTestLinuxVirtualMachine_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachine",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) PutGalleryImageReference(value *DevTestLinuxVirtualMachineGalleryImageReference) {
	if err := d.validatePutGalleryImageReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGalleryImageReference",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) PutInboundNatRule(value interface{}) {
	if err := d.validatePutInboundNatRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInboundNatRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) PutTimeouts(value *DevTestLinuxVirtualMachineTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) ResetAllowClaim() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowClaim",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) ResetDisallowPublicIpAddress() {
	_jsii_.InvokeVoid(
		d,
		"resetDisallowPublicIpAddress",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) ResetInboundNatRule() {
	_jsii_.InvokeVoid(
		d,
		"resetInboundNatRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) ResetNotes() {
	_jsii_.InvokeVoid(
		d,
		"resetNotes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) ResetSshKey() {
	_jsii_.InvokeVoid(
		d,
		"resetSshKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

