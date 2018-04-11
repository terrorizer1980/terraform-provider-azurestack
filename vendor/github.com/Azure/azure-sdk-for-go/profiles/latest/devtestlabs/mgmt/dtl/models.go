// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package dtl

import original "github.com/Azure/azure-sdk-for-go/services/devtestlabs/mgmt/2016-05-15/dtl"

type ArmTemplatesClient = original.ArmTemplatesClient

func NewArmTemplatesClient(subscriptionID string) ArmTemplatesClient {
	return original.NewArmTemplatesClient(subscriptionID)
}
func NewArmTemplatesClientWithBaseURI(baseURI string, subscriptionID string) ArmTemplatesClient {
	return original.NewArmTemplatesClientWithBaseURI(baseURI, subscriptionID)
}

type ArtifactsClient = original.ArtifactsClient

func NewArtifactsClient(subscriptionID string) ArtifactsClient {
	return original.NewArtifactsClient(subscriptionID)
}
func NewArtifactsClientWithBaseURI(baseURI string, subscriptionID string) ArtifactsClient {
	return original.NewArtifactsClientWithBaseURI(baseURI, subscriptionID)
}

type ArtifactSourcesClient = original.ArtifactSourcesClient

func NewArtifactSourcesClient(subscriptionID string) ArtifactSourcesClient {
	return original.NewArtifactSourcesClient(subscriptionID)
}
func NewArtifactSourcesClientWithBaseURI(baseURI string, subscriptionID string) ArtifactSourcesClient {
	return original.NewArtifactSourcesClientWithBaseURI(baseURI, subscriptionID)
}

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}

type CostsClient = original.CostsClient

func NewCostsClient(subscriptionID string) CostsClient {
	return original.NewCostsClient(subscriptionID)
}
func NewCostsClientWithBaseURI(baseURI string, subscriptionID string) CostsClient {
	return original.NewCostsClientWithBaseURI(baseURI, subscriptionID)
}

type CustomImagesClient = original.CustomImagesClient

func NewCustomImagesClient(subscriptionID string) CustomImagesClient {
	return original.NewCustomImagesClient(subscriptionID)
}
func NewCustomImagesClientWithBaseURI(baseURI string, subscriptionID string) CustomImagesClient {
	return original.NewCustomImagesClientWithBaseURI(baseURI, subscriptionID)
}

type DisksClient = original.DisksClient

func NewDisksClient(subscriptionID string) DisksClient {
	return original.NewDisksClient(subscriptionID)
}
func NewDisksClientWithBaseURI(baseURI string, subscriptionID string) DisksClient {
	return original.NewDisksClientWithBaseURI(baseURI, subscriptionID)
}

type EnvironmentsClient = original.EnvironmentsClient

func NewEnvironmentsClient(subscriptionID string) EnvironmentsClient {
	return original.NewEnvironmentsClient(subscriptionID)
}
func NewEnvironmentsClientWithBaseURI(baseURI string, subscriptionID string) EnvironmentsClient {
	return original.NewEnvironmentsClientWithBaseURI(baseURI, subscriptionID)
}

type FormulasClient = original.FormulasClient

func NewFormulasClient(subscriptionID string) FormulasClient {
	return original.NewFormulasClient(subscriptionID)
}
func NewFormulasClientWithBaseURI(baseURI string, subscriptionID string) FormulasClient {
	return original.NewFormulasClientWithBaseURI(baseURI, subscriptionID)
}

type GalleryImagesClient = original.GalleryImagesClient

func NewGalleryImagesClient(subscriptionID string) GalleryImagesClient {
	return original.NewGalleryImagesClient(subscriptionID)
}
func NewGalleryImagesClientWithBaseURI(baseURI string, subscriptionID string) GalleryImagesClient {
	return original.NewGalleryImagesClientWithBaseURI(baseURI, subscriptionID)
}

type GlobalSchedulesClient = original.GlobalSchedulesClient

func NewGlobalSchedulesClient(subscriptionID string) GlobalSchedulesClient {
	return original.NewGlobalSchedulesClient(subscriptionID)
}
func NewGlobalSchedulesClientWithBaseURI(baseURI string, subscriptionID string) GlobalSchedulesClient {
	return original.NewGlobalSchedulesClientWithBaseURI(baseURI, subscriptionID)
}

type LabsClient = original.LabsClient

func NewLabsClient(subscriptionID string) LabsClient {
	return original.NewLabsClient(subscriptionID)
}
func NewLabsClientWithBaseURI(baseURI string, subscriptionID string) LabsClient {
	return original.NewLabsClientWithBaseURI(baseURI, subscriptionID)
}

