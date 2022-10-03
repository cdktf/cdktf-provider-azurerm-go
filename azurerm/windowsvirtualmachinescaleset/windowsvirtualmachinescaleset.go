package windowsvirtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/windowsvirtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set azurerm_windows_virtual_machine_scale_set}.
type WindowsVirtualMachineScaleSet interface {
	cdktf.TerraformResource
	AdditionalCapabilities() WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference
	AdditionalCapabilitiesInput() *WindowsVirtualMachineScaleSetAdditionalCapabilities
	AdditionalUnattendContent() WindowsVirtualMachineScaleSetAdditionalUnattendContentList
	AdditionalUnattendContentInput() interface{}
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	AdminUsername() *string
	SetAdminUsername(val *string)
	AdminUsernameInput() *string
	AutomaticInstanceRepair() WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference
	AutomaticInstanceRepairInput() *WindowsVirtualMachineScaleSetAutomaticInstanceRepair
	AutomaticOsUpgradePolicy() WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference
	AutomaticOsUpgradePolicyInput() *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy
	BootDiagnostics() WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference
	BootDiagnosticsInput() *WindowsVirtualMachineScaleSetBootDiagnostics
	CapacityReservationGroupId() *string
	SetCapacityReservationGroupId(val *string)
	CapacityReservationGroupIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComputerNamePrefix() *string
	SetComputerNamePrefix(val *string)
	ComputerNamePrefixInput() *string
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
	CustomData() *string
	SetCustomData(val *string)
	CustomDataInput() *string
	DataDisk() WindowsVirtualMachineScaleSetDataDiskList
	DataDiskInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DoNotRunExtensionsOnOverprovisionedMachines() interface{}
	SetDoNotRunExtensionsOnOverprovisionedMachines(val interface{})
	DoNotRunExtensionsOnOverprovisionedMachinesInput() interface{}
	EdgeZone() *string
	SetEdgeZone(val *string)
	EdgeZoneInput() *string
	EnableAutomaticUpdates() interface{}
	SetEnableAutomaticUpdates(val interface{})
	EnableAutomaticUpdatesInput() interface{}
	EncryptionAtHostEnabled() interface{}
	SetEncryptionAtHostEnabled(val interface{})
	EncryptionAtHostEnabledInput() interface{}
	EvictionPolicy() *string
	SetEvictionPolicy(val *string)
	EvictionPolicyInput() *string
	Extension() WindowsVirtualMachineScaleSetExtensionList
	ExtensionInput() interface{}
	ExtensionOperationsEnabled() interface{}
	SetExtensionOperationsEnabled(val interface{})
	ExtensionOperationsEnabledInput() interface{}
	ExtensionsTimeBudget() *string
	SetExtensionsTimeBudget(val *string)
	ExtensionsTimeBudgetInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GalleryApplications() WindowsVirtualMachineScaleSetGalleryApplicationsList
	GalleryApplicationsInput() interface{}
	HealthProbeId() *string
	SetHealthProbeId(val *string)
	HealthProbeIdInput() *string
	HostGroupId() *string
	SetHostGroupId(val *string)
	HostGroupIdInput() *string
	Id() *string
	SetId(val *string)
	Identity() WindowsVirtualMachineScaleSetIdentityOutputReference
	IdentityInput() *WindowsVirtualMachineScaleSetIdentity
	IdInput() *string
	Instances() *float64
	SetInstances(val *float64)
	InstancesInput() *float64
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
	MaxBidPrice() *float64
	SetMaxBidPrice(val *float64)
	MaxBidPriceInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterface() WindowsVirtualMachineScaleSetNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	// The tree node.
	Node() constructs.Node
	OsDisk() WindowsVirtualMachineScaleSetOsDiskOutputReference
	OsDiskInput() *WindowsVirtualMachineScaleSetOsDisk
	Overprovision() interface{}
	SetOverprovision(val interface{})
	OverprovisionInput() interface{}
	Plan() WindowsVirtualMachineScaleSetPlanOutputReference
	PlanInput() *WindowsVirtualMachineScaleSetPlan
	PlatformFaultDomainCount() *float64
	SetPlatformFaultDomainCount(val *float64)
	PlatformFaultDomainCountInput() *float64
	Priority() *string
	SetPriority(val *string)
	PriorityInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProvisionVmAgent() interface{}
	SetProvisionVmAgent(val interface{})
	ProvisionVmAgentInput() interface{}
	ProximityPlacementGroupId() *string
	SetProximityPlacementGroupId(val *string)
	ProximityPlacementGroupIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RollingUpgradePolicy() WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference
	RollingUpgradePolicyInput() *WindowsVirtualMachineScaleSetRollingUpgradePolicy
	ScaleIn() WindowsVirtualMachineScaleSetScaleInOutputReference
	ScaleInInput() *WindowsVirtualMachineScaleSetScaleIn
	ScaleInPolicy() *string
	SetScaleInPolicy(val *string)
	ScaleInPolicyInput() *string
	Secret() WindowsVirtualMachineScaleSetSecretList
	SecretInput() interface{}
	SecureBootEnabled() interface{}
	SetSecureBootEnabled(val interface{})
	SecureBootEnabledInput() interface{}
	SinglePlacementGroup() interface{}
	SetSinglePlacementGroup(val interface{})
	SinglePlacementGroupInput() interface{}
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	SourceImageId() *string
	SetSourceImageId(val *string)
	SourceImageIdInput() *string
	SourceImageReference() WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference
	SourceImageReferenceInput() *WindowsVirtualMachineScaleSetSourceImageReference
	SpotRestore() WindowsVirtualMachineScaleSetSpotRestoreOutputReference
	SpotRestoreInput() *WindowsVirtualMachineScaleSetSpotRestore
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerminateNotification() WindowsVirtualMachineScaleSetTerminateNotificationOutputReference
	TerminateNotificationInput() *WindowsVirtualMachineScaleSetTerminateNotification
	TerminationNotification() WindowsVirtualMachineScaleSetTerminationNotificationOutputReference
	TerminationNotificationInput() *WindowsVirtualMachineScaleSetTerminationNotification
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() WindowsVirtualMachineScaleSetTimeoutsOutputReference
	TimeoutsInput() interface{}
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	UniqueId() *string
	UpgradeMode() *string
	SetUpgradeMode(val *string)
	UpgradeModeInput() *string
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	VtpmEnabled() interface{}
	SetVtpmEnabled(val interface{})
	VtpmEnabledInput() interface{}
	WinrmListener() WindowsVirtualMachineScaleSetWinrmListenerList
	WinrmListenerInput() interface{}
	ZoneBalance() interface{}
	SetZoneBalance(val interface{})
	ZoneBalanceInput() interface{}
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
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
	PutAdditionalCapabilities(value *WindowsVirtualMachineScaleSetAdditionalCapabilities)
	PutAdditionalUnattendContent(value interface{})
	PutAutomaticInstanceRepair(value *WindowsVirtualMachineScaleSetAutomaticInstanceRepair)
	PutAutomaticOsUpgradePolicy(value *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy)
	PutBootDiagnostics(value *WindowsVirtualMachineScaleSetBootDiagnostics)
	PutDataDisk(value interface{})
	PutExtension(value interface{})
	PutGalleryApplications(value interface{})
	PutIdentity(value *WindowsVirtualMachineScaleSetIdentity)
	PutNetworkInterface(value interface{})
	PutOsDisk(value *WindowsVirtualMachineScaleSetOsDisk)
	PutPlan(value *WindowsVirtualMachineScaleSetPlan)
	PutRollingUpgradePolicy(value *WindowsVirtualMachineScaleSetRollingUpgradePolicy)
	PutScaleIn(value *WindowsVirtualMachineScaleSetScaleIn)
	PutSecret(value interface{})
	PutSourceImageReference(value *WindowsVirtualMachineScaleSetSourceImageReference)
	PutSpotRestore(value *WindowsVirtualMachineScaleSetSpotRestore)
	PutTerminateNotification(value *WindowsVirtualMachineScaleSetTerminateNotification)
	PutTerminationNotification(value *WindowsVirtualMachineScaleSetTerminationNotification)
	PutTimeouts(value *WindowsVirtualMachineScaleSetTimeouts)
	PutWinrmListener(value interface{})
	ResetAdditionalCapabilities()
	ResetAdditionalUnattendContent()
	ResetAutomaticInstanceRepair()
	ResetAutomaticOsUpgradePolicy()
	ResetBootDiagnostics()
	ResetCapacityReservationGroupId()
	ResetComputerNamePrefix()
	ResetCustomData()
	ResetDataDisk()
	ResetDoNotRunExtensionsOnOverprovisionedMachines()
	ResetEdgeZone()
	ResetEnableAutomaticUpdates()
	ResetEncryptionAtHostEnabled()
	ResetEvictionPolicy()
	ResetExtension()
	ResetExtensionOperationsEnabled()
	ResetExtensionsTimeBudget()
	ResetGalleryApplications()
	ResetHealthProbeId()
	ResetHostGroupId()
	ResetId()
	ResetIdentity()
	ResetLicenseType()
	ResetMaxBidPrice()
	ResetOverprovision()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlan()
	ResetPlatformFaultDomainCount()
	ResetPriority()
	ResetProvisionVmAgent()
	ResetProximityPlacementGroupId()
	ResetRollingUpgradePolicy()
	ResetScaleIn()
	ResetScaleInPolicy()
	ResetSecret()
	ResetSecureBootEnabled()
	ResetSinglePlacementGroup()
	ResetSourceImageId()
	ResetSourceImageReference()
	ResetSpotRestore()
	ResetTags()
	ResetTerminateNotification()
	ResetTerminationNotification()
	ResetTimeouts()
	ResetTimezone()
	ResetUpgradeMode()
	ResetUserData()
	ResetVtpmEnabled()
	ResetWinrmListener()
	ResetZoneBalance()
	ResetZones()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for WindowsVirtualMachineScaleSet
type jsiiProxy_WindowsVirtualMachineScaleSet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdditionalCapabilities() WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference {
	var returns WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"additionalCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdditionalCapabilitiesInput() *WindowsVirtualMachineScaleSetAdditionalCapabilities {
	var returns *WindowsVirtualMachineScaleSetAdditionalCapabilities
	_jsii_.Get(
		j,
		"additionalCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdditionalUnattendContent() WindowsVirtualMachineScaleSetAdditionalUnattendContentList {
	var returns WindowsVirtualMachineScaleSetAdditionalUnattendContentList
	_jsii_.Get(
		j,
		"additionalUnattendContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdditionalUnattendContentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalUnattendContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdminUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AutomaticInstanceRepair() WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference {
	var returns WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference
	_jsii_.Get(
		j,
		"automaticInstanceRepair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AutomaticInstanceRepairInput() *WindowsVirtualMachineScaleSetAutomaticInstanceRepair {
	var returns *WindowsVirtualMachineScaleSetAutomaticInstanceRepair
	_jsii_.Get(
		j,
		"automaticInstanceRepairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AutomaticOsUpgradePolicy() WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference {
	var returns WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference
	_jsii_.Get(
		j,
		"automaticOsUpgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AutomaticOsUpgradePolicyInput() *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy {
	var returns *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy
	_jsii_.Get(
		j,
		"automaticOsUpgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) BootDiagnostics() WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference {
	var returns WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference
	_jsii_.Get(
		j,
		"bootDiagnostics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) BootDiagnosticsInput() *WindowsVirtualMachineScaleSetBootDiagnostics {
	var returns *WindowsVirtualMachineScaleSetBootDiagnostics
	_jsii_.Get(
		j,
		"bootDiagnosticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) CapacityReservationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) CapacityReservationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ComputerNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ComputerNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) CustomData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) CustomDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) DataDisk() WindowsVirtualMachineScaleSetDataDiskList {
	var returns WindowsVirtualMachineScaleSetDataDiskList
	_jsii_.Get(
		j,
		"dataDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) DataDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) DoNotRunExtensionsOnOverprovisionedMachines() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"doNotRunExtensionsOnOverprovisionedMachines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) DoNotRunExtensionsOnOverprovisionedMachinesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"doNotRunExtensionsOnOverprovisionedMachinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EdgeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EdgeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EnableAutomaticUpdates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EnableAutomaticUpdatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EncryptionAtHostEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtHostEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EncryptionAtHostEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtHostEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EvictionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EvictionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Extension() WindowsVirtualMachineScaleSetExtensionList {
	var returns WindowsVirtualMachineScaleSetExtensionList
	_jsii_.Get(
		j,
		"extension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ExtensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ExtensionOperationsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionOperationsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ExtensionOperationsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionOperationsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ExtensionsTimeBudget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionsTimeBudget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ExtensionsTimeBudgetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionsTimeBudgetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) GalleryApplications() WindowsVirtualMachineScaleSetGalleryApplicationsList {
	var returns WindowsVirtualMachineScaleSetGalleryApplicationsList
	_jsii_.Get(
		j,
		"galleryApplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) GalleryApplicationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"galleryApplicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) HealthProbeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthProbeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) HealthProbeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthProbeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) HostGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) HostGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Identity() WindowsVirtualMachineScaleSetIdentityOutputReference {
	var returns WindowsVirtualMachineScaleSetIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) IdentityInput() *WindowsVirtualMachineScaleSetIdentity {
	var returns *WindowsVirtualMachineScaleSetIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Instances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) InstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) MaxBidPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBidPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) MaxBidPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBidPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) NetworkInterface() WindowsVirtualMachineScaleSetNetworkInterfaceList {
	var returns WindowsVirtualMachineScaleSetNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) OsDisk() WindowsVirtualMachineScaleSetOsDiskOutputReference {
	var returns WindowsVirtualMachineScaleSetOsDiskOutputReference
	_jsii_.Get(
		j,
		"osDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) OsDiskInput() *WindowsVirtualMachineScaleSetOsDisk {
	var returns *WindowsVirtualMachineScaleSetOsDisk
	_jsii_.Get(
		j,
		"osDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Overprovision() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overprovision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) OverprovisionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overprovisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Plan() WindowsVirtualMachineScaleSetPlanOutputReference {
	var returns WindowsVirtualMachineScaleSetPlanOutputReference
	_jsii_.Get(
		j,
		"plan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) PlanInput() *WindowsVirtualMachineScaleSetPlan {
	var returns *WindowsVirtualMachineScaleSetPlan
	_jsii_.Get(
		j,
		"planInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) PlatformFaultDomainCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"platformFaultDomainCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) PlatformFaultDomainCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"platformFaultDomainCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ProvisionVmAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ProvisionVmAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ProximityPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ProximityPlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) RollingUpgradePolicy() WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference {
	var returns WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference
	_jsii_.Get(
		j,
		"rollingUpgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) RollingUpgradePolicyInput() *WindowsVirtualMachineScaleSetRollingUpgradePolicy {
	var returns *WindowsVirtualMachineScaleSetRollingUpgradePolicy
	_jsii_.Get(
		j,
		"rollingUpgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ScaleIn() WindowsVirtualMachineScaleSetScaleInOutputReference {
	var returns WindowsVirtualMachineScaleSetScaleInOutputReference
	_jsii_.Get(
		j,
		"scaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ScaleInInput() *WindowsVirtualMachineScaleSetScaleIn {
	var returns *WindowsVirtualMachineScaleSetScaleIn
	_jsii_.Get(
		j,
		"scaleInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ScaleInPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleInPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ScaleInPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleInPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Secret() WindowsVirtualMachineScaleSetSecretList {
	var returns WindowsVirtualMachineScaleSetSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SecureBootEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureBootEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SecureBootEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureBootEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SinglePlacementGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singlePlacementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SinglePlacementGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singlePlacementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SourceImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SourceImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SourceImageReference() WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference {
	var returns WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference
	_jsii_.Get(
		j,
		"sourceImageReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SourceImageReferenceInput() *WindowsVirtualMachineScaleSetSourceImageReference {
	var returns *WindowsVirtualMachineScaleSetSourceImageReference
	_jsii_.Get(
		j,
		"sourceImageReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SpotRestore() WindowsVirtualMachineScaleSetSpotRestoreOutputReference {
	var returns WindowsVirtualMachineScaleSetSpotRestoreOutputReference
	_jsii_.Get(
		j,
		"spotRestore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SpotRestoreInput() *WindowsVirtualMachineScaleSetSpotRestore {
	var returns *WindowsVirtualMachineScaleSetSpotRestore
	_jsii_.Get(
		j,
		"spotRestoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TerminateNotification() WindowsVirtualMachineScaleSetTerminateNotificationOutputReference {
	var returns WindowsVirtualMachineScaleSetTerminateNotificationOutputReference
	_jsii_.Get(
		j,
		"terminateNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TerminateNotificationInput() *WindowsVirtualMachineScaleSetTerminateNotification {
	var returns *WindowsVirtualMachineScaleSetTerminateNotification
	_jsii_.Get(
		j,
		"terminateNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TerminationNotification() WindowsVirtualMachineScaleSetTerminationNotificationOutputReference {
	var returns WindowsVirtualMachineScaleSetTerminationNotificationOutputReference
	_jsii_.Get(
		j,
		"terminationNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TerminationNotificationInput() *WindowsVirtualMachineScaleSetTerminationNotification {
	var returns *WindowsVirtualMachineScaleSetTerminationNotification
	_jsii_.Get(
		j,
		"terminationNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Timeouts() WindowsVirtualMachineScaleSetTimeoutsOutputReference {
	var returns WindowsVirtualMachineScaleSetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) UniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) UpgradeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) UpgradeModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) VtpmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vtpmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) VtpmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vtpmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) WinrmListener() WindowsVirtualMachineScaleSetWinrmListenerList {
	var returns WindowsVirtualMachineScaleSetWinrmListenerList
	_jsii_.Get(
		j,
		"winrmListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) WinrmListenerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"winrmListenerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ZoneBalance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneBalance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ZoneBalanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneBalanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set azurerm_windows_virtual_machine_scale_set} Resource.
func NewWindowsVirtualMachineScaleSet(scope constructs.Construct, id *string, config *WindowsVirtualMachineScaleSetConfig) WindowsVirtualMachineScaleSet {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSet{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set azurerm_windows_virtual_machine_scale_set} Resource.
func NewWindowsVirtualMachineScaleSet_Override(w WindowsVirtualMachineScaleSet, scope constructs.Construct, id *string, config *WindowsVirtualMachineScaleSetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSet",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetAdminPassword(val *string) {
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetAdminUsername(val *string) {
	_jsii_.Set(
		j,
		"adminUsername",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetCapacityReservationGroupId(val *string) {
	_jsii_.Set(
		j,
		"capacityReservationGroupId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetComputerNamePrefix(val *string) {
	_jsii_.Set(
		j,
		"computerNamePrefix",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetCustomData(val *string) {
	_jsii_.Set(
		j,
		"customData",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetDoNotRunExtensionsOnOverprovisionedMachines(val interface{}) {
	_jsii_.Set(
		j,
		"doNotRunExtensionsOnOverprovisionedMachines",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetEdgeZone(val *string) {
	_jsii_.Set(
		j,
		"edgeZone",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetEnableAutomaticUpdates(val interface{}) {
	_jsii_.Set(
		j,
		"enableAutomaticUpdates",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetEncryptionAtHostEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionAtHostEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetEvictionPolicy(val *string) {
	_jsii_.Set(
		j,
		"evictionPolicy",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetExtensionOperationsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"extensionOperationsEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetExtensionsTimeBudget(val *string) {
	_jsii_.Set(
		j,
		"extensionsTimeBudget",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetHealthProbeId(val *string) {
	_jsii_.Set(
		j,
		"healthProbeId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetHostGroupId(val *string) {
	_jsii_.Set(
		j,
		"hostGroupId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetInstances(val *float64) {
	_jsii_.Set(
		j,
		"instances",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetLicenseType(val *string) {
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetMaxBidPrice(val *float64) {
	_jsii_.Set(
		j,
		"maxBidPrice",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetOverprovision(val interface{}) {
	_jsii_.Set(
		j,
		"overprovision",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetPlatformFaultDomainCount(val *float64) {
	_jsii_.Set(
		j,
		"platformFaultDomainCount",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetPriority(val *string) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetProvisionVmAgent(val interface{}) {
	_jsii_.Set(
		j,
		"provisionVmAgent",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetProximityPlacementGroupId(val *string) {
	_jsii_.Set(
		j,
		"proximityPlacementGroupId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetScaleInPolicy(val *string) {
	_jsii_.Set(
		j,
		"scaleInPolicy",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetSecureBootEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"secureBootEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetSinglePlacementGroup(val interface{}) {
	_jsii_.Set(
		j,
		"singlePlacementGroup",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetSku(val *string) {
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetSourceImageId(val *string) {
	_jsii_.Set(
		j,
		"sourceImageId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetTimezone(val *string) {
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetUpgradeMode(val *string) {
	_jsii_.Set(
		j,
		"upgradeMode",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetUserData(val *string) {
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetVtpmEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"vtpmEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetZoneBalance(val interface{}) {
	_jsii_.Set(
		j,
		"zoneBalance",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SetZones(val *[]*string) {
	_jsii_.Set(
		j,
		"zones",
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
func WindowsVirtualMachineScaleSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WindowsVirtualMachineScaleSet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutAdditionalCapabilities(value *WindowsVirtualMachineScaleSetAdditionalCapabilities) {
	_jsii_.InvokeVoid(
		w,
		"putAdditionalCapabilities",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutAdditionalUnattendContent(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putAdditionalUnattendContent",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutAutomaticInstanceRepair(value *WindowsVirtualMachineScaleSetAutomaticInstanceRepair) {
	_jsii_.InvokeVoid(
		w,
		"putAutomaticInstanceRepair",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutAutomaticOsUpgradePolicy(value *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy) {
	_jsii_.InvokeVoid(
		w,
		"putAutomaticOsUpgradePolicy",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutBootDiagnostics(value *WindowsVirtualMachineScaleSetBootDiagnostics) {
	_jsii_.InvokeVoid(
		w,
		"putBootDiagnostics",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutDataDisk(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putDataDisk",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutExtension(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putExtension",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutGalleryApplications(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putGalleryApplications",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutIdentity(value *WindowsVirtualMachineScaleSetIdentity) {
	_jsii_.InvokeVoid(
		w,
		"putIdentity",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutNetworkInterface(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutOsDisk(value *WindowsVirtualMachineScaleSetOsDisk) {
	_jsii_.InvokeVoid(
		w,
		"putOsDisk",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutPlan(value *WindowsVirtualMachineScaleSetPlan) {
	_jsii_.InvokeVoid(
		w,
		"putPlan",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutRollingUpgradePolicy(value *WindowsVirtualMachineScaleSetRollingUpgradePolicy) {
	_jsii_.InvokeVoid(
		w,
		"putRollingUpgradePolicy",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutScaleIn(value *WindowsVirtualMachineScaleSetScaleIn) {
	_jsii_.InvokeVoid(
		w,
		"putScaleIn",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutSecret(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putSecret",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutSourceImageReference(value *WindowsVirtualMachineScaleSetSourceImageReference) {
	_jsii_.InvokeVoid(
		w,
		"putSourceImageReference",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutSpotRestore(value *WindowsVirtualMachineScaleSetSpotRestore) {
	_jsii_.InvokeVoid(
		w,
		"putSpotRestore",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutTerminateNotification(value *WindowsVirtualMachineScaleSetTerminateNotification) {
	_jsii_.InvokeVoid(
		w,
		"putTerminateNotification",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutTerminationNotification(value *WindowsVirtualMachineScaleSetTerminationNotification) {
	_jsii_.InvokeVoid(
		w,
		"putTerminationNotification",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutTimeouts(value *WindowsVirtualMachineScaleSetTimeouts) {
	_jsii_.InvokeVoid(
		w,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutWinrmListener(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putWinrmListener",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetAdditionalCapabilities() {
	_jsii_.InvokeVoid(
		w,
		"resetAdditionalCapabilities",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetAdditionalUnattendContent() {
	_jsii_.InvokeVoid(
		w,
		"resetAdditionalUnattendContent",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetAutomaticInstanceRepair() {
	_jsii_.InvokeVoid(
		w,
		"resetAutomaticInstanceRepair",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetAutomaticOsUpgradePolicy() {
	_jsii_.InvokeVoid(
		w,
		"resetAutomaticOsUpgradePolicy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetBootDiagnostics() {
	_jsii_.InvokeVoid(
		w,
		"resetBootDiagnostics",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetCapacityReservationGroupId() {
	_jsii_.InvokeVoid(
		w,
		"resetCapacityReservationGroupId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetComputerNamePrefix() {
	_jsii_.InvokeVoid(
		w,
		"resetComputerNamePrefix",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetCustomData() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomData",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetDataDisk() {
	_jsii_.InvokeVoid(
		w,
		"resetDataDisk",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetDoNotRunExtensionsOnOverprovisionedMachines() {
	_jsii_.InvokeVoid(
		w,
		"resetDoNotRunExtensionsOnOverprovisionedMachines",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetEdgeZone() {
	_jsii_.InvokeVoid(
		w,
		"resetEdgeZone",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetEnableAutomaticUpdates() {
	_jsii_.InvokeVoid(
		w,
		"resetEnableAutomaticUpdates",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetEncryptionAtHostEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetEncryptionAtHostEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetEvictionPolicy() {
	_jsii_.InvokeVoid(
		w,
		"resetEvictionPolicy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetExtension() {
	_jsii_.InvokeVoid(
		w,
		"resetExtension",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetExtensionOperationsEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetExtensionOperationsEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetExtensionsTimeBudget() {
	_jsii_.InvokeVoid(
		w,
		"resetExtensionsTimeBudget",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetGalleryApplications() {
	_jsii_.InvokeVoid(
		w,
		"resetGalleryApplications",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetHealthProbeId() {
	_jsii_.InvokeVoid(
		w,
		"resetHealthProbeId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetHostGroupId() {
	_jsii_.InvokeVoid(
		w,
		"resetHostGroupId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetIdentity() {
	_jsii_.InvokeVoid(
		w,
		"resetIdentity",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetLicenseType() {
	_jsii_.InvokeVoid(
		w,
		"resetLicenseType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetMaxBidPrice() {
	_jsii_.InvokeVoid(
		w,
		"resetMaxBidPrice",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetOverprovision() {
	_jsii_.InvokeVoid(
		w,
		"resetOverprovision",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetPlan() {
	_jsii_.InvokeVoid(
		w,
		"resetPlan",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetPlatformFaultDomainCount() {
	_jsii_.InvokeVoid(
		w,
		"resetPlatformFaultDomainCount",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetPriority() {
	_jsii_.InvokeVoid(
		w,
		"resetPriority",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetProvisionVmAgent() {
	_jsii_.InvokeVoid(
		w,
		"resetProvisionVmAgent",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetProximityPlacementGroupId() {
	_jsii_.InvokeVoid(
		w,
		"resetProximityPlacementGroupId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetRollingUpgradePolicy() {
	_jsii_.InvokeVoid(
		w,
		"resetRollingUpgradePolicy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetScaleIn() {
	_jsii_.InvokeVoid(
		w,
		"resetScaleIn",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetScaleInPolicy() {
	_jsii_.InvokeVoid(
		w,
		"resetScaleInPolicy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetSecret() {
	_jsii_.InvokeVoid(
		w,
		"resetSecret",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetSecureBootEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetSecureBootEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetSinglePlacementGroup() {
	_jsii_.InvokeVoid(
		w,
		"resetSinglePlacementGroup",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetSourceImageId() {
	_jsii_.InvokeVoid(
		w,
		"resetSourceImageId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetSourceImageReference() {
	_jsii_.InvokeVoid(
		w,
		"resetSourceImageReference",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetSpotRestore() {
	_jsii_.InvokeVoid(
		w,
		"resetSpotRestore",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetTerminateNotification() {
	_jsii_.InvokeVoid(
		w,
		"resetTerminateNotification",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetTerminationNotification() {
	_jsii_.InvokeVoid(
		w,
		"resetTerminationNotification",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetTimeouts() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetTimezone() {
	_jsii_.InvokeVoid(
		w,
		"resetTimezone",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetUpgradeMode() {
	_jsii_.InvokeVoid(
		w,
		"resetUpgradeMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetUserData() {
	_jsii_.InvokeVoid(
		w,
		"resetUserData",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetVtpmEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetVtpmEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetWinrmListener() {
	_jsii_.InvokeVoid(
		w,
		"resetWinrmListener",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetZoneBalance() {
	_jsii_.InvokeVoid(
		w,
		"resetZoneBalance",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetZones() {
	_jsii_.InvokeVoid(
		w,
		"resetZones",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetAdditionalCapabilities struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#ultra_ssd_enabled WindowsVirtualMachineScaleSet#ultra_ssd_enabled}.
	UltraSsdEnabled interface{} `field:"optional" json:"ultraSsdEnabled" yaml:"ultraSsdEnabled"`
}

type WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference interface {
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
	InternalValue() *WindowsVirtualMachineScaleSetAdditionalCapabilities
	SetInternalValue(val *WindowsVirtualMachineScaleSetAdditionalCapabilities)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UltraSsdEnabled() interface{}
	SetUltraSsdEnabled(val interface{})
	UltraSsdEnabledInput() interface{}
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
	ResetUltraSsdEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) InternalValue() *WindowsVirtualMachineScaleSetAdditionalCapabilities {
	var returns *WindowsVirtualMachineScaleSetAdditionalCapabilities
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) UltraSsdEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ultraSsdEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) UltraSsdEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ultraSsdEnabledInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference_Override(w WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) SetInternalValue(val *WindowsVirtualMachineScaleSetAdditionalCapabilities) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) SetUltraSsdEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"ultraSsdEnabled",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) ResetUltraSsdEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetUltraSsdEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetAdditionalUnattendContent struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#content WindowsVirtualMachineScaleSet#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#setting WindowsVirtualMachineScaleSet#setting}.
	Setting *string `field:"required" json:"setting" yaml:"setting"`
}

type WindowsVirtualMachineScaleSetAdditionalUnattendContentList interface {
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
	Get(index *float64) WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetAdditionalUnattendContentList
type jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetAdditionalUnattendContentList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsVirtualMachineScaleSetAdditionalUnattendContentList {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetAdditionalUnattendContentList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetAdditionalUnattendContentList_Override(w WindowsVirtualMachineScaleSetAdditionalUnattendContentList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetAdditionalUnattendContentList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList) Get(index *float64) WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference {
	var returns WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference interface {
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
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Setting() *string
	SetSetting(val *string)
	SettingInput() *string
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

// The jsii proxy struct for WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) Setting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) SettingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference_Override(w WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) SetContent(val *string) {
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) SetSetting(val *string) {
	_jsii_.Set(
		j,
		"setting",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAdditionalUnattendContentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetAutomaticInstanceRepair struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#enabled WindowsVirtualMachineScaleSet#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#grace_period WindowsVirtualMachineScaleSet#grace_period}.
	GracePeriod *string `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
}

type WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	GracePeriod() *string
	SetGracePeriod(val *string)
	GracePeriodInput() *string
	InternalValue() *WindowsVirtualMachineScaleSetAutomaticInstanceRepair
	SetInternalValue(val *WindowsVirtualMachineScaleSetAutomaticInstanceRepair)
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
	ResetGracePeriod()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) GracePeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) GracePeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) InternalValue() *WindowsVirtualMachineScaleSetAutomaticInstanceRepair {
	var returns *WindowsVirtualMachineScaleSetAutomaticInstanceRepair
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference_Override(w WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) SetGracePeriod(val *string) {
	_jsii_.Set(
		j,
		"gracePeriod",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) SetInternalValue(val *WindowsVirtualMachineScaleSetAutomaticInstanceRepair) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) ResetGracePeriod() {
	_jsii_.InvokeVoid(
		w,
		"resetGracePeriod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#disable_automatic_rollback WindowsVirtualMachineScaleSet#disable_automatic_rollback}.
	DisableAutomaticRollback interface{} `field:"required" json:"disableAutomaticRollback" yaml:"disableAutomaticRollback"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#enable_automatic_os_upgrade WindowsVirtualMachineScaleSet#enable_automatic_os_upgrade}.
	EnableAutomaticOsUpgrade interface{} `field:"required" json:"enableAutomaticOsUpgrade" yaml:"enableAutomaticOsUpgrade"`
}

type WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference interface {
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
	DisableAutomaticRollback() interface{}
	SetDisableAutomaticRollback(val interface{})
	DisableAutomaticRollbackInput() interface{}
	EnableAutomaticOsUpgrade() interface{}
	SetEnableAutomaticOsUpgrade(val interface{})
	EnableAutomaticOsUpgradeInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy
	SetInternalValue(val *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy)
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

// The jsii proxy struct for WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) DisableAutomaticRollback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomaticRollback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) DisableAutomaticRollbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomaticRollbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) EnableAutomaticOsUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticOsUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) EnableAutomaticOsUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticOsUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) InternalValue() *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy {
	var returns *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference_Override(w WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) SetDisableAutomaticRollback(val interface{}) {
	_jsii_.Set(
		j,
		"disableAutomaticRollback",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) SetEnableAutomaticOsUpgrade(val interface{}) {
	_jsii_.Set(
		j,
		"enableAutomaticOsUpgrade",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) SetInternalValue(val *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetBootDiagnostics struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#storage_account_uri WindowsVirtualMachineScaleSet#storage_account_uri}.
	StorageAccountUri *string `field:"optional" json:"storageAccountUri" yaml:"storageAccountUri"`
}

type WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference interface {
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
	InternalValue() *WindowsVirtualMachineScaleSetBootDiagnostics
	SetInternalValue(val *WindowsVirtualMachineScaleSetBootDiagnostics)
	StorageAccountUri() *string
	SetStorageAccountUri(val *string)
	StorageAccountUriInput() *string
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
	ResetStorageAccountUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) InternalValue() *WindowsVirtualMachineScaleSetBootDiagnostics {
	var returns *WindowsVirtualMachineScaleSetBootDiagnostics
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) StorageAccountUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) StorageAccountUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetBootDiagnosticsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetBootDiagnosticsOutputReference_Override(w WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) SetInternalValue(val *WindowsVirtualMachineScaleSetBootDiagnostics) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) SetStorageAccountUri(val *string) {
	_jsii_.Set(
		j,
		"storageAccountUri",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) ResetStorageAccountUri() {
	_jsii_.InvokeVoid(
		w,
		"resetStorageAccountUri",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#admin_password WindowsVirtualMachineScaleSet#admin_password}.
	AdminPassword *string `field:"required" json:"adminPassword" yaml:"adminPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#admin_username WindowsVirtualMachineScaleSet#admin_username}.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#instances WindowsVirtualMachineScaleSet#instances}.
	Instances *float64 `field:"required" json:"instances" yaml:"instances"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#location WindowsVirtualMachineScaleSet#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#name WindowsVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#network_interface WindowsVirtualMachineScaleSet#network_interface}
	NetworkInterface interface{} `field:"required" json:"networkInterface" yaml:"networkInterface"`
	// os_disk block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#os_disk WindowsVirtualMachineScaleSet#os_disk}
	OsDisk *WindowsVirtualMachineScaleSetOsDisk `field:"required" json:"osDisk" yaml:"osDisk"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#resource_group_name WindowsVirtualMachineScaleSet#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#sku WindowsVirtualMachineScaleSet#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// additional_capabilities block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#additional_capabilities WindowsVirtualMachineScaleSet#additional_capabilities}
	AdditionalCapabilities *WindowsVirtualMachineScaleSetAdditionalCapabilities `field:"optional" json:"additionalCapabilities" yaml:"additionalCapabilities"`
	// additional_unattend_content block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#additional_unattend_content WindowsVirtualMachineScaleSet#additional_unattend_content}
	AdditionalUnattendContent interface{} `field:"optional" json:"additionalUnattendContent" yaml:"additionalUnattendContent"`
	// automatic_instance_repair block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#automatic_instance_repair WindowsVirtualMachineScaleSet#automatic_instance_repair}
	AutomaticInstanceRepair *WindowsVirtualMachineScaleSetAutomaticInstanceRepair `field:"optional" json:"automaticInstanceRepair" yaml:"automaticInstanceRepair"`
	// automatic_os_upgrade_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#automatic_os_upgrade_policy WindowsVirtualMachineScaleSet#automatic_os_upgrade_policy}
	AutomaticOsUpgradePolicy *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy `field:"optional" json:"automaticOsUpgradePolicy" yaml:"automaticOsUpgradePolicy"`
	// boot_diagnostics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#boot_diagnostics WindowsVirtualMachineScaleSet#boot_diagnostics}
	BootDiagnostics *WindowsVirtualMachineScaleSetBootDiagnostics `field:"optional" json:"bootDiagnostics" yaml:"bootDiagnostics"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#capacity_reservation_group_id WindowsVirtualMachineScaleSet#capacity_reservation_group_id}.
	CapacityReservationGroupId *string `field:"optional" json:"capacityReservationGroupId" yaml:"capacityReservationGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#computer_name_prefix WindowsVirtualMachineScaleSet#computer_name_prefix}.
	ComputerNamePrefix *string `field:"optional" json:"computerNamePrefix" yaml:"computerNamePrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#custom_data WindowsVirtualMachineScaleSet#custom_data}.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
	// data_disk block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#data_disk WindowsVirtualMachineScaleSet#data_disk}
	DataDisk interface{} `field:"optional" json:"dataDisk" yaml:"dataDisk"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#do_not_run_extensions_on_overprovisioned_machines WindowsVirtualMachineScaleSet#do_not_run_extensions_on_overprovisioned_machines}.
	DoNotRunExtensionsOnOverprovisionedMachines interface{} `field:"optional" json:"doNotRunExtensionsOnOverprovisionedMachines" yaml:"doNotRunExtensionsOnOverprovisionedMachines"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#edge_zone WindowsVirtualMachineScaleSet#edge_zone}.
	EdgeZone *string `field:"optional" json:"edgeZone" yaml:"edgeZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#enable_automatic_updates WindowsVirtualMachineScaleSet#enable_automatic_updates}.
	EnableAutomaticUpdates interface{} `field:"optional" json:"enableAutomaticUpdates" yaml:"enableAutomaticUpdates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#encryption_at_host_enabled WindowsVirtualMachineScaleSet#encryption_at_host_enabled}.
	EncryptionAtHostEnabled interface{} `field:"optional" json:"encryptionAtHostEnabled" yaml:"encryptionAtHostEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#eviction_policy WindowsVirtualMachineScaleSet#eviction_policy}.
	EvictionPolicy *string `field:"optional" json:"evictionPolicy" yaml:"evictionPolicy"`
	// extension block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#extension WindowsVirtualMachineScaleSet#extension}
	Extension interface{} `field:"optional" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#extension_operations_enabled WindowsVirtualMachineScaleSet#extension_operations_enabled}.
	ExtensionOperationsEnabled interface{} `field:"optional" json:"extensionOperationsEnabled" yaml:"extensionOperationsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#extensions_time_budget WindowsVirtualMachineScaleSet#extensions_time_budget}.
	ExtensionsTimeBudget *string `field:"optional" json:"extensionsTimeBudget" yaml:"extensionsTimeBudget"`
	// gallery_applications block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#gallery_applications WindowsVirtualMachineScaleSet#gallery_applications}
	GalleryApplications interface{} `field:"optional" json:"galleryApplications" yaml:"galleryApplications"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#health_probe_id WindowsVirtualMachineScaleSet#health_probe_id}.
	HealthProbeId *string `field:"optional" json:"healthProbeId" yaml:"healthProbeId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#host_group_id WindowsVirtualMachineScaleSet#host_group_id}.
	HostGroupId *string `field:"optional" json:"hostGroupId" yaml:"hostGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#id WindowsVirtualMachineScaleSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#identity WindowsVirtualMachineScaleSet#identity}
	Identity *WindowsVirtualMachineScaleSetIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#license_type WindowsVirtualMachineScaleSet#license_type}.
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#max_bid_price WindowsVirtualMachineScaleSet#max_bid_price}.
	MaxBidPrice *float64 `field:"optional" json:"maxBidPrice" yaml:"maxBidPrice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#overprovision WindowsVirtualMachineScaleSet#overprovision}.
	Overprovision interface{} `field:"optional" json:"overprovision" yaml:"overprovision"`
	// plan block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#plan WindowsVirtualMachineScaleSet#plan}
	Plan *WindowsVirtualMachineScaleSetPlan `field:"optional" json:"plan" yaml:"plan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#platform_fault_domain_count WindowsVirtualMachineScaleSet#platform_fault_domain_count}.
	PlatformFaultDomainCount *float64 `field:"optional" json:"platformFaultDomainCount" yaml:"platformFaultDomainCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#priority WindowsVirtualMachineScaleSet#priority}.
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#provision_vm_agent WindowsVirtualMachineScaleSet#provision_vm_agent}.
	ProvisionVmAgent interface{} `field:"optional" json:"provisionVmAgent" yaml:"provisionVmAgent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#proximity_placement_group_id WindowsVirtualMachineScaleSet#proximity_placement_group_id}.
	ProximityPlacementGroupId *string `field:"optional" json:"proximityPlacementGroupId" yaml:"proximityPlacementGroupId"`
	// rolling_upgrade_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#rolling_upgrade_policy WindowsVirtualMachineScaleSet#rolling_upgrade_policy}
	RollingUpgradePolicy *WindowsVirtualMachineScaleSetRollingUpgradePolicy `field:"optional" json:"rollingUpgradePolicy" yaml:"rollingUpgradePolicy"`
	// scale_in block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#scale_in WindowsVirtualMachineScaleSet#scale_in}
	ScaleIn *WindowsVirtualMachineScaleSetScaleIn `field:"optional" json:"scaleIn" yaml:"scaleIn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#scale_in_policy WindowsVirtualMachineScaleSet#scale_in_policy}.
	ScaleInPolicy *string `field:"optional" json:"scaleInPolicy" yaml:"scaleInPolicy"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#secret WindowsVirtualMachineScaleSet#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#secure_boot_enabled WindowsVirtualMachineScaleSet#secure_boot_enabled}.
	SecureBootEnabled interface{} `field:"optional" json:"secureBootEnabled" yaml:"secureBootEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#single_placement_group WindowsVirtualMachineScaleSet#single_placement_group}.
	SinglePlacementGroup interface{} `field:"optional" json:"singlePlacementGroup" yaml:"singlePlacementGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#source_image_id WindowsVirtualMachineScaleSet#source_image_id}.
	SourceImageId *string `field:"optional" json:"sourceImageId" yaml:"sourceImageId"`
	// source_image_reference block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#source_image_reference WindowsVirtualMachineScaleSet#source_image_reference}
	SourceImageReference *WindowsVirtualMachineScaleSetSourceImageReference `field:"optional" json:"sourceImageReference" yaml:"sourceImageReference"`
	// spot_restore block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#spot_restore WindowsVirtualMachineScaleSet#spot_restore}
	SpotRestore *WindowsVirtualMachineScaleSetSpotRestore `field:"optional" json:"spotRestore" yaml:"spotRestore"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#tags WindowsVirtualMachineScaleSet#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// terminate_notification block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#terminate_notification WindowsVirtualMachineScaleSet#terminate_notification}
	TerminateNotification *WindowsVirtualMachineScaleSetTerminateNotification `field:"optional" json:"terminateNotification" yaml:"terminateNotification"`
	// termination_notification block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#termination_notification WindowsVirtualMachineScaleSet#termination_notification}
	TerminationNotification *WindowsVirtualMachineScaleSetTerminationNotification `field:"optional" json:"terminationNotification" yaml:"terminationNotification"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#timeouts WindowsVirtualMachineScaleSet#timeouts}
	Timeouts *WindowsVirtualMachineScaleSetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#timezone WindowsVirtualMachineScaleSet#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#upgrade_mode WindowsVirtualMachineScaleSet#upgrade_mode}.
	UpgradeMode *string `field:"optional" json:"upgradeMode" yaml:"upgradeMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#user_data WindowsVirtualMachineScaleSet#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#vtpm_enabled WindowsVirtualMachineScaleSet#vtpm_enabled}.
	VtpmEnabled interface{} `field:"optional" json:"vtpmEnabled" yaml:"vtpmEnabled"`
	// winrm_listener block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#winrm_listener WindowsVirtualMachineScaleSet#winrm_listener}
	WinrmListener interface{} `field:"optional" json:"winrmListener" yaml:"winrmListener"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#zone_balance WindowsVirtualMachineScaleSet#zone_balance}.
	ZoneBalance interface{} `field:"optional" json:"zoneBalance" yaml:"zoneBalance"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#zones WindowsVirtualMachineScaleSet#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

type WindowsVirtualMachineScaleSetDataDisk struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#caching WindowsVirtualMachineScaleSet#caching}.
	Caching *string `field:"required" json:"caching" yaml:"caching"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#disk_size_gb WindowsVirtualMachineScaleSet#disk_size_gb}.
	DiskSizeGb *float64 `field:"required" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#lun WindowsVirtualMachineScaleSet#lun}.
	Lun *float64 `field:"required" json:"lun" yaml:"lun"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#storage_account_type WindowsVirtualMachineScaleSet#storage_account_type}.
	StorageAccountType *string `field:"required" json:"storageAccountType" yaml:"storageAccountType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#create_option WindowsVirtualMachineScaleSet#create_option}.
	CreateOption *string `field:"optional" json:"createOption" yaml:"createOption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#disk_encryption_set_id WindowsVirtualMachineScaleSet#disk_encryption_set_id}.
	DiskEncryptionSetId *string `field:"optional" json:"diskEncryptionSetId" yaml:"diskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#name WindowsVirtualMachineScaleSet#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#ultra_ssd_disk_iops_read_write WindowsVirtualMachineScaleSet#ultra_ssd_disk_iops_read_write}.
	UltraSsdDiskIopsReadWrite *float64 `field:"optional" json:"ultraSsdDiskIopsReadWrite" yaml:"ultraSsdDiskIopsReadWrite"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#ultra_ssd_disk_mbps_read_write WindowsVirtualMachineScaleSet#ultra_ssd_disk_mbps_read_write}.
	UltraSsdDiskMbpsReadWrite *float64 `field:"optional" json:"ultraSsdDiskMbpsReadWrite" yaml:"ultraSsdDiskMbpsReadWrite"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#write_accelerator_enabled WindowsVirtualMachineScaleSet#write_accelerator_enabled}.
	WriteAcceleratorEnabled interface{} `field:"optional" json:"writeAcceleratorEnabled" yaml:"writeAcceleratorEnabled"`
}

type WindowsVirtualMachineScaleSetDataDiskList interface {
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
	Get(index *float64) WindowsVirtualMachineScaleSetDataDiskOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetDataDiskList
type jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetDataDiskList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsVirtualMachineScaleSetDataDiskList {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetDataDiskList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetDataDiskList_Override(w WindowsVirtualMachineScaleSetDataDiskList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetDataDiskList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList) Get(index *float64) WindowsVirtualMachineScaleSetDataDiskOutputReference {
	var returns WindowsVirtualMachineScaleSetDataDiskOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetDataDiskOutputReference interface {
	cdktf.ComplexObject
	Caching() *string
	SetCaching(val *string)
	CachingInput() *string
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
	CreateOption() *string
	SetCreateOption(val *string)
	CreateOptionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DiskEncryptionSetId() *string
	SetDiskEncryptionSetId(val *string)
	DiskEncryptionSetIdInput() *string
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Lun() *float64
	SetLun(val *float64)
	LunInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	StorageAccountType() *string
	SetStorageAccountType(val *string)
	StorageAccountTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UltraSsdDiskIopsReadWrite() *float64
	SetUltraSsdDiskIopsReadWrite(val *float64)
	UltraSsdDiskIopsReadWriteInput() *float64
	UltraSsdDiskMbpsReadWrite() *float64
	SetUltraSsdDiskMbpsReadWrite(val *float64)
	UltraSsdDiskMbpsReadWriteInput() *float64
	WriteAcceleratorEnabled() interface{}
	SetWriteAcceleratorEnabled(val interface{})
	WriteAcceleratorEnabledInput() interface{}
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
	ResetCreateOption()
	ResetDiskEncryptionSetId()
	ResetName()
	ResetUltraSsdDiskIopsReadWrite()
	ResetUltraSsdDiskMbpsReadWrite()
	ResetWriteAcceleratorEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetDataDiskOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) Caching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) CachingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) CreateOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) CreateOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) DiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) DiskEncryptionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) Lun() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) LunInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) StorageAccountType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) StorageAccountTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) UltraSsdDiskIopsReadWrite() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ultraSsdDiskIopsReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) UltraSsdDiskIopsReadWriteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ultraSsdDiskIopsReadWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) UltraSsdDiskMbpsReadWrite() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ultraSsdDiskMbpsReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) UltraSsdDiskMbpsReadWriteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ultraSsdDiskMbpsReadWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) WriteAcceleratorEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeAcceleratorEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) WriteAcceleratorEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeAcceleratorEnabledInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetDataDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsVirtualMachineScaleSetDataDiskOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetDataDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetDataDiskOutputReference_Override(w WindowsVirtualMachineScaleSetDataDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetDataDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) SetCaching(val *string) {
	_jsii_.Set(
		j,
		"caching",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) SetCreateOption(val *string) {
	_jsii_.Set(
		j,
		"createOption",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) SetDiskEncryptionSetId(val *string) {
	_jsii_.Set(
		j,
		"diskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) SetDiskSizeGb(val *float64) {
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) SetLun(val *float64) {
	_jsii_.Set(
		j,
		"lun",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) SetStorageAccountType(val *string) {
	_jsii_.Set(
		j,
		"storageAccountType",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) SetUltraSsdDiskIopsReadWrite(val *float64) {
	_jsii_.Set(
		j,
		"ultraSsdDiskIopsReadWrite",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) SetUltraSsdDiskMbpsReadWrite(val *float64) {
	_jsii_.Set(
		j,
		"ultraSsdDiskMbpsReadWrite",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) SetWriteAcceleratorEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"writeAcceleratorEnabled",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) ResetCreateOption() {
	_jsii_.InvokeVoid(
		w,
		"resetCreateOption",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) ResetDiskEncryptionSetId() {
	_jsii_.InvokeVoid(
		w,
		"resetDiskEncryptionSetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		w,
		"resetName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) ResetUltraSsdDiskIopsReadWrite() {
	_jsii_.InvokeVoid(
		w,
		"resetUltraSsdDiskIopsReadWrite",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) ResetUltraSsdDiskMbpsReadWrite() {
	_jsii_.InvokeVoid(
		w,
		"resetUltraSsdDiskMbpsReadWrite",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) ResetWriteAcceleratorEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetWriteAcceleratorEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetDataDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetExtension struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#name WindowsVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#publisher WindowsVirtualMachineScaleSet#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#type WindowsVirtualMachineScaleSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#type_handler_version WindowsVirtualMachineScaleSet#type_handler_version}.
	TypeHandlerVersion *string `field:"required" json:"typeHandlerVersion" yaml:"typeHandlerVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#automatic_upgrade_enabled WindowsVirtualMachineScaleSet#automatic_upgrade_enabled}.
	AutomaticUpgradeEnabled interface{} `field:"optional" json:"automaticUpgradeEnabled" yaml:"automaticUpgradeEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#auto_upgrade_minor_version WindowsVirtualMachineScaleSet#auto_upgrade_minor_version}.
	AutoUpgradeMinorVersion interface{} `field:"optional" json:"autoUpgradeMinorVersion" yaml:"autoUpgradeMinorVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#force_update_tag WindowsVirtualMachineScaleSet#force_update_tag}.
	ForceUpdateTag *string `field:"optional" json:"forceUpdateTag" yaml:"forceUpdateTag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#protected_settings WindowsVirtualMachineScaleSet#protected_settings}.
	ProtectedSettings *string `field:"optional" json:"protectedSettings" yaml:"protectedSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#provision_after_extensions WindowsVirtualMachineScaleSet#provision_after_extensions}.
	ProvisionAfterExtensions *[]*string `field:"optional" json:"provisionAfterExtensions" yaml:"provisionAfterExtensions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#settings WindowsVirtualMachineScaleSet#settings}.
	Settings *string `field:"optional" json:"settings" yaml:"settings"`
}

type WindowsVirtualMachineScaleSetExtensionList interface {
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
	Get(index *float64) WindowsVirtualMachineScaleSetExtensionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetExtensionList
type jsiiProxy_WindowsVirtualMachineScaleSetExtensionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetExtensionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsVirtualMachineScaleSetExtensionList {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetExtensionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetExtensionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetExtensionList_Override(w WindowsVirtualMachineScaleSetExtensionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetExtensionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionList) Get(index *float64) WindowsVirtualMachineScaleSetExtensionOutputReference {
	var returns WindowsVirtualMachineScaleSetExtensionOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetExtensionOutputReference interface {
	cdktf.ComplexObject
	AutomaticUpgradeEnabled() interface{}
	SetAutomaticUpgradeEnabled(val interface{})
	AutomaticUpgradeEnabledInput() interface{}
	AutoUpgradeMinorVersion() interface{}
	SetAutoUpgradeMinorVersion(val interface{})
	AutoUpgradeMinorVersionInput() interface{}
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
	ForceUpdateTag() *string
	SetForceUpdateTag(val *string)
	ForceUpdateTagInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	ProtectedSettings() *string
	SetProtectedSettings(val *string)
	ProtectedSettingsInput() *string
	ProvisionAfterExtensions() *[]*string
	SetProvisionAfterExtensions(val *[]*string)
	ProvisionAfterExtensionsInput() *[]*string
	Publisher() *string
	SetPublisher(val *string)
	PublisherInput() *string
	Settings() *string
	SetSettings(val *string)
	SettingsInput() *string
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
	TypeHandlerVersion() *string
	SetTypeHandlerVersion(val *string)
	TypeHandlerVersionInput() *string
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
	ResetAutomaticUpgradeEnabled()
	ResetAutoUpgradeMinorVersion()
	ResetForceUpdateTag()
	ResetProtectedSettings()
	ResetProvisionAfterExtensions()
	ResetSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetExtensionOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) AutomaticUpgradeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUpgradeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) AutomaticUpgradeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUpgradeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) AutoUpgradeMinorVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpgradeMinorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) AutoUpgradeMinorVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpgradeMinorVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ForceUpdateTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceUpdateTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ForceUpdateTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceUpdateTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ProtectedSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectedSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ProtectedSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectedSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ProvisionAfterExtensions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"provisionAfterExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ProvisionAfterExtensionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"provisionAfterExtensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) Publisher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) PublisherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) Settings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) TypeHandlerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeHandlerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) TypeHandlerVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeHandlerVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetExtensionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsVirtualMachineScaleSetExtensionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetExtensionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetExtensionOutputReference_Override(w WindowsVirtualMachineScaleSetExtensionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetExtensionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SetAutomaticUpgradeEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"automaticUpgradeEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SetAutoUpgradeMinorVersion(val interface{}) {
	_jsii_.Set(
		j,
		"autoUpgradeMinorVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SetForceUpdateTag(val *string) {
	_jsii_.Set(
		j,
		"forceUpdateTag",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SetProtectedSettings(val *string) {
	_jsii_.Set(
		j,
		"protectedSettings",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SetProvisionAfterExtensions(val *[]*string) {
	_jsii_.Set(
		j,
		"provisionAfterExtensions",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SetPublisher(val *string) {
	_jsii_.Set(
		j,
		"publisher",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SetSettings(val *string) {
	_jsii_.Set(
		j,
		"settings",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) SetTypeHandlerVersion(val *string) {
	_jsii_.Set(
		j,
		"typeHandlerVersion",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ResetAutomaticUpgradeEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetAutomaticUpgradeEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ResetAutoUpgradeMinorVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetAutoUpgradeMinorVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ResetForceUpdateTag() {
	_jsii_.InvokeVoid(
		w,
		"resetForceUpdateTag",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ResetProtectedSettings() {
	_jsii_.InvokeVoid(
		w,
		"resetProtectedSettings",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ResetProvisionAfterExtensions() {
	_jsii_.InvokeVoid(
		w,
		"resetProvisionAfterExtensions",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ResetSettings() {
	_jsii_.InvokeVoid(
		w,
		"resetSettings",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetExtensionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetGalleryApplications struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#package_reference_id WindowsVirtualMachineScaleSet#package_reference_id}.
	PackageReferenceId *string `field:"required" json:"packageReferenceId" yaml:"packageReferenceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#configuration_reference_blob_uri WindowsVirtualMachineScaleSet#configuration_reference_blob_uri}.
	ConfigurationReferenceBlobUri *string `field:"optional" json:"configurationReferenceBlobUri" yaml:"configurationReferenceBlobUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#order WindowsVirtualMachineScaleSet#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#tag WindowsVirtualMachineScaleSet#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

type WindowsVirtualMachineScaleSetGalleryApplicationsList interface {
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
	Get(index *float64) WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetGalleryApplicationsList
type jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetGalleryApplicationsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsVirtualMachineScaleSetGalleryApplicationsList {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetGalleryApplicationsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetGalleryApplicationsList_Override(w WindowsVirtualMachineScaleSetGalleryApplicationsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetGalleryApplicationsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList) Get(index *float64) WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference {
	var returns WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference interface {
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
	ConfigurationReferenceBlobUri() *string
	SetConfigurationReferenceBlobUri(val *string)
	ConfigurationReferenceBlobUriInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Order() *float64
	SetOrder(val *float64)
	OrderInput() *float64
	PackageReferenceId() *string
	SetPackageReferenceId(val *string)
	PackageReferenceIdInput() *string
	Tag() *string
	SetTag(val *string)
	TagInput() *string
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
	ResetConfigurationReferenceBlobUri()
	ResetOrder()
	ResetTag()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) ConfigurationReferenceBlobUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationReferenceBlobUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) ConfigurationReferenceBlobUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationReferenceBlobUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) PackageReferenceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageReferenceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) PackageReferenceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageReferenceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) Tag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) TagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetGalleryApplicationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetGalleryApplicationsOutputReference_Override(w WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) SetConfigurationReferenceBlobUri(val *string) {
	_jsii_.Set(
		j,
		"configurationReferenceBlobUri",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) SetOrder(val *float64) {
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) SetPackageReferenceId(val *string) {
	_jsii_.Set(
		j,
		"packageReferenceId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) SetTag(val *string) {
	_jsii_.Set(
		j,
		"tag",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) ResetConfigurationReferenceBlobUri() {
	_jsii_.InvokeVoid(
		w,
		"resetConfigurationReferenceBlobUri",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		w,
		"resetOrder",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) ResetTag() {
	_jsii_.InvokeVoid(
		w,
		"resetTag",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetGalleryApplicationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#type WindowsVirtualMachineScaleSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#identity_ids WindowsVirtualMachineScaleSet#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

type WindowsVirtualMachineScaleSetIdentityOutputReference interface {
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
	IdentityIds() *[]*string
	SetIdentityIds(val *[]*string)
	IdentityIdsInput() *[]*string
	InternalValue() *WindowsVirtualMachineScaleSetIdentity
	SetInternalValue(val *WindowsVirtualMachineScaleSetIdentity)
	PrincipalId() *string
	TenantId() *string
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
	ResetIdentityIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetIdentityOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) IdentityIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) IdentityIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) InternalValue() *WindowsVirtualMachineScaleSetIdentity {
	var returns *WindowsVirtualMachineScaleSetIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetIdentityOutputReference_Override(w WindowsVirtualMachineScaleSetIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) SetIdentityIds(val *[]*string) {
	_jsii_.Set(
		j,
		"identityIds",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) SetInternalValue(val *WindowsVirtualMachineScaleSetIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) ResetIdentityIds() {
	_jsii_.InvokeVoid(
		w,
		"resetIdentityIds",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetNetworkInterface struct {
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#ip_configuration WindowsVirtualMachineScaleSet#ip_configuration}
	IpConfiguration interface{} `field:"required" json:"ipConfiguration" yaml:"ipConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#name WindowsVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#dns_servers WindowsVirtualMachineScaleSet#dns_servers}.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#enable_accelerated_networking WindowsVirtualMachineScaleSet#enable_accelerated_networking}.
	EnableAcceleratedNetworking interface{} `field:"optional" json:"enableAcceleratedNetworking" yaml:"enableAcceleratedNetworking"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#enable_ip_forwarding WindowsVirtualMachineScaleSet#enable_ip_forwarding}.
	EnableIpForwarding interface{} `field:"optional" json:"enableIpForwarding" yaml:"enableIpForwarding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#network_security_group_id WindowsVirtualMachineScaleSet#network_security_group_id}.
	NetworkSecurityGroupId *string `field:"optional" json:"networkSecurityGroupId" yaml:"networkSecurityGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#primary WindowsVirtualMachineScaleSet#primary}.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
}

type WindowsVirtualMachineScaleSetNetworkInterfaceIpConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#name WindowsVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#application_gateway_backend_address_pool_ids WindowsVirtualMachineScaleSet#application_gateway_backend_address_pool_ids}.
	ApplicationGatewayBackendAddressPoolIds *[]*string `field:"optional" json:"applicationGatewayBackendAddressPoolIds" yaml:"applicationGatewayBackendAddressPoolIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#application_security_group_ids WindowsVirtualMachineScaleSet#application_security_group_ids}.
	ApplicationSecurityGroupIds *[]*string `field:"optional" json:"applicationSecurityGroupIds" yaml:"applicationSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#load_balancer_backend_address_pool_ids WindowsVirtualMachineScaleSet#load_balancer_backend_address_pool_ids}.
	LoadBalancerBackendAddressPoolIds *[]*string `field:"optional" json:"loadBalancerBackendAddressPoolIds" yaml:"loadBalancerBackendAddressPoolIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#load_balancer_inbound_nat_rules_ids WindowsVirtualMachineScaleSet#load_balancer_inbound_nat_rules_ids}.
	LoadBalancerInboundNatRulesIds *[]*string `field:"optional" json:"loadBalancerInboundNatRulesIds" yaml:"loadBalancerInboundNatRulesIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#primary WindowsVirtualMachineScaleSet#primary}.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// public_ip_address block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#public_ip_address WindowsVirtualMachineScaleSet#public_ip_address}
	PublicIpAddress interface{} `field:"optional" json:"publicIpAddress" yaml:"publicIpAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#subnet_id WindowsVirtualMachineScaleSet#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#version WindowsVirtualMachineScaleSet#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

type WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList interface {
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
	Get(index *float64) WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList
type jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList_Override(w WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList) Get(index *float64) WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference {
	var returns WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference interface {
	cdktf.ComplexObject
	ApplicationGatewayBackendAddressPoolIds() *[]*string
	SetApplicationGatewayBackendAddressPoolIds(val *[]*string)
	ApplicationGatewayBackendAddressPoolIdsInput() *[]*string
	ApplicationSecurityGroupIds() *[]*string
	SetApplicationSecurityGroupIds(val *[]*string)
	ApplicationSecurityGroupIdsInput() *[]*string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoadBalancerBackendAddressPoolIds() *[]*string
	SetLoadBalancerBackendAddressPoolIds(val *[]*string)
	LoadBalancerBackendAddressPoolIdsInput() *[]*string
	LoadBalancerInboundNatRulesIds() *[]*string
	SetLoadBalancerInboundNatRulesIds(val *[]*string)
	LoadBalancerInboundNatRulesIdsInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Primary() interface{}
	SetPrimary(val interface{})
	PrimaryInput() interface{}
	PublicIpAddress() WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList
	PublicIpAddressInput() interface{}
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
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
	PutPublicIpAddress(value interface{})
	ResetApplicationGatewayBackendAddressPoolIds()
	ResetApplicationSecurityGroupIds()
	ResetLoadBalancerBackendAddressPoolIds()
	ResetLoadBalancerInboundNatRulesIds()
	ResetPrimary()
	ResetPublicIpAddress()
	ResetSubnetId()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ApplicationGatewayBackendAddressPoolIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationGatewayBackendAddressPoolIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ApplicationGatewayBackendAddressPoolIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationGatewayBackendAddressPoolIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ApplicationSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ApplicationSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) LoadBalancerBackendAddressPoolIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerBackendAddressPoolIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) LoadBalancerBackendAddressPoolIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerBackendAddressPoolIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) LoadBalancerInboundNatRulesIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerInboundNatRulesIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) LoadBalancerInboundNatRulesIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerInboundNatRulesIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) Primary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) PrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) PublicIpAddress() WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList {
	var returns WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList
	_jsii_.Get(
		j,
		"publicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) PublicIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference_Override(w WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) SetApplicationGatewayBackendAddressPoolIds(val *[]*string) {
	_jsii_.Set(
		j,
		"applicationGatewayBackendAddressPoolIds",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) SetApplicationSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"applicationSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) SetLoadBalancerBackendAddressPoolIds(val *[]*string) {
	_jsii_.Set(
		j,
		"loadBalancerBackendAddressPoolIds",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) SetLoadBalancerInboundNatRulesIds(val *[]*string) {
	_jsii_.Set(
		j,
		"loadBalancerInboundNatRulesIds",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) SetPrimary(val interface{}) {
	_jsii_.Set(
		j,
		"primary",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) PutPublicIpAddress(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putPublicIpAddress",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ResetApplicationGatewayBackendAddressPoolIds() {
	_jsii_.InvokeVoid(
		w,
		"resetApplicationGatewayBackendAddressPoolIds",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ResetApplicationSecurityGroupIds() {
	_jsii_.InvokeVoid(
		w,
		"resetApplicationSecurityGroupIds",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ResetLoadBalancerBackendAddressPoolIds() {
	_jsii_.InvokeVoid(
		w,
		"resetLoadBalancerBackendAddressPoolIds",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ResetLoadBalancerInboundNatRulesIds() {
	_jsii_.InvokeVoid(
		w,
		"resetLoadBalancerInboundNatRulesIds",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ResetPrimary() {
	_jsii_.InvokeVoid(
		w,
		"resetPrimary",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ResetPublicIpAddress() {
	_jsii_.InvokeVoid(
		w,
		"resetPublicIpAddress",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		w,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddress struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#name WindowsVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#domain_name_label WindowsVirtualMachineScaleSet#domain_name_label}.
	DomainNameLabel *string `field:"optional" json:"domainNameLabel" yaml:"domainNameLabel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#idle_timeout_in_minutes WindowsVirtualMachineScaleSet#idle_timeout_in_minutes}.
	IdleTimeoutInMinutes *float64 `field:"optional" json:"idleTimeoutInMinutes" yaml:"idleTimeoutInMinutes"`
	// ip_tag block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#ip_tag WindowsVirtualMachineScaleSet#ip_tag}
	IpTag interface{} `field:"optional" json:"ipTag" yaml:"ipTag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#public_ip_prefix_id WindowsVirtualMachineScaleSet#public_ip_prefix_id}.
	PublicIpPrefixId *string `field:"optional" json:"publicIpPrefixId" yaml:"publicIpPrefixId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#version WindowsVirtualMachineScaleSet#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

type WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTag struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#tag WindowsVirtualMachineScaleSet#tag}.
	Tag *string `field:"required" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#type WindowsVirtualMachineScaleSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

type WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList interface {
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
	Get(index *float64) WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList
type jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList_Override(w WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList) Get(index *float64) WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference {
	var returns WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Tag() *string
	SetTag(val *string)
	TagInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) Tag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) TagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference_Override(w WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) SetTag(val *string) {
	_jsii_.Set(
		j,
		"tag",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList interface {
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
	Get(index *float64) WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList
type jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList_Override(w WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList) Get(index *float64) WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference {
	var returns WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference interface {
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
	DomainNameLabel() *string
	SetDomainNameLabel(val *string)
	DomainNameLabelInput() *string
	// Experimental.
	Fqn() *string
	IdleTimeoutInMinutes() *float64
	SetIdleTimeoutInMinutes(val *float64)
	IdleTimeoutInMinutesInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpTag() WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList
	IpTagInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	PublicIpPrefixId() *string
	SetPublicIpPrefixId(val *string)
	PublicIpPrefixIdInput() *string
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
	PutIpTag(value interface{})
	ResetDomainNameLabel()
	ResetIdleTimeoutInMinutes()
	ResetIpTag()
	ResetPublicIpPrefixId()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) DomainNameLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) DomainNameLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) IdleTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) IdleTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) IpTag() WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList {
	var returns WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList
	_jsii_.Get(
		j,
		"ipTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) IpTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) PublicIpPrefixId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpPrefixId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) PublicIpPrefixIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpPrefixIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference_Override(w WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) SetDomainNameLabel(val *string) {
	_jsii_.Set(
		j,
		"domainNameLabel",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) SetIdleTimeoutInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"idleTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) SetPublicIpPrefixId(val *string) {
	_jsii_.Set(
		j,
		"publicIpPrefixId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) PutIpTag(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putIpTag",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ResetDomainNameLabel() {
	_jsii_.InvokeVoid(
		w,
		"resetDomainNameLabel",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ResetIdleTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		w,
		"resetIdleTimeoutInMinutes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ResetIpTag() {
	_jsii_.InvokeVoid(
		w,
		"resetIpTag",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ResetPublicIpPrefixId() {
	_jsii_.InvokeVoid(
		w,
		"resetPublicIpPrefixId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetNetworkInterfaceList interface {
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
	Get(index *float64) WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetNetworkInterfaceList
type jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetNetworkInterfaceList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsVirtualMachineScaleSetNetworkInterfaceList {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetNetworkInterfaceList_Override(w WindowsVirtualMachineScaleSetNetworkInterfaceList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList) Get(index *float64) WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference {
	var returns WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference interface {
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
	DnsServers() *[]*string
	SetDnsServers(val *[]*string)
	DnsServersInput() *[]*string
	EnableAcceleratedNetworking() interface{}
	SetEnableAcceleratedNetworking(val interface{})
	EnableAcceleratedNetworkingInput() interface{}
	EnableIpForwarding() interface{}
	SetEnableIpForwarding(val interface{})
	EnableIpForwardingInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpConfiguration() WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList
	IpConfigurationInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkSecurityGroupId() *string
	SetNetworkSecurityGroupId(val *string)
	NetworkSecurityGroupIdInput() *string
	Primary() interface{}
	SetPrimary(val interface{})
	PrimaryInput() interface{}
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
	PutIpConfiguration(value interface{})
	ResetDnsServers()
	ResetEnableAcceleratedNetworking()
	ResetEnableIpForwarding()
	ResetNetworkSecurityGroupId()
	ResetPrimary()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) DnsServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) EnableAcceleratedNetworking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceleratedNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) EnableAcceleratedNetworkingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceleratedNetworkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) EnableIpForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIpForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) EnableIpForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIpForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) IpConfiguration() WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList {
	var returns WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationList
	_jsii_.Get(
		j,
		"ipConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) IpConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) NetworkSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) NetworkSecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) Primary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) PrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetNetworkInterfaceOutputReference_Override(w WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) SetDnsServers(val *[]*string) {
	_jsii_.Set(
		j,
		"dnsServers",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) SetEnableAcceleratedNetworking(val interface{}) {
	_jsii_.Set(
		j,
		"enableAcceleratedNetworking",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) SetEnableIpForwarding(val interface{}) {
	_jsii_.Set(
		j,
		"enableIpForwarding",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) SetNetworkSecurityGroupId(val *string) {
	_jsii_.Set(
		j,
		"networkSecurityGroupId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) SetPrimary(val interface{}) {
	_jsii_.Set(
		j,
		"primary",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) PutIpConfiguration(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putIpConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) ResetDnsServers() {
	_jsii_.InvokeVoid(
		w,
		"resetDnsServers",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) ResetEnableAcceleratedNetworking() {
	_jsii_.InvokeVoid(
		w,
		"resetEnableAcceleratedNetworking",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) ResetEnableIpForwarding() {
	_jsii_.InvokeVoid(
		w,
		"resetEnableIpForwarding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) ResetNetworkSecurityGroupId() {
	_jsii_.InvokeVoid(
		w,
		"resetNetworkSecurityGroupId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) ResetPrimary() {
	_jsii_.InvokeVoid(
		w,
		"resetPrimary",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetOsDisk struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#caching WindowsVirtualMachineScaleSet#caching}.
	Caching *string `field:"required" json:"caching" yaml:"caching"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#storage_account_type WindowsVirtualMachineScaleSet#storage_account_type}.
	StorageAccountType *string `field:"required" json:"storageAccountType" yaml:"storageAccountType"`
	// diff_disk_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#diff_disk_settings WindowsVirtualMachineScaleSet#diff_disk_settings}
	DiffDiskSettings *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings `field:"optional" json:"diffDiskSettings" yaml:"diffDiskSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#disk_encryption_set_id WindowsVirtualMachineScaleSet#disk_encryption_set_id}.
	DiskEncryptionSetId *string `field:"optional" json:"diskEncryptionSetId" yaml:"diskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#disk_size_gb WindowsVirtualMachineScaleSet#disk_size_gb}.
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#secure_vm_disk_encryption_set_id WindowsVirtualMachineScaleSet#secure_vm_disk_encryption_set_id}.
	SecureVmDiskEncryptionSetId *string `field:"optional" json:"secureVmDiskEncryptionSetId" yaml:"secureVmDiskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#security_encryption_type WindowsVirtualMachineScaleSet#security_encryption_type}.
	SecurityEncryptionType *string `field:"optional" json:"securityEncryptionType" yaml:"securityEncryptionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#write_accelerator_enabled WindowsVirtualMachineScaleSet#write_accelerator_enabled}.
	WriteAcceleratorEnabled interface{} `field:"optional" json:"writeAcceleratorEnabled" yaml:"writeAcceleratorEnabled"`
}

type WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#option WindowsVirtualMachineScaleSet#option}.
	Option *string `field:"required" json:"option" yaml:"option"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#placement WindowsVirtualMachineScaleSet#placement}.
	Placement *string `field:"optional" json:"placement" yaml:"placement"`
}

type WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference interface {
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
	InternalValue() *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings
	SetInternalValue(val *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings)
	Option() *string
	SetOption(val *string)
	OptionInput() *string
	Placement() *string
	SetPlacement(val *string)
	PlacementInput() *string
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
	ResetPlacement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) InternalValue() *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings {
	var returns *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) Option() *string {
	var returns *string
	_jsii_.Get(
		j,
		"option",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) OptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) Placement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) PlacementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference_Override(w WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) SetInternalValue(val *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) SetOption(val *string) {
	_jsii_.Set(
		j,
		"option",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) SetPlacement(val *string) {
	_jsii_.Set(
		j,
		"placement",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) ResetPlacement() {
	_jsii_.InvokeVoid(
		w,
		"resetPlacement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetOsDiskOutputReference interface {
	cdktf.ComplexObject
	Caching() *string
	SetCaching(val *string)
	CachingInput() *string
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
	DiffDiskSettings() WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference
	DiffDiskSettingsInput() *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings
	DiskEncryptionSetId() *string
	SetDiskEncryptionSetId(val *string)
	DiskEncryptionSetIdInput() *string
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsVirtualMachineScaleSetOsDisk
	SetInternalValue(val *WindowsVirtualMachineScaleSetOsDisk)
	SecureVmDiskEncryptionSetId() *string
	SetSecureVmDiskEncryptionSetId(val *string)
	SecureVmDiskEncryptionSetIdInput() *string
	SecurityEncryptionType() *string
	SetSecurityEncryptionType(val *string)
	SecurityEncryptionTypeInput() *string
	StorageAccountType() *string
	SetStorageAccountType(val *string)
	StorageAccountTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WriteAcceleratorEnabled() interface{}
	SetWriteAcceleratorEnabled(val interface{})
	WriteAcceleratorEnabledInput() interface{}
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
	PutDiffDiskSettings(value *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings)
	ResetDiffDiskSettings()
	ResetDiskEncryptionSetId()
	ResetDiskSizeGb()
	ResetSecureVmDiskEncryptionSetId()
	ResetSecurityEncryptionType()
	ResetWriteAcceleratorEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetOsDiskOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) Caching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) CachingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) DiffDiskSettings() WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference {
	var returns WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference
	_jsii_.Get(
		j,
		"diffDiskSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) DiffDiskSettingsInput() *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings {
	var returns *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings
	_jsii_.Get(
		j,
		"diffDiskSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) DiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) DiskEncryptionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) InternalValue() *WindowsVirtualMachineScaleSetOsDisk {
	var returns *WindowsVirtualMachineScaleSetOsDisk
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SecureVmDiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secureVmDiskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SecureVmDiskEncryptionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secureVmDiskEncryptionSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SecurityEncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityEncryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SecurityEncryptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityEncryptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) StorageAccountType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) StorageAccountTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) WriteAcceleratorEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeAcceleratorEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) WriteAcceleratorEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeAcceleratorEnabledInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetOsDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetOsDiskOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetOsDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetOsDiskOutputReference_Override(w WindowsVirtualMachineScaleSetOsDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetOsDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SetCaching(val *string) {
	_jsii_.Set(
		j,
		"caching",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SetDiskEncryptionSetId(val *string) {
	_jsii_.Set(
		j,
		"diskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SetDiskSizeGb(val *float64) {
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SetInternalValue(val *WindowsVirtualMachineScaleSetOsDisk) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SetSecureVmDiskEncryptionSetId(val *string) {
	_jsii_.Set(
		j,
		"secureVmDiskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SetSecurityEncryptionType(val *string) {
	_jsii_.Set(
		j,
		"securityEncryptionType",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SetStorageAccountType(val *string) {
	_jsii_.Set(
		j,
		"storageAccountType",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SetWriteAcceleratorEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"writeAcceleratorEnabled",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) PutDiffDiskSettings(value *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings) {
	_jsii_.InvokeVoid(
		w,
		"putDiffDiskSettings",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ResetDiffDiskSettings() {
	_jsii_.InvokeVoid(
		w,
		"resetDiffDiskSettings",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ResetDiskEncryptionSetId() {
	_jsii_.InvokeVoid(
		w,
		"resetDiskEncryptionSetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		w,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ResetSecureVmDiskEncryptionSetId() {
	_jsii_.InvokeVoid(
		w,
		"resetSecureVmDiskEncryptionSetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ResetSecurityEncryptionType() {
	_jsii_.InvokeVoid(
		w,
		"resetSecurityEncryptionType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ResetWriteAcceleratorEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetWriteAcceleratorEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetPlan struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#name WindowsVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#product WindowsVirtualMachineScaleSet#product}.
	Product *string `field:"required" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#publisher WindowsVirtualMachineScaleSet#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
}

type WindowsVirtualMachineScaleSetPlanOutputReference interface {
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
	InternalValue() *WindowsVirtualMachineScaleSetPlan
	SetInternalValue(val *WindowsVirtualMachineScaleSetPlan)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Product() *string
	SetProduct(val *string)
	ProductInput() *string
	Publisher() *string
	SetPublisher(val *string)
	PublisherInput() *string
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

// The jsii proxy struct for WindowsVirtualMachineScaleSetPlanOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) InternalValue() *WindowsVirtualMachineScaleSetPlan {
	var returns *WindowsVirtualMachineScaleSetPlan
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) Product() *string {
	var returns *string
	_jsii_.Get(
		j,
		"product",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) ProductInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) Publisher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) PublisherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetPlanOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetPlanOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetPlanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetPlanOutputReference_Override(w WindowsVirtualMachineScaleSetPlanOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetPlanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) SetInternalValue(val *WindowsVirtualMachineScaleSetPlan) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) SetProduct(val *string) {
	_jsii_.Set(
		j,
		"product",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) SetPublisher(val *string) {
	_jsii_.Set(
		j,
		"publisher",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetPlanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetRollingUpgradePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#max_batch_instance_percent WindowsVirtualMachineScaleSet#max_batch_instance_percent}.
	MaxBatchInstancePercent *float64 `field:"required" json:"maxBatchInstancePercent" yaml:"maxBatchInstancePercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#max_unhealthy_instance_percent WindowsVirtualMachineScaleSet#max_unhealthy_instance_percent}.
	MaxUnhealthyInstancePercent *float64 `field:"required" json:"maxUnhealthyInstancePercent" yaml:"maxUnhealthyInstancePercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#max_unhealthy_upgraded_instance_percent WindowsVirtualMachineScaleSet#max_unhealthy_upgraded_instance_percent}.
	MaxUnhealthyUpgradedInstancePercent *float64 `field:"required" json:"maxUnhealthyUpgradedInstancePercent" yaml:"maxUnhealthyUpgradedInstancePercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#pause_time_between_batches WindowsVirtualMachineScaleSet#pause_time_between_batches}.
	PauseTimeBetweenBatches *string `field:"required" json:"pauseTimeBetweenBatches" yaml:"pauseTimeBetweenBatches"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#cross_zone_upgrades_enabled WindowsVirtualMachineScaleSet#cross_zone_upgrades_enabled}.
	CrossZoneUpgradesEnabled interface{} `field:"optional" json:"crossZoneUpgradesEnabled" yaml:"crossZoneUpgradesEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#prioritize_unhealthy_instances_enabled WindowsVirtualMachineScaleSet#prioritize_unhealthy_instances_enabled}.
	PrioritizeUnhealthyInstancesEnabled interface{} `field:"optional" json:"prioritizeUnhealthyInstancesEnabled" yaml:"prioritizeUnhealthyInstancesEnabled"`
}

type WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference interface {
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
	CrossZoneUpgradesEnabled() interface{}
	SetCrossZoneUpgradesEnabled(val interface{})
	CrossZoneUpgradesEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsVirtualMachineScaleSetRollingUpgradePolicy
	SetInternalValue(val *WindowsVirtualMachineScaleSetRollingUpgradePolicy)
	MaxBatchInstancePercent() *float64
	SetMaxBatchInstancePercent(val *float64)
	MaxBatchInstancePercentInput() *float64
	MaxUnhealthyInstancePercent() *float64
	SetMaxUnhealthyInstancePercent(val *float64)
	MaxUnhealthyInstancePercentInput() *float64
	MaxUnhealthyUpgradedInstancePercent() *float64
	SetMaxUnhealthyUpgradedInstancePercent(val *float64)
	MaxUnhealthyUpgradedInstancePercentInput() *float64
	PauseTimeBetweenBatches() *string
	SetPauseTimeBetweenBatches(val *string)
	PauseTimeBetweenBatchesInput() *string
	PrioritizeUnhealthyInstancesEnabled() interface{}
	SetPrioritizeUnhealthyInstancesEnabled(val interface{})
	PrioritizeUnhealthyInstancesEnabledInput() interface{}
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
	ResetCrossZoneUpgradesEnabled()
	ResetPrioritizeUnhealthyInstancesEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) CrossZoneUpgradesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossZoneUpgradesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) CrossZoneUpgradesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossZoneUpgradesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) InternalValue() *WindowsVirtualMachineScaleSetRollingUpgradePolicy {
	var returns *WindowsVirtualMachineScaleSetRollingUpgradePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxBatchInstancePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchInstancePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxBatchInstancePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchInstancePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxUnhealthyInstancePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyInstancePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxUnhealthyInstancePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyInstancePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxUnhealthyUpgradedInstancePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyUpgradedInstancePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxUnhealthyUpgradedInstancePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyUpgradedInstancePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) PauseTimeBetweenBatches() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pauseTimeBetweenBatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) PauseTimeBetweenBatchesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pauseTimeBetweenBatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) PrioritizeUnhealthyInstancesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prioritizeUnhealthyInstancesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) PrioritizeUnhealthyInstancesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prioritizeUnhealthyInstancesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference_Override(w WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) SetCrossZoneUpgradesEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"crossZoneUpgradesEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) SetInternalValue(val *WindowsVirtualMachineScaleSetRollingUpgradePolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) SetMaxBatchInstancePercent(val *float64) {
	_jsii_.Set(
		j,
		"maxBatchInstancePercent",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) SetMaxUnhealthyInstancePercent(val *float64) {
	_jsii_.Set(
		j,
		"maxUnhealthyInstancePercent",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) SetMaxUnhealthyUpgradedInstancePercent(val *float64) {
	_jsii_.Set(
		j,
		"maxUnhealthyUpgradedInstancePercent",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) SetPauseTimeBetweenBatches(val *string) {
	_jsii_.Set(
		j,
		"pauseTimeBetweenBatches",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) SetPrioritizeUnhealthyInstancesEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"prioritizeUnhealthyInstancesEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ResetCrossZoneUpgradesEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetCrossZoneUpgradesEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ResetPrioritizeUnhealthyInstancesEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetPrioritizeUnhealthyInstancesEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetScaleIn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#force_deletion_enabled WindowsVirtualMachineScaleSet#force_deletion_enabled}.
	ForceDeletionEnabled interface{} `field:"optional" json:"forceDeletionEnabled" yaml:"forceDeletionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#rule WindowsVirtualMachineScaleSet#rule}.
	Rule *string `field:"optional" json:"rule" yaml:"rule"`
}

type WindowsVirtualMachineScaleSetScaleInOutputReference interface {
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
	ForceDeletionEnabled() interface{}
	SetForceDeletionEnabled(val interface{})
	ForceDeletionEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsVirtualMachineScaleSetScaleIn
	SetInternalValue(val *WindowsVirtualMachineScaleSetScaleIn)
	Rule() *string
	SetRule(val *string)
	RuleInput() *string
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
	ResetForceDeletionEnabled()
	ResetRule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetScaleInOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) ForceDeletionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeletionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) ForceDeletionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeletionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) InternalValue() *WindowsVirtualMachineScaleSetScaleIn {
	var returns *WindowsVirtualMachineScaleSetScaleIn
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) Rule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) RuleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetScaleInOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetScaleInOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetScaleInOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetScaleInOutputReference_Override(w WindowsVirtualMachineScaleSetScaleInOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetScaleInOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) SetForceDeletionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"forceDeletionEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) SetInternalValue(val *WindowsVirtualMachineScaleSetScaleIn) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) SetRule(val *string) {
	_jsii_.Set(
		j,
		"rule",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) ResetForceDeletionEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetForceDeletionEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) ResetRule() {
	_jsii_.InvokeVoid(
		w,
		"resetRule",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetScaleInOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetSecret struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#certificate WindowsVirtualMachineScaleSet#certificate}
	Certificate interface{} `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#key_vault_id WindowsVirtualMachineScaleSet#key_vault_id}.
	KeyVaultId *string `field:"required" json:"keyVaultId" yaml:"keyVaultId"`
}

type WindowsVirtualMachineScaleSetSecretCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#store WindowsVirtualMachineScaleSet#store}.
	Store *string `field:"required" json:"store" yaml:"store"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#url WindowsVirtualMachineScaleSet#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

type WindowsVirtualMachineScaleSetSecretCertificateList interface {
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
	Get(index *float64) WindowsVirtualMachineScaleSetSecretCertificateOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetSecretCertificateList
type jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetSecretCertificateList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsVirtualMachineScaleSetSecretCertificateList {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetSecretCertificateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetSecretCertificateList_Override(w WindowsVirtualMachineScaleSetSecretCertificateList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetSecretCertificateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList) Get(index *float64) WindowsVirtualMachineScaleSetSecretCertificateOutputReference {
	var returns WindowsVirtualMachineScaleSetSecretCertificateOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetSecretCertificateOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Store() *string
	SetStore(val *string)
	StoreInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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

// The jsii proxy struct for WindowsVirtualMachineScaleSetSecretCertificateOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) Store() *string {
	var returns *string
	_jsii_.Get(
		j,
		"store",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) StoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetSecretCertificateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsVirtualMachineScaleSetSecretCertificateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetSecretCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetSecretCertificateOutputReference_Override(w WindowsVirtualMachineScaleSetSecretCertificateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetSecretCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) SetStore(val *string) {
	_jsii_.Set(
		j,
		"store",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretCertificateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetSecretList interface {
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
	Get(index *float64) WindowsVirtualMachineScaleSetSecretOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetSecretList
type jsiiProxy_WindowsVirtualMachineScaleSetSecretList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetSecretList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsVirtualMachineScaleSetSecretList {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetSecretList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetSecretList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetSecretList_Override(w WindowsVirtualMachineScaleSetSecretList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetSecretList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretList) Get(index *float64) WindowsVirtualMachineScaleSetSecretOutputReference {
	var returns WindowsVirtualMachineScaleSetSecretOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetSecretOutputReference interface {
	cdktf.ComplexObject
	Certificate() WindowsVirtualMachineScaleSetSecretCertificateList
	CertificateInput() interface{}
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyVaultId() *string
	SetKeyVaultId(val *string)
	KeyVaultIdInput() *string
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
	PutCertificate(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetSecretOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) Certificate() WindowsVirtualMachineScaleSetSecretCertificateList {
	var returns WindowsVirtualMachineScaleSetSecretCertificateList
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) CertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) KeyVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) KeyVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetSecretOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsVirtualMachineScaleSetSecretOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetSecretOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetSecretOutputReference_Override(w WindowsVirtualMachineScaleSetSecretOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetSecretOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) SetKeyVaultId(val *string) {
	_jsii_.Set(
		j,
		"keyVaultId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) PutCertificate(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putCertificate",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSecretOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetSourceImageReference struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#offer WindowsVirtualMachineScaleSet#offer}.
	Offer *string `field:"required" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#publisher WindowsVirtualMachineScaleSet#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#sku WindowsVirtualMachineScaleSet#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#version WindowsVirtualMachineScaleSet#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
}

type WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference interface {
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
	InternalValue() *WindowsVirtualMachineScaleSetSourceImageReference
	SetInternalValue(val *WindowsVirtualMachineScaleSetSourceImageReference)
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

// The jsii proxy struct for WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) InternalValue() *WindowsVirtualMachineScaleSetSourceImageReference {
	var returns *WindowsVirtualMachineScaleSetSourceImageReference
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) Offer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) OfferInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) Publisher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) PublisherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetSourceImageReferenceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetSourceImageReferenceOutputReference_Override(w WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) SetInternalValue(val *WindowsVirtualMachineScaleSetSourceImageReference) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) SetOffer(val *string) {
	_jsii_.Set(
		j,
		"offer",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) SetPublisher(val *string) {
	_jsii_.Set(
		j,
		"publisher",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) SetSku(val *string) {
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetSpotRestore struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#enabled WindowsVirtualMachineScaleSet#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#timeout WindowsVirtualMachineScaleSet#timeout}.
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

type WindowsVirtualMachineScaleSetSpotRestoreOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsVirtualMachineScaleSetSpotRestore
	SetInternalValue(val *WindowsVirtualMachineScaleSetSpotRestore)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
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
	ResetEnabled()
	ResetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetSpotRestoreOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) InternalValue() *WindowsVirtualMachineScaleSetSpotRestore {
	var returns *WindowsVirtualMachineScaleSetSpotRestore
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetSpotRestoreOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetSpotRestoreOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetSpotRestoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetSpotRestoreOutputReference_Override(w WindowsVirtualMachineScaleSetSpotRestoreOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetSpotRestoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) SetInternalValue(val *WindowsVirtualMachineScaleSetSpotRestore) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) SetTimeout(val *string) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeout",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetSpotRestoreOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetTerminateNotification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#enabled WindowsVirtualMachineScaleSet#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#timeout WindowsVirtualMachineScaleSet#timeout}.
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

type WindowsVirtualMachineScaleSetTerminateNotificationOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsVirtualMachineScaleSetTerminateNotification
	SetInternalValue(val *WindowsVirtualMachineScaleSetTerminateNotification)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
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
	ResetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetTerminateNotificationOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) InternalValue() *WindowsVirtualMachineScaleSetTerminateNotification {
	var returns *WindowsVirtualMachineScaleSetTerminateNotification
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetTerminateNotificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetTerminateNotificationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetTerminateNotificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetTerminateNotificationOutputReference_Override(w WindowsVirtualMachineScaleSetTerminateNotificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetTerminateNotificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) SetInternalValue(val *WindowsVirtualMachineScaleSetTerminateNotification) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) SetTimeout(val *string) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeout",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminateNotificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetTerminationNotification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#enabled WindowsVirtualMachineScaleSet#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#timeout WindowsVirtualMachineScaleSet#timeout}.
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

type WindowsVirtualMachineScaleSetTerminationNotificationOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsVirtualMachineScaleSetTerminationNotification
	SetInternalValue(val *WindowsVirtualMachineScaleSetTerminationNotification)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
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
	ResetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetTerminationNotificationOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) InternalValue() *WindowsVirtualMachineScaleSetTerminationNotification {
	var returns *WindowsVirtualMachineScaleSetTerminationNotification
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetTerminationNotificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetTerminationNotificationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetTerminationNotificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetTerminationNotificationOutputReference_Override(w WindowsVirtualMachineScaleSetTerminationNotificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetTerminationNotificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) SetInternalValue(val *WindowsVirtualMachineScaleSetTerminationNotification) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) SetTimeout(val *string) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeout",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTerminationNotificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#create WindowsVirtualMachineScaleSet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#delete WindowsVirtualMachineScaleSet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#read WindowsVirtualMachineScaleSet#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#update WindowsVirtualMachineScaleSet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type WindowsVirtualMachineScaleSetTimeoutsOutputReference interface {
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

// The jsii proxy struct for WindowsVirtualMachineScaleSetTimeoutsOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetTimeoutsOutputReference_Override(w WindowsVirtualMachineScaleSetTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		w,
		"resetCreate",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		w,
		"resetDelete",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		w,
		"resetRead",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		w,
		"resetUpdate",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetWinrmListener struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#protocol WindowsVirtualMachineScaleSet#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#certificate_url WindowsVirtualMachineScaleSet#certificate_url}.
	CertificateUrl *string `field:"optional" json:"certificateUrl" yaml:"certificateUrl"`
}

type WindowsVirtualMachineScaleSetWinrmListenerList interface {
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
	Get(index *float64) WindowsVirtualMachineScaleSetWinrmListenerOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetWinrmListenerList
type jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetWinrmListenerList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsVirtualMachineScaleSetWinrmListenerList {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetWinrmListenerList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetWinrmListenerList_Override(w WindowsVirtualMachineScaleSetWinrmListenerList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetWinrmListenerList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList) Get(index *float64) WindowsVirtualMachineScaleSetWinrmListenerOutputReference {
	var returns WindowsVirtualMachineScaleSetWinrmListenerOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsVirtualMachineScaleSetWinrmListenerOutputReference interface {
	cdktf.ComplexObject
	CertificateUrl() *string
	SetCertificateUrl(val *string)
	CertificateUrlInput() *string
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
	ResetCertificateUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetWinrmListenerOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) CertificateUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) CertificateUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetWinrmListenerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsVirtualMachineScaleSetWinrmListenerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetWinrmListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetWinrmListenerOutputReference_Override(w WindowsVirtualMachineScaleSetWinrmListenerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetWinrmListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) SetCertificateUrl(val *string) {
	_jsii_.Set(
		j,
		"certificateUrl",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) ResetCertificateUrl() {
	_jsii_.InvokeVoid(
		w,
		"resetCertificateUrl",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetWinrmListenerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

