package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PhaseStatus string

const (
	StatusReady    PhaseStatus = "Ready"
	StatusNotReady PhaseStatus = "NotReady"
)

type ConditionType string

const (
	ConditionInitialized ConditionType = "Initializing"
	ConditionProcessing  ConditionType = "Processing"
	ConditionRunning     ConditionType = "Running"
	ConditionUnknown     ConditionType = "Unknown"
)

// ProxyStatus defines the observed state of Proxy
type ProxyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// TODO:description
	Phase PhaseStatus `json:"phase"`
	// TODO:description
	Conditions Conditions `json:"conditions"`
	// TODO:description
	ReadyNodes int32 `json:"readyNodes"`
	// TODO:description
	Version string `json:"version"`
}

type Conditions []Condition

type Condition struct {
	Type           ConditionType      `json:"type"`
	Status         v1.ConditionStatus `json:"status"`
	LastUpdateTime metav1.Time        `json:"lastUpdateTime,omitempty"`
}

func (p *Proxy) SetInitializedStatus() {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append(p.Status.Conditions, Condition{
		Type:           ConditionInitialized,
		Status:         v1.ConditionTrue,
		LastUpdateTime: metav1.Now(),
	})
	p.Status.ReadyNodes = 0
	p.Status.Version = p.Spec.Version
}

func (p *Proxy) SetInitializationFailed() {
	p.Status.Conditions = append(p.Status.Conditions, Condition{
		Type:           ConditionInitialized,
		Status:         v1.ConditionFalse,
		LastUpdateTime: metav1.Now(),
	})
}

func (p *Proxy) SetInitializationSuccess() {
	p.Status.Conditions = append(p.Status.Conditions, Condition{
		Type:           ConditionProcessing,
		Status:         v1.ConditionTrue,
		LastUpdateTime: metav1.Now(),
	})
}

func (p *Proxy) SetNotRunning() {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append([]Condition{}, Condition{
		Type:           ConditionRunning,
		Status:         v1.ConditionFalse,
		LastUpdateTime: metav1.Now(),
	})
	p.Status.ReadyNodes = 0
	p.Status.Version = p.Spec.Version
}

func (p *Proxy) SetRunningButNotReady(readyNodes int32) {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append(p.Status.Conditions, Condition{
		Type:           ConditionRunning,
		Status:         v1.ConditionTrue,
		LastUpdateTime: metav1.Now(),
	})
	p.Status.ReadyNodes = readyNodes
	p.Status.Version = p.Spec.Version
}

func (p *Proxy) SetReady() {
	p.Status.Phase = StatusReady
	p.Status.Conditions = append(p.Status.Conditions, Condition{
		Type:           ConditionRunning,
		Status:         v1.ConditionTrue,
		LastUpdateTime: metav1.Now(),
	})
	p.Status.ReadyNodes = p.Spec.Replicas
	p.Status.Version = p.Spec.Version
}