type CostThresholdStatus = original.CostThresholdStatus

const (
	Disabled CostThresholdStatus = original.Disabled
	Enabled  CostThresholdStatus = original.Enabled
)

func PossibleCostThresholdStatusValues() []CostThresholdStatus {
	return original.PossibleCostThresholdStatusValues()
}

type CostType = original.CostType

const (
	Projected   CostType = original.Projected
	Reported    CostType = original.Reported
	Unavailable CostType = original.Unavailable
)

func PossibleCostTypeValues() []CostType {
	return original.PossibleCostTypeValues()
}

type CustomImageOsType = original.CustomImageOsType

const (
	Linux   CustomImageOsType = original.Linux
	None    CustomImageOsType = original.None
	Windows CustomImageOsType = original.Windows
)

func PossibleCustomImageOsTypeValues() []CustomImageOsType {
	return original.PossibleCustomImageOsTypeValues()
}

type EnableStatus = original.EnableStatus

const (
	EnableStatusDisabled EnableStatus = original.EnableStatusDisabled
	EnableStatusEnabled  EnableStatus = original.EnableStatusEnabled
)

func PossibleEnableStatusValues() []EnableStatus {
	return original.PossibleEnableStatusValues()
}

type FileUploadOptions = original.FileUploadOptions

const (
	FileUploadOptionsNone                            FileUploadOptions = original.FileUploadOptionsNone
	FileUploadOptionsUploadFilesAndGenerateSasTokens FileUploadOptions = original.FileUploadOptionsUploadFilesAndGenerateSasTokens
)

func PossibleFileUploadOptionsValues() []FileUploadOptions {
	return original.PossibleFileUploadOptionsValues()
}

type HostCachingOptions = original.HostCachingOptions

const (
	HostCachingOptionsNone      HostCachingOptions = original.HostCachingOptionsNone
	HostCachingOptionsReadOnly  HostCachingOptions = original.HostCachingOptionsReadOnly
	HostCachingOptionsReadWrite HostCachingOptions = original.HostCachingOptionsReadWrite
)

func PossibleHostCachingOptionsValues() []HostCachingOptions {
	return original.PossibleHostCachingOptionsValues()
}

type HTTPStatusCode = original.HTTPStatusCode

const (
	Accepted                     HTTPStatusCode = original.Accepted
	BadGateway                   HTTPStatusCode = original.BadGateway
	BadRequest                   HTTPStatusCode = original.BadRequest
	Conflict                     HTTPStatusCode = original.Conflict
	Continue                     HTTPStatusCode = original.Continue
	Created                      HTTPStatusCode = original.Created
	ExpectationFailed            HTTPStatusCode = original.ExpectationFailed
	Forbidden                    HTTPStatusCode = original.Forbidden
	GatewayTimeout               HTTPStatusCode = original.GatewayTimeout
	Gone                         HTTPStatusCode = original.Gone
	HTTPVersionNotSupported      HTTPStatusCode = original.HTTPVersionNotSupported
	InternalServerError          HTTPStatusCode = original.InternalServerError
	LengthRequired               HTTPStatusCode = original.LengthRequired
	MethodNotAllowed             HTTPStatusCode = original.MethodNotAllowed
	MovedPermanently             HTTPStatusCode = original.MovedPermanently
	MultipleChoices              HTTPStatusCode = original.MultipleChoices
	NoContent                    HTTPStatusCode = original.NoContent
	NonAuthoritativeInformation  HTTPStatusCode = original.NonAuthoritativeInformation
	NotAcceptable                HTTPStatusCode = original.NotAcceptable
	NotFound                     HTTPStatusCode = original.NotFound
	NotImplemented               HTTPStatusCode = original.NotImplemented
	NotModified                  HTTPStatusCode = original.NotModified
	OK                           HTTPStatusCode = original.OK
	PartialContent               HTTPStatusCode = original.PartialContent
	PaymentRequired              HTTPStatusCode = original.PaymentRequired
	PreconditionFailed           HTTPStatusCode = original.PreconditionFailed
	ProxyAuthenticationRequired  HTTPStatusCode = original.ProxyAuthenticationRequired
	Redirect                     HTTPStatusCode = original.Redirect
	RequestedRangeNotSatisfiable HTTPStatusCode = original.RequestedRangeNotSatisfiable
	RequestEntityTooLarge        HTTPStatusCode = original.RequestEntityTooLarge
	RequestTimeout               HTTPStatusCode = original.RequestTimeout
	RequestURITooLong            HTTPStatusCode = original.RequestURITooLong
	ResetContent                 HTTPStatusCode = original.ResetContent
	SeeOther                     HTTPStatusCode = original.SeeOther
	ServiceUnavailable           HTTPStatusCode = original.ServiceUnavailable
	SwitchingProtocols           HTTPStatusCode = original.SwitchingProtocols
	TemporaryRedirect            HTTPStatusCode = original.TemporaryRedirect
	Unauthorized                 HTTPStatusCode = original.Unauthorized
	UnsupportedMediaType         HTTPStatusCode = original.UnsupportedMediaType
	Unused                       HTTPStatusCode = original.Unused
	UpgradeRequired              HTTPStatusCode = original.UpgradeRequired
	UseProxy                     HTTPStatusCode = original.UseProxy
)

