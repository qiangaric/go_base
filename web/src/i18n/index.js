import { createI18n } from 'vue-i18n' 
import field from './field.js'
import pkg from './pkg.js';

const language = (
  (navigator.language ? navigator.language : navigator.userLanguage) || "zh"
).toLowerCase();

const messages = {
  en: {
    kuboardspray: 'Kuboard Spray',
    loginRequired: 'Login Required.',
    isRequiredField: 'is required.',
    field: field.en,
    pkg: pkg.en,
    unit: {
      second: '秒',
    },
    obj: {
      cluster: 'Kubernetes Cluster',
      resource: 'Resource Package',
      bastion: 'Bastion',
      localhost: 'KuboardSpray',
      addon: 'Addon: {name}',
    },
    node: {
      kube_control_plane: 'control plane',
      kube_node: 'worker node',
      etcd: 'etcd node',
      k8s_cluster: 'K8s Cluster',
    },
    cis: {
      conflict_warn: 'Conflict with CIS WARN rules',
    },
    msg: {
      feature_doesnot_support_selected_resource_package: 'This feature dosenot support the current resource package.',
      preview_yaml: 'Preview Yaml',
      name: 'Name',
      ok: 'OK',
      add: 'Add',
      save: 'Save',
      close: 'Close',
      cancel: 'Cancel',
      upload: 'Upload',
      edit: 'Edit',
      view: 'View',
      refresh: 'Refresh',
      prompt: 'Prompt',
      save_succeeded: 'Succeeded in saving.',
      save_failed: 'Failed in saving. {msg}',
      delete_succeeded: 'Succeeded in deleting.',
      delete_failed: 'Failed in deleting. {msg}',
      status: 'Status',
      confirmToCancel: "Modifications will be lost to proceed, do you confirm ?",
      operations: 'Operations',
      help: 'Help',
    },
    upgrade: {
      ignoreThisVersion: 'Ignore this version.',
      ignored: 'ignored version {version}, will keep in silence until we get a newer upgrade.',
    }
  },
  zh: { 
    kuboardspray: 'Kuboard Spray',
    loginRequired: '请先登录。',
    isRequiredField: '为必填字段',
    field: field.zh,
    pkg: pkg.zh,
    unit: {
      second: '秒',
    },
    obj: {
      cluster: 'Kubernetes 集群',
      resource: '资源包',
      bastion: '跳板机',
      localhost: 'KuboardSpray',
      addon: '可选组件：{name}',
    },
    node: {
      kube_control_plane: '控制节点',
      kube_node: '工作节点',
      etcd: 'ETCD节点',
      k8s_cluster: 'K8s 集群',
    },
    cis: {
      conflict_warn: 'CIS 扫描将出现 WARN',
    },
    msg: {
      feature_doesnot_support_selected_resource_package: '此特性不支持当前版本的资源包',
      preview_yaml: '预览 YAML',
      name: '名 称',
      ok: '确 定',
      add: '添 加',
      save: '保 存',
      delete: '删 除',
      close: '关 闭',
      cancel: '取 消',
      upload: '上 传',
      edit: '编 辑',
      view: '查 看',
      refresh: '刷 新',
      prompt: '提 示',
      save_succeeded: '保存成功',
      save_failed: '保存失败： {msg}',
      delete_succeeded: '删除成功',
      delete_failed: '删除失败： {msg}',
      status: '状 态',
      confirmToCancel: '将丢失已修改内容，确认取消编辑？',
      operations: '操 作',
      help: '帮 助',
    },
    upgrade: {
      ignoreThisVersion: '忽略这个版本',
      ignored: '已经忽略了版本 {version}，不会再提示您更新到该版本。',
    }
  },
}

const i18n = createI18n({
  fallbackLocale: 'zh',
  globalInjection:true,
  legacy: true,
  silentFallbackWarn: true,
  silentTranslationWarn: true,
  locale: language.split("-")[0] || "zh",
  messages,
});

export default i18n