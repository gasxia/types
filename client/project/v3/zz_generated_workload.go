package client

import (
	"github.com/rancher/norman/types"
)

const (
	WorkloadType                               = "workload"
	WorkloadFieldActiveDeadlineSeconds         = "activeDeadlineSeconds"
	WorkloadFieldAnnotations                   = "annotations"
	WorkloadFieldAutomountServiceAccountToken  = "automountServiceAccountToken"
	WorkloadFieldContainers                    = "containers"
	WorkloadFieldCreated                       = "created"
	WorkloadFieldCreatorID                     = "creatorId"
	WorkloadFieldCronJobConfig                 = "cronJobConfig"
	WorkloadFieldCronJobStatus                 = "cronJobStatus"
	WorkloadFieldDNSConfig                     = "dnsConfig"
	WorkloadFieldDNSPolicy                     = "dnsPolicy"
	WorkloadFieldDaemonSetConfig               = "daemonSetConfig"
	WorkloadFieldDaemonSetStatus               = "daemonSetStatus"
	WorkloadFieldDeploymentConfig              = "deploymentConfig"
	WorkloadFieldDeploymentStatus              = "deploymentStatus"
	WorkloadFieldEnableServiceLinks            = "enableServiceLinks"
	WorkloadFieldEphemeralContainers           = "ephemeralContainers"
	WorkloadFieldFsgid                         = "fsgid"
	WorkloadFieldGids                          = "gids"
	WorkloadFieldHostAliases                   = "hostAliases"
	WorkloadFieldHostIPC                       = "hostIPC"
	WorkloadFieldHostNetwork                   = "hostNetwork"
	WorkloadFieldHostPID                       = "hostPID"
	WorkloadFieldHostname                      = "hostname"
	WorkloadFieldImagePullSecrets              = "imagePullSecrets"
	WorkloadFieldJobConfig                     = "jobConfig"
	WorkloadFieldJobStatus                     = "jobStatus"
	WorkloadFieldLabels                        = "labels"
	WorkloadFieldName                          = "name"
	WorkloadFieldNamespaceId                   = "namespaceId"
	WorkloadFieldNodeID                        = "nodeId"
	WorkloadFieldOverhead                      = "overhead"
	WorkloadFieldOwnerReferences               = "ownerReferences"
	WorkloadFieldPaused                        = "paused"
	WorkloadFieldPreemptionPolicy              = "preemptionPolicy"
	WorkloadFieldProjectID                     = "projectId"
	WorkloadFieldPublicEndpoints               = "publicEndpoints"
	WorkloadFieldReadinessGates                = "readinessGates"
	WorkloadFieldRemoved                       = "removed"
	WorkloadFieldReplicaSetConfig              = "replicaSetConfig"
	WorkloadFieldReplicaSetStatus              = "replicaSetStatus"
	WorkloadFieldReplicationControllerConfig   = "replicationControllerConfig"
	WorkloadFieldReplicationControllerStatus   = "replicationControllerStatus"
	WorkloadFieldRestartPolicy                 = "restartPolicy"
	WorkloadFieldRunAsGroup                    = "runAsGroup"
	WorkloadFieldRunAsNonRoot                  = "runAsNonRoot"
	WorkloadFieldRuntimeClassName              = "runtimeClassName"
	WorkloadFieldScale                         = "scale"
	WorkloadFieldScheduling                    = "scheduling"
	WorkloadFieldSelector                      = "selector"
	WorkloadFieldServiceAccountName            = "serviceAccountName"
	WorkloadFieldShareProcessNamespace         = "shareProcessNamespace"
	WorkloadFieldState                         = "state"
	WorkloadFieldStatefulSetConfig             = "statefulSetConfig"
	WorkloadFieldStatefulSetStatus             = "statefulSetStatus"
	WorkloadFieldSubdomain                     = "subdomain"
	WorkloadFieldSysctls                       = "sysctls"
	WorkloadFieldTTLSecondsAfterFinished       = "ttlSecondsAfterFinished"
	WorkloadFieldTerminationGracePeriodSeconds = "terminationGracePeriodSeconds"
	WorkloadFieldTopologySpreadConstraints     = "topologySpreadConstraints"
	WorkloadFieldTransitioning                 = "transitioning"
	WorkloadFieldTransitioningMessage          = "transitioningMessage"
	WorkloadFieldUUID                          = "uuid"
	WorkloadFieldUid                           = "uid"
	WorkloadFieldVolumes                       = "volumes"
	WorkloadFieldWindowsOptions                = "windowsOptions"
	WorkloadFieldWorkloadAnnotations           = "workloadAnnotations"
	WorkloadFieldWorkloadLabels                = "workloadLabels"
	WorkloadFieldWorkloadMetrics               = "workloadMetrics"
)