func PossibleHTTPStatusCodeValues() []HTTPStatusCode {
	return original.PossibleHTTPStatusCodeValues()
}

type LinuxOsState = original.LinuxOsState

const (
	DeprovisionApplied   LinuxOsState = original.DeprovisionApplied
	DeprovisionRequested LinuxOsState = original.DeprovisionRequested
	NonDeprovisioned     LinuxOsState = original.NonDeprovisioned
)

func PossibleLinuxOsStateValues() []LinuxOsState {
	return original.PossibleLinuxOsStateValues()
}

type NotificationChannelEventType = original.NotificationChannelEventType

const (
	AutoShutdown NotificationChannelEventType = original.AutoShutdown
	Cost         NotificationChannelEventType = original.Cost
)

func PossibleNotificationChannelEventTypeValues() []NotificationChannelEventType {
	return original.PossibleNotificationChannelEventTypeValues()
}

type NotificationStatus = original.NotificationStatus

const (
	NotificationStatusDisabled NotificationStatus = original.NotificationStatusDisabled
	NotificationStatusEnabled  NotificationStatus = original.NotificationStatusEnabled
)

func PossibleNotificationStatusValues() []NotificationStatus {
	return original.PossibleNotificationStatusValues()
}

type PolicyEvaluatorType = original.PolicyEvaluatorType

const (
	AllowedValuesPolicy PolicyEvaluatorType = original.AllowedValuesPolicy
	MaxValuePolicy      PolicyEvaluatorType = original.MaxValuePolicy
)

func PossiblePolicyEvaluatorTypeValues() []PolicyEvaluatorType {
	return original.PossiblePolicyEvaluatorTypeValues()
}

type PolicyFactName = original.PolicyFactName

const (
	PolicyFactNameGalleryImage                PolicyFactName = original.PolicyFactNameGalleryImage
	PolicyFactNameLabPremiumVMCount           PolicyFactName = original.PolicyFactNameLabPremiumVMCount
	PolicyFactNameLabTargetCost               PolicyFactName = original.PolicyFactNameLabTargetCost
	PolicyFactNameLabVMCount                  PolicyFactName = original.PolicyFactNameLabVMCount
	PolicyFactNameLabVMSize                   PolicyFactName = original.PolicyFactNameLabVMSize
	PolicyFactNameUserOwnedLabPremiumVMCount  PolicyFactName = original.PolicyFactNameUserOwnedLabPremiumVMCount
	PolicyFactNameUserOwnedLabVMCount         PolicyFactName = original.PolicyFactNameUserOwnedLabVMCount
	PolicyFactNameUserOwnedLabVMCountInSubnet PolicyFactName = original.PolicyFactNameUserOwnedLabVMCountInSubnet
)

func PossiblePolicyFactNameValues() []PolicyFactName {
	return original.PossiblePolicyFactNameValues()
}

type PolicyStatus = original.PolicyStatus

const (
	PolicyStatusDisabled PolicyStatus = original.PolicyStatusDisabled
	PolicyStatusEnabled  PolicyStatus = original.PolicyStatusEnabled
)

func PossiblePolicyStatusValues() []PolicyStatus {
	return original.PossiblePolicyStatusValues()
}

type PremiumDataDisk = original.PremiumDataDisk

const (
	PremiumDataDiskDisabled PremiumDataDisk = original.PremiumDataDiskDisabled
	PremiumDataDiskEnabled  PremiumDataDisk = original.PremiumDataDiskEnabled
)

func PossiblePremiumDataDiskValues() []PremiumDataDisk {
	return original.PossiblePremiumDataDiskValues()
}

type ReportingCycleType = original.ReportingCycleType

const (
	CalendarMonth ReportingCycleType = original.CalendarMonth
	Custom        ReportingCycleType = original.Custom
)

func PossibleReportingCycleTypeValues() []ReportingCycleType {
	return original.PossibleReportingCycleTypeValues()
}

type SourceControlType = original.SourceControlType

