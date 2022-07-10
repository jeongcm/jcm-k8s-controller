### What is JCM_K8S_CONTROLLER?

---

- jcm_k8s_controller is controller of kubernetes which use kubebuilder

- jcm_k8s_controller watch custrom resource which name is jcm.

- jcm will get resources(deployments, replicaSet, pods, configMap, sercret...) informations which name start jcm(ex jcm_deployment),
  and stored in store interface.(ex. kv store, or database)
  
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
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Phase JcmPhase `json:"phase,omitempty"`

	Condition []JcmCondition `json:"conditions,omitempty"`

	Message string `json:"message,omitempty"`
	// +optional
	Reason string `json:"reason,omitempty"`
}

```
