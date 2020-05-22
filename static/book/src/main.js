import Vue from 'vue'
import 'bootstrap-css-only/css/bootstrap.min.css';
import VueRouter from 'vue-router'
import router from './router/index'
import VueAxios from 'vue-axios'
import 'mdbvue/build/css/mdb.css';
import App from './App.vue'

Vue.config.productionTip = false;
Vue.use(VueRouter, VueAxios);
new Vue({
  router,
  render: h => h(App),
}).$mount('#app');