const (
	GitHub SourceControlType = original.GitHub
	VsoGit SourceControlType = original.VsoGit
)

func PossibleSourceControlTypeValues() []SourceControlType {
	return original.PossibleSourceControlTypeValues()
}

type StorageType = original.StorageType

const (
	Premium  StorageType = original.Premium
	Standard StorageType = original.Standard
)

func PossibleStorageTypeValues() []StorageType {
	return original.PossibleStorageTypeValues()
}

type TargetCostStatus = original.TargetCostStatus

const (
	TargetCostStatusDisabled TargetCostStatus = original.TargetCostStatusDisabled
	TargetCostStatusEnabled  TargetCostStatus = original.TargetCostStatusEnabled
)

func PossibleTargetCostStatusValues() []TargetCostStatus {
	return original.PossibleTargetCostStatusValues()
}

type TransportProtocol = original.TransportProtocol

const (
	TCP TransportProtocol = original.TCP
	UDP TransportProtocol = original.UDP
)

func PossibleTransportProtocolValues() []TransportProtocol {
	return original.PossibleTransportProtocolValues()
}

type UsagePermissionType = original.UsagePermissionType

const (
	Allow   UsagePermissionType = original.Allow
	Default UsagePermissionType = original.Default
	Deny    UsagePermissionType = original.Deny
)

func PossibleUsagePermissionTypeValues() []UsagePermissionType {
	return original.PossibleUsagePermissionTypeValues()
}

type VirtualMachineCreationSource = original.VirtualMachineCreationSource

const (
	FromCustomImage  VirtualMachineCreationSource = original.FromCustomImage
	FromGalleryImage VirtualMachineCreationSource = original.FromGalleryImage
)

func PossibleVirtualMachineCreationSourceValues() []VirtualMachineCreationSource {
	return original.PossibleVirtualMachineCreationSourceValues()
}

type WindowsOsState = original.WindowsOsState

const (
	NonSysprepped    WindowsOsState = original.NonSysprepped
	SysprepApplied   WindowsOsState = original.SysprepApplied
	SysprepRequested WindowsOsState = original.SysprepRequested
)

func PossibleWindowsOsStateValues() []WindowsOsState {
	return original.PossibleWindowsOsStateValues()
}

