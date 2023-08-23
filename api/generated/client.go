// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package generated

import (
	"context"
	"net/http"
	"time"

	"github.com/Yamashou/gqlgenc/clientv2"
)

type GQLClient interface {
	GetCloudProfile(ctx context.Context, where CloudProfileWhereUniqueInput, interceptors ...clientv2.RequestInterceptor) (*GetCloudProfile, error)
	GetCloudProfiles(ctx context.Context, whereAccount AccountWhereUniqueInput, whereCloudProfile *CloudProfileWhereInput, interceptors ...clientv2.RequestInterceptor) (*GetCloudProfiles, error)
	CreateCloudProfile(ctx context.Context, whereAccount AccountWhereUniqueInput, data CreateCloudProfileInput, interceptors ...clientv2.RequestInterceptor) (*CreateCloudProfile, error)
	GetDataStorage(ctx context.Context, where DataStorageWhereUniqueInput, interceptors ...clientv2.RequestInterceptor) (*GetDataStorage, error)
	GetDataStorageContainer(ctx context.Context, where DataStorageContainerWhereUniqueInput, interceptors ...clientv2.RequestInterceptor) (*GetDataStorageContainer, error)
	GetProjects(ctx context.Context, whereAccount AccountWhereUniqueInput, whereProject *ProjectWhereInput, interceptors ...clientv2.RequestInterceptor) (*GetProjects, error)
	CreateProject(ctx context.Context, whereAccount AccountWhereUniqueInput, data CreateProjectInput, interceptors ...clientv2.RequestInterceptor) (*CreateProject, error)
	GetAccounts(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*GetAccounts, error)
}

type Client struct {
	Client *clientv2.Client
}

func NewClient(cli *http.Client, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) GQLClient {
	return &Client{Client: clientv2.NewClient(cli, baseURL, options, interceptors...)}
}

