"use strict";

import Vue from 'vue';
import VueRouter from 'vue-router'
import App from './components/app.vue';
import Thermals from './components/thermals.vue';
 
Vue.use(VueRouter);

const router = new VueRouter({
  routes: [
    {
      path: '/',
      component: Thermals
    },
    {
      path: '/thermals',
      component: Thermals
    },
  ]
});

const BaseVue = Vue.extend({ router })

new BaseVue({
  el: '#app',
  router,
  render: createEle => createEle(App)
});