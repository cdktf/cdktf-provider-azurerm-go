// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSet",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "additionalCapabilities", GoGetter: "AdditionalCapabilities"},
			_jsii_.MemberProperty{JsiiProperty: "additionalCapabilitiesInput", GoGetter: "AdditionalCapabilitiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "automaticInstanceRepair", GoGetter: "AutomaticInstanceRepair"},
			_jsii_.MemberProperty{JsiiProperty: "automaticInstanceRepairInput", GoGetter: "AutomaticInstanceRepairInput"},
			_jsii_.MemberProperty{JsiiProperty: "bootDiagnostics", GoGetter: "BootDiagnostics"},
			_jsii_.MemberProperty{JsiiProperty: "bootDiagnosticsInput", GoGetter: "BootDiagnosticsInput"},
			_jsii_.MemberProperty{JsiiProperty: "capacityReservationGroupId", GoGetter: "CapacityReservationGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "capacityReservationGroupIdInput", GoGetter: "CapacityReservationGroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dataDisk", GoGetter: "DataDisk"},
			_jsii_.MemberProperty{JsiiProperty: "dataDiskInput", GoGetter: "DataDiskInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionAtHostEnabled", GoGetter: "EncryptionAtHostEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionAtHostEnabledInput", GoGetter: "EncryptionAtHostEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "evictionPolicy", GoGetter: "EvictionPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "evictionPolicyInput", GoGetter: "EvictionPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "extension", GoGetter: "Extension"},
			_jsii_.MemberProperty{JsiiProperty: "extensionInput", GoGetter: "ExtensionInput"},
			_jsii_.MemberProperty{JsiiProperty: "extensionOperationsEnabled", GoGetter: "ExtensionOperationsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "extensionOperationsEnabledInput", GoGetter: "ExtensionOperationsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "extensionsTimeBudget", GoGetter: "ExtensionsTimeBudget"},
			_jsii_.MemberProperty{JsiiProperty: "extensionsTimeBudgetInput", GoGetter: "ExtensionsTimeBudgetInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "identity", GoGetter: "Identity"},
			_jsii_.MemberProperty{JsiiProperty: "identityInput", GoGetter: "IdentityInput"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "instances", GoGetter: "Instances"},
			_jsii_.MemberProperty{JsiiProperty: "instancesInput", GoGetter: "InstancesInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "licenseType", GoGetter: "LicenseType"},
			_jsii_.MemberProperty{JsiiProperty: "licenseTypeInput", GoGetter: "LicenseTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxBidPrice", GoGetter: "MaxBidPrice"},
			_jsii_.MemberProperty{JsiiProperty: "maxBidPriceInput", GoGetter: "MaxBidPriceInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterface", GoGetter: "NetworkInterface"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceInput", GoGetter: "NetworkInterfaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "osDisk", GoGetter: "OsDisk"},
			_jsii_.MemberProperty{JsiiProperty: "osDiskInput", GoGetter: "OsDiskInput"},
			_jsii_.MemberProperty{JsiiProperty: "osProfile", GoGetter: "OsProfile"},
			_jsii_.MemberProperty{JsiiProperty: "osProfileInput", GoGetter: "OsProfileInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "plan", GoGetter: "Plan"},
			_jsii_.MemberProperty{JsiiProperty: "planInput", GoGetter: "PlanInput"},
			_jsii_.MemberProperty{JsiiProperty: "platformFaultDomainCount", GoGetter: "PlatformFaultDomainCount"},
			_jsii_.MemberProperty{JsiiProperty: "platformFaultDomainCountInput", GoGetter: "PlatformFaultDomainCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "priorityInput", GoGetter: "PriorityInput"},
			_jsii_.MemberProperty{JsiiProperty: "priorityMix", GoGetter: "PriorityMix"},
			_jsii_.MemberProperty{JsiiProperty: "priorityMixInput", GoGetter: "PriorityMixInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "proximityPlacementGroupId", GoGetter: "ProximityPlacementGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "proximityPlacementGroupIdInput", GoGetter: "ProximityPlacementGroupIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAdditionalCapabilities", GoMethod: "PutAdditionalCapabilities"},
			_jsii_.MemberMethod{JsiiMethod: "putAutomaticInstanceRepair", GoMethod: "PutAutomaticInstanceRepair"},
			_jsii_.MemberMethod{JsiiMethod: "putBootDiagnostics", GoMethod: "PutBootDiagnostics"},
			_jsii_.MemberMethod{JsiiMethod: "putDataDisk", GoMethod: "PutDataDisk"},
			_jsii_.MemberMethod{JsiiMethod: "putExtension", GoMethod: "PutExtension"},
			_jsii_.MemberMethod{JsiiMethod: "putIdentity", GoMethod: "PutIdentity"},
			_jsii_.MemberMethod{JsiiMethod: "putNetworkInterface", GoMethod: "PutNetworkInterface"},
			_jsii_.MemberMethod{JsiiMethod: "putOsDisk", GoMethod: "PutOsDisk"},
			_jsii_.MemberMethod{JsiiMethod: "putOsProfile", GoMethod: "PutOsProfile"},
			_jsii_.MemberMethod{JsiiMethod: "putPlan", GoMethod: "PutPlan"},
			_jsii_.MemberMethod{JsiiMethod: "putPriorityMix", GoMethod: "PutPriorityMix"},
			_jsii_.MemberMethod{JsiiMethod: "putSourceImageReference", GoMethod: "PutSourceImageReference"},
			_jsii_.MemberMethod{JsiiMethod: "putTerminationNotification", GoMethod: "PutTerminationNotification"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdditionalCapabilities", GoMethod: "ResetAdditionalCapabilities"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutomaticInstanceRepair", GoMethod: "ResetAutomaticInstanceRepair"},
			_jsii_.MemberMethod{JsiiMethod: "resetBootDiagnostics", GoMethod: "ResetBootDiagnostics"},
			_jsii_.MemberMethod{JsiiMethod: "resetCapacityReservationGroupId", GoMethod: "ResetCapacityReservationGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataDisk", GoMethod: "ResetDataDisk"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionAtHostEnabled", GoMethod: "ResetEncryptionAtHostEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetEvictionPolicy", GoMethod: "ResetEvictionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetExtension", GoMethod: "ResetExtension"},
			_jsii_.MemberMethod{JsiiMethod: "resetExtensionOperationsEnabled", GoMethod: "ResetExtensionOperationsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetExtensionsTimeBudget", GoMethod: "ResetExtensionsTimeBudget"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentity", GoMethod: "ResetIdentity"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstances", GoMethod: "ResetInstances"},
			_jsii_.MemberMethod{JsiiMethod: "resetLicenseType", GoMethod: "ResetLicenseType"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxBidPrice", GoMethod: "ResetMaxBidPrice"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkInterface", GoMethod: "ResetNetworkInterface"},
			_jsii_.MemberMethod{JsiiMethod: "resetOsDisk", GoMethod: "ResetOsDisk"},
			_jsii_.MemberMethod{JsiiMethod: "resetOsProfile", GoMethod: "ResetOsProfile"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlan", GoMethod: "ResetPlan"},
			_jsii_.MemberMethod{JsiiMethod: "resetPriority", GoMethod: "ResetPriority"},
			_jsii_.MemberMethod{JsiiMethod: "resetPriorityMix", GoMethod: "ResetPriorityMix"},
			_jsii_.MemberMethod{JsiiMethod: "resetProximityPlacementGroupId", GoMethod: "ResetProximityPlacementGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSinglePlacementGroup", GoMethod: "ResetSinglePlacementGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkuName", GoMethod: "ResetSkuName"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceImageId", GoMethod: "ResetSourceImageId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceImageReference", GoMethod: "ResetSourceImageReference"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTerminationNotification", GoMethod: "ResetTerminationNotification"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserDataBase64", GoMethod: "ResetUserDataBase64"},
			_jsii_.MemberMethod{JsiiMethod: "resetZoneBalance", GoMethod: "ResetZoneBalance"},
			_jsii_.MemberMethod{JsiiMethod: "resetZones", GoMethod: "ResetZones"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupName", GoGetter: "ResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupNameInput", GoGetter: "ResourceGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "singlePlacementGroup", GoGetter: "SinglePlacementGroup"},
			_jsii_.MemberProperty{JsiiProperty: "singlePlacementGroupInput", GoGetter: "SinglePlacementGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "skuName", GoGetter: "SkuName"},
			_jsii_.MemberProperty{JsiiProperty: "skuNameInput", GoGetter: "SkuNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceImageId", GoGetter: "SourceImageId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceImageIdInput", GoGetter: "SourceImageIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceImageReference", GoGetter: "SourceImageReference"},
			_jsii_.MemberProperty{JsiiProperty: "sourceImageReferenceInput", GoGetter: "SourceImageReferenceInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terminationNotification", GoGetter: "TerminationNotification"},
			_jsii_.MemberProperty{JsiiProperty: "terminationNotificationInput", GoGetter: "TerminationNotificationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "uniqueId", GoGetter: "UniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "userDataBase64", GoGetter: "UserDataBase64"},
			_jsii_.MemberProperty{JsiiProperty: "userDataBase64Input", GoGetter: "UserDataBase64Input"},
			_jsii_.MemberProperty{JsiiProperty: "zoneBalance", GoGetter: "ZoneBalance"},
			_jsii_.MemberProperty{JsiiProperty: "zoneBalanceInput", GoGetter: "ZoneBalanceInput"},
			_jsii_.MemberProperty{JsiiProperty: "zones", GoGetter: "Zones"},
			_jsii_.MemberProperty{JsiiProperty: "zonesInput", GoGetter: "ZonesInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSet{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetAdditionalCapabilities",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetAdditionalCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetAdditionalCapabilitiesOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetAdditionalCapabilitiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetUltraSsdEnabled", GoMethod: "ResetUltraSsdEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "ultraSsdEnabled", GoGetter: "UltraSsdEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "ultraSsdEnabledInput", GoGetter: "UltraSsdEnabledInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetAdditionalCapabilitiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetAutomaticInstanceRepair",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetAutomaticInstanceRepair)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetAutomaticInstanceRepairOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetAutomaticInstanceRepairOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "gracePeriod", GoGetter: "GracePeriod"},
			_jsii_.MemberProperty{JsiiProperty: "gracePeriodInput", GoGetter: "GracePeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetGracePeriod", GoMethod: "ResetGracePeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetAutomaticInstanceRepairOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetBootDiagnostics",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetBootDiagnostics)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetBootDiagnosticsOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetBootDiagnosticsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageAccountUri", GoMethod: "ResetStorageAccountUri"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountUri", GoGetter: "StorageAccountUri"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountUriInput", GoGetter: "StorageAccountUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetBootDiagnosticsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetConfig",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetDataDisk",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetDataDisk)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetDataDiskList",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetDataDiskList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetDataDiskList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetDataDiskOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetDataDiskOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "caching", GoGetter: "Caching"},
			_jsii_.MemberProperty{JsiiProperty: "cachingInput", GoGetter: "CachingInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "createOption", GoGetter: "CreateOption"},
			_jsii_.MemberProperty{JsiiProperty: "createOptionInput", GoGetter: "CreateOptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diskEncryptionSetId", GoGetter: "DiskEncryptionSetId"},
			_jsii_.MemberProperty{JsiiProperty: "diskEncryptionSetIdInput", GoGetter: "DiskEncryptionSetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "diskSizeGb", GoGetter: "DiskSizeGb"},
			_jsii_.MemberProperty{JsiiProperty: "diskSizeGbInput", GoGetter: "DiskSizeGbInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lun", GoGetter: "Lun"},
			_jsii_.MemberProperty{JsiiProperty: "lunInput", GoGetter: "LunInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreateOption", GoMethod: "ResetCreateOption"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskEncryptionSetId", GoMethod: "ResetDiskEncryptionSetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetUltraSsdDiskIopsReadWrite", GoMethod: "ResetUltraSsdDiskIopsReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "resetUltraSsdDiskMbpsReadWrite", GoMethod: "ResetUltraSsdDiskMbpsReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "resetWriteAcceleratorEnabled", GoMethod: "ResetWriteAcceleratorEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountType", GoGetter: "StorageAccountType"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountTypeInput", GoGetter: "StorageAccountTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "ultraSsdDiskIopsReadWrite", GoGetter: "UltraSsdDiskIopsReadWrite"},
			_jsii_.MemberProperty{JsiiProperty: "ultraSsdDiskIopsReadWriteInput", GoGetter: "UltraSsdDiskIopsReadWriteInput"},
			_jsii_.MemberProperty{JsiiProperty: "ultraSsdDiskMbpsReadWrite", GoGetter: "UltraSsdDiskMbpsReadWrite"},
			_jsii_.MemberProperty{JsiiProperty: "ultraSsdDiskMbpsReadWriteInput", GoGetter: "UltraSsdDiskMbpsReadWriteInput"},
			_jsii_.MemberProperty{JsiiProperty: "writeAcceleratorEnabled", GoGetter: "WriteAcceleratorEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "writeAcceleratorEnabledInput", GoGetter: "WriteAcceleratorEnabledInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetDataDiskOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetExtension",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetExtension)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetExtensionList",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetExtensionList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetExtensionList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetExtensionOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetExtensionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoUpgradeMinorVersionEnabled", GoGetter: "AutoUpgradeMinorVersionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "autoUpgradeMinorVersionEnabledInput", GoGetter: "AutoUpgradeMinorVersionEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "extensionsToProvisionAfterVmCreation", GoGetter: "ExtensionsToProvisionAfterVmCreation"},
			_jsii_.MemberProperty{JsiiProperty: "extensionsToProvisionAfterVmCreationInput", GoGetter: "ExtensionsToProvisionAfterVmCreationInput"},
			_jsii_.MemberProperty{JsiiProperty: "failureSuppressionEnabled", GoGetter: "FailureSuppressionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "failureSuppressionEnabledInput", GoGetter: "FailureSuppressionEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "forceExtensionExecutionOnChange", GoGetter: "ForceExtensionExecutionOnChange"},
			_jsii_.MemberProperty{JsiiProperty: "forceExtensionExecutionOnChangeInput", GoGetter: "ForceExtensionExecutionOnChangeInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "protectedSettings", GoGetter: "ProtectedSettings"},
			_jsii_.MemberProperty{JsiiProperty: "protectedSettingsFromKeyVault", GoGetter: "ProtectedSettingsFromKeyVault"},
			_jsii_.MemberProperty{JsiiProperty: "protectedSettingsFromKeyVaultInput", GoGetter: "ProtectedSettingsFromKeyVaultInput"},
			_jsii_.MemberProperty{JsiiProperty: "protectedSettingsInput", GoGetter: "ProtectedSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "publisher", GoGetter: "Publisher"},
			_jsii_.MemberProperty{JsiiProperty: "publisherInput", GoGetter: "PublisherInput"},
			_jsii_.MemberMethod{JsiiMethod: "putProtectedSettingsFromKeyVault", GoMethod: "PutProtectedSettingsFromKeyVault"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoUpgradeMinorVersionEnabled", GoMethod: "ResetAutoUpgradeMinorVersionEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetExtensionsToProvisionAfterVmCreation", GoMethod: "ResetExtensionsToProvisionAfterVmCreation"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailureSuppressionEnabled", GoMethod: "ResetFailureSuppressionEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetForceExtensionExecutionOnChange", GoMethod: "ResetForceExtensionExecutionOnChange"},
			_jsii_.MemberMethod{JsiiMethod: "resetProtectedSettings", GoMethod: "ResetProtectedSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetProtectedSettingsFromKeyVault", GoMethod: "ResetProtectedSettingsFromKeyVault"},
			_jsii_.MemberMethod{JsiiMethod: "resetSettings", GoMethod: "ResetSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberProperty{JsiiProperty: "settingsInput", GoGetter: "SettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeHandlerVersion", GoGetter: "TypeHandlerVersion"},
			_jsii_.MemberProperty{JsiiProperty: "typeHandlerVersionInput", GoGetter: "TypeHandlerVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetExtensionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetExtensionProtectedSettingsFromKeyVault",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetExtensionProtectedSettingsFromKeyVault)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetExtensionProtectedSettingsFromKeyVaultOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetExtensionProtectedSettingsFromKeyVaultOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secretUrl", GoGetter: "SecretUrl"},
			_jsii_.MemberProperty{JsiiProperty: "secretUrlInput", GoGetter: "SecretUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceVaultId", GoGetter: "SourceVaultId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceVaultIdInput", GoGetter: "SourceVaultIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetExtensionProtectedSettingsFromKeyVaultOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetIdentity",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetIdentityOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetIdentityOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "identityIds", GoGetter: "IdentityIds"},
			_jsii_.MemberProperty{JsiiProperty: "identityIdsInput", GoGetter: "IdentityIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetIdentityOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterface",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetNetworkInterface)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfiguration",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationList",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applicationGatewayBackendAddressPoolIds", GoGetter: "ApplicationGatewayBackendAddressPoolIds"},
			_jsii_.MemberProperty{JsiiProperty: "applicationGatewayBackendAddressPoolIdsInput", GoGetter: "ApplicationGatewayBackendAddressPoolIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "applicationSecurityGroupIds", GoGetter: "ApplicationSecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "applicationSecurityGroupIdsInput", GoGetter: "ApplicationSecurityGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerBackendAddressPoolIds", GoGetter: "LoadBalancerBackendAddressPoolIds"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerBackendAddressPoolIdsInput", GoGetter: "LoadBalancerBackendAddressPoolIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "primary", GoGetter: "Primary"},
			_jsii_.MemberProperty{JsiiProperty: "primaryInput", GoGetter: "PrimaryInput"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpAddress", GoGetter: "PublicIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpAddressInput", GoGetter: "PublicIpAddressInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPublicIpAddress", GoMethod: "PutPublicIpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetApplicationGatewayBackendAddressPoolIds", GoMethod: "ResetApplicationGatewayBackendAddressPoolIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetApplicationSecurityGroupIds", GoMethod: "ResetApplicationSecurityGroupIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoadBalancerBackendAddressPoolIds", GoMethod: "ResetLoadBalancerBackendAddressPoolIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimary", GoMethod: "ResetPrimary"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicIpAddress", GoMethod: "ResetPublicIpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubnetId", GoMethod: "ResetSubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetVersion", GoMethod: "ResetVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdInput", GoGetter: "SubnetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "versionInput", GoGetter: "VersionInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddress",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTag",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTag)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "tag", GoGetter: "Tag"},
			_jsii_.MemberProperty{JsiiProperty: "tagInput", GoGetter: "TagInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domainNameLabel", GoGetter: "DomainNameLabel"},
			_jsii_.MemberProperty{JsiiProperty: "domainNameLabelInput", GoGetter: "DomainNameLabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idleTimeoutInMinutes", GoGetter: "IdleTimeoutInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "idleTimeoutInMinutesInput", GoGetter: "IdleTimeoutInMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipTag", GoGetter: "IpTag"},
			_jsii_.MemberProperty{JsiiProperty: "ipTagInput", GoGetter: "IpTagInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpPrefixId", GoGetter: "PublicIpPrefixId"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpPrefixIdInput", GoGetter: "PublicIpPrefixIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "putIpTag", GoMethod: "PutIpTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomainNameLabel", GoMethod: "ResetDomainNameLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdleTimeoutInMinutes", GoMethod: "ResetIdleTimeoutInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpTag", GoMethod: "ResetIpTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicIpPrefixId", GoMethod: "ResetPublicIpPrefixId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkuName", GoMethod: "ResetSkuName"},
			_jsii_.MemberMethod{JsiiMethod: "resetVersion", GoMethod: "ResetVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "skuName", GoGetter: "SkuName"},
			_jsii_.MemberProperty{JsiiProperty: "skuNameInput", GoGetter: "SkuNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "versionInput", GoGetter: "VersionInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterfaceList",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetNetworkInterfaceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dnsServers", GoGetter: "DnsServers"},
			_jsii_.MemberProperty{JsiiProperty: "dnsServersInput", GoGetter: "DnsServersInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableAcceleratedNetworking", GoGetter: "EnableAcceleratedNetworking"},
			_jsii_.MemberProperty{JsiiProperty: "enableAcceleratedNetworkingInput", GoGetter: "EnableAcceleratedNetworkingInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableIpForwarding", GoGetter: "EnableIpForwarding"},
			_jsii_.MemberProperty{JsiiProperty: "enableIpForwardingInput", GoGetter: "EnableIpForwardingInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipConfiguration", GoGetter: "IpConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "ipConfigurationInput", GoGetter: "IpConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkSecurityGroupId", GoGetter: "NetworkSecurityGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "networkSecurityGroupIdInput", GoGetter: "NetworkSecurityGroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "primary", GoGetter: "Primary"},
			_jsii_.MemberProperty{JsiiProperty: "primaryInput", GoGetter: "PrimaryInput"},
			_jsii_.MemberMethod{JsiiMethod: "putIpConfiguration", GoMethod: "PutIpConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetDnsServers", GoMethod: "ResetDnsServers"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableAcceleratedNetworking", GoMethod: "ResetEnableAcceleratedNetworking"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableIpForwarding", GoMethod: "ResetEnableIpForwarding"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkSecurityGroupId", GoMethod: "ResetNetworkSecurityGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimary", GoMethod: "ResetPrimary"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsDisk",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsDiskDiffDiskSettings",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsDiskDiffDiskSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "option", GoGetter: "Option"},
			_jsii_.MemberProperty{JsiiProperty: "optionInput", GoGetter: "OptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "placement", GoGetter: "Placement"},
			_jsii_.MemberProperty{JsiiProperty: "placementInput", GoGetter: "PlacementInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlacement", GoMethod: "ResetPlacement"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsDiskOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsDiskOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "caching", GoGetter: "Caching"},
			_jsii_.MemberProperty{JsiiProperty: "cachingInput", GoGetter: "CachingInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diffDiskSettings", GoGetter: "DiffDiskSettings"},
			_jsii_.MemberProperty{JsiiProperty: "diffDiskSettingsInput", GoGetter: "DiffDiskSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "diskEncryptionSetId", GoGetter: "DiskEncryptionSetId"},
			_jsii_.MemberProperty{JsiiProperty: "diskEncryptionSetIdInput", GoGetter: "DiskEncryptionSetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "diskSizeGb", GoGetter: "DiskSizeGb"},
			_jsii_.MemberProperty{JsiiProperty: "diskSizeGbInput", GoGetter: "DiskSizeGbInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putDiffDiskSettings", GoMethod: "PutDiffDiskSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiffDiskSettings", GoMethod: "ResetDiffDiskSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskEncryptionSetId", GoMethod: "ResetDiskEncryptionSetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskSizeGb", GoMethod: "ResetDiskSizeGb"},
			_jsii_.MemberMethod{JsiiMethod: "resetWriteAcceleratorEnabled", GoMethod: "ResetWriteAcceleratorEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountType", GoGetter: "StorageAccountType"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountTypeInput", GoGetter: "StorageAccountTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "writeAcceleratorEnabled", GoGetter: "WriteAcceleratorEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "writeAcceleratorEnabledInput", GoGetter: "WriteAcceleratorEnabledInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfile",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfiguration",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileLinuxConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationAdminSshKey",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationAdminSshKey)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationAdminSshKeyList",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationAdminSshKeyList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationAdminSshKeyList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationAdminSshKeyOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationAdminSshKeyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "publicKey", GoGetter: "PublicKey"},
			_jsii_.MemberProperty{JsiiProperty: "publicKeyInput", GoGetter: "PublicKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
			_jsii_.MemberProperty{JsiiProperty: "usernameInput", GoGetter: "UsernameInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationAdminSshKeyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "adminPassword", GoGetter: "AdminPassword"},
			_jsii_.MemberProperty{JsiiProperty: "adminPasswordInput", GoGetter: "AdminPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "adminSshKey", GoGetter: "AdminSshKey"},
			_jsii_.MemberProperty{JsiiProperty: "adminSshKeyInput", GoGetter: "AdminSshKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "adminUsername", GoGetter: "AdminUsername"},
			_jsii_.MemberProperty{JsiiProperty: "adminUsernameInput", GoGetter: "AdminUsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "computerNamePrefix", GoGetter: "ComputerNamePrefix"},
			_jsii_.MemberProperty{JsiiProperty: "computerNamePrefixInput", GoGetter: "ComputerNamePrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "disablePasswordAuthentication", GoGetter: "DisablePasswordAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "disablePasswordAuthenticationInput", GoGetter: "DisablePasswordAuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "patchAssessmentMode", GoGetter: "PatchAssessmentMode"},
			_jsii_.MemberProperty{JsiiProperty: "patchAssessmentModeInput", GoGetter: "PatchAssessmentModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "patchMode", GoGetter: "PatchMode"},
			_jsii_.MemberProperty{JsiiProperty: "patchModeInput", GoGetter: "PatchModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "provisionVmAgent", GoGetter: "ProvisionVmAgent"},
			_jsii_.MemberProperty{JsiiProperty: "provisionVmAgentInput", GoGetter: "ProvisionVmAgentInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAdminSshKey", GoMethod: "PutAdminSshKey"},
			_jsii_.MemberMethod{JsiiMethod: "putSecret", GoMethod: "PutSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdminPassword", GoMethod: "ResetAdminPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdminSshKey", GoMethod: "ResetAdminSshKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetComputerNamePrefix", GoMethod: "ResetComputerNamePrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisablePasswordAuthentication", GoMethod: "ResetDisablePasswordAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetPatchAssessmentMode", GoMethod: "ResetPatchAssessmentMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetPatchMode", GoMethod: "ResetPatchMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetProvisionVmAgent", GoMethod: "ResetProvisionVmAgent"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecret", GoMethod: "ResetSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
			_jsii_.MemberProperty{JsiiProperty: "secretInput", GoGetter: "SecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecret",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretCertificate",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretCertificate)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretCertificateList",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretCertificateList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretCertificateList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretCertificateOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretCertificateOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretCertificateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretList",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "certificateInput", GoGetter: "CertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyVaultId", GoGetter: "KeyVaultId"},
			_jsii_.MemberProperty{JsiiProperty: "keyVaultIdInput", GoGetter: "KeyVaultIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCertificate", GoMethod: "PutCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customData", GoGetter: "CustomData"},
			_jsii_.MemberProperty{JsiiProperty: "customDataInput", GoGetter: "CustomDataInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "linuxConfiguration", GoGetter: "LinuxConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "linuxConfigurationInput", GoGetter: "LinuxConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putLinuxConfiguration", GoMethod: "PutLinuxConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putWindowsConfiguration", GoMethod: "PutWindowsConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomData", GoMethod: "ResetCustomData"},
			_jsii_.MemberMethod{JsiiMethod: "resetLinuxConfiguration", GoMethod: "ResetLinuxConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetWindowsConfiguration", GoMethod: "ResetWindowsConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "windowsConfiguration", GoGetter: "WindowsConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "windowsConfigurationInput", GoGetter: "WindowsConfigurationInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileWindowsConfiguration",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileWindowsConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "adminPassword", GoGetter: "AdminPassword"},
			_jsii_.MemberProperty{JsiiProperty: "adminPasswordInput", GoGetter: "AdminPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "adminUsername", GoGetter: "AdminUsername"},
			_jsii_.MemberProperty{JsiiProperty: "adminUsernameInput", GoGetter: "AdminUsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "computerNamePrefix", GoGetter: "ComputerNamePrefix"},
			_jsii_.MemberProperty{JsiiProperty: "computerNamePrefixInput", GoGetter: "ComputerNamePrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableAutomaticUpdates", GoGetter: "EnableAutomaticUpdates"},
			_jsii_.MemberProperty{JsiiProperty: "enableAutomaticUpdatesInput", GoGetter: "EnableAutomaticUpdatesInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hotpatchingEnabled", GoGetter: "HotpatchingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "hotpatchingEnabledInput", GoGetter: "HotpatchingEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "patchAssessmentMode", GoGetter: "PatchAssessmentMode"},
			_jsii_.MemberProperty{JsiiProperty: "patchAssessmentModeInput", GoGetter: "PatchAssessmentModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "patchMode", GoGetter: "PatchMode"},
			_jsii_.MemberProperty{JsiiProperty: "patchModeInput", GoGetter: "PatchModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "provisionVmAgent", GoGetter: "ProvisionVmAgent"},
			_jsii_.MemberProperty{JsiiProperty: "provisionVmAgentInput", GoGetter: "ProvisionVmAgentInput"},
			_jsii_.MemberMethod{JsiiMethod: "putSecret", GoMethod: "PutSecret"},
			_jsii_.MemberMethod{JsiiMethod: "putWinrmListener", GoMethod: "PutWinrmListener"},
			_jsii_.MemberMethod{JsiiMethod: "resetComputerNamePrefix", GoMethod: "ResetComputerNamePrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableAutomaticUpdates", GoMethod: "ResetEnableAutomaticUpdates"},
			_jsii_.MemberMethod{JsiiMethod: "resetHotpatchingEnabled", GoMethod: "ResetHotpatchingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetPatchAssessmentMode", GoMethod: "ResetPatchAssessmentMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetPatchMode", GoMethod: "ResetPatchMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetProvisionVmAgent", GoMethod: "ResetProvisionVmAgent"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecret", GoMethod: "ResetSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimezone", GoMethod: "ResetTimezone"},
			_jsii_.MemberMethod{JsiiMethod: "resetWinrmListener", GoMethod: "ResetWinrmListener"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
			_jsii_.MemberProperty{JsiiProperty: "secretInput", GoGetter: "SecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timezone", GoGetter: "Timezone"},
			_jsii_.MemberProperty{JsiiProperty: "timezoneInput", GoGetter: "TimezoneInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "winrmListener", GoGetter: "WinrmListener"},
			_jsii_.MemberProperty{JsiiProperty: "winrmListenerInput", GoGetter: "WinrmListenerInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecret",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretCertificate",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretCertificate)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretCertificateList",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretCertificateList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretCertificateList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretCertificateOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretCertificateOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "store", GoGetter: "Store"},
			_jsii_.MemberProperty{JsiiProperty: "storeInput", GoGetter: "StoreInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretCertificateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretList",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "certificateInput", GoGetter: "CertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyVaultId", GoGetter: "KeyVaultId"},
			_jsii_.MemberProperty{JsiiProperty: "keyVaultIdInput", GoGetter: "KeyVaultIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCertificate", GoMethod: "PutCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationWinrmListener",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationWinrmListener)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationWinrmListenerList",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationWinrmListenerList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationWinrmListenerList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationWinrmListenerOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationWinrmListenerOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateUrl", GoGetter: "CertificateUrl"},
			_jsii_.MemberProperty{JsiiProperty: "certificateUrlInput", GoGetter: "CertificateUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "protocolInput", GoGetter: "ProtocolInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificateUrl", GoMethod: "ResetCertificateUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationWinrmListenerOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetPlan",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetPlan)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetPlanOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetPlanOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "product", GoGetter: "Product"},
			_jsii_.MemberProperty{JsiiProperty: "productInput", GoGetter: "ProductInput"},
			_jsii_.MemberProperty{JsiiProperty: "publisher", GoGetter: "Publisher"},
			_jsii_.MemberProperty{JsiiProperty: "publisherInput", GoGetter: "PublisherInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetPlanOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetPriorityMix",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetPriorityMix)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetPriorityMixOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetPriorityMixOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "baseRegularCount", GoGetter: "BaseRegularCount"},
			_jsii_.MemberProperty{JsiiProperty: "baseRegularCountInput", GoGetter: "BaseRegularCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "regularPercentageAboveBase", GoGetter: "RegularPercentageAboveBase"},
			_jsii_.MemberProperty{JsiiProperty: "regularPercentageAboveBaseInput", GoGetter: "RegularPercentageAboveBaseInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBaseRegularCount", GoMethod: "ResetBaseRegularCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegularPercentageAboveBase", GoMethod: "ResetRegularPercentageAboveBase"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetPriorityMixOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetSourceImageReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetSourceImageReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetSourceImageReferenceOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetSourceImageReferenceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "offer", GoGetter: "Offer"},
			_jsii_.MemberProperty{JsiiProperty: "offerInput", GoGetter: "OfferInput"},
			_jsii_.MemberProperty{JsiiProperty: "publisher", GoGetter: "Publisher"},
			_jsii_.MemberProperty{JsiiProperty: "publisherInput", GoGetter: "PublisherInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sku", GoGetter: "Sku"},
			_jsii_.MemberProperty{JsiiProperty: "skuInput", GoGetter: "SkuInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "versionInput", GoGetter: "VersionInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetSourceImageReferenceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetTerminationNotification",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetTerminationNotification)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetTerminationNotificationOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetTerminationNotificationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeout", GoMethod: "ResetTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInput", GoGetter: "TimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetTerminationNotificationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetTimeouts",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetTimeoutsOutputReference",
		reflect.TypeOf((*OrchestratedVirtualMachineScaleSetTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrchestratedVirtualMachineScaleSetTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