type Query struct {
	AWSAppConfig                         *AWSAppConfig                         "json:\"aWSAppConfig,omitempty\" graphql:\"aWSAppConfig\""
	AWSContainerRegistryConfig           *AWSContainerRegistryConfig           "json:\"aWSContainerRegistryConfig,omitempty\" graphql:\"aWSContainerRegistryConfig\""
	AWSFalconConfig                      *AWSFalconConfig                      "json:\"aWSFalconConfig,omitempty\" graphql:\"aWSFalconConfig\""
	AWSNebulonConfig                     *AWSNebulonConfig                     "json:\"aWSNebulonConfig,omitempty\" graphql:\"aWSNebulonConfig\""
	AWSYodaConfig                        *AWSYodaConfig                        "json:\"aWSYodaConfig,omitempty\" graphql:\"aWSYodaConfig\""
	Application                          *Application                          "json:\"application,omitempty\" graphql:\"application\""
	ApplicationContainer                 *ApplicationContainer                 "json:\"applicationContainer,omitempty\" graphql:\"applicationContainer\""
	AzureAppConfig                       *AzureAppConfig                       "json:\"azureAppConfig,omitempty\" graphql:\"azureAppConfig\""
	AzureContainerRegistryConfig         *AzureContainerRegistryConfig         "json:\"azureContainerRegistryConfig,omitempty\" graphql:\"azureContainerRegistryConfig\""
	AzureFalconConfig                    *AzureFalconConfig                    "json:\"azureFalconConfig,omitempty\" graphql:\"azureFalconConfig\""
	AzureNebulonConfig                   *AzureNebulonConfig                   "json:\"azureNebulonConfig,omitempty\" graphql:\"azureNebulonConfig\""
	AzureYodaConfig                      *AzureYodaConfig                      "json:\"azureYodaConfig,omitempty\" graphql:\"azureYodaConfig\""
	CloudProfile                         *CloudProfile                         "json:\"cloudProfile,omitempty\" graphql:\"cloudProfile\""
	CloudProviderAppConfig               *CloudProviderAppConfig               "json:\"cloudProviderAppConfig,omitempty\" graphql:\"cloudProviderAppConfig\""
	CloudProviderContainerRegistryConfig *CloudProviderContainerRegistryConfig "json:\"cloudProviderContainerRegistryConfig,omitempty\" graphql:\"cloudProviderContainerRegistryConfig\""
	CloudProviderFalconConfig            *CloudProviderFalconConfig            "json:\"cloudProviderFalconConfig,omitempty\" graphql:\"cloudProviderFalconConfig\""
	CloudProviderNebulonConfig           *CloudProviderNebulonConfig           "json:\"cloudProviderNebulonConfig,omitempty\" graphql:\"cloudProviderNebulonConfig\""
	CloudProviderYodaConfig              *CloudProviderYodaConfig              "json:\"cloudProviderYodaConfig,omitempty\" graphql:\"cloudProviderYodaConfig\""
	ContainerRegistry                    *ContainerRegistry                    "json:\"containerRegistry,omitempty\" graphql:\"containerRegistry\""
	DataStorage                          *DataStorage                          "json:\"dataStorage,omitempty\" graphql:\"dataStorage\""
	DataStorageContainer                 *DataStorageContainer                 "json:\"dataStorageContainer,omitempty\" graphql:\"dataStorageContainer\""
	EnvironmentVariable                  *EnvironmentVariable                  "json:\"environmentVariable,omitempty\" graphql:\"environmentVariable\""
	Experiment                           *Experiment                           "json:\"experiment,omitempty\" graphql:\"experiment\""
	ExperimentRun                        *ExperimentRun                        "json:\"experimentRun,omitempty\" graphql:\"experimentRun\""
	FalconMLConfigs                      []*FalconMLConfig                     "json:\"falconMLConfigs\" graphql:\"falconMLConfigs\""
	FalconMLConfig                       *FalconMLConfig                       "json:\"falconMLConfig,omitempty\" graphql:\"falconMLConfig\""
	GCPAppConfig                         *GCPAppConfig                         "json:\"gCPAppConfig,omitempty\" graphql:\"gCPAppConfig\""
	GCPContainerRegistryConfig           *GCPContainerRegistryConfig           "json:\"gCPContainerRegistryConfig,omitempty\" graphql:\"gCPContainerRegistryConfig\""
	GCPFalconConfig                      *GCPFalconConfig                      "json:\"gCPFalconConfig,omitempty\" graphql:\"gCPFalconConfig\""
	GCPNebulonConfig                     *GCPNebulonConfig                     "json:\"gCPNebulonConfig,omitempty\" graphql:\"gCPNebulonConfig\""
	GCPYodaConfig                        *GCPYodaConfig                        "json:\"gCPYodaConfig,omitempty\" graphql:\"gCPYodaConfig\""
	KubernetesProfile                    *KubernetesProfile                    "json:\"kubernetesProfile,omitempty\" graphql:\"kubernetesProfile\""
	PersonalAccessTokens                 []*PersonalAccessToken                "json:\"personalAccessTokens\" graphql:\"personalAccessTokens\""
	PersonalAccessToken                  *PersonalAccessToken                  "json:\"personalAccessToken,omitempty\" graphql:\"personalAccessToken\""
	AggregateProject                     AggregateProject                      "json:\"aggregateProject\" graphql:\"aggregateProject\""
	Project                              *Project                              "json:\"project,omitempty\" graphql:\"project\""
	RayCluster                           *RayCluster                           "json:\"rayCluster,omitempty\" graphql:\"rayCluster\""
	RayClusterConfig                     *RayClusterConfig                     "json:\"rayClusterConfig,omitempty\" graphql:\"rayClusterConfig\""
	ResourceAccessToken                  *ResourceAccessToken                  "json:\"resourceAccessToken,omitempty\" graphql:\"resourceAccessToken\""
	TaskAction                           *TaskAction                           "json:\"taskAction,omitempty\" graphql:\"taskAction\""
	TaskStep                             *TaskStep                             "json:\"taskStep,omitempty\" graphql:\"taskStep\""
	Training                             *Training                             "json:\"training,omitempty\" graphql:\"training\""
	TrainingAnsibleTask                  *TrainingAnsibleTask                  "json:\"trainingAnsibleTask,omitempty\" graphql:\"trainingAnsibleTask\""
	VPc                                  *Vpc                                  "json:\"vPC,omitempty\" graphql:\"vPC\""
	VPCAttachment                        *VPCAttachment                        "json:\"vPCAttachment,omitempty\" graphql:\"vPCAttachment\""
	Account                              Account                               "json:\"account\" graphql:\"account\""
	ResourceTypes                        []ResourceType                        "json:\"resourceTypes\" graphql:\"resourceTypes\""
	BillingAccount                       *BillingAccount                       "json:\"billingAccount,omitempty\" graphql:\"billingAccount\""
	CloudProfiles                        []*CloudProfile                       "json:\"cloudProfiles\" graphql:\"cloudProfiles\""
	KubernetesProfiles                   []*KubernetesProfile                  "json:\"kubernetesProfiles\" graphql:\"kubernetesProfiles\""
	Team                                 Team                                  "json:\"team\" graphql:\"team\""
	TeamMemberInvitation                 TeamMemberInvitation                  "json:\"teamMemberInvitation\" graphql:\"teamMemberInvitation\""
	Me                                   User                                  "json:\"me\" graphql:\"me\""
	Users                                []*User                               "json:\"users\" graphql:\"users\""
	Vpcs                                 []*Vpc                                "json:\"vpcs\" graphql:\"vpcs\""
	Trainings                            []*Training                           "json:\"trainings\" graphql:\"trainings\""
	CanCreateTraining                    bool                                  "json:\"canCreateTraining\" graphql:\"canCreateTraining\""
	CanCreateDataStorage                 bool                                  "json:\"canCreateDataStorage\" graphql:\"canCreateDataStorage\""
	DataStorages                         []*DataStorage                        "json:\"dataStorages\" graphql:\"dataStorages\""
	Projects                             []*Project                            "json:\"projects\" graphql:\"projects\""
	ProjectByName                        Project                               "json:\"projectByName\" graphql:\"projectByName\""
	CanCreateApplication                 bool                                  "json:\"canCreateApplication\" graphql:\"canCreateApplication\""
	Applications                         []*Application                        "json:\"applications\" graphql:\"applications\""
	ApplicationInfrastructurePlans       []*ApplicationInfrastructurePlan      "json:\"applicationInfrastructurePlans\" graphql:\"applicationInfrastructurePlans\""
	TrainingInfrastructurePlans          []*TrainingInfrastructurePlan         "json:\"trainingInfrastructurePlans\" graphql:\"trainingInfrastructurePlans\""
	DataStorageContainerBrowser          *DataStorageContainerBrowser          "json:\"dataStorageContainerBrowser,omitempty\" graphql:\"dataStorageContainerBrowser\""
	MlflowExperiment                     *MLFlowExperiment                     "json:\"mlflowExperiment,omitempty\" graphql:\"mlflowExperiment\""
	CanCreateExperiment                  bool                                  "json:\"canCreateExperiment\" graphql:\"canCreateExperiment\""
	Experiments                          []*Experiment                         "json:\"experiments\" graphql:\"experiments\""
	ExperimentRuns                       []*ExperimentRun                      "json:\"experimentRuns\" graphql:\"experimentRuns\""
	RayClusters                          []*RayCluster                         "json:\"rayClusters\" graphql:\"rayClusters\""
	CanCreateRayCluster                  bool                                  "json:\"canCreateRayCluster\" graphql:\"canCreateRayCluster\""
	ContainerRegistries                  []*ContainerRegistry                  "json:\"containerRegistries\" graphql:\"containerRegistries\""
}
type Mutation struct {
	DeleteOnePersonalAccessToken    *PersonalAccessToken      "json:\"deleteOnePersonalAccessToken,omitempty\" graphql:\"deleteOnePersonalAccessToken\""
	UpdateAccount                   Account                   "json:\"updateAccount\" graphql:\"updateAccount\""
	SetGithubAppInstallationState   string                    "json:\"setGithubAppInstallationState\" graphql:\"setGithubAppInstallationState\""
	UpdateOnboarding                Onboarding                "json:\"updateOnboarding\" graphql:\"updateOnboarding\""
	AddBillingAccountDetails        BillingAccount            "json:\"addBillingAccountDetails\" graphql:\"addBillingAccountDetails\""
	UpdateBillingAccount            BillingAccount            "json:\"updateBillingAccount\" graphql:\"updateBillingAccount\""
	CreateBillingSubscription       BillingSubscription       "json:\"createBillingSubscription\" graphql:\"createBillingSubscription\""
	CancelBillingSubscription       BillingSubscription       "json:\"cancelBillingSubscription\" graphql:\"cancelBillingSubscription\""
	CreateCloudProfile              CloudProfile              "json:\"createCloudProfile\" graphql:\"createCloudProfile\""
	UpdateCloudProfile              CloudProfile              "json:\"updateCloudProfile\" graphql:\"updateCloudProfile\""
	UpdateCloudCredentials          CloudProfile              "json:\"updateCloudCredentials\" graphql:\"updateCloudCredentials\""
	DeleteCloudProfile              CloudProfile              "json:\"deleteCloudProfile\" graphql:\"deleteCloudProfile\""
	CreateKubernetesProfile         KubernetesProfile         "json:\"createKubernetesProfile\" graphql:\"createKubernetesProfile\""
	DeleteKubernetesProfile         KubernetesProfile         "json:\"deleteKubernetesProfile\" graphql:\"deleteKubernetesProfile\""
	CreateTeam                      Team                      "json:\"createTeam\" graphql:\"createTeam\""
	InviteUsersToTeam               []*TeamMemberInvitation   "json:\"inviteUsersToTeam\" graphql:\"inviteUsersToTeam\""
	AcceptTeamMemberInvitation      TeamMemberInvitation      "json:\"acceptTeamMemberInvitation\" graphql:\"acceptTeamMemberInvitation\""
	MatchTeamMemberInvitations      []*TeamMemberInvitation   "json:\"matchTeamMemberInvitations\" graphql:\"matchTeamMemberInvitations\""
	UpdateTeam                      Team                      "json:\"updateTeam\" graphql:\"updateTeam\""
	RemoveUserFromTeam              Team                      "json:\"removeUserFromTeam\" graphql:\"removeUserFromTeam\""
	UpdateUser                      User                      "json:\"updateUser\" graphql:\"updateUser\""
	CreatePersonalAccessToken       CreatePersonalAccessToken "json:\"createPersonalAccessToken\" graphql:\"createPersonalAccessToken\""
	CreateVpc                       Vpc                       "json:\"createVPC\" graphql:\"createVPC\""
	DeleteVpc                       Vpc                       "json:\"deleteVPC\" graphql:\"deleteVPC\""
	CreateTraining                  Training                  "json:\"createTraining\" graphql:\"createTraining\""
	DeleteTraining                  Training                  "json:\"deleteTraining\" graphql:\"deleteTraining\""
	StartTraining                   Training                  "json:\"startTraining\" graphql:\"startTraining\""
	StopTraining                    Training                  "json:\"stopTraining\" graphql:\"stopTraining\""
	MountDataStorageOnTraining      Training                  "json:\"mountDataStorageOnTraining\" graphql:\"mountDataStorageOnTraining\""
	UnmountDataStorageOnTraining    Training                  "json:\"unmountDataStorageOnTraining\" graphql:\"unmountDataStorageOnTraining\""
	CreateDataStorage               DataStorage               "json:\"createDataStorage\" graphql:\"createDataStorage\""
	DeleteDataStorage               DataStorage               "json:\"deleteDataStorage\" graphql:\"deleteDataStorage\""
	CreateProject                   Project                   "json:\"createProject\" graphql:\"createProject\""
	UpdateProject                   Project                   "json:\"updateProject\" graphql:\"updateProject\""
	DeleteProject                   Project                   "json:\"deleteProject\" graphql:\"deleteProject\""
	ConnectCloudProfile             Project                   "json:\"connectCloudProfile\" graphql:\"connectCloudProfile\""
	ConnectProjectRepository        Project                   "json:\"connectProjectRepository\" graphql:\"connectProjectRepository\""
	DisconnectProjectRepository     Project                   "json:\"disconnectProjectRepository\" graphql:\"disconnectProjectRepository\""
	DeleteApplication               Application               "json:\"deleteApplication\" graphql:\"deleteApplication\""
	CreateApplication               Application               "json:\"createApplication\" graphql:\"createApplication\""
	UpdateApplication               Application               "json:\"updateApplication\" graphql:\"updateApplication\""
	ForceNewDeploymentOnApplication Application               "json:\"forceNewDeploymentOnApplication\" graphql:\"forceNewDeploymentOnApplication\""
	CreateExperiment                Experiment                "json:\"createExperiment\" graphql:\"createExperiment\""
	DeleteExperiment                Experiment                "json:\"deleteExperiment\" graphql:\"deleteExperiment\""
	StartExperimentRun              ExperimentRun             "json:\"startExperimentRun\" graphql:\"startExperimentRun\""
	UpdateExperimentRun             ExperimentRun             "json:\"updateExperimentRun\" graphql:\"updateExperimentRun\""
	CreateRayCluster                RayCluster                "json:\"createRayCluster\" graphql:\"createRayCluster\""
	DeleteRayCluster                RayCluster                "json:\"deleteRayCluster\" graphql:\"deleteRayCluster\""
	CreateContainerRegistry         ContainerRegistry         "json:\"createContainerRegistry\" graphql:\"createContainerRegistry\""
	DeleteContainerRegistry         ContainerRegistry         "json:\"deleteContainerRegistry\" graphql:\"deleteContainerRegistry\""
	AddGithubAppInstallation        *GithubAppInstallation    "json:\"addGithubAppInstallation,omitempty\" graphql:\"addGithubAppInstallation\""
	RemoveGithubAppInstallation     bool                      "json:\"removeGithubAppInstallation\" graphql:\"removeGithubAppInstallation\""
}
type AWSYodaConfigFragment struct {
	ID                 string  "json:\"id\" graphql:\"id\""
	AwsAccessKey       *string "json:\"awsAccessKey,omitempty\" graphql:\"awsAccessKey\""
	AwsSecretAccessKey *string "json:\"awsSecretAccessKey,omitempty\" graphql:\"awsSecretAccessKey\""
	AwsRegion          string  "json:\"awsRegion\" graphql:\"awsRegion\""
}

