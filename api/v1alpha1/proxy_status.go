package v1alpha1

import (
	"fmt"
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
	ConditionProcessing ConditionType = "Processing"
	ConditionRunning    ConditionType = "Running"
	ConditionUnknow     ConditionType = "Unknow"
)

// ProxyStatus defines the observed state of Proxy
type ProxyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Phase          PhaseStatus `json:"phase"`
	Conditions     Conditions  `json:"conditions"`
	AvailableNodes string      `json:"availableNodes"`
	Version        string      `json:"version"`
}

type Conditions []Condition

type Condition struct {
	Type           ConditionType      `json:"type"`
	Status         v1.ConditionStatus `json:"status"`
	LastUpdateTime metav1.Time        `json:"lastUpdateTime,omitempty"`
}

func (p *Proxy) SetInitStatus() {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append(p.Status.Conditions, Condition{
		Type:           ConditionProcessing,
		Status:         v1.ConditionTrue,
		LastUpdateTime: metav1.Now(),
	})
	p.Status.AvailableNodes = fmt.Sprintf("0/%d", p.Spec.Replicas)
	p.Status.Version = p.Spec.Version
}

func (p *Proxy) SetInitFailed() {
	p.Status.Conditions = append(p.Status.Conditions, Condition{
		Type:           ConditionProcessing,
		Status:         v1.ConditionFalse,
		LastUpdateTime: metav1.Now(),
	})
}

func (p *Proxy) SetRunningButNotready(readyCount int32) {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append([]Condition{}, Condition{
		Type:           ConditionRunning,
		Status:         v1.ConditionFalse,
		LastUpdateTime: metav1.Now(),
	})
	p.Status.AvailableNodes = fmt.Sprintf("%d/%d", readyCount, p.Spec.Replicas)
	p.Status.Version = p.Spec.Version
}
