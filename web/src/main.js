import { createApp } from 'vue'
import router from './router/index.js'
import {initKuboardMfe} from './micro-front-end.js'

import ElementPlus from 'element-plus';
import './styles/element-variables.scss'
import './styles/element-ui.css'
import icons from './styles/el-icons'

import zhCn from 'element-plus/es/locale/lang/zh-cn'

import store from './store'
import i18n from './i18n'
import initAxios from './utils/axios.js'
import openUrlInBlank from './utils/open-in-blank.js'
import validators from './utils/validators.js'

import components from './components/index.js'

import { VueClipboard } from '@soerenmartius/vue3-clipboard'

import App from './App.vue'

import axios from 'axios'
import checkNewVersion from './utils/version-checker.js'

axios.get("./version.json?nocache=" + new Date().getTime()).then(resp => {
  window.KuboardSpray = { version: resp.data }
  window.KuboardSpray.version.trimed = window.KuboardSpray.version.version.slice(0, window.KuboardSpray.version.version.length - 6)
  window.KuboardSpray.version.arch = window.KuboardSpray.version.version.slice(window.KuboardSpray.version.version.length - 5)
  const app = createApp(App)
  app.use(ElementPlus, {size: 'small', locale: zhCn, zIndex: 3000})
  app.use(store)
  app.use(router)
  app.use(i18n)
  icons(app)
  app.use(components)
  app.use(initAxios)
  app.use(validators)
  app.use(openUrlInBlank)
  app.use(VueClipboard)
  app.config.unwrapInjectedRef = true
  initKuboardMfe(app)
  app.mount('#app')
  checkNewVersion.init()
})