func (t *AWSYodaConfigFragment) GetID() string {
	if t == nil {
		t = &AWSYodaConfigFragment{}
	}
	return t.ID
}
func (t *AWSYodaConfigFragment) GetAwsAccessKey() *string {
	if t == nil {
		t = &AWSYodaConfigFragment{}
	}
	return t.AwsAccessKey
}
func (t *AWSYodaConfigFragment) GetAwsSecretAccessKey() *string {
	if t == nil {
		t = &AWSYodaConfigFragment{}
	}
	return t.AwsSecretAccessKey
}
func (t *AWSYodaConfigFragment) GetAwsRegion() string {
	if t == nil {
		t = &AWSYodaConfigFragment{}
	}
	return t.AwsRegion
}

type AzureYodaConfigFragment struct {
	ID               string  "json:\"id\" graphql:\"id\""
	StorageAccount   *string "json:\"storageAccount,omitempty\" graphql:\"storageAccount\""
	StorageAccessKey *string "json:\"storageAccessKey,omitempty\" graphql:\"storageAccessKey\""
}

func (t *AzureYodaConfigFragment) GetID() string {
	if t == nil {
		t = &AzureYodaConfigFragment{}
	}
	return t.ID
}
func (t *AzureYodaConfigFragment) GetStorageAccount() *string {
	if t == nil {
		t = &AzureYodaConfigFragment{}
	}
	return t.StorageAccount
}
func (t *AzureYodaConfigFragment) GetStorageAccessKey() *string {
	if t == nil {
		t = &AzureYodaConfigFragment{}
	}
	return t.StorageAccessKey
}

