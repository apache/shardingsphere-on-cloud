package v1alpha1

// User TODO: description
type User struct {
	// +optional
	UserConfig string `json:"-" yaml:"user"`
	UserName   string `json:"userName" yaml:"-"`
	PassWord   string `json:"passWord" yaml:"-"`
	// +optional
	HostName string `json:"hostName,omitempty" yaml:"-"`
}

// Provider TODO: description
type Provider struct {
	Type string `json:"type" yaml:"type"`
}

// Auth TODO: description
type Auth struct {
	Users []User `json:"users" yaml:"users"`
	// +optional
	Provider Provider `json:"provider,omitempty"`
}

// Props TODO: description
type Props struct {
	// +optional
	KernelExecutorSize int `json:"kernel-executor-size,omitempty" yaml:"kernel-executor-size"`
	// +optional
	CheckTableMetadataEnabled bool `json:"check-table-metadata-enabled,omitempty" yaml:"check-table-metadata-enabled"`
	// +optional
	ProxyBackendQueryFetchSize int `json:"proxy-backend-query-fetch-size,omitempty" yaml:"proxy-backend-query-fetch-size"`
	// +optional
	CheckDuplicateTableEnabled bool `json:"check-duplicate-table-enabled,omitempty" yaml:"check-duplicate-table-enabled"`
	// +optional
	ProxyFrontendExecutorSize int `json:"proxy-frontend-executor-size,omitempty" yaml:"proxy-frontend-executor-size"`
	// +optional
	ProxyBackendExecutorSuitable string `json:"proxy-backend-executor-suitable,omitempty" yaml:"proxy-backend-executor-suitable"`
}

type ClusterProps struct {
	NameSpace   string `json:"namespace" yaml:"namespace"`
	ServerLists string `json:"server-lists" yaml:"server-lists"`
	// +optional
	RetryIntervalMilliseconds int `json:"retryIntervalMilliseconds,omitempty" yaml:"retryIntervalMilliseconds,omitempty"`
	// +optional
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`
	// +optional
	TimeToLiveSeconds int `json:"timeToLiveSeconds,omitempty" yaml:"timeToLiveSeconds,omitempty"`
	// +optional
	OperationTimeoutMilliseconds int `json:"operationTimeoutMilliseconds,omitempty" yaml:"operationTimeoutMilliseconds,omitempty"`
	// +optional
	Digest string `json:"digest,omitempty" yaml:"digest,omitempty"`
}

type RepositoryConfig struct {
	Type  string       `json:"type" yaml:"type"`
	Props ClusterProps `json:"props" yaml:"props"`
}

type ClusterConfig struct {
	Type       string           `json:"type" yaml:"type"`
	Repository RepositoryConfig `json:"repository" yaml:"repository"`
	Overwrite  bool             `json:"overwrite" yaml:"overwrite"`
}
