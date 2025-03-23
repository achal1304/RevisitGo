
# Kubernetes YAML Configuration Files and Components

## 1. Pod YAML

A **Pod** is the smallest and simplest unit in Kubernetes. It represents a single instance of a running process in the cluster and can contain one or more containers.

### YAML Example:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: mypod
  labels:
    app: myapp
spec:
  containers:
  - name: mycontainer
    image: nginx:latest
    ports:
    - containerPort: 80
```

### Key Components of Pod YAML:
- **apiVersion**: Specifies the API version that should be used to create the object.
- **kind**: Defines the type of object being created. In this case, it is a "Pod".
- **metadata**: Contains data that helps to identify the object.
  - **name**: Name of the Pod.
  - **labels**: Set of key-value pairs to organize and select the Pod.
- **spec**: Defines the desired state of the Pod.
  - **containers**: List of containers within the Pod.
    - **name**: Name of the container.
    - **image**: Container image to be used.
    - **ports**: Specifies the ports exposed by the container.
    - **containerPort**: The port on which the container is listening.

---

## 2. Deployment YAML

A **Deployment** is a higher-level controller that manages the scaling, updating, and rollbacks of Pods.

### YAML Example:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp-container
        image: myapp:latest
        ports:
        - containerPort: 8080
```

### Key Components of Deployment YAML:
- **apiVersion**: Specifies the API version to create the object.
- **kind**: Defines the type of object. In this case, it is a "Deployment".
- **metadata**: Contains data about the object.
  - **name**: Name of the Deployment.
- **spec**: Defines the desired state of the Deployment.
  - **replicas**: Number of replicas (Pods) that should be running.
  - **selector**: Defines how the Deployment will select which Pods to manage.
  - **template**: Describes the Pod template used for the Deployment.
    - **metadata**: Contains labels for Pods.
    - **spec**: Defines the containers within the Pods created by the Deployment.
      - **containers**: List of containers.
        - **name**: Name of the container.
        - **image**: Container image to be used.
        - **ports**: List of ports exposed by the container.

---

## 3. Service YAML

A **Service** is an abstraction that defines a logical set of Pods and provides a stable endpoint for accessing them.

### YAML Example:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: myservice
spec:
  selector:
    app: myapp
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
  type: LoadBalancer
```

### Key Components of Service YAML:
- **apiVersion**: Specifies the API version to create the object.
- **kind**: Defines the type of object. In this case, it is a "Service".
- **metadata**: Contains data about the object.
  - **name**: Name of the Service.
- **spec**: Defines the desired state of the Service.
  - **selector**: Defines the label selector to find the Pods to expose.
  - **ports**: Defines the ports exposed by the service.
    - **port**: The port on which the service is exposed.
    - **targetPort**: The port to which the traffic is forwarded inside the container.
  - **type**: Specifies the type of service. `LoadBalancer` will create an external load balancer.

---

## 4. ConfigMap YAML

A **ConfigMap** is used to store configuration data in the form of key-value pairs. These configurations can be referenced by Pods.

### YAML Example:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: my-config
data:
  mykey: myvalue
```

### Key Components of ConfigMap YAML:
- **apiVersion**: Specifies the API version to create the object.
- **kind**: Defines the type of object. In this case, it is a "ConfigMap".
- **metadata**: Contains data about the object.
  - **name**: Name of the ConfigMap.
- **data**: Contains key-value pairs used as configuration data.
  - **mykey**: A key in the ConfigMap with its associated value "myvalue".

---

## 5. PersistentVolume and PersistentVolumeClaim YAML

Persistent storage is managed through PersistentVolumes (PVs) and PersistentVolumeClaims (PVCs). A PVC is a request for storage from a user, and a PV represents the actual storage resource.

### PersistentVolume YAML Example:

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: my-pv
spec:
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  hostPath:
    path: "/mnt/data"
```

### PersistentVolumeClaim YAML Example:

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi
```

### Key Components of PersistentVolume and PersistentVolumeClaim YAML:
- **apiVersion**: Specifies the API version to create the object.
- **kind**: Defines the type of object. For PV it’s "PersistentVolume" and for PVC it’s "PersistentVolumeClaim".
- **metadata**: Contains data about the object.
  - **name**: Name of the PV or PVC.
- **spec**: Defines the desired state of the PV or PVC.
  - **capacity (PV)**: The amount of storage available.
  - **accessModes (PV and PVC)**: Defines the access modes. `ReadWriteOnce` means only one node can access it at a time.
  - **persistentVolumeReclaimPolicy (PV)**: Defines the reclaim policy. `Retain` keeps the volume even if the PVC is deleted.
  - **hostPath (PV)**: Defines where the volume is located on the host.
  - **resources (PVC)**: Requests storage from the PV.
  - **storage (PVC)**: Specifies how much storage is requested.