type GCPYodaConfigFragment struct {
	ID                   string  "json:\"id\" graphql:\"id\""
	GcpServiceAccountKey *string "json:\"gcpServiceAccountKey,omitempty\" graphql:\"gcpServiceAccountKey\""
}

func (t *GCPYodaConfigFragment) GetID() string {
	if t == nil {
		t = &GCPYodaConfigFragment{}
	}
	return t.ID
}
func (t *GCPYodaConfigFragment) GetGcpServiceAccountKey() *string {
	if t == nil {
		t = &GCPYodaConfigFragment{}
	}
	return t.GcpServiceAccountKey
}

type CloudProviderYodaConfigFragment struct {
	ID          string                   "json:\"id\" graphql:\"id\""
	AwsConfig   *AWSYodaConfigFragment   "json:\"awsConfig,omitempty\" graphql:\"awsConfig\""
	AzureConfig *AzureYodaConfigFragment "json:\"azureConfig,omitempty\" graphql:\"azureConfig\""
	GcpConfig   *GCPYodaConfigFragment   "json:\"gcpConfig,omitempty\" graphql:\"gcpConfig\""
}

func (t *CloudProviderYodaConfigFragment) GetID() string {
	if t == nil {
		t = &CloudProviderYodaConfigFragment{}
	}
	return t.ID
}
func (t *CloudProviderYodaConfigFragment) GetAwsConfig() *AWSYodaConfigFragment {
	if t == nil {
		t = &CloudProviderYodaConfigFragment{}
	}
	return t.AwsConfig
}
func (t *CloudProviderYodaConfigFragment) GetAzureConfig() *AzureYodaConfigFragment {
	if t == nil {
		t = &CloudProviderYodaConfigFragment{}
	}
	return t.AzureConfig
}
func (t *CloudProviderYodaConfigFragment) GetGcpConfig() *GCPYodaConfigFragment {
	if t == nil {
		t = &CloudProviderYodaConfigFragment{}
	}
	return t.GcpConfig
}