type ApplicableSchedule = original.ApplicableSchedule
type ApplicableScheduleFragment = original.ApplicableScheduleFragment
type ApplicableScheduleProperties = original.ApplicableScheduleProperties
type ApplicableSchedulePropertiesFragment = original.ApplicableSchedulePropertiesFragment
type ApplyArtifactsRequest = original.ApplyArtifactsRequest
type ArmTemplate = original.ArmTemplate
type ArmTemplateInfo = original.ArmTemplateInfo
type ArmTemplateParameterProperties = original.ArmTemplateParameterProperties
type ArmTemplateProperties = original.ArmTemplateProperties
type Artifact = original.Artifact
type ArtifactDeploymentStatusProperties = original.ArtifactDeploymentStatusProperties
type ArtifactDeploymentStatusPropertiesFragment = original.ArtifactDeploymentStatusPropertiesFragment
type ArtifactInstallProperties = original.ArtifactInstallProperties
type ArtifactInstallPropertiesFragment = original.ArtifactInstallPropertiesFragment
type ArtifactParameterProperties = original.ArtifactParameterProperties
type ArtifactParameterPropertiesFragment = original.ArtifactParameterPropertiesFragment
type ArtifactProperties = original.ArtifactProperties
type ArtifactSource = original.ArtifactSource
type ArtifactSourceFragment = original.ArtifactSourceFragment
type ArtifactSourceProperties = original.ArtifactSourceProperties
type ArtifactSourcePropertiesFragment = original.ArtifactSourcePropertiesFragment
type AttachDiskProperties = original.AttachDiskProperties
type AttachNewDataDiskOptions = original.AttachNewDataDiskOptions
type BulkCreationParameters = original.BulkCreationParameters
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type ComputeDataDisk = original.ComputeDataDisk
type ComputeDataDiskFragment = original.ComputeDataDiskFragment
type ComputeVMInstanceViewStatus = original.ComputeVMInstanceViewStatus
type ComputeVMInstanceViewStatusFragment = original.ComputeVMInstanceViewStatusFragment
type ComputeVMProperties = original.ComputeVMProperties
type ComputeVMPropertiesFragment = original.ComputeVMPropertiesFragment
type CostThresholdProperties = original.CostThresholdProperties
type CustomImage = original.CustomImage
type CustomImageProperties = original.CustomImageProperties
type CustomImagePropertiesCustom = original.CustomImagePropertiesCustom
type CustomImagePropertiesFromVM = original.CustomImagePropertiesFromVM
type CustomImagesCreateOrUpdateFuture = original.CustomImagesCreateOrUpdateFuture
type CustomImagesDeleteFuture = original.CustomImagesDeleteFuture
type DataDiskProperties = original.DataDiskProperties
type DayDetails = original.DayDetails
type DayDetailsFragment = original.DayDetailsFragment
type DetachDataDiskProperties = original.DetachDataDiskProperties
type DetachDiskProperties = original.DetachDiskProperties
type Disk = original.Disk
type DiskProperties = original.DiskProperties
type DisksAttachFuture = original.DisksAttachFuture
type DisksCreateOrUpdateFuture = original.DisksCreateOrUpdateFuture
type DisksDeleteFuture = original.DisksDeleteFuture
type DisksDetachFuture = original.DisksDetachFuture
type Environment = original.Environment
type EnvironmentDeploymentProperties = original.EnvironmentDeploymentProperties
type EnvironmentProperties = original.EnvironmentProperties
type EnvironmentsCreateOrUpdateFuture = original.EnvironmentsCreateOrUpdateFuture
type EnvironmentsDeleteFuture = original.EnvironmentsDeleteFuture
type EvaluatePoliciesProperties = original.EvaluatePoliciesProperties
type EvaluatePoliciesRequest = original.EvaluatePoliciesRequest
type EvaluatePoliciesResponse = original.EvaluatePoliciesResponse
type Event = original.Event
type EventFragment = original.EventFragment
type ExportResourceUsageParameters = original.ExportResourceUsageParameters
type ExternalSubnet = original.ExternalSubnet
type ExternalSubnetFragment = original.ExternalSubnetFragment
type Formula = original.Formula
type FormulaProperties = original.FormulaProperties
type FormulaPropertiesFromVM = original.FormulaPropertiesFromVM
type FormulasCreateOrUpdateFuture = original.FormulasCreateOrUpdateFuture
type GalleryImage = original.GalleryImage
type GalleryImageProperties = original.GalleryImageProperties
type GalleryImageReference = original.GalleryImageReference
type GalleryImageReferenceFragment = original.GalleryImageReferenceFragment
type GenerateArmTemplateRequest = original.GenerateArmTemplateRequest
type GenerateUploadURIParameter = original.GenerateUploadURIParameter
type GenerateUploadURIResponse = original.GenerateUploadURIResponse
type GlobalSchedulesExecuteFuture = original.GlobalSchedulesExecuteFuture
type GlobalSchedulesRetargetFuture = original.GlobalSchedulesRetargetFuture
type HourDetails = original.HourDetails
type HourDetailsFragment = original.HourDetailsFragment
type IdentityProperties = original.IdentityProperties
type InboundNatRule = original.InboundNatRule
type InboundNatRuleFragment = original.InboundNatRuleFragment
type Lab = original.Lab
type LabCost = original.LabCost
type LabCostDetailsProperties = original.LabCostDetailsProperties
type LabCostProperties = original.LabCostProperties
type LabCostSummaryProperties = original.LabCostSummaryProperties
type LabFragment = original.LabFragment
type LabProperties = original.LabProperties
type LabPropertiesFragment = original.LabPropertiesFragment
type LabResourceCostProperties = original.LabResourceCostProperties
type LabsClaimAnyVMFuture = original.LabsClaimAnyVMFuture
type LabsCreateEnvironmentFuture = original.LabsCreateEnvironmentFuture
type LabsCreateOrUpdateFuture = original.LabsCreateOrUpdateFuture
type LabsDeleteFuture = original.LabsDeleteFuture
type LabsExportResourceUsageFuture = original.LabsExportResourceUsageFuture
type LabVhd = original.LabVhd
type LabVirtualMachine = original.LabVirtualMachine
type LabVirtualMachineCreationParameter = original.LabVirtualMachineCreationParameter
type LabVirtualMachineCreationParameterProperties = original.LabVirtualMachineCreationParameterProperties
type LabVirtualMachineFragment = original.LabVirtualMachineFragment
type LabVirtualMachineProperties = original.LabVirtualMachineProperties
type LabVirtualMachinePropertiesFragment = original.LabVirtualMachinePropertiesFragment
type LinuxOsInfo = original.LinuxOsInfo
type NetworkInterfaceProperties = original.NetworkInterfaceProperties
type NetworkInterfacePropertiesFragment = original.NetworkInterfacePropertiesFragment
type NotificationChannel = original.NotificationChannel
type NotificationChannelFragment = original.NotificationChannelFragment
type NotificationChannelProperties = original.NotificationChannelProperties
type NotificationChannelPropertiesFragment = original.NotificationChannelPropertiesFragment
type NotificationSettings = original.NotificationSettings
type NotificationSettingsFragment = original.NotificationSettingsFragment
type NotifyParameters = original.NotifyParameters
type OperationError = original.OperationError
type OperationMetadata = original.OperationMetadata
type OperationMetadataDisplay = original.OperationMetadataDisplay
type OperationResult = original.OperationResult
type ParameterInfo = original.ParameterInfo
type ParametersValueFileInfo = original.ParametersValueFileInfo
type PercentageCostThresholdProperties = original.PercentageCostThresholdProperties
type Policy = original.Policy
type PolicyFragment = original.PolicyFragment
type PolicyProperties = original.PolicyProperties
type PolicyPropertiesFragment = original.PolicyPropertiesFragment
type PolicySetResult = original.PolicySetResult
type PolicyViolation = original.PolicyViolation
type Port = original.Port
type PortFragment = original.PortFragment
type ProviderOperationResult = original.ProviderOperationResult
type ProviderOperationResultIterator = original.ProviderOperationResultIterator
type ProviderOperationResultPage = original.ProviderOperationResultPage
type Resource = original.Resource
type ResponseWithContinuationArmTemplate = original.ResponseWithContinuationArmTemplate
type ResponseWithContinuationArmTemplateIterator = original.ResponseWithContinuationArmTemplateIterator
type ResponseWithContinuationArmTemplatePage = original.ResponseWithContinuationArmTemplatePage
type ResponseWithContinuationArtifact = original.ResponseWithContinuationArtifact
type ResponseWithContinuationArtifactIterator = original.ResponseWithContinuationArtifactIterator
type ResponseWithContinuationArtifactPage = original.ResponseWithContinuationArtifactPage
type ResponseWithContinuationArtifactSource = original.ResponseWithContinuationArtifactSource
type ResponseWithContinuationArtifactSourceIterator = original.ResponseWithContinuationArtifactSourceIterator
type ResponseWithContinuationArtifactSourcePage = original.ResponseWithContinuationArtifactSourcePage
type ResponseWithContinuationCustomImage = original.ResponseWithContinuationCustomImage
type ResponseWithContinuationCustomImageIterator = original.ResponseWithContinuationCustomImageIterator
type ResponseWithContinuationCustomImagePage = original.ResponseWithContinuationCustomImagePage
type ResponseWithContinuationDisk = original.ResponseWithContinuationDisk
type ResponseWithContinuationDiskIterator = original.ResponseWithContinuationDiskIterator
type ResponseWithContinuationDiskPage = original.ResponseWithContinuationDiskPage
type ResponseWithContinuationDtlEnvironment = original.ResponseWithContinuationDtlEnvironment
type ResponseWithContinuationDtlEnvironmentIterator = original.ResponseWithContinuationDtlEnvironmentIterator
type ResponseWithContinuationDtlEnvironmentPage = original.ResponseWithContinuationDtlEnvironmentPage
type ResponseWithContinuationFormula = original.ResponseWithContinuationFormula
type ResponseWithContinuationFormulaIterator = original.ResponseWithContinuationFormulaIterator
type ResponseWithContinuationFormulaPage = original.ResponseWithContinuationFormulaPage
type ResponseWithContinuationGalleryImage = original.ResponseWithContinuationGalleryImage
type ResponseWithContinuationGalleryImageIterator = original.ResponseWithContinuationGalleryImageIterator
type ResponseWithContinuationGalleryImagePage = original.ResponseWithContinuationGalleryImagePage
type ResponseWithContinuationLab = original.ResponseWithContinuationLab
type ResponseWithContinuationLabIterator = original.ResponseWithContinuationLabIterator
type ResponseWithContinuationLabPage = original.ResponseWithContinuationLabPage
type ResponseWithContinuationLabVhd = original.ResponseWithContinuationLabVhd
type ResponseWithContinuationLabVhdIterator = original.ResponseWithContinuationLabVhdIterator
type ResponseWithContinuationLabVhdPage = original.ResponseWithContinuationLabVhdPage
type ResponseWithContinuationLabVirtualMachine = original.ResponseWithContinuationLabVirtualMachine
type ResponseWithContinuationLabVirtualMachineIterator = original.ResponseWithContinuationLabVirtualMachineIterator
type ResponseWithContinuationLabVirtualMachinePage = original.ResponseWithContinuationLabVirtualMachinePage
type ResponseWithContinuationNotificationChannel = original.ResponseWithContinuationNotificationChannel
type ResponseWithContinuationNotificationChannelIterator = original.ResponseWithContinuationNotificationChannelIterator
type ResponseWithContinuationNotificationChannelPage = original.ResponseWithContinuationNotificationChannelPage
type ResponseWithContinuationPolicy = original.ResponseWithContinuationPolicy
type ResponseWithContinuationPolicyIterator = original.ResponseWithContinuationPolicyIterator
type ResponseWithContinuationPolicyPage = original.ResponseWithContinuationPolicyPage
type ResponseWithContinuationSchedule = original.ResponseWithContinuationSchedule
type ResponseWithContinuationScheduleIterator = original.ResponseWithContinuationScheduleIterator
type ResponseWithContinuationSchedulePage = original.ResponseWithContinuationSchedulePage
type ResponseWithContinuationSecret = original.ResponseWithContinuationSecret
type ResponseWithContinuationSecretIterator = original.ResponseWithContinuationSecretIterator
type ResponseWithContinuationSecretPage = original.ResponseWithContinuationSecretPage
type ResponseWithContinuationServiceRunner = original.ResponseWithContinuationServiceRunner
type ResponseWithContinuationServiceRunnerIterator = original.ResponseWithContinuationServiceRunnerIterator
type ResponseWithContinuationServiceRunnerPage = original.ResponseWithContinuationServiceRunnerPage
type ResponseWithContinuationUser = original.ResponseWithContinuationUser
type ResponseWithContinuationUserIterator = original.ResponseWithContinuationUserIterator
type ResponseWithContinuationUserPage = original.ResponseWithContinuationUserPage
type ResponseWithContinuationVirtualNetwork = original.ResponseWithContinuationVirtualNetwork
type ResponseWithContinuationVirtualNetworkIterator = original.ResponseWithContinuationVirtualNetworkIterator
type ResponseWithContinuationVirtualNetworkPage = original.ResponseWithContinuationVirtualNetworkPage
type RetargetScheduleProperties = original.RetargetScheduleProperties
type Schedule = original.Schedule
type ScheduleFragment = original.ScheduleFragment
type ScheduleProperties = original.ScheduleProperties
type SchedulePropertiesFragment = original.SchedulePropertiesFragment
type SchedulesExecuteFuture = original.SchedulesExecuteFuture
type Secret = original.Secret
type SecretProperties = original.SecretProperties
type ServiceRunner = original.ServiceRunner
type SharedPublicIPAddressConfiguration = original.SharedPublicIPAddressConfiguration
type SharedPublicIPAddressConfigurationFragment = original.SharedPublicIPAddressConfigurationFragment
type ShutdownNotificationContent = original.ShutdownNotificationContent
type Subnet = original.Subnet
type SubnetFragment = original.SubnetFragment
type SubnetOverride = original.SubnetOverride
type SubnetOverrideFragment = original.SubnetOverrideFragment
type SubnetSharedPublicIPAddressConfiguration = original.SubnetSharedPublicIPAddressConfiguration
type SubnetSharedPublicIPAddressConfigurationFragment = original.SubnetSharedPublicIPAddressConfigurationFragment
type TargetCostProperties = original.TargetCostProperties
type User = original.User
type UserFragment = original.UserFragment
type UserIdentity = original.UserIdentity
type UserIdentityFragment = original.UserIdentityFragment
type UserProperties = original.UserProperties
type UserPropertiesFragment = original.UserPropertiesFragment
type UsersDeleteFuture = original.UsersDeleteFuture
type UserSecretStore = original.UserSecretStore
type UserSecretStoreFragment = original.UserSecretStoreFragment
type VirtualMachinesAddDataDiskFuture = original.VirtualMachinesAddDataDiskFuture
type VirtualMachinesApplyArtifactsFuture = original.VirtualMachinesApplyArtifactsFuture
type VirtualMachineSchedulesExecuteFuture = original.VirtualMachineSchedulesExecuteFuture
type VirtualMachinesClaimFuture = original.VirtualMachinesClaimFuture
type VirtualMachinesCreateOrUpdateFuture = original.VirtualMachinesCreateOrUpdateFuture
type VirtualMachinesDeleteFuture = original.VirtualMachinesDeleteFuture
type VirtualMachinesDetachDataDiskFuture = original.VirtualMachinesDetachDataDiskFuture
type VirtualMachinesStartFuture = original.VirtualMachinesStartFuture
type VirtualMachinesStopFuture = original.VirtualMachinesStopFuture
type VirtualNetwork = original.VirtualNetwork
type VirtualNetworkFragment = original.VirtualNetworkFragment
type VirtualNetworkProperties = original.VirtualNetworkProperties
type VirtualNetworkPropertiesFragment = original.VirtualNetworkPropertiesFragment
type VirtualNetworksCreateOrUpdateFuture = original.VirtualNetworksCreateOrUpdateFuture
type VirtualNetworksDeleteFuture = original.VirtualNetworksDeleteFuture
type WeekDetails = original.WeekDetails
type WeekDetailsFragment = original.WeekDetailsFragment
type WindowsOsInfo = original.WindowsOsInfo
type NotificationChannelsClient = original.NotificationChannelsClient

