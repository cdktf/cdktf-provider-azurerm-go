package devtestlinuxvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/devtestlinuxvirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine azurerm_dev_test_linux_virtual_machine}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
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

func (j *jsiiProxy_DevTestLinuxVirtualMachine) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine azurerm_dev_test_linux_virtual_machine} Resource.
func NewDevTestLinuxVirtualMachine(scope constructs.Construct, id *string, config *DevTestLinuxVirtualMachineConfig) DevTestLinuxVirtualMachine {
	_init_.Initialize()

	j := jsiiProxy_DevTestLinuxVirtualMachine{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachine",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine azurerm_dev_test_linux_virtual_machine} Resource.
func NewDevTestLinuxVirtualMachine_Override(d DevTestLinuxVirtualMachine, scope constructs.Construct, id *string, config *DevTestLinuxVirtualMachineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachine",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetAllowClaim(val interface{}) {
	_jsii_.Set(
		j,
		"allowClaim",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetDisallowPublicIpAddress(val interface{}) {
	_jsii_.Set(
		j,
		"disallowPublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetLabName(val *string) {
	_jsii_.Set(
		j,
		"labName",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetLabSubnetName(val *string) {
	_jsii_.Set(
		j,
		"labSubnetName",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetLabVirtualNetworkId(val *string) {
	_jsii_.Set(
		j,
		"labVirtualNetworkId",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetNotes(val *string) {
	_jsii_.Set(
		j,
		"notes",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetSize(val *string) {
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetSshKey(val *string) {
	_jsii_.Set(
		j,
		"sshKey",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetStorageType(val *string) {
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachine) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachine",
		"isConstruct",
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

func (d *jsiiProxy_DevTestLinuxVirtualMachine) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) PutGalleryImageReference(value *DevTestLinuxVirtualMachineGalleryImageReference) {
	_jsii_.InvokeVoid(
		d,
		"putGalleryImageReference",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) PutInboundNatRule(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putInboundNatRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachine) PutTimeouts(value *DevTestLinuxVirtualMachineTimeouts) {
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

type DevTestLinuxVirtualMachineConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// gallery_image_reference block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#gallery_image_reference DevTestLinuxVirtualMachine#gallery_image_reference}
	GalleryImageReference *DevTestLinuxVirtualMachineGalleryImageReference `field:"required" json:"galleryImageReference" yaml:"galleryImageReference"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#lab_name DevTestLinuxVirtualMachine#lab_name}.
	LabName *string `field:"required" json:"labName" yaml:"labName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#lab_subnet_name DevTestLinuxVirtualMachine#lab_subnet_name}.
	LabSubnetName *string `field:"required" json:"labSubnetName" yaml:"labSubnetName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#lab_virtual_network_id DevTestLinuxVirtualMachine#lab_virtual_network_id}.
	LabVirtualNetworkId *string `field:"required" json:"labVirtualNetworkId" yaml:"labVirtualNetworkId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#location DevTestLinuxVirtualMachine#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#name DevTestLinuxVirtualMachine#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#resource_group_name DevTestLinuxVirtualMachine#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#size DevTestLinuxVirtualMachine#size}.
	Size *string `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#storage_type DevTestLinuxVirtualMachine#storage_type}.
	StorageType *string `field:"required" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#username DevTestLinuxVirtualMachine#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#allow_claim DevTestLinuxVirtualMachine#allow_claim}.
	AllowClaim interface{} `field:"optional" json:"allowClaim" yaml:"allowClaim"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#disallow_public_ip_address DevTestLinuxVirtualMachine#disallow_public_ip_address}.
	DisallowPublicIpAddress interface{} `field:"optional" json:"disallowPublicIpAddress" yaml:"disallowPublicIpAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#id DevTestLinuxVirtualMachine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// inbound_nat_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#inbound_nat_rule DevTestLinuxVirtualMachine#inbound_nat_rule}
	InboundNatRule interface{} `field:"optional" json:"inboundNatRule" yaml:"inboundNatRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#notes DevTestLinuxVirtualMachine#notes}.
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#password DevTestLinuxVirtualMachine#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#ssh_key DevTestLinuxVirtualMachine#ssh_key}.
	SshKey *string `field:"optional" json:"sshKey" yaml:"sshKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#tags DevTestLinuxVirtualMachine#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#timeouts DevTestLinuxVirtualMachine#timeouts}
	Timeouts *DevTestLinuxVirtualMachineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type DevTestLinuxVirtualMachineGalleryImageReference struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#offer DevTestLinuxVirtualMachine#offer}.
	Offer *string `field:"required" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#publisher DevTestLinuxVirtualMachine#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#sku DevTestLinuxVirtualMachine#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#version DevTestLinuxVirtualMachine#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
}

type DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *DevTestLinuxVirtualMachineGalleryImageReference
	SetInternalValue(val *DevTestLinuxVirtualMachineGalleryImageReference)
	Offer() *string
	SetOffer(val *string)
	OfferInput() *string
	Publisher() *string
	SetPublisher(val *string)
	PublisherInput() *string
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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

// The jsii proxy struct for DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference
type jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) InternalValue() *DevTestLinuxVirtualMachineGalleryImageReference {
	var returns *DevTestLinuxVirtualMachineGalleryImageReference
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) Offer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) OfferInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) Publisher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) PublisherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewDevTestLinuxVirtualMachineGalleryImageReferenceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDevTestLinuxVirtualMachineGalleryImageReferenceOutputReference_Override(d DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) SetInternalValue(val *DevTestLinuxVirtualMachineGalleryImageReference) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) SetOffer(val *string) {
	_jsii_.Set(
		j,
		"offer",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) SetPublisher(val *string) {
	_jsii_.Set(
		j,
		"publisher",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) SetSku(val *string) {
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineGalleryImageReferenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DevTestLinuxVirtualMachineInboundNatRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#backend_port DevTestLinuxVirtualMachine#backend_port}.
	BackendPort *float64 `field:"required" json:"backendPort" yaml:"backendPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#protocol DevTestLinuxVirtualMachine#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

type DevTestLinuxVirtualMachineInboundNatRuleList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DevTestLinuxVirtualMachineInboundNatRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DevTestLinuxVirtualMachineInboundNatRuleList
type jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDevTestLinuxVirtualMachineInboundNatRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DevTestLinuxVirtualMachineInboundNatRuleList {
	_init_.Initialize()

	j := jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachineInboundNatRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDevTestLinuxVirtualMachineInboundNatRuleList_Override(d DevTestLinuxVirtualMachineInboundNatRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachineInboundNatRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList) Get(index *float64) DevTestLinuxVirtualMachineInboundNatRuleOutputReference {
	var returns DevTestLinuxVirtualMachineInboundNatRuleOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DevTestLinuxVirtualMachineInboundNatRuleOutputReference interface {
	cdktf.ComplexObject
	BackendPort() *float64
	SetBackendPort(val *float64)
	BackendPortInput() *float64
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
	// Experimental.
	Fqn() *string
	FrontendPort() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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

// The jsii proxy struct for DevTestLinuxVirtualMachineInboundNatRuleOutputReference
type jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) BackendPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backendPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) BackendPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backendPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) FrontendPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frontendPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDevTestLinuxVirtualMachineInboundNatRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DevTestLinuxVirtualMachineInboundNatRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachineInboundNatRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDevTestLinuxVirtualMachineInboundNatRuleOutputReference_Override(d DevTestLinuxVirtualMachineInboundNatRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachineInboundNatRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) SetBackendPort(val *float64) {
	_jsii_.Set(
		j,
		"backendPort",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineInboundNatRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DevTestLinuxVirtualMachineTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#create DevTestLinuxVirtualMachine#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#delete DevTestLinuxVirtualMachine#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#read DevTestLinuxVirtualMachine#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#update DevTestLinuxVirtualMachine#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type DevTestLinuxVirtualMachineTimeoutsOutputReference interface {
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
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Read() *string
	SetRead(val *string)
	ReadInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetCreate()
	ResetDelete()
	ResetRead()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DevTestLinuxVirtualMachineTimeoutsOutputReference
type jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewDevTestLinuxVirtualMachineTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DevTestLinuxVirtualMachineTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachineTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDevTestLinuxVirtualMachineTimeoutsOutputReference_Override(d DevTestLinuxVirtualMachineTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.devTestLinuxVirtualMachine.DevTestLinuxVirtualMachineTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		d,
		"resetCreate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		d,
		"resetDelete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		d,
		"resetRead",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestLinuxVirtualMachineTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