type CloudProfileFragment struct {
	ID       string        "json:\"id\" graphql:\"id\""
	Name     string        "json:\"name\" graphql:\"name\""
	Provider CloudProvider "json:\"provider\" graphql:\"provider\""
}

func (t *CloudProfileFragment) GetID() string {
	if t == nil {
		t = &CloudProfileFragment{}
	}
	return t.ID
}
func (t *CloudProfileFragment) GetName() string {
	if t == nil {
		t = &CloudProfileFragment{}
	}
	return t.Name
}
func (t *CloudProfileFragment) GetProvider() *CloudProvider {
	if t == nil {
		t = &CloudProfileFragment{}
	}
	return &t.Provider
}

type DataStorageFragment struct {
	ID                      string                           "json:\"id\" graphql:\"id\""
	Name                    string                           "json:\"name\" graphql:\"name\""
	CloudProfile            *CloudProfileFragment            "json:\"cloudProfile,omitempty\" graphql:\"cloudProfile\""
	CloudProviderYodaConfig *CloudProviderYodaConfigFragment "json:\"cloudProviderYodaConfig,omitempty\" graphql:\"cloudProviderYodaConfig\""
}

func (t *DataStorageFragment) GetID() string {
	if t == nil {
		t = &DataStorageFragment{}
	}
	return t.ID
}
func (t *DataStorageFragment) GetName() string {
	if t == nil {
		t = &DataStorageFragment{}
	}
	return t.Name
}
func (t *DataStorageFragment) GetCloudProfile() *CloudProfileFragment {
	if t == nil {
		t = &DataStorageFragment{}
	}
	return t.CloudProfile
}
func (t *DataStorageFragment) GetCloudProviderYodaConfig() *CloudProviderYodaConfigFragment {
	if t == nil {
		t = &DataStorageFragment{}
	}
	return t.CloudProviderYodaConfig
}

