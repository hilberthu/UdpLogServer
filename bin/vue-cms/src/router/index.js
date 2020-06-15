import Layout from '@/layout/layout'
import Login from '@/views/login'
import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/**
 * alwaysShow: true           if set true, will always show the root menu,
 * whatever its child routes length if not set alwaysShow, only more than ont
 * route under the children it will becomes nested mode, otherwise not show the
 * root menu alwaysShow: true
 * 如果设置为true,它将总是现在在根目录。如果不设置的话，当它只有1个子路由的时候，会把
 *                            它的唯一子路由放到跟目录上来，而不显示它自己本身。
 */

export const constantRouterMap =
    [
      {
        path: '/login',
        name: 'login',
        hidden: true,
        component: Login,
        meta: {title: '登录'}
      },
      {
        path: '/404',
        component: () => import('@/views/error-page/404'),
        hidden: true
      },
      {
        path: '/',
        hidden: true,
        component: Layout,
        redirect: '/home',
        children: [{
          path: 'home',
          name: 'home',
          component: () => import('@/views/homepage'),
          meta: {title: '首页'}
        }]
      },
      {
        path: '/user',
        component: Layout,
        hidden: true,
        meta: {icon: 'tickets', title: '个人中心'},
        children: [
          {
            path: 'profile',
            name: 'Profile',
            component: () => import('@/views/user/profile'),
            meta: {icon: 'warning', title: '个人中心'}
          },
          {
            path: 'avatar',
            name: 'Avatar',
            component: () => import('@/views/user/profile'),
            meta: {icon: 'warning', title: '修改头像'}
          }
        ]
      },
      {
        path: '/gm',
        component: Layout,
        redirect: '/gm/logquery',
        alwaysShow: true,
        meta: {icon: 'svg-aperture', title: '日志查询'},
        children: [{
          path: 'logquery',
          name: 'logquery',
          component: () => import('@/views/gm/logquery'),
          meta: {icon: 'svg-aperture', title: '日志查询'}
        }]
      }
    ]

    export default new Router({
      // mode: 'history',  require service support
      // scrollBehavior: () => ({ y: 0 }),
      routes: constantRouterMap
    })

export const asyncRouterMap = [

]