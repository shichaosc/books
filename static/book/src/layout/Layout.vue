<template>
  <div id="app">
    <mdb-navbar class="white sticky-top navbar-light">
      <!-- Navbar brand -->
      <mdb-navbar-brand href="https://mdbootstrap.com/">
        <!--<mdb-icon class="black-text" icon="yin-yang" size="lg"/>-->
        <span class="black-text" v-on:click="goHome()">I Love Reading</span>
      </mdb-navbar-brand>
      <mdb-navbar-toggler>
        <mdb-navbar-nav>
          <mdb-btn size="sm" color="danger" v-on:click="goHome()">分类</mdb-btn>
          <mdb-btn size="sm" color="info" v-on:click="goShelf()">书架</mdb-btn>
          <mdb-btn size="sm" color="secondary" v-on:click="goSearch()">搜索</mdb-btn>
          <!-- Dropdown -->
        </mdb-navbar-nav>
      </mdb-navbar-toggler>
    </mdb-navbar>
    <app-main></app-main>
  </div>
</template>

<script>
  import AppMain from '../components/AppMain.vue'
  import {
    mdbBtn,
    mdbDropdown,
    mdbDropdownItem,
    mdbDropdownMenu,
    mdbDropdownToggle,
    mdbIcon,
    mdbMask,
    mdbNavbar,
    mdbNavbarBrand,
    mdbNavbarNav,
    mdbNavbarToggler,
    mdbNavItem,
  } from 'mdbvue';

  export default {
    name: 'Layout',
    methods: {
      goHome() {
        this.$router.push({'name': 'Main'})
      },
      goShelf() {
        this.$router.push({'name': 'Shelf'})
      },
      goSearch() {
        this.$router.push({'name': 'Search'})
      }
    },
    components: {
      mdbBtn,
      mdbMask,
      mdbIcon,
      mdbNavbar,
      mdbNavbarBrand,
      mdbNavbarToggler,
      mdbNavbarNav,
      mdbNavItem,
      mdbDropdown,
      mdbDropdownMenu,
      mdbDropdownToggle,
      mdbDropdownItem,
      AppMain
    }
  }
</script>

<style>
  #app {
    /*margin-top: 60px;*/
    /*padding: 5px;*/
  }
</style>
