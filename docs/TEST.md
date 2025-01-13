# Testing

Типы тестов, которые есть в этой репе:
- unit
- e2e

Для всех тестов можно использовать тестовый контейнер, образ для которого можно собрать командой:
```
make test-image
```

## unit

Запускаются через ```make test```. Используют встроенную в Go поддержку тестирования.

Для запуска тестов можно использовать тестовый контейнер, чтобы тесты запускались в linux-окружении.
- запустить тестовый контейнер:
  ```bash 
  make docker-unit-test-container
  ```
- внутри контейнера запустить тесты:
  ```
  cd /local-code
  make test
  ```

## e2e

В директории tests есть e2e/e2e-kubernetes/integration/sanity субдиректории. Основная масса тестов находится в e2e. 
Тесты в е2е функциональные, в основном работют в API Kubernetes и API облака.

Для выполнения тестов потребуется:
- Для single-az 1 нода мастер и 1 воркер
- Для multi-az 3 ноды мастеров в разных аз и 1 воркер в одной аз

### Как запустить тесты
Для запуска тестов понадобится:
- [создать](https://docs.k2.cloud/ru/services/kubernetes/eks_cluster.html#createekscluster) Kubernetes кластер в облаке с установленным aws-ebs-csi-driver
- задать переменные окружения:
  - `KUBECONFIG_FILE` - местонахождения kubeconfig для созданного кластера на локальном диске
  - `AWS_ACCESS_KEY_ID` - ключ для доступа к API
  - `AWS_SECRET_ACCESS_KEY` - серкрет для доступа к API
  - `EC2_URL` - URL сервиса EC2
  - `AWS_REGION` - регион
  - `AWS_AVAILABILITY_ZONES` - перечислинные через запятую зоны доступности, в которых развернут кластер

- запустить тестовый контейнер командой: 
```bash 
make docker-e2e-container
```
- внутри контейнера запустить тесты:
  - для single az:
  ```
  ~/go/bin/ginkgo -v --focus="\\[ebs-csi-e2e\] \\[single-az\\]" --skip="\[modify-volume\]"  /local-code/tests/e2e
  ```
  - для multi az:
  ```
  ~/go/bin/ginkgo -v --focus="\\[ebs-csi-e2e\] \\[multi-az\\]" --skip="\[modify-volume\]"  /local-code/tests/e2e
  ```
Список тестов, которые должны проходить:
- [ebs-csi-e2e] [single-az] Pre-Provisioned [env] should write and read to a pre-provisioned volume
- [ebs-csi-e2e] [single-az] Pre-Provisioned [env] should use a pre-defined snapshot and create pv from that
- [ebs-csi-e2e] [single-az] Pre-Provisioned [env] should use a pre-provisioned volume and mount it as readOnly in a pod
- [ebs-csi-e2e] [single-az] Pre-Provisioned [env] should use a pre-provisioned volume and retain PV with reclaimPolicy "Retain"
- [ebs-csi-e2e] [single-az] Pre-Provisioned [env] should use a pre-provisioned volume and delete PV with reclaimPolicy "Delete"
- [ebs-csi-e2e] [single-az] [format-options] Formatting a volume using an ext3 filesystem with a custom blocksize parameter successfully mounts and is resizable
- [ebs-csi-e2e] [single-az] [format-options] Formatting a volume using an ext3 filesystem with a custom bytesperinode parameter successfully mounts and is resizable
- [ebs-csi-e2e] [single-az] [format-options] Formatting a volume using an ext3 filesystem with a custom inodesize parameter successfully mounts and is resizable
- [ebs-csi-e2e] [single-az] [format-options] Formatting a volume using an ext3 filesystem with a custom numberofinodes parameter successfully mounts and is resizable
- [ebs-csi-e2e] [single-az] [format-options] Formatting a volume using an ext4 filesystem with a custom blocksize parameter successfully mounts and is resizable
- [ebs-csi-e2e] [single-az] [format-options] Formatting a volume using an ext4 filesystem with a custom bytesperinode parameter successfully mounts and is resizable
- [ebs-csi-e2e] [single-az] [format-options] Formatting a volume using an ext4 filesystem with a custom ext4bigalloc parameter successfully mounts and is resizable
- [ebs-csi-e2e] [single-az] [format-options] Formatting a volume using an ext4 filesystem with a custom ext4clustersize parameter successfully mounts and is resizable
- [ebs-csi-e2e] [single-az] [format-options] Formatting a volume using an ext4 filesystem with a custom inodesize parameter successfully mounts and is resizable
- [ebs-csi-e2e] [single-az] [format-options] Formatting a volume using an ext4 filesystem with a custom numberofinodes parameter successfully mounts and is resizable
- [ebs-csi-e2e] [single-az] [format-options] Formatting a volume using an xfs filesystem with a custom blocksize parameter successfully mounts and is resizable
- [ebs-csi-e2e] [single-az] [format-options] Formatting a volume using an xfs filesystem with a custom inodesize parameter successfully mounts and is resizable
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should create a volume on demand with volume type "gp2" and fs type "xfs"
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should create a volume on demand with volume type "io2" and fs type "xfs"
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should create a volume on demand with volume type "st2" and fs type "xfs"
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should create a volume on demand with volumeType "gp2" and encryption
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should create a volume on demand with volumeType "io2" and encryption
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should create a volume on demand with volumeType "st2" and encryption
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should create a volume on demand with provided mountOptions
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should create multiple PV objects, bind to PVCs and attach all to a single pod
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should create multiple PV objects, bind to PVCs and attach all to different pods
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should create a raw block volume on demand
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should create a raw block volume and a filesystem volume on demand and bind to the same pod
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should create multiple PV objects, bind to PVCs and attach all to different pods on the same node
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should create a volume on demand and mount it as readOnly in a pod
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should delete PV with reclaimPolicy "Delete"
- [ebs-csi-e2e] [single-az] Dynamic Provisioning [env] should retain PV with reclaimPolicy "Retain"
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should create a deployment object, write and read to it, delete the pod and write and read to it again
- [ebs-csi-e2e] [single-az] Dynamic Provisioning should create a volume on demand and resize it
- [ebs-csi-e2e] [single-az] Snapshot should create a pod, write and read to it, take a volume snapshot, and create another pod from the snapshot
- [ebs-csi-e2e] [single-az] [requires-aws-api] Dynamic Provisioning should create a volume with additional tags
- [ebs-csi-e2e] [single-az] [requires-aws-api] Dynamic Provisioning should create a snapshot with additional tags