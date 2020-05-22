<style scoped="scoped">
  @import url(https://fonts.googleapis.com/css?family=Roboto:400,500);
  @import url(https://fonts.googleapis.com/icon?family=Material+Icons);
  @import url(https://cdnjs.cloudflare.com/ajax/libs/normalize/5.0.0/normalize.min.css);

  [ripple] {
    position: relative;
    overflow: hidden;
  }

  [ripple] .ripple--container {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
  }

  [ripple] .ripple--container span {
    -webkit-transform: scale(0);
    transform: scale(0);
    border-radius: 100%;
    position: absolute;
    opacity: 0.75;
    background-color: rgba(33, 150, 243, 0.5);
    -webkit-animation: ripple 1000ms;
    animation: ripple 1000ms;
  }

  @-webkit-keyframes ripple {
    to {
      opacity: 0;
      -webkit-transform: scale(2);
      transform: scale(2);
    }
  }

  @keyframes ripple {
    to {
      opacity: 0;
      -webkit-transform: scale(2);
      transform: scale(2);
    }
  }

  html, body {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-family: 'Roboto', sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  body {
    height: 100%;
    transition: padding-top 150ms;
    padding-top: 56px;
  }

  .action-bar {
    transition: height 150ms;
    width: 100%;
    position: fixed;
    top: 0;
    margin-left: -15px;
    z-index: 999;
    color: rgba(0, 0, 0, 0.9);
    height: 56px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    box-shadow: 0 3px 5px -1px rgba(0, 0, 0, 0.2), 0 5px 8px 0 rgba(0, 0, 0, 0.14), 0 1px 14px 0 rgba(0, 0, 0, 0.12);
  }

  .action-bar .title {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .action-bar .title h3 {
    font-weight: 400;
    margin: 0;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    cursor: pointer;
  }

  .action-bar .title i {
    color: rgba(0, 0, 0, 0.7);
    display: flex;
    justify-content: center;
    align-items: center;
    width: 32px;
    height: 32px;
    border-radius: 50%;
    margin: 16px 32px 16px 16px;
    cursor: pointer;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
  }

  .action-bar .icons {
    display: flex;
    justify-content: space-between;
    align-content: center;
  }

  .action-bar .icons i {
    transition: margin 150ms;
    color: rgba(0, 0, 0, 0.7);
    display: flex;
    justify-content: center;
    align-items: center;
    width: 32px;
    height: 32px;
    border-radius: 50%;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    margin: 16px;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .side-nav {
    -webkit-transform: translateX(-320px);
    transform: translateX(-320px);
    overflow-x: hidden;
    z-index: 1000;
    position: fixed;
    top: 0;
    left: 0;
    height: 100vh;
    transition: max-width 200ms, -webkit-transform 200ms cubic-bezier(0.4, 0, 0.6, 1);
    transition: max-width 200ms, transform 200ms cubic-bezier(0.4, 0, 0.6, 1);
    transition: max-width 200ms, transform 200ms cubic-bezier(0.4, 0, 0.6, 1), -webkit-transform 200ms cubic-bezier(0.4, 0, 0.6, 1);
    min-width: 240px;
    max-width: 280px;
    box-shadow: 0 8px 10px -5px rgba(0, 0, 0, 0.2), 0 16px 24px 2px rgba(0, 0, 0, 0.14), 0 6px 30px 5px rgba(0, 0, 0, 0.12);
  }

  .side-nav header {
    transition: padding 150ms;
    background: url(https://lh3.googleusercontent.com/-OpzXLwUodco/VUzSNKm5x5I/AAAAAAAANtw/QlXckPJDTxs/w1366-h768/Wallrox%2B%25281%2529.png) no-repeat;
    background-size: 320px 180px;
    width: auto;
  }

  .side-nav header .rich-area {
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding: 0 16px 8px 16px;
    height: 56px;
  }

  .side-nav header .rich-area .email {
    display: flex;
    justify-content: space-between;
    align-items: center;
    color: #fff;
  }

  .side-nav header .rich-area .email i {
    border-radius: 50%;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    height: 32px;
    width: 32px;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
  }

  .side-nav header .rich-area a {
    font-size: 15px;
    text-decoration: none;
    color: #fff;
    /*box-shadow: 0, 0, 0;*/
  }

  .side-nav header .avatars-container {
    padding: 16px 16px 0 16px;
    display: flex;
    justify-content: space-between;
  }

  .side-nav header .avatars-container img {
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    border-radius: 50%;
  }

  .side-nav nav ul {
    background-color: #fff;
    padding: 8px 0 0 0;
    margin: 0;
    list-style: none;
  }

  .side-nav nav ul hr {
    background-color: rgba(0, 0, 0, 0.15);
    height: 1.2px;
    border: 0;
  }

  .side-nav nav ul #subheader {
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    pointer-events: none;
    color: rgba(0, 0, 0, 0.5);
  }

  .side-nav nav ul li {
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    color: rgba(0, 0, 0, 0.9);
    padding: 0 16px;
    font-weight: 500;
    font-size: 14px;
    cursor: pointer;
    height: 48px;
    display: flex;
    align-items: center;
    margin: 0;
    padding: 0 16px;
  }

  .side-nav nav ul li.active {
    color: #2196F3;
  }

  .side-nav nav ul li.active i {
    color: #2196F3;
  }

  .side-nav nav ul li i {
    color: rgba(0, 0, 0, 0.6);
    padding-right: 32px;
  }

  @media all and (max-width: 335px) {
    .action-bar .icons i {
      margin: 8px;
    }
  }

  @media all and (min-width: 960px) {
    body {
      padding-top: 64px;
    }

    .action-bar {
      height: 64px;
    }

    .action-bar .title i {
      margin: 24px 32px 24px 24px;
    }

    .action-bar .icons i {
      margin: 24px;
    }

    .side-nav {
      max-width: 320px;
    }

    .side-nav header .rich-area {
      padding: 0 24px 8px 24px;
    }

    .side-nav header .avatars-container {
      padding: 24px 24px 0 24px;
    }
  }

  @media all and (max-width: 320px) {
    .action-bar {
      height: 56px;
    }

    .action-bar .title i {
      margin: 0 8px;
    }

    .action-bar .icons i {
      margin: 0 6px;
    }
  }

</style>

<template>
  <mdb-container fluid class="z-depth-1 yellow lighten-5">
    <div v-show="actionShow" class="action-bar">
      <div class="title">
        <i class="material-icons" ripple id='menu' v-on:click="openSideNav()">menu</i>
        <h2 style="margin-top: 15px">{{name}}</h2>
      </div>
      <div class="icons">
        <!--<i class="material-icons" ripple>search</i>-->
        <i class="material-icons" ripple>favorite</i>
        <i class="material-icons" ripple>more_vert</i>
      </div>
    </div>

    <div class="side-nav" v-on:click="closeSideNav()">
      <header>
        <div class="rich-area">
          <h5 style="color:white;margin-top: 15px;">{{name}}</h5>
        </div>
      </header>
      <nav>
        <ul>
          <li ripple v-on:click="getChapterContent(chapter.Id)" v-for="chapter in chapters">
            {{chapter.Name}}
          </li>
        </ul>
      </nav>
    </div>
    <div style="margin-top: 56px;">
      <mdb-row>
        <mdb-col>
          <h5 style="margin-top: 15px">{{contentName}}</h5>
          <span v-html="content"></span>
        </mdb-col>
      </mdb-row>
    </div>
  </mdb-container>
</template>

<script>
  import ajax from '../utils/requests'
  import {
    mdbBtn,
    mdbCol,
    mdbContainer,
    mdbModal,
    mdbModalBody,
    mdbModalFooter,
    mdbModalHeader,
    mdbModalTitle,
    mdbPopover,
    mdbRow,
    mdbTooltip
  } from 'mdbvue';

  export default {
    data() {
      return {
        actionShow: true,
        id: 0,
        content: '',
        contentName: '',
        name: '',
        sideNav: '',
        chapters: [],
      }
    },
    name: "Read",
    components: {
      mdbCol,
      mdbContainer,
      mdbRow,
      mdbBtn,
      mdbPopover,
      mdbTooltip, mdbModal, mdbModalHeader, mdbModalTitle, mdbModalBody, mdbModalFooter
    },
    methods: {
      openSideNav() {
        this.sideNav.style.transform = 'translateX(0)'
      },
      closeSideNav() {
        this.sideNav.style.transform = 'translateX(-320px)'
      },
      getChapterContent(id) {
        ajax.get('/chapter', {'chapterId': id, 'bookId': this.id}).then(
          (res) => {
            this.content = res.Data[0].Content;
            this.contentName = res.Data[0].Name
          }
        );
      }
    },
    mounted() {

      this.sideNav = document.querySelector('.side-nav');
      window.addEventListener('resize', this.closeSideNav);

    },
    updated() {
      var cleanUp, debounce, i, len, ripple, rippleContainer, ripples, showRipple;

      debounce = function (func, delay) {
        var inDebounce;
        inDebounce = undefined;
        return function () {
          var args, context;
          context = this;
          args = arguments;
          clearTimeout(inDebounce);
          return inDebounce = setTimeout(function () {
            return func.apply(context, args);
          }, delay);
        };
      };

      showRipple = function (e) {
        var pos, ripple, rippler, size, style, x, y;
        ripple = this;
        rippler = document.createElement('span');
        size = ripple.offsetWidth;
        pos = ripple.getBoundingClientRect();
        x = e.pageX - pos.left - (size / 2);
        y = e.pageY - pos.top - (size / 2);
        style = 'top:' + y + 'px; left: ' + x + 'px; height: ' + size + 'px; width: ' + size + 'px;';
        ripple.rippleContainer.appendChild(rippler);
        return rippler.setAttribute('style', style);
      };

      cleanUp = function () {
        while (this.rippleContainer.firstChild) {
          this.rippleContainer.removeChild(this.rippleContainer.firstChild);
        }
      };

      ripples = document.querySelectorAll('[ripple]');

      for (i = 0, len = ripples.length; i < len; i++) {
        ripple = ripples[i];
        rippleContainer = document.createElement('div');
        rippleContainer.className = 'ripple--container';
        ripple.addEventListener('mousedown', showRipple);
        ripple.addEventListener('mouseup', debounce(cleanUp, 2000));
        ripple.rippleContainer = rippleContainer;
        ripple.appendChild(rippleContainer);
      }

      var actionBar = document.querySelector('.action-bar')
      var sideNav = document.querySelector('.side-nav')
      var width = document.documentElement.clientWidth

      sideNav.style.width = width - actionBar.clientHeight + 'px'

    },
    created() {
      let id = this.$route.params.id;
      this.name = this.$route.params.name;
      this.id = id;
      ajax.get('/chapter/list', {'bookId': id}).then(
        res => {
          this.chapters = res.Data;
          ajax.get('/chapter', {'chapterId': res.Data[0].Id, 'bookId': this.id}).then(
            (res) => {
              this.content = res.Data[0].Content;
              this.contentName = res.Data[0].Name
            }
          );
        }
      );

    }
  }
</script>
