# 镜像转换工具

## 前提
- 需要 docker hub 有对应 Repositories
- 例如拉取 quay.io/jetstack/cert-manager-cainjector:v1.10.0，需要在docker hub 上面创建 cert-manager-cainjector 仓库
## 使用 
### 推送镜像

- 
    `./image-convert --docker-hub-user {docker hub user} --push --s-image quay.io/jetstack/cert-manager-cainjector:v1.10.0`
### 拉取镜像

    `./image-convert --docker-hub-user {docker hub user} --pull --s-image quay.io/jetstack/cert-manager-cainjector:v1.10.0`