### What is JCM_K8S_CONTROLLER?

---

- jcm_k8s_controller is controller of kubernetes which use kubebuilder

- jcm_k8s_controller watch custrom resource which name is jcm.

- jcm will get resources(deployments, replicaSet, pods, configMap, sercret...) informations, and stored in store interface.(ex. kv store, or database)
  

### JCM_K8S_CONTROLLER environment
---

- cluster: kind (kubernetes version: 1.22.0)
- golang version: 1.16
- db: cockroach db
- kv store: etcd
