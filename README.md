### What is JCM_K8S_CONTROLLER?

---

- jcm_k8s_controller is controller of kubernetes which use kubebuilder

- jcm_k8s_controller watch custrom resource which is name is jcm.

- jcm will create resources(deployments, replicaSet, pods, configMap, sercret...) informations which name start jcm(ex jcm_deployment_1),
  from configMap, annotation(from MutateWebhookConfiguration)
  
- soon, it will be monitored by jcm_monitor_service that use prometheus.
### JCM_K8S_CONTROLLER environment
---

- cluster: kind (kubernetes version: 1.22.0)
- golang version: 1.16
- db: cockroach db
- kv store: etcd


### CustomResource JCM Spec, Status

```go
// JcmSpec defines the desired state of Jcm
type JcmSpec struct {
	Prefix string `json:"prefix,omitempty"`
}

// JcmStatus defines the observed state of Jcm
type JcmStatus struct {
	Phase JcmPhase `json:"phase,omitempty"`

	Condition []JcmCondition `json:"conditions,omitempty"`

	Message string `json:"message,omitempty"`
	// +optional
	Reason string `json:"reason,omitempty"`
}

```