type DataStorageContainerFragment struct {
	ID            string  "json:\"id\" graphql:\"id\""
	DirectoryName string  "json:\"directoryName\" graphql:\"directoryName\""
	CloudName     *string "json:\"cloudName,omitempty\" graphql:\"cloudName\""
	DataStorageID string  "json:\"dataStorageId\" graphql:\"dataStorageId\""
}

func (t *DataStorageContainerFragment) GetID() string {
	if t == nil {
		t = &DataStorageContainerFragment{}
	}
	return t.ID
}
func (t *DataStorageContainerFragment) GetDirectoryName() string {
	if t == nil {
		t = &DataStorageContainerFragment{}
	}
	return t.DirectoryName
}
func (t *DataStorageContainerFragment) GetCloudName() *string {
	if t == nil {
		t = &DataStorageContainerFragment{}
	}
	return t.CloudName
}
func (t *DataStorageContainerFragment) GetDataStorageID() string {
	if t == nil {
		t = &DataStorageContainerFragment{}
	}
	return t.DataStorageID
}

type ProjectFragment struct {
	ID             string    "json:\"id\" graphql:\"id\""
	Name           string    "json:\"name\" graphql:\"name\""
	CreatedAt      time.Time "json:\"createdAt\" graphql:\"createdAt\""
	CloudProfileID *string   "json:\"cloudProfileId,omitempty\" graphql:\"cloudProfileId\""
}

func (t *ProjectFragment) GetID() string {
	if t == nil {
		t = &ProjectFragment{}
	}
	return t.ID
}
func (t *ProjectFragment) GetName() string {
	if t == nil {
		t = &ProjectFragment{}
	}
	return t.Name
}
func (t *ProjectFragment) GetCreatedAt() *time.Time {
	if t == nil {
		t = &ProjectFragment{}
	}
	return &t.CreatedAt
}
func (t *ProjectFragment) GetCloudProfileID() *string {
	if t == nil {
		t = &ProjectFragment{}
	}
	return t.CloudProfileID
}

type AccountFragment struct {
	ID       string  "json:\"id\" graphql:\"id\""
	Username string  "json:\"username\" graphql:\"username\""
	Email    string  "json:\"email\" graphql:\"email\""
	Picture  *string "json:\"picture,omitempty\" graphql:\"picture\""
}

func (t *AccountFragment) GetID() string {
	if t == nil {
		t = &AccountFragment{}
	}
	return t.ID
}
func (t *AccountFragment) GetUsername() string {
	if t == nil {
		t = &AccountFragment{}
	}
	return t.Username
}
func (t *AccountFragment) GetEmail() string {
	if t == nil {
		t = &AccountFragment{}
	}
	return t.Email
}
func (t *AccountFragment) GetPicture() *string {
	if t == nil {
		t = &AccountFragment{}
	}
	return t.Picture
}

type GetAccounts_Me_Teams struct {
	Account *AccountFragment "json:\"account\" graphql:\"account\""
}

func (t *GetAccounts_Me_Teams) GetAccount() *AccountFragment {
	if t == nil {
		t = &GetAccounts_Me_Teams{}
	}
	return t.Account
}

type GetAccounts_Me struct {
	Account *AccountFragment        "json:\"account\" graphql:\"account\""
	Teams   []*GetAccounts_Me_Teams "json:\"teams\" graphql:\"teams\""
}

func (t *GetAccounts_Me) GetAccount() *AccountFragment {
	if t == nil {
		t = &GetAccounts_Me{}
	}
	return t.Account
}
func (t *GetAccounts_Me) GetTeams() []*GetAccounts_Me_Teams {
	if t == nil {
		t = &GetAccounts_Me{}
	}
	return t.Teams
}

