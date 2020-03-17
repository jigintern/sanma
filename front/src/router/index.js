import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [{
    path: '/',
    name: 'Top',
    component: () => import('../views/Top.vue')
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('../views/About.vue')
  },
  {
    path: '/signin',
    name: 'Signin',
    component: () => import('../views/Signin.vue')
  },
  {
    path: '/signup',
    name: 'Signup',
    component: () => import('../views/Signup.vue')
  },
  {
    path: '/read-article/:id',
    name: 'Read-article',
    component: () => import('../views/Read-article.vue')
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: () => import('../views/Dashboard.vue')
  },
  {
    path: '/post-article',
    name: 'Post-article',
    component: () => import('../views/Post-article.vue')
  },
  {
    path: '/edit-article',
    name: 'Edit-article',
    component: () => import('../views/Edit-article.vue')
  },
  {
    path: '/add-crawl',
    name: 'Add-crawl',
    component: () => import('../views/Add-crawl.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router