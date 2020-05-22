import Vue from 'vue';
import VueRouter from 'vue-router';
import Layout from '../layout/Layout';
import EmptyLayout from '../layout/EmptyLayout';
import Home from '../components/Home'
import BookList from '../components/BookList'
import BookItem from '../components/BookItem'
import Read from '../components/Read'
import Shelf from '../components/Shelf'
import Search from '../components/Search'

Vue.use(VueRouter);


export const constantRouterMap = [
  {
    path: '/',
    component: Layout,
    redirect: '/main/index'
  },
  {
    path: '/main',
    component: Layout,
    children: [
      {
        path: 'index',
        name: 'Main',
        component: Home,
      },
    ]
  },
  {
    path: '/book',
    component: Layout,
    children: [
      {
        path: 'list/:id',
        name: 'BookList',
        component: BookList,
      },
      {
        path: 'item',
        name: 'BookItem',
        component: BookItem,
      }
    ]
  },
  {
    path: '/search',
    component: Layout,
    children: [
      {
        path: 'index',
        name: 'Search',
        component: Search,
      }
    ]
  },
  {
    path: '/shelf',
    component: Layout,
    children: [
      {
        path: 'index',
        name: 'Shelf',
        component: Shelf,
      }
    ]
  },
  {
    path: '/read',
    component: EmptyLayout,
    children: [
      {
        path: 'index/:id/:name',
        name: 'Read',
        component: Read,
      }
    ]
  },
];

export default new VueRouter({
  routes: constantRouterMap
});
