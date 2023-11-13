import axios from 'axios'
import Cookies from 'js-cookie'

// let kuboardSprayId = 'default'
// let splitedPath = location.pathname.split('/')
// if (splitedPath[0] === 'kuboardspray' && splitedPath[1] !== undefined) {
//   kuboardSprayId = splitedPath[1]
// }

var kuboardSprayApi

const baseURL = `./api`

var vueapp

const comp = {
  install(app) {
    vueapp = app
    kuboardSprayApi = axios.create({
      baseURL: baseURL,
      timeout: 120000,
      headers: {
        Authorization: 'Bearer ' + Cookies.get('KuboardSprayToken')
      }
    })
    kuboardSprayApi.interceptors.response.use(function (response) {
      return response;
    }, function (error) {
      if (error.response && error.response.status === 401) {
        if (error.response.request && error.response.request.responseURL && error.request.responseURL.indexOf('/api/login') >= 0) {
          return Promise.reject(error)
        }
        window.VueAppComponent.$alert(window.VueAppComponent.$t('loginRequired'), window.VueAppComponent.$t('loginRequired'), {
          callback: () => {
            clearAllCookie()
            window.VueAppComponent.$router.push('/login')
          }
        })
      }
      return Promise.reject(error);
    });
    app.config.globalProperties.kuboardSprayApi = kuboardSprayApi
  }
}

export default comp

export {baseURL}

export function clearAllCookie() {
  Cookies.remove('KuboardSprayToken', { path: location.pathname })
  Cookies.remove('KuboardSprayLogin', { path: location.pathname })
  comp.install(vueapp)
}

export function setupCookie(token, expires) {
  console.log(token)
  Cookies.set('KuboardSprayToken', token, { path: location.pathname, expires: expires } )
  comp.install(vueapp)
}