type Workload struct {
	types.Resource
	ActiveDeadlineSeconds         *int64                         `json:"activeDeadlineSeconds,omitempty" yaml:"activeDeadlineSeconds,omitempty"`
	Annotations                   map[string]string              `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AutomountServiceAccountToken  *bool                          `json:"automountServiceAccountToken,omitempty" yaml:"automountServiceAccountToken,omitempty"`
	Containers                    []Container                    `json:"containers,omitempty" yaml:"containers,omitempty"`
	Created                       string                         `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                     string                         `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	CronJobConfig                 *CronJobConfig                 `json:"cronJobConfig,omitempty" yaml:"cronJobConfig,omitempty"`
	CronJobStatus                 *CronJobStatus                 `json:"cronJobStatus,omitempty" yaml:"cronJobStatus,omitempty"`
	DNSConfig                     *PodDNSConfig                  `json:"dnsConfig,omitempty" yaml:"dnsConfig,omitempty"`
	DNSPolicy                     string                         `json:"dnsPolicy,omitempty" yaml:"dnsPolicy,omitempty"`
	DaemonSetConfig               *DaemonSetConfig               `json:"daemonSetConfig,omitempty" yaml:"daemonSetConfig,omitempty"`
	DaemonSetStatus               *DaemonSetStatus               `json:"daemonSetStatus,omitempty" yaml:"daemonSetStatus,omitempty"`
	DeploymentConfig              *DeploymentConfig              `json:"deploymentConfig,omitempty" yaml:"deploymentConfig,omitempty"`
	DeploymentStatus              *DeploymentStatus              `json:"deploymentStatus,omitempty" yaml:"deploymentStatus,omitempty"`
	EnableServiceLinks            *bool                          `json:"enableServiceLinks,omitempty" yaml:"enableServiceLinks,omitempty"`
	EphemeralContainers           []EphemeralContainer           `json:"ephemeralContainers,omitempty" yaml:"ephemeralContainers,omitempty"`
	Fsgid                         *int64                         `json:"fsgid,omitempty" yaml:"fsgid,omitempty"`
	Gids                          []int64                        `json:"gids,omitempty" yaml:"gids,omitempty"`
	HostAliases                   []HostAlias                    `json:"hostAliases,omitempty" yaml:"hostAliases,omitempty"`
	HostIPC                       bool                           `json:"hostIPC,omitempty" yaml:"hostIPC,omitempty"`
	HostNetwork                   bool                           `json:"hostNetwork,omitempty" yaml:"hostNetwork,omitempty"`
	HostPID                       bool                           `json:"hostPID,omitempty" yaml:"hostPID,omitempty"`
	Hostname                      string                         `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	ImagePullSecrets              []LocalObjectReference         `json:"imagePullSecrets,omitempty" yaml:"imagePullSecrets,omitempty"`
	JobConfig                     *JobConfig                     `json:"jobConfig,omitempty" yaml:"jobConfig,omitempty"`
	JobStatus                     *JobStatus                     `json:"jobStatus,omitempty" yaml:"jobStatus,omitempty"`
	Labels                        map[string]string              `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                          string                         `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId                   string                         `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	NodeID                        string                         `json:"nodeId,omitempty" yaml:"nodeId,omitempty"`
	Overhead                      map[string]string              `json:"overhead,omitempty" yaml:"overhead,omitempty"`
	OwnerReferences               []OwnerReference               `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Paused                        bool                           `json:"paused,omitempty" yaml:"paused,omitempty"`
	PreemptionPolicy              string                         `json:"preemptionPolicy,omitempty" yaml:"preemptionPolicy,omitempty"`
	ProjectID                     string                         `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	PublicEndpoints               []PublicEndpoint               `json:"publicEndpoints,omitempty" yaml:"publicEndpoints,omitempty"`
	ReadinessGates                []PodReadinessGate             `json:"readinessGates,omitempty" yaml:"readinessGates,omitempty"`
	Removed                       string                         `json:"removed,omitempty" yaml:"removed,omitempty"`
	ReplicaSetConfig              *ReplicaSetConfig              `json:"replicaSetConfig,omitempty" yaml:"replicaSetConfig,omitempty"`
	ReplicaSetStatus              *ReplicaSetStatus              `json:"replicaSetStatus,omitempty" yaml:"replicaSetStatus,omitempty"`
	ReplicationControllerConfig   *ReplicationControllerConfig   `json:"replicationControllerConfig,omitempty" yaml:"replicationControllerConfig,omitempty"`
	ReplicationControllerStatus   *ReplicationControllerStatus   `json:"replicationControllerStatus,omitempty" yaml:"replicationControllerStatus,omitempty"`
	RestartPolicy                 string                         `json:"restartPolicy,omitempty" yaml:"restartPolicy,omitempty"`
	RunAsGroup                    *int64                         `json:"runAsGroup,omitempty" yaml:"runAsGroup,omitempty"`
	RunAsNonRoot                  *bool                          `json:"runAsNonRoot,omitempty" yaml:"runAsNonRoot,omitempty"`
	RuntimeClassName              string                         `json:"runtimeClassName,omitempty" yaml:"runtimeClassName,omitempty"`
	Scale                         *int64                         `json:"scale,omitempty" yaml:"scale,omitempty"`
	Scheduling                    *Scheduling                    `json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	Selector                      *LabelSelector                 `json:"selector,omitempty" yaml:"selector,omitempty"`
	ServiceAccountName            string                         `json:"serviceAccountName,omitempty" yaml:"serviceAccountName,omitempty"`
	ShareProcessNamespace         *bool                          `json:"shareProcessNamespace,omitempty" yaml:"shareProcessNamespace,omitempty"`
	State                         string                         `json:"state,omitempty" yaml:"state,omitempty"`
	StatefulSetConfig             *StatefulSetConfig             `json:"statefulSetConfig,omitempty" yaml:"statefulSetConfig,omitempty"`
	StatefulSetStatus             *StatefulSetStatus             `json:"statefulSetStatus,omitempty" yaml:"statefulSetStatus,omitempty"`
	Subdomain                     string                         `json:"subdomain,omitempty" yaml:"subdomain,omitempty"`
	Sysctls                       []Sysctl                       `json:"sysctls,omitempty" yaml:"sysctls,omitempty"`
	TTLSecondsAfterFinished       *int64                         `json:"ttlSecondsAfterFinished,omitempty" yaml:"ttlSecondsAfterFinished,omitempty"`
	TerminationGracePeriodSeconds *int64                         `json:"terminationGracePeriodSeconds,omitempty" yaml:"terminationGracePeriodSeconds,omitempty"`
	TopologySpreadConstraints     []TopologySpreadConstraint     `json:"topologySpreadConstraints,omitempty" yaml:"topologySpreadConstraints,omitempty"`
	Transitioning                 string                         `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage          string                         `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                          string                         `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Uid                           *int64                         `json:"uid,omitempty" yaml:"uid,omitempty"`
	Volumes                       []Volume                       `json:"volumes,omitempty" yaml:"volumes,omitempty"`
	WindowsOptions                *WindowsSecurityContextOptions `json:"windowsOptions,omitempty" yaml:"windowsOptions,omitempty"`
	WorkloadAnnotations           map[string]string              `json:"workloadAnnotations,omitempty" yaml:"workloadAnnotations,omitempty"`
	WorkloadLabels                map[string]string              `json:"workloadLabels,omitempty" yaml:"workloadLabels,omitempty"`
	WorkloadMetrics               []WorkloadMetric               `json:"workloadMetrics,omitempty" yaml:"workloadMetrics,omitempty"`
}

type WorkloadCollection struct {
	types.Collection
	Data   []Workload `json:"data,omitempty"`
	client *WorkloadClient
}

type WorkloadClient struct {
	apiClient *Client
}

type WorkloadOperations interface {
	List(opts *types.ListOpts) (*WorkloadCollection, error)
	Create(opts *Workload) (*Workload, error)
	Update(existing *Workload, updates interface{}) (*Workload, error)
	Replace(existing *Workload) (*Workload, error)
	ByID(id string) (*Workload, error)
	Delete(container *Workload) error

	ActionPause(resource *Workload) error

	ActionRedeploy(resource *Workload) error

	ActionResume(resource *Workload) error

	ActionRollback(resource *Workload, input *RollbackRevision) error
}

func newWorkloadClient(apiClient *Client) *WorkloadClient {
	return &WorkloadClient{
		apiClient: apiClient,
	}
}

func (c *WorkloadClient) Create(container *Workload) (*Workload, error) {
	resp := &Workload{}
	err := c.apiClient.Ops.DoCreate(WorkloadType, container, resp)
	return resp, err
}

func (c *WorkloadClient) Update(existing *Workload, updates interface{}) (*Workload, error) {
	resp := &Workload{}
	err := c.apiClient.Ops.DoUpdate(WorkloadType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *WorkloadClient) Replace(obj *Workload) (*Workload, error) {
	resp := &Workload{}
	err := c.apiClient.Ops.DoReplace(WorkloadType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *WorkloadClient) List(opts *types.ListOpts) (*WorkloadCollection, error) {
	resp := &WorkloadCollection{}
	err := c.apiClient.Ops.DoList(WorkloadType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *WorkloadCollection) Next() (*WorkloadCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &WorkloadCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *WorkloadClient) ByID(id string) (*Workload, error) {
	resp := &Workload{}
	err := c.apiClient.Ops.DoByID(WorkloadType, id, resp)
	return resp, err
}

func (c *WorkloadClient) Delete(container *Workload) error {
	return c.apiClient.Ops.DoResourceDelete(WorkloadType, &container.Resource)
}

func (c *WorkloadClient) ActionPause(resource *Workload) error {
	err := c.apiClient.Ops.DoAction(WorkloadType, "pause", &resource.Resource, nil, nil)
	return err
}

func (c *WorkloadClient) ActionRedeploy(resource *Workload) error {
	err := c.apiClient.Ops.DoAction(WorkloadType, "redeploy", &resource.Resource, nil, nil)
	return err
}

func (c *WorkloadClient) ActionResume(resource *Workload) error {
	err := c.apiClient.Ops.DoAction(WorkloadType, "resume", &resource.Resource, nil, nil)
	return err
}

func (c *WorkloadClient) ActionRollback(resource *Workload, input *RollbackRevision) error {
	err := c.apiClient.Ops.DoAction(WorkloadType, "rollback", &resource.Resource, input, nil)
	return err
}