func NewNotificationChannelsClient(subscriptionID string) NotificationChannelsClient {
	return original.NewNotificationChannelsClient(subscriptionID)
}
func NewNotificationChannelsClientWithBaseURI(baseURI string, subscriptionID string) NotificationChannelsClient {
	return original.NewNotificationChannelsClientWithBaseURI(baseURI, subscriptionID)
}

type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}

type PoliciesClient = original.PoliciesClient

func NewPoliciesClient(subscriptionID string) PoliciesClient {
	return original.NewPoliciesClient(subscriptionID)
}
func NewPoliciesClientWithBaseURI(baseURI string, subscriptionID string) PoliciesClient {
	return original.NewPoliciesClientWithBaseURI(baseURI, subscriptionID)
}

type PolicySetsClient = original.PolicySetsClient

func NewPolicySetsClient(subscriptionID string) PolicySetsClient {
	return original.NewPolicySetsClient(subscriptionID)
}
func NewPolicySetsClientWithBaseURI(baseURI string, subscriptionID string) PolicySetsClient {
	return original.NewPolicySetsClientWithBaseURI(baseURI, subscriptionID)
}

type ProviderOperationsClient = original.ProviderOperationsClient

func NewProviderOperationsClient(subscriptionID string) ProviderOperationsClient {
	return original.NewProviderOperationsClient(subscriptionID)
}
func NewProviderOperationsClientWithBaseURI(baseURI string, subscriptionID string) ProviderOperationsClient {
	return original.NewProviderOperationsClientWithBaseURI(baseURI, subscriptionID)
}

