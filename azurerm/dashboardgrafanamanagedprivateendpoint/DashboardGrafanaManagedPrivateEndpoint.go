// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboardgrafanamanagedprivateendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/dashboardgrafanamanagedprivateendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/dashboard_grafana_managed_private_endpoint azurerm_dashboard_grafana_managed_private_endpoint}.
type DashboardGrafanaManagedPrivateEndpoint interface {
	cdktf.TerraformResource
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GrafanaId() *string
	SetGrafanaId(val *string)
	GrafanaIdInput() *string
	GroupIds() *[]*string
	SetGroupIds(val *[]*string)
	GroupIdsInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
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
	PrivateLinkResourceId() *string
	SetPrivateLinkResourceId(val *string)
	PrivateLinkResourceIdInput() *string
	PrivateLinkResourceRegion() *string
	SetPrivateLinkResourceRegion(val *string)
	PrivateLinkResourceRegionInput() *string
	PrivateLinkServiceUrl() *string
	SetPrivateLinkServiceUrl(val *string)
	PrivateLinkServiceUrlInput() *string
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
	RequestMessage() *string
	SetRequestMessage(val *string)
	RequestMessageInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DashboardGrafanaManagedPrivateEndpointTimeoutsOutputReference
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
	PutTimeouts(value *DashboardGrafanaManagedPrivateEndpointTimeouts)
	ResetGroupIds()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateLinkResourceRegion()
	ResetPrivateLinkServiceUrl()
	ResetRequestMessage()
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

// The jsii proxy struct for DashboardGrafanaManagedPrivateEndpoint
type jsiiProxy_DashboardGrafanaManagedPrivateEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) GrafanaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) GrafanaIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) GroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) GroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) PrivateLinkResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) PrivateLinkResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) PrivateLinkResourceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkResourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) PrivateLinkResourceRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkResourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) PrivateLinkServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkServiceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) PrivateLinkServiceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkServiceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) RequestMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) RequestMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) Timeouts() DashboardGrafanaManagedPrivateEndpointTimeoutsOutputReference {
	var returns DashboardGrafanaManagedPrivateEndpointTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/dashboard_grafana_managed_private_endpoint azurerm_dashboard_grafana_managed_private_endpoint} Resource.
func NewDashboardGrafanaManagedPrivateEndpoint(scope constructs.Construct, id *string, config *DashboardGrafanaManagedPrivateEndpointConfig) DashboardGrafanaManagedPrivateEndpoint {
	_init_.Initialize()

	if err := validateNewDashboardGrafanaManagedPrivateEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardGrafanaManagedPrivateEndpoint{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dashboardGrafanaManagedPrivateEndpoint.DashboardGrafanaManagedPrivateEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/dashboard_grafana_managed_private_endpoint azurerm_dashboard_grafana_managed_private_endpoint} Resource.
func NewDashboardGrafanaManagedPrivateEndpoint_Override(d DashboardGrafanaManagedPrivateEndpoint, scope constructs.Construct, id *string, config *DashboardGrafanaManagedPrivateEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dashboardGrafanaManagedPrivateEndpoint.DashboardGrafanaManagedPrivateEndpoint",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetGrafanaId(val *string) {
	if err := j.validateSetGrafanaIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grafanaId",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetGroupIds(val *[]*string) {
	if err := j.validateSetGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupIds",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetPrivateLinkResourceId(val *string) {
	if err := j.validateSetPrivateLinkResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateLinkResourceId",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetPrivateLinkResourceRegion(val *string) {
	if err := j.validateSetPrivateLinkResourceRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateLinkResourceRegion",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetPrivateLinkServiceUrl(val *string) {
	if err := j.validateSetPrivateLinkServiceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateLinkServiceUrl",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetRequestMessage(val *string) {
	if err := j.validateSetRequestMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestMessage",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a DashboardGrafanaManagedPrivateEndpoint resource upon running "cdktf plan <stack-name>".
func DashboardGrafanaManagedPrivateEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDashboardGrafanaManagedPrivateEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dashboardGrafanaManagedPrivateEndpoint.DashboardGrafanaManagedPrivateEndpoint",
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
func DashboardGrafanaManagedPrivateEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDashboardGrafanaManagedPrivateEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dashboardGrafanaManagedPrivateEndpoint.DashboardGrafanaManagedPrivateEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DashboardGrafanaManagedPrivateEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDashboardGrafanaManagedPrivateEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dashboardGrafanaManagedPrivateEndpoint.DashboardGrafanaManagedPrivateEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DashboardGrafanaManagedPrivateEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDashboardGrafanaManagedPrivateEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dashboardGrafanaManagedPrivateEndpoint.DashboardGrafanaManagedPrivateEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DashboardGrafanaManagedPrivateEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dashboardGrafanaManagedPrivateEndpoint.DashboardGrafanaManagedPrivateEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) PutTimeouts(value *DashboardGrafanaManagedPrivateEndpointTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) ResetGroupIds() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) ResetPrivateLinkResourceRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateLinkResourceRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) ResetPrivateLinkServiceUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateLinkServiceUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) ResetRequestMessage() {
	_jsii_.InvokeVoid(
		d,
		"resetRequestMessage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaManagedPrivateEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

