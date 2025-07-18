<i18n>
en:
  resourceVersion: Resource Package Version
  not_support_cluster_version: Current resource package doesnot support the operation to view cluster version, if you want to upgrade the cluster, please select a new resource package with the above button
  chooseNewResourcePackage: Upgrade to New ResourcePackage
  finished_upgrade: Finished upgrading cluster, previous version is
  finished_upgrade_title: Finished upgrading cluster
zh:
  resourceVersion: 资源包版本
  not_support_cluster_version: 当前资源包不支持查询集群版本的操作，如果您想要升级集群，请点击上面的按钮并选择新的资源包
  chooseNewResourcePackage: 选择新的资源包（升级集群）
  finished_upgrade: 已完成集群升级，升级前的资源包版本是
  finished_upgrade_title: 已完成集群升级

</i18n>

<template>
  <el-scrollbar height="calc(100vh - 220px)">
    <div>
      <el-form label-position="left" label-width="120px">
        <el-form-item :label="t('resourceVersion')">
          <div class="app_text_mono version_no">
            <span style="margin-right: 20px;">{{ resource.metadata.version }}</span>
            <template v-if="!cluster.history.processing">
              <UpgradeTask v-if="pendingUpgrade" :cluster="cluster"
                :controlPlanePendingUpgrade="controlPlanePendingUpgrade" @refresh="$emit('refresh')"></UpgradeTask>
              <template v-else>
                <el-button type="danger" icon="el-icon-upload" @click="$refs.choose.show()"
                  style="margin-left: 20px;">{{ t('chooseNewResourcePackage') }}</el-button>
                <template v-if="cluster.inventory.all.hosts.localhost.pangeecluster_resource_package_previous">
                  <el-popover placement="top-start" :title="t('finished_upgrade_title')" :width="350" trigger="hover">
                    <template #reference>
                      <el-button type="success" round icon="el-icon-circle-check">{{ t('finished_upgrade_title')
                      }}</el-button>
                    </template>
                    <div>{{ t('finished_upgrade') }}</div>
                    <div class="app_text_mono" style="line-height: 36px; color: var(--el-color-primary);">
                      {{ cluster.inventory.all.hosts.localhost.pangeecluster_resource_package_previous }}</div>
                  </el-popover>
                </template>
              </template>
            </template>
            <PangeeClusterLink href="https://pangee-cluster.cn/guide/maintain/upgrade.html" style="margin-left: 20px;"
              :size="12">
            </PangeeClusterLink>
          </div>
        </el-form-item>
      </el-form>
    </div>
    <el-alert
      v-if="resource.data.supported_playbooks['cluster_version_' + cluster.inventory.all.children.target.vars.container_manager] === undefined"
      type="warning" :closable="false">
      {{ t('not_support_cluster_version') }}
    </el-alert>
    <el-alert v-else-if="errMsg" type="error" :closable="false">
      <pre>{{ errMsg }}</pre>
    </el-alert>
    <CompareVersion v-else :cluster="cluster" :version="version"
      :controlPlanePendingUpgrade="controlPlanePendingUpgrade" @refresh="$emit('refresh')"></CompareVersion>
    <ChooseNewResourcePackage ref="choose" :cluster="cluster" @refresh="$emit('refresh')"></ChooseNewResourcePackage>
  </el-scrollbar>
</template>

<script>
import CompareVersion from './CompareVersion.vue'
import ChooseNewResourcePackage from './ChooseNewResourcePackage.vue'
import UpgradeTask from './UpgradeTask.vue'

export default {
  props: {
    cluster: { type: Object, required: true },
  },
  data() {
    return {
      version: undefined,
      errMsg: undefined,
      loading: false,
    }
  },
  computed: {
    pendingUpgrade() {
      for (let k in this.cluster.inventory.all.hosts) {
        let host = this.cluster.inventory.all.hosts[k]
        if (host.pangeecluster_node_action === 'upgrade_node') {
          return true
        }
      }
      return false
    },
    controlPlanePendingUpgrade() {
      for (let key in this.cluster.inventory.all.children.target.children.k8s_cluster.children.kube_control_plane.hosts) {
        if (this.cluster.inventory.all.hosts[key].pangeecluster_node_action === 'upgrade_node') {
          return true
        }
      }
      for (let key in this.cluster.inventory.all.children.target.children.etcd.hosts) {
        if (this.cluster.inventory.all.hosts[key].pangeecluster_node_action === 'upgrade_node') {
          return true
        }
      }
      return false
    },
    resource() {
      return this.cluster.resourcePackage
    }
  },
  components: { CompareVersion, ChooseNewResourcePackage, UpgradeTask },
  mounted() {
    this.loadClusterVersion()
  },
  methods: {
    async loadClusterVersion() {
      if (this.resource.data.supported_playbooks['cluster_version_' + this.cluster.inventory.all.children.target.vars.container_manager] === undefined) {
        return
      }
      this.loading = true
      this.errMsg = undefined
      await this.pangeeClusterApi.get(`/clusters/${this.cluster.name}/state/version`).then(resp => {
        this.version = resp.data.data
      }).catch(e => {
        console.log(e)
        if (e.response && e.response.data.message) {
          this.errMsg = e.response.data.message
        } else {
          this.errMsg = '' + e
        }
      })
      this.loading = false
    },
  }
}
</script>

<style scoped lang="css">
.version_no {
  font-weight: bolder;
  font-size: 14px;
}
</style>
