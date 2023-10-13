// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlmanagedinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/mssqlmanagedinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/mssql_managed_instance azurerm_mssql_managed_instance}.
type MssqlManagedInstance interface {
	cdktf.TerraformResource
	AdministratorLogin() *string
	SetAdministratorLogin(val *string)
	AdministratorLoginInput() *string
	AdministratorLoginPassword() *string
	SetAdministratorLoginPassword(val *string)
	AdministratorLoginPasswordInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Collation() *string
	SetCollation(val *string)
	CollationInput() *string
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
	DnsZonePartnerId() *string
	SetDnsZonePartnerId(val *string)
	DnsZonePartnerIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Fqdn() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	Identity() MssqlManagedInstanceIdentityOutputReference
	IdentityInput() *MssqlManagedInstanceIdentity
	IdInput() *string
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenanceConfigurationName() *string
	SetMaintenanceConfigurationName(val *string)
	MaintenanceConfigurationNameInput() *string
	MinimumTlsVersion() *string
	SetMinimumTlsVersion(val *string)
	MinimumTlsVersionInput() *string
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
	ProxyOverride() *string
	SetProxyOverride(val *string)
	ProxyOverrideInput() *string
	PublicDataEndpointEnabled() interface{}
	SetPublicDataEndpointEnabled(val interface{})
	PublicDataEndpointEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	StorageAccountType() *string
	SetStorageAccountType(val *string)
	StorageAccountTypeInput() *string
	StorageSizeInGb() *float64
	SetStorageSizeInGb(val *float64)
	StorageSizeInGbInput() *float64
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MssqlManagedInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimezoneId() *string
	SetTimezoneId(val *string)
	TimezoneIdInput() *string
	Vcores() *float64
	SetVcores(val *float64)
	VcoresInput() *float64
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
	PutIdentity(value *MssqlManagedInstanceIdentity)
	PutTimeouts(value *MssqlManagedInstanceTimeouts)
	ResetCollation()
	ResetDnsZonePartnerId()
	ResetId()
	ResetIdentity()
	ResetMaintenanceConfigurationName()
	ResetMinimumTlsVersion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProxyOverride()
	ResetPublicDataEndpointEnabled()
	ResetStorageAccountType()
	ResetTags()
	ResetTimeouts()
	ResetTimezoneId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MssqlManagedInstance
type jsiiProxy_MssqlManagedInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MssqlManagedInstance) AdministratorLogin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) AdministratorLoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) AdministratorLoginPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) AdministratorLoginPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Collation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) CollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) DnsZonePartnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsZonePartnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) DnsZonePartnerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsZonePartnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Identity() MssqlManagedInstanceIdentityOutputReference {
	var returns MssqlManagedInstanceIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) IdentityInput() *MssqlManagedInstanceIdentity {
	var returns *MssqlManagedInstanceIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) MaintenanceConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) MaintenanceConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) MinimumTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) MinimumTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) ProxyOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) ProxyOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) PublicDataEndpointEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicDataEndpointEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) PublicDataEndpointEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicDataEndpointEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) StorageAccountType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) StorageAccountTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) StorageSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) StorageSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSizeInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Timeouts() MssqlManagedInstanceTimeoutsOutputReference {
	var returns MssqlManagedInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) TimezoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) TimezoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) Vcores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstance) VcoresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcoresInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/mssql_managed_instance azurerm_mssql_managed_instance} Resource.
func NewMssqlManagedInstance(scope constructs.Construct, id *string, config *MssqlManagedInstanceConfig) MssqlManagedInstance {
	_init_.Initialize()

	if err := validateNewMssqlManagedInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlManagedInstance{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlManagedInstance.MssqlManagedInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/mssql_managed_instance azurerm_mssql_managed_instance} Resource.
func NewMssqlManagedInstance_Override(m MssqlManagedInstance, scope constructs.Construct, id *string, config *MssqlManagedInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlManagedInstance.MssqlManagedInstance",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetAdministratorLogin(val *string) {
	if err := j.validateSetAdministratorLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorLogin",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetAdministratorLoginPassword(val *string) {
	if err := j.validateSetAdministratorLoginPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorLoginPassword",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetCollation(val *string) {
	if err := j.validateSetCollationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collation",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetDnsZonePartnerId(val *string) {
	if err := j.validateSetDnsZonePartnerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsZonePartnerId",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetMaintenanceConfigurationName(val *string) {
	if err := j.validateSetMaintenanceConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceConfigurationName",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetMinimumTlsVersion(val *string) {
	if err := j.validateSetMinimumTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumTlsVersion",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetProxyOverride(val *string) {
	if err := j.validateSetProxyOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyOverride",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetPublicDataEndpointEnabled(val interface{}) {
	if err := j.validateSetPublicDataEndpointEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicDataEndpointEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetStorageAccountType(val *string) {
	if err := j.validateSetStorageAccountTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountType",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetStorageSizeInGb(val *float64) {
	if err := j.validateSetStorageSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSizeInGb",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetTimezoneId(val *string) {
	if err := j.validateSetTimezoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezoneId",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstance)SetVcores(val *float64) {
	if err := j.validateSetVcoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vcores",
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
func MssqlManagedInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMssqlManagedInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlManagedInstance.MssqlManagedInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MssqlManagedInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMssqlManagedInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlManagedInstance.MssqlManagedInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MssqlManagedInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMssqlManagedInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlManagedInstance.MssqlManagedInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MssqlManagedInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.mssqlManagedInstance.MssqlManagedInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MssqlManagedInstance) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MssqlManagedInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstance) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MssqlManagedInstance) PutIdentity(value *MssqlManagedInstanceIdentity) {
	if err := m.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIdentity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlManagedInstance) PutTimeouts(value *MssqlManagedInstanceTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlManagedInstance) ResetCollation() {
	_jsii_.InvokeVoid(
		m,
		"resetCollation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstance) ResetDnsZonePartnerId() {
	_jsii_.InvokeVoid(
		m,
		"resetDnsZonePartnerId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstance) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstance) ResetIdentity() {
	_jsii_.InvokeVoid(
		m,
		"resetIdentity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstance) ResetMaintenanceConfigurationName() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenanceConfigurationName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstance) ResetMinimumTlsVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetMinimumTlsVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstance) ResetProxyOverride() {
	_jsii_.InvokeVoid(
		m,
		"resetProxyOverride",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstance) ResetPublicDataEndpointEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetPublicDataEndpointEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstance) ResetStorageAccountType() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageAccountType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstance) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstance) ResetTimezoneId() {
	_jsii_.InvokeVoid(
		m,
		"resetTimezoneId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