type GetCloudProfile struct {
	CloudProfile *CloudProfileFragment "json:\"cloudProfile,omitempty\" graphql:\"cloudProfile\""
}

func (t *GetCloudProfile) GetCloudProfile() *CloudProfileFragment {
	if t == nil {
		t = &GetCloudProfile{}
	}
	return t.CloudProfile
}

type GetCloudProfiles struct {
	CloudProfiles []*CloudProfileFragment "json:\"cloudProfiles\" graphql:\"cloudProfiles\""
}

func (t *GetCloudProfiles) GetCloudProfiles() []*CloudProfileFragment {
	if t == nil {
		t = &GetCloudProfiles{}
	}
	return t.CloudProfiles
}

type CreateCloudProfile struct {
	CreateCloudProfile *CloudProfileFragment "json:\"createCloudProfile\" graphql:\"createCloudProfile\""
}

func (t *CreateCloudProfile) GetCreateCloudProfile() *CloudProfileFragment {
	if t == nil {
		t = &CreateCloudProfile{}
	}
	return t.CreateCloudProfile
}

type GetDataStorage struct {
	DataStorage *DataStorageFragment "json:\"dataStorage,omitempty\" graphql:\"dataStorage\""
}

func (t *GetDataStorage) GetDataStorage() *DataStorageFragment {
	if t == nil {
		t = &GetDataStorage{}
	}
	return t.DataStorage
}

type GetDataStorageContainer struct {
	DataStorageContainer *DataStorageContainerFragment "json:\"dataStorageContainer,omitempty\" graphql:\"dataStorageContainer\""
}

func (t *GetDataStorageContainer) GetDataStorageContainer() *DataStorageContainerFragment {
	if t == nil {
		t = &GetDataStorageContainer{}
	}
	return t.DataStorageContainer
}

type GetProjects struct {
	Projects []*ProjectFragment "json:\"projects\" graphql:\"projects\""
}

func (t *GetProjects) GetProjects() []*ProjectFragment {
	if t == nil {
		t = &GetProjects{}
	}
	return t.Projects
}

type CreateProject struct {
	CreateProject *ProjectFragment "json:\"createProject\" graphql:\"createProject\""
}

func (t *CreateProject) GetCreateProject() *ProjectFragment {
	if t == nil {
		t = &CreateProject{}
	}
	return t.CreateProject
}

type GetAccounts struct {
	Me GetAccounts_Me "json:\"me\" graphql:\"me\""
}

func (t *GetAccounts) GetMe() *GetAccounts_Me {
	if t == nil {
		t = &GetAccounts{}
	}
	return &t.Me
}

const GetCloudProfileDocument = `query GetCloudProfile ($where: CloudProfileWhereUniqueInput!) {
	cloudProfile(where: $where) {
		... CloudProfileFragment
	}
}
fragment CloudProfileFragment on CloudProfile {
	id
	name
	provider
}
`

