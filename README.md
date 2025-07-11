# PangeeCluster

基于 [kubespray](https://github.com/kubernetes-sigs/kubespray) 提供图形化的 K8S 集群离线安装、维护工具。

PangeeCluster 的在线文档地址为 [https://pangee-cluster.cn](https://pangee-cluster.cn)

## 快速安装

找一台不低于 1 核 2G，不少于 10G 剩余磁盘空间，已经安装好 docker 的服务器，执行如下指令，即可完成 PangeeCluster 的安装：

```sh
docker run -d \
  --restart=unless-stopped \
  --name=pangee-cluster \
  -p 80:80/tcp \
  -e TZ=Asia/Shanghai \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v ~/pangee-cluster-data:/data \
  eipwork/pangee-cluster:latest-amd64
  # 如果抓不到这个镜像，可以尝试一下这个备用地址：
  # swr.cn-east-2.myhuaweicloud.com/kuboard/pangee-cluster:latest-amd64
```

在浏览器地址栏中输入 `http://这台机器的IP地址`，输入用户名 `admin`，默认密码 `Kuboard123`，即可登录 pangee-cluster 界面，剩下的事情，在界面上操作一下，您就会啦。如果有困难，试试这篇文档 <a href="https://pangee-cluster.cn/guide/install-k8s.html" target="_blank">使用 PangeeCluster 安装 Kubernetes 集群</a>

**常见问题**

- 导入资源包时，可能会碰到 `no such file or directory` 或者 `permission denied` 之类的错误提示，通常是因为您开启了 SELinux，导致 pangee-cluster 不能读取映射到容器 `/data` 的路径
- pangee-cluster 所在机器不能当做 K8S 集群的一个节点，因为安装过程中会重启集群节点的容器引擎，这会导致 pangee-cluster 被重启掉。

## 配置开发测试环境

[配置开发测试环境](./docs/setup-dev/dev.md)

## 自制资源包

Kuboard-Spray 将定期提供最新版本的资源包，可以在 kubord-spray 资源包管理界面中查到，如果您是离线环境，也可以在 [https://kuboard.cn/support/pangee-cluster/](https://pangee-cluster.cn/support/)找到最新的资源包。您也可以自己制作资源包，资源包的项目地址在 [pangee-cluster-resource](https://github.com/opencmit/pangee-cluster-resource)，资源包的制作方法会在 2022 年 2 月提供文档。

## 社区

对此项目感兴趣的同学，请添加本项目的 Star 以后，扫码加入群聊，二维码在网站的页尾。

[pangee-cluster.cn](https://pangee-cluster.cn)
