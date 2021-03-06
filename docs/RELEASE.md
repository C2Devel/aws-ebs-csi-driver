# Инструкция по релизу новой версии

Инструкция протестирована на:
```sh
# uname -r
5.6.13-100.fc30.x86_64
# cat /etc/os-release 
NAME=Fedora
VERSION="30 (Thirty)"
ID=fedora
VERSION_ID=30
VERSION_CODENAME=""
PLATFORM_ID="platform:f30"
PRETTY_NAME="Fedora 30 (Thirty)"
ANSI_COLOR="0;34"
LOGO=fedora-logo-icon
CPE_NAME="cpe:/o:fedoraproject:fedora:30"
HOME_URL="https://fedoraproject.org/"
DOCUMENTATION_URL="https://docs.fedoraproject.org/en-US/fedora/f30/system-administrators-guide/"
SUPPORT_URL="https://fedoraproject.org/wiki/Communicating_and_getting_help"
BUG_REPORT_URL="https://bugzilla.redhat.com/"
REDHAT_BUGZILLA_PRODUCT="Fedora"
REDHAT_BUGZILLA_PRODUCT_VERSION=30
REDHAT_SUPPORT_PRODUCT="Fedora"
REDHAT_SUPPORT_PRODUCT_VERSION=30
PRIVACY_POLICY_URL="https://fedoraproject.org/wiki/Legal:PrivacyPolicy"
# docker --version
Docker version 19.03.12, build 48a66213fe
# ./kustomize version
{Version:kustomize/v3.6.1 GitCommit:c97fa946d576eb6ed559f17f2ac43b3b5a8d5dbd BuildDate:2020-05-27T20:47:35Z GoOs:linux GoArch:amd64}
```
## Версионирование

Используется следующая схема версионирования - <upstream_version>-CROC<X>. Где X - инкрементируется с каждым новым релизом. Например при текущей версии апстрима v0.5.0 и текущей версии этой репы v0.5.0-CROC1 следующая версия будет v0.5.0-CROC2. При обновлении версии апстрима, например до v0.6.0, успешный ребейз на новый апстрим будет результирован в версию v0.6.0-CROC2. Предполагается суппорт только актуальных версий.

Версии обозначаются гит тегами. Тегируется мастер ветка используя механизм релизов гитхаба. При создании нового релиза, описание релиза заполняется краткой сводкой изменений в новом релизе. После создания нового релиза (и тега), тег забирается на локалку (git pull upstream master --tags) и выполняется ручная сборка и публикация артефактов.

## Артефакты

Релизными артефактами этой репы является докер имадж и deployment конфиги для бубернетеса. При любом новом релизе необходимо обновлять kustomization.yaml и генерить бандл (например при релизе v0.5.0-CROC1):
- в файле deployment/kubernetes/stable/kustomization.yaml изменить ```newTag``` на новый актуальный (v0.5.0-CROC1)
- используя утилиту [kustomize](https://github.com/kubernetes-sigs/kustomize) собрать сингл-yaml-файл бандл для деплоймента:
```
kustomize build ./deployment/kubernetes/stable/ > ./deployment/kubernetes/stable/k_bundle.yaml
```

Для создания докер имаджа необходимы установленный и настроенный докер демон - https://docs.docker.com/get-docker/ . Для сборки имаджа необходимо:
- находясь в руте репы выполнить:
```docker build -t aws-ebs-csi-driver```
- после успешной сборки протегировать имадж:
```docker tag aws-ebs-csi-driver dhub.c2.croc.ru/kaas/aws-ebs-csi-driver:<version>```
- запушить имадж в регистри (необходимы врайт права в регистри неймспейсе):
```docker push dhub.c2.croc.ru/kaas/aws-ebs-csi-driver:<version>```