func (c *Client) GetCloudProfile(ctx context.Context, where CloudProfileWhereUniqueInput, interceptors ...clientv2.RequestInterceptor) (*GetCloudProfile, error) {
	vars := map[string]interface{}{
		"where": where,
	}

	var res GetCloudProfile
	if err := c.Client.Post(ctx, "GetCloudProfile", GetCloudProfileDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetCloudProfilesDocument = `query GetCloudProfiles ($whereAccount: AccountWhereUniqueInput!, $whereCloudProfile: CloudProfileWhereInput) {
	cloudProfiles(whereAccount: $whereAccount, whereCloudProfile: $whereCloudProfile) {
		... CloudProfileFragment
	}
}
fragment CloudProfileFragment on CloudProfile {
	id
	name
	provider
}
`

func (c *Client) GetCloudProfiles(ctx context.Context, whereAccount AccountWhereUniqueInput, whereCloudProfile *CloudProfileWhereInput, interceptors ...clientv2.RequestInterceptor) (*GetCloudProfiles, error) {
	vars := map[string]interface{}{
		"whereAccount":      whereAccount,
		"whereCloudProfile": whereCloudProfile,
	}

	var res GetCloudProfiles
	if err := c.Client.Post(ctx, "GetCloudProfiles", GetCloudProfilesDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const CreateCloudProfileDocument = `mutation CreateCloudProfile ($whereAccount: AccountWhereUniqueInput!, $data: CreateCloudProfileInput!) {
	createCloudProfile(whereAccount: $whereAccount, data: $data) {
		... CloudProfileFragment
	}
}
fragment CloudProfileFragment on CloudProfile {
	id
	name
	provider
}
`

func (c *Client) CreateCloudProfile(ctx context.Context, whereAccount AccountWhereUniqueInput, data CreateCloudProfileInput, interceptors ...clientv2.RequestInterceptor) (*CreateCloudProfile, error) {
	vars := map[string]interface{}{
		"whereAccount": whereAccount,
		"data":         data,
	}

	var res CreateCloudProfile
	if err := c.Client.Post(ctx, "CreateCloudProfile", CreateCloudProfileDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetDataStorageDocument = `query GetDataStorage ($where: DataStorageWhereUniqueInput!) {
	dataStorage(where: $where) {
		... DataStorageFragment
	}
}
fragment DataStorageFragment on DataStorage {
	id
	name
	cloudProfile {
		... CloudProfileFragment
	}
	cloudProviderYodaConfig {
		... CloudProviderYodaConfigFragment
	}
}
fragment CloudProfileFragment on CloudProfile {
	id
	name
	provider
}
fragment CloudProviderYodaConfigFragment on CloudProviderYodaConfig {
	id
	awsConfig {
		... AWSYodaConfigFragment
	}
	azureConfig {
		... AzureYodaConfigFragment
	}
	gcpConfig {
		... GCPYodaConfigFragment
	}
}
fragment AWSYodaConfigFragment on AWSYodaConfig {
	id
	awsAccessKey
	awsSecretAccessKey
	awsRegion
}
fragment AzureYodaConfigFragment on AzureYodaConfig {
	id
	storageAccount
	storageAccessKey
}
fragment GCPYodaConfigFragment on GCPYodaConfig {
	id
	gcpServiceAccountKey
}
`

func (c *Client) GetDataStorage(ctx context.Context, where DataStorageWhereUniqueInput, interceptors ...clientv2.RequestInterceptor) (*GetDataStorage, error) {
	vars := map[string]interface{}{
		"where": where,
	}

	var res GetDataStorage
	if err := c.Client.Post(ctx, "GetDataStorage", GetDataStorageDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetDataStorageContainerDocument = `query GetDataStorageContainer ($where: DataStorageContainerWhereUniqueInput!) {
	dataStorageContainer(where: $where) {
		... DataStorageContainerFragment
	}
}
fragment DataStorageContainerFragment on DataStorageContainer {
	id
	directoryName
	cloudName
	dataStorageId
}
`

func (c *Client) GetDataStorageContainer(ctx context.Context, where DataStorageContainerWhereUniqueInput, interceptors ...clientv2.RequestInterceptor) (*GetDataStorageContainer, error) {
	vars := map[string]interface{}{
		"where": where,
	}

	var res GetDataStorageContainer
	if err := c.Client.Post(ctx, "GetDataStorageContainer", GetDataStorageContainerDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetProjectsDocument = `query GetProjects ($whereAccount: AccountWhereUniqueInput!, $whereProject: ProjectWhereInput) {
	projects(whereAccount: $whereAccount, whereProject: $whereProject) {
		... ProjectFragment
	}
}
fragment ProjectFragment on Project {
	id
	name
	createdAt
	cloudProfileId
}
`

func (c *Client) GetProjects(ctx context.Context, whereAccount AccountWhereUniqueInput, whereProject *ProjectWhereInput, interceptors ...clientv2.RequestInterceptor) (*GetProjects, error) {
	vars := map[string]interface{}{
		"whereAccount": whereAccount,
		"whereProject": whereProject,
	}

	var res GetProjects
	if err := c.Client.Post(ctx, "GetProjects", GetProjectsDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const CreateProjectDocument = `mutation CreateProject ($whereAccount: AccountWhereUniqueInput!, $data: CreateProjectInput!) {
	createProject(data: $data, whereAccount: $whereAccount) {
		... ProjectFragment
	}
}
fragment ProjectFragment on Project {
	id
	name
	createdAt
	cloudProfileId
}
`

func (c *Client) CreateProject(ctx context.Context, whereAccount AccountWhereUniqueInput, data CreateProjectInput, interceptors ...clientv2.RequestInterceptor) (*CreateProject, error) {
	vars := map[string]interface{}{
		"whereAccount": whereAccount,
		"data":         data,
	}

	var res CreateProject
	if err := c.Client.Post(ctx, "CreateProject", CreateProjectDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetAccountsDocument = `query GetAccounts {
	me {
		account {
			... AccountFragment
		}
		teams {
			account {
				... AccountFragment
			}
		}
	}
}
fragment AccountFragment on Account {
	id
	username
	email
	picture
}
`

func (c *Client) GetAccounts(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*GetAccounts, error) {
	vars := map[string]interface{}{}

	var res GetAccounts
	if err := c.Client.Post(ctx, "GetAccounts", GetAccountsDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}
