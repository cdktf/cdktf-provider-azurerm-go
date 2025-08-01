// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nginxdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/nginxdeployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/nginx_deployment azurerm_nginx_deployment}.
type NginxDeployment interface {
	cdktf.TerraformResource
	AutomaticUpgradeChannel() *string
	SetAutomaticUpgradeChannel(val *string)
	AutomaticUpgradeChannelInput() *string
	AutoScaleProfile() NginxDeploymentAutoScaleProfileList
	AutoScaleProfileInput() interface{}
	Capacity() *float64
	SetCapacity(val *float64)
	CapacityInput() *float64
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
	DataplaneApiEndpoint() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiagnoseSupportEnabled() interface{}
	SetDiagnoseSupportEnabled(val interface{})
	DiagnoseSupportEnabledInput() interface{}
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FrontendPrivate() NginxDeploymentFrontendPrivateList
	FrontendPrivateInput() interface{}
	FrontendPublic() NginxDeploymentFrontendPublicOutputReference
	FrontendPublicInput() *NginxDeploymentFrontendPublic
	Id() *string
	SetId(val *string)
	Identity() NginxDeploymentIdentityOutputReference
	IdentityInput() *NginxDeploymentIdentity
	IdInput() *string
	IpAddress() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LoggingStorageAccount() NginxDeploymentLoggingStorageAccountList
	LoggingStorageAccountInput() interface{}
	ManagedResourceGroup() *string
	SetManagedResourceGroup(val *string)
	ManagedResourceGroupInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterface() NginxDeploymentNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	NginxVersion() *string
	// The tree node.
	Node() constructs.Node
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
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NginxDeploymentTimeoutsOutputReference
	TimeoutsInput() interface{}
	WebApplicationFirewall() NginxDeploymentWebApplicationFirewallOutputReference
	WebApplicationFirewallInput() *NginxDeploymentWebApplicationFirewall
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
	PutAutoScaleProfile(value interface{})
	PutFrontendPrivate(value interface{})
	PutFrontendPublic(value *NginxDeploymentFrontendPublic)
	PutIdentity(value *NginxDeploymentIdentity)
	PutLoggingStorageAccount(value interface{})
	PutNetworkInterface(value interface{})
	PutTimeouts(value *NginxDeploymentTimeouts)
	PutWebApplicationFirewall(value *NginxDeploymentWebApplicationFirewall)
	ResetAutomaticUpgradeChannel()
	ResetAutoScaleProfile()
	ResetCapacity()
	ResetDiagnoseSupportEnabled()
	ResetEmail()
	ResetFrontendPrivate()
	ResetFrontendPublic()
	ResetId()
	ResetIdentity()
	ResetLoggingStorageAccount()
	ResetManagedResourceGroup()
	ResetNetworkInterface()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTimeouts()
	ResetWebApplicationFirewall()
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

// The jsii proxy struct for NginxDeployment
type jsiiProxy_NginxDeployment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NginxDeployment) AutomaticUpgradeChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticUpgradeChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) AutomaticUpgradeChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticUpgradeChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) AutoScaleProfile() NginxDeploymentAutoScaleProfileList {
	var returns NginxDeploymentAutoScaleProfileList
	_jsii_.Get(
		j,
		"autoScaleProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) AutoScaleProfileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScaleProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Capacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) CapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) DataplaneApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataplaneApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) DiagnoseSupportEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diagnoseSupportEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) DiagnoseSupportEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diagnoseSupportEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) FrontendPrivate() NginxDeploymentFrontendPrivateList {
	var returns NginxDeploymentFrontendPrivateList
	_jsii_.Get(
		j,
		"frontendPrivate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) FrontendPrivateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frontendPrivateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) FrontendPublic() NginxDeploymentFrontendPublicOutputReference {
	var returns NginxDeploymentFrontendPublicOutputReference
	_jsii_.Get(
		j,
		"frontendPublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) FrontendPublicInput() *NginxDeploymentFrontendPublic {
	var returns *NginxDeploymentFrontendPublic
	_jsii_.Get(
		j,
		"frontendPublicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Identity() NginxDeploymentIdentityOutputReference {
	var returns NginxDeploymentIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) IdentityInput() *NginxDeploymentIdentity {
	var returns *NginxDeploymentIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) LoggingStorageAccount() NginxDeploymentLoggingStorageAccountList {
	var returns NginxDeploymentLoggingStorageAccountList
	_jsii_.Get(
		j,
		"loggingStorageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) LoggingStorageAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingStorageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) ManagedResourceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedResourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) ManagedResourceGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedResourceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) NetworkInterface() NginxDeploymentNetworkInterfaceList {
	var returns NginxDeploymentNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) NginxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nginxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) Timeouts() NginxDeploymentTimeoutsOutputReference {
	var returns NginxDeploymentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) WebApplicationFirewall() NginxDeploymentWebApplicationFirewallOutputReference {
	var returns NginxDeploymentWebApplicationFirewallOutputReference
	_jsii_.Get(
		j,
		"webApplicationFirewall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeployment) WebApplicationFirewallInput() *NginxDeploymentWebApplicationFirewall {
	var returns *NginxDeploymentWebApplicationFirewall
	_jsii_.Get(
		j,
		"webApplicationFirewallInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/nginx_deployment azurerm_nginx_deployment} Resource.
func NewNginxDeployment(scope constructs.Construct, id *string, config *NginxDeploymentConfig) NginxDeployment {
	_init_.Initialize()

	if err := validateNewNginxDeploymentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NginxDeployment{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.nginxDeployment.NginxDeployment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/nginx_deployment azurerm_nginx_deployment} Resource.
func NewNginxDeployment_Override(n NginxDeployment, scope constructs.Construct, id *string, config *NginxDeploymentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.nginxDeployment.NginxDeployment",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NginxDeployment)SetAutomaticUpgradeChannel(val *string) {
	if err := j.validateSetAutomaticUpgradeChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticUpgradeChannel",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetCapacity(val *float64) {
	if err := j.validateSetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetDiagnoseSupportEnabled(val interface{}) {
	if err := j.validateSetDiagnoseSupportEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diagnoseSupportEnabled",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetManagedResourceGroup(val *string) {
	if err := j.validateSetManagedResourceGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedResourceGroup",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetSku(val *string) {
	if err := j.validateSetSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_NginxDeployment)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a NginxDeployment resource upon running "cdktf plan <stack-name>".
func NginxDeployment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNginxDeployment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.nginxDeployment.NginxDeployment",
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
func NginxDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNginxDeployment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.nginxDeployment.NginxDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NginxDeployment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNginxDeployment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.nginxDeployment.NginxDeployment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NginxDeployment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNginxDeployment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.nginxDeployment.NginxDeployment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NginxDeployment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.nginxDeployment.NginxDeployment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NginxDeployment) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NginxDeployment) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NginxDeployment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NginxDeployment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NginxDeployment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NginxDeployment) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NginxDeployment) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NginxDeployment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NginxDeployment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NginxDeployment) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NginxDeployment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NginxDeployment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NginxDeployment) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NginxDeployment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NginxDeployment) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NginxDeployment) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NginxDeployment) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NginxDeployment) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NginxDeployment) PutAutoScaleProfile(value interface{}) {
	if err := n.validatePutAutoScaleProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putAutoScaleProfile",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NginxDeployment) PutFrontendPrivate(value interface{}) {
	if err := n.validatePutFrontendPrivateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putFrontendPrivate",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NginxDeployment) PutFrontendPublic(value *NginxDeploymentFrontendPublic) {
	if err := n.validatePutFrontendPublicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putFrontendPublic",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NginxDeployment) PutIdentity(value *NginxDeploymentIdentity) {
	if err := n.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putIdentity",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NginxDeployment) PutLoggingStorageAccount(value interface{}) {
	if err := n.validatePutLoggingStorageAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putLoggingStorageAccount",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NginxDeployment) PutNetworkInterface(value interface{}) {
	if err := n.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NginxDeployment) PutTimeouts(value *NginxDeploymentTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NginxDeployment) PutWebApplicationFirewall(value *NginxDeploymentWebApplicationFirewall) {
	if err := n.validatePutWebApplicationFirewallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putWebApplicationFirewall",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NginxDeployment) ResetAutomaticUpgradeChannel() {
	_jsii_.InvokeVoid(
		n,
		"resetAutomaticUpgradeChannel",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) ResetAutoScaleProfile() {
	_jsii_.InvokeVoid(
		n,
		"resetAutoScaleProfile",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) ResetCapacity() {
	_jsii_.InvokeVoid(
		n,
		"resetCapacity",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) ResetDiagnoseSupportEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetDiagnoseSupportEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) ResetEmail() {
	_jsii_.InvokeVoid(
		n,
		"resetEmail",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) ResetFrontendPrivate() {
	_jsii_.InvokeVoid(
		n,
		"resetFrontendPrivate",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) ResetFrontendPublic() {
	_jsii_.InvokeVoid(
		n,
		"resetFrontendPublic",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) ResetIdentity() {
	_jsii_.InvokeVoid(
		n,
		"resetIdentity",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) ResetLoggingStorageAccount() {
	_jsii_.InvokeVoid(
		n,
		"resetLoggingStorageAccount",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) ResetManagedResourceGroup() {
	_jsii_.InvokeVoid(
		n,
		"resetManagedResourceGroup",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		n,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) ResetTags() {
	_jsii_.InvokeVoid(
		n,
		"resetTags",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) ResetWebApplicationFirewall() {
	_jsii_.InvokeVoid(
		n,
		"resetWebApplicationFirewall",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeployment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NginxDeployment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NginxDeployment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NginxDeployment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NginxDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NginxDeployment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