type SchedulesClient = original.SchedulesClient

func NewSchedulesClient(subscriptionID string) SchedulesClient {
	return original.NewSchedulesClient(subscriptionID)
}
func NewSchedulesClientWithBaseURI(baseURI string, subscriptionID string) SchedulesClient {
	return original.NewSchedulesClientWithBaseURI(baseURI, subscriptionID)
}

type SecretsClient = original.SecretsClient

func NewSecretsClient(subscriptionID string) SecretsClient {
	return original.NewSecretsClient(subscriptionID)
}
func NewSecretsClientWithBaseURI(baseURI string, subscriptionID string) SecretsClient {
	return original.NewSecretsClientWithBaseURI(baseURI, subscriptionID)
}

type ServiceRunnersClient = original.ServiceRunnersClient

func NewServiceRunnersClient(subscriptionID string) ServiceRunnersClient {
	return original.NewServiceRunnersClient(subscriptionID)
}
func NewServiceRunnersClientWithBaseURI(baseURI string, subscriptionID string) ServiceRunnersClient {
	return original.NewServiceRunnersClientWithBaseURI(baseURI, subscriptionID)
}

type UsersClient = original.UsersClient

func NewUsersClient(subscriptionID string) UsersClient {
	return original.NewUsersClient(subscriptionID)
}
func NewUsersClientWithBaseURI(baseURI string, subscriptionID string) UsersClient {
	return original.NewUsersClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}

type VirtualMachinesClient = original.VirtualMachinesClient

func NewVirtualMachinesClient(subscriptionID string) VirtualMachinesClient {
	return original.NewVirtualMachinesClient(subscriptionID)
}
func NewVirtualMachinesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachinesClient {
	return original.NewVirtualMachinesClientWithBaseURI(baseURI, subscriptionID)
}

type VirtualMachineSchedulesClient = original.VirtualMachineSchedulesClient

func NewVirtualMachineSchedulesClient(subscriptionID string) VirtualMachineSchedulesClient {
	return original.NewVirtualMachineSchedulesClient(subscriptionID)
}
func NewVirtualMachineSchedulesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineSchedulesClient {
	return original.NewVirtualMachineSchedulesClientWithBaseURI(baseURI, subscriptionID)
}

type VirtualNetworksClient = original.VirtualNetworksClient

func NewVirtualNetworksClient(subscriptionID string) VirtualNetworksClient {
	return original.NewVirtualNetworksClient(subscriptionID)
}
func NewVirtualNetworksClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworksClient {
	return original.NewVirtualNetworksClientWithBaseURI(baseURI, subscriptionID)
}