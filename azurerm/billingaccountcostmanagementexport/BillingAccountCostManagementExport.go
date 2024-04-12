// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package billingaccountcostmanagementexport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/billingaccountcostmanagementexport/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/billing_account_cost_management_export azurerm_billing_account_cost_management_export}.
type BillingAccountCostManagementExport interface {
	cdktf.TerraformResource
	Active() interface{}
	SetActive(val interface{})
	ActiveInput() interface{}
	BillingAccountId() *string
	SetBillingAccountId(val *string)
	BillingAccountIdInput() *string
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
	ExportDataOptions() BillingAccountCostManagementExportExportDataOptionsOutputReference
	ExportDataOptionsInput() *BillingAccountCostManagementExportExportDataOptions
	ExportDataStorageLocation() BillingAccountCostManagementExportExportDataStorageLocationOutputReference
	ExportDataStorageLocationInput() *BillingAccountCostManagementExportExportDataStorageLocation
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
	RecurrencePeriodEndDate() *string
	SetRecurrencePeriodEndDate(val *string)
	RecurrencePeriodEndDateInput() *string
	RecurrencePeriodStartDate() *string
	SetRecurrencePeriodStartDate(val *string)
	RecurrencePeriodStartDateInput() *string
	RecurrenceType() *string
	SetRecurrenceType(val *string)
	RecurrenceTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BillingAccountCostManagementExportTimeoutsOutputReference
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
	PutExportDataOptions(value *BillingAccountCostManagementExportExportDataOptions)
	PutExportDataStorageLocation(value *BillingAccountCostManagementExportExportDataStorageLocation)
	PutTimeouts(value *BillingAccountCostManagementExportTimeouts)
	ResetActive()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for BillingAccountCostManagementExport
type jsiiProxy_BillingAccountCostManagementExport struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BillingAccountCostManagementExport) Active() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) ActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) BillingAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) BillingAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) ExportDataOptions() BillingAccountCostManagementExportExportDataOptionsOutputReference {
	var returns BillingAccountCostManagementExportExportDataOptionsOutputReference
	_jsii_.Get(
		j,
		"exportDataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) ExportDataOptionsInput() *BillingAccountCostManagementExportExportDataOptions {
	var returns *BillingAccountCostManagementExportExportDataOptions
	_jsii_.Get(
		j,
		"exportDataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) ExportDataStorageLocation() BillingAccountCostManagementExportExportDataStorageLocationOutputReference {
	var returns BillingAccountCostManagementExportExportDataStorageLocationOutputReference
	_jsii_.Get(
		j,
		"exportDataStorageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) ExportDataStorageLocationInput() *BillingAccountCostManagementExportExportDataStorageLocation {
	var returns *BillingAccountCostManagementExportExportDataStorageLocation
	_jsii_.Get(
		j,
		"exportDataStorageLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) RecurrencePeriodEndDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrencePeriodEndDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) RecurrencePeriodEndDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrencePeriodEndDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) RecurrencePeriodStartDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrencePeriodStartDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) RecurrencePeriodStartDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrencePeriodStartDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) RecurrenceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrenceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) RecurrenceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrenceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) Timeouts() BillingAccountCostManagementExportTimeoutsOutputReference {
	var returns BillingAccountCostManagementExportTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingAccountCostManagementExport) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/billing_account_cost_management_export azurerm_billing_account_cost_management_export} Resource.
func NewBillingAccountCostManagementExport(scope constructs.Construct, id *string, config *BillingAccountCostManagementExportConfig) BillingAccountCostManagementExport {
	_init_.Initialize()

	if err := validateNewBillingAccountCostManagementExportParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BillingAccountCostManagementExport{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.billingAccountCostManagementExport.BillingAccountCostManagementExport",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/billing_account_cost_management_export azurerm_billing_account_cost_management_export} Resource.
func NewBillingAccountCostManagementExport_Override(b BillingAccountCostManagementExport, scope constructs.Construct, id *string, config *BillingAccountCostManagementExportConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.billingAccountCostManagementExport.BillingAccountCostManagementExport",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BillingAccountCostManagementExport)SetActive(val interface{}) {
	if err := j.validateSetActiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"active",
		val,
	)
}

func (j *jsiiProxy_BillingAccountCostManagementExport)SetBillingAccountId(val *string) {
	if err := j.validateSetBillingAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingAccountId",
		val,
	)
}

func (j *jsiiProxy_BillingAccountCostManagementExport)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BillingAccountCostManagementExport)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BillingAccountCostManagementExport)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BillingAccountCostManagementExport)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BillingAccountCostManagementExport)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BillingAccountCostManagementExport)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BillingAccountCostManagementExport)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BillingAccountCostManagementExport)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BillingAccountCostManagementExport)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BillingAccountCostManagementExport)SetRecurrencePeriodEndDate(val *string) {
	if err := j.validateSetRecurrencePeriodEndDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recurrencePeriodEndDate",
		val,
	)
}

func (j *jsiiProxy_BillingAccountCostManagementExport)SetRecurrencePeriodStartDate(val *string) {
	if err := j.validateSetRecurrencePeriodStartDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recurrencePeriodStartDate",
		val,
	)
}

func (j *jsiiProxy_BillingAccountCostManagementExport)SetRecurrenceType(val *string) {
	if err := j.validateSetRecurrenceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recurrenceType",
		val,
	)
}

// Generates CDKTF code for importing a BillingAccountCostManagementExport resource upon running "cdktf plan <stack-name>".
func BillingAccountCostManagementExport_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBillingAccountCostManagementExport_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.billingAccountCostManagementExport.BillingAccountCostManagementExport",
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
func BillingAccountCostManagementExport_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBillingAccountCostManagementExport_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.billingAccountCostManagementExport.BillingAccountCostManagementExport",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BillingAccountCostManagementExport_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBillingAccountCostManagementExport_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.billingAccountCostManagementExport.BillingAccountCostManagementExport",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BillingAccountCostManagementExport_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBillingAccountCostManagementExport_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.billingAccountCostManagementExport.BillingAccountCostManagementExport",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BillingAccountCostManagementExport_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.billingAccountCostManagementExport.BillingAccountCostManagementExport",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BillingAccountCostManagementExport) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BillingAccountCostManagementExport) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BillingAccountCostManagementExport) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BillingAccountCostManagementExport) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BillingAccountCostManagementExport) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BillingAccountCostManagementExport) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BillingAccountCostManagementExport) PutExportDataOptions(value *BillingAccountCostManagementExportExportDataOptions) {
	if err := b.validatePutExportDataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putExportDataOptions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BillingAccountCostManagementExport) PutExportDataStorageLocation(value *BillingAccountCostManagementExportExportDataStorageLocation) {
	if err := b.validatePutExportDataStorageLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putExportDataStorageLocation",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BillingAccountCostManagementExport) PutTimeouts(value *BillingAccountCostManagementExportTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BillingAccountCostManagementExport) ResetActive() {
	_jsii_.InvokeVoid(
		b,
		"resetActive",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingAccountCostManagementExport) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingAccountCostManagementExport) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingAccountCostManagementExport) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingAccountCostManagementExport) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingAccountCostManagementExport) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

