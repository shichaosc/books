<style scoped>
  /*
*  Imports
*/
  /*
*  Imports
*/
  @import url(https://cdnjs.cloudflare.com/ajax/libs/normalize/5.0.0/normalize.min.css);
  @import url(https://cdn.jsdelivr.net/foundation/6.2.0/foundation.min.css);
  @import url(https://fonts.googleapis.com/css?family=Lato:400,300,700,900);
  @import url(https://fonts.googleapis.com/css?family=Roboto+Slab:400,700);
  @import url("https://netdna.bootstrapcdn.com/font-awesome/3.2.1/css/font-awesome.min.css");

  /*
   *  Variables
  */
  /* Single Line truncation */
  /* Multiline truncation */
  /*
   *  Global styles
  */
  html,
  .container-fluid,
  .main,
  .main-container {
    height: 100%;
  }

  .container-fluid {
    font-family: "Lato", "Helvetica Neue", Helvetica, Arial, sans-serif;
    color: #313131;
    background: #ecf0f1;
    height: 1000px;
  }

  .row {
    max-width: 1024px;
  }

  /*
   *  Typography
  */
  .container-fluid,
  p,
  a,
  li {
    font-family: "Lato", "Helvetica Neue", Helvetica, Arial, sans-serif;
    font-size: 15px;
  }

  a {
    color: #ffffff;
  }

  h2,
  h3 {
    margin-top: 6px;
    margin-bottom: 6px;
    font-size: 28px;
    font-weight: bold;
    letter-spacing: -1px;
    color: #313131;
  }

  h2 {
    font-size: 28px;
  }

  h3 {
    font-size: 24px;
  }

  h4 {
    margin-bottom: 12px;
    font-size: 22px;
    line-height: 40px;
    color: rgba(49, 49, 49, 0.7);
  }

  p:not(.author) {
    font-size: 15px;
    font-weight: 300;
    line-height: 1.4;
  }

  p.author {
    margin-bottom: 10px;
    letter-spacing: -1px;
    font-weight: 400;
    color: rgba(49, 49, 49, 0.5);
  }

  /*
   *  Button
  */
  a.button {
    margin-bottom: 0;
    padding: 8px 14px;
    font-size: 14px;
    font-weight: 600;
    border-radius: 3px;
    background-color: rgba(49, 49, 49, 0.5);
  }

  a.button:hover, a.button:focus {
    background-color: #219e9a;
  }

  /*
   *  Header
  */
  .page-header {
    position: relative;
    margin-bottom: 55px;
    width: 100%;
    border-bottom: 1px solid #d2d6d5;
    background-color: #fff;
  }

  .main-logo {
    display: inline-block;
    padding: 1em;
    width: auto;
  }

  .main-logo a.logo {
    display: block;
    width: 150px;
    height: 40px;
    text-indent: -9999px;
    background: url("http://interactivejoe.com/book-viewer/assets/images/logo.svg");
    background-color: #fff;
    transition: background-color 250ms ease-out;
  }

  .main-logo a.logo:hover, .main-logo a.logo:focus {
    background-color: #ecf0f1;
  }

  .menu-search {
    float: right;
    width: calc(100% - 200px);
  }

  /*
   *  Search Input
  */
  .catalog-search {
    position: relative;
    margin: 0 20px;
    max-width: 275px;
    width: calc(100% - 2em);
    vertical-align: top;
    overflow: hidden;
    float: right;
  }

  .input_field {
    position: relative;
    display: block;
    float: right;
    margin-top: 10px;
    padding: 14px 7px 0;
    width: 100%;
    border: none;
    border-radius: 0;
    color: #313131;
    font-weight: normal;
    font-family: "Lato", "Helvetica Neue", Helvetica, Arial, sans-serif;
    background: none;
    box-shadow: none;
    -webkit-appearance: none;
    /* for box shadows to show on iOS */
  }

  .input_field:focus {
    outline: none;
    border: none;
    background: none;
    box-shadow: none;
    -webkit-appearance: none;
  }

  .input_field:focus .input_label-content {
    bottom: 20px;
  }

  .input_label {
    position: absolute;
    display: inline-block;
    bottom: 0;
    left: 0;
    padding: 0 0.25em;
    width: 100%;
    height: calc(100% - 1em);
    color: #d2d6d5;
    font-weight: light;
    font-size: 15px;
    text-align: left;
    pointer-events: none;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    -webkit-touch-callout: none;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
  }

  .input_label:before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: calc(100% - 10px);
    border-bottom: 1px solid rgba(49, 49, 49, 0.45);
  }

  .input_label:after {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    margin-top: 2px;
    width: 100%;
    height: calc(100% - 10px);
    border-bottom: 3px solid #219e9a;
    -webkit-transform: translate3d(-100%, 0, 0);
    transform: translate3d(-100%, 0, 0);
    transition: -webkit-transform 0.3s;
    transition: transform 0.3s;
    transition: transform 0.3s, -webkit-transform 0.3s;
  }

  .input_label-content {
    position: absolute;
    width: 100%;
    bottom: 14px;
  }

  .input_label-search {
    position: relative;
    display: block;
    color: rgba(49, 49, 49, 0.45);
  }

  .input_label-search:after {
    content: '\f002';
    position: absolute;
    top: 7px;
    right: 5px;
    font-family: 'fontawesome';
  }

  .input_field:focus + .input_label::after,
  .input--filled .input_label::after {
    -webkit-transform: translate3d(0, 0, 0);
    transform: translate3d(0, 0, 0);
  }

  .input_field:focus + .input_label .input_label-content,
  .input--filled .input_label-content {
    -webkit-animation: anim-1 0.3s forwards;
    animation: anim-1 0.3s forwards;
  }

  @-webkit-keyframes anim-1 {
    50% {
      opacity: 0;
      -webkit-transform: translate3d(1em, 0, 0);
      transform: translate3d(1em, 0, 0);
    }
    51% {
      opacity: 0;
      -webkit-transform: translate3d(-1em, -40%, 0);
      transform: translate3d(-1em, -40%, 0);
    }
    100% {
      opacity: 1;
      -webkit-transform: translate3d(0, -40%, 0);
      transform: translate3d(0, -40%, 0);
    }
  }

  @keyframes anim-1 {
    50% {
      opacity: 0;
      bottom: 24px;
      -webkit-transform: translate3d(5em, 0, 0);
      transform: translate3d(5em, 0, 0);
    }
    51% {
      opacity: 0;
      bottom: 24px;
      -webkit-transform: translate3d(-5em, -40%, 0);
      transform: translate3d(-5em, -40%, 0);
    }
    100% {
      opacity: 1;
      bottom: 24px;
      -webkit-transform: translate3d(0, -40%, 0);
      transform: translate3d(0, -40%, 0);
    }
  }

  /*
   *  Menu
  */
  .main-navigation {
    position: relative;
    float: right;
    margin: 16px 0;
    padding: 0 20px;
    height: 40px;
    border-left: 1px solid #d2d6d5;
  }

  .main-navigation a {
    display: inline-block;
    line-height: 40px;
    vertical-align: middle;
    font-size: 14px;
    font-weight: bold;
    text-transform: uppercase;
    color: #313131;
  }

  .main-navigation a:before {
    content: '\f0c9';
    position: relative;
    display: inline-block;
    padding-right: 10px;
    font-family: "fontawesome";
    font-size: 18px;
    font-weight: 400;
    color: rgba(49, 49, 49, 0.45);
    vertical-align: middle;
  }

  /*
   *  Off Canvas Menu
  */
  .main-container {
    position: relative;
    overflow-x: hidden;
  }

  .main-container.prevent-scroll, .main-container.nav-menu-open {
    overflow: hidden;
  }

  .nav-menu {
    position: absolute;
    top: 0;
    right: 0;
    z-index: 100;
    visibility: visible;
    width: 300px;
    height: 100%;
    background: #219e9a;
    transition: all 0.5s;
    -webkit-transform: translate3d(100%, 0, 0);
    transform: translate3d(100%, 0, 0);
  }

  .nav-menu:after {
    position: absolute;
    top: 0;
    right: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.2);
    content: '';
    opacity: 1;
    transition: opacity 0.5s;
    display: none;
  }

  .nav-menu h2 {
    margin: 0;
    padding: 1em;
    color: rgba(0, 0, 0, 0.4);
    text-shadow: 0 0 1px rgba(0, 0, 0, 0.1);
    font-weight: 300;
    font-size: 2em;
  }

  .nav-menu ul {
    margin: 0;
    padding: 0;
    list-style-type: none;
  }

  .nav-menu li a {
    display: block;
    padding: 1em 1em 1em 1.2em;
    outline: none;
    box-shadow: inset 0 -1px rgba(0, 0, 0, 0.2);
    color: #f3efe0;
    text-transform: uppercase;
    text-shadow: 0 0 1px rgba(255, 255, 255, 0.1);
    letter-spacing: 1px;
    font-weight: 400;
    transition: background 0.3s, box-shadow 0.3s;
  }

  .nav-menu li:first-child a {
    box-shadow: inset 0 -1px rgba(0, 0, 0, 0.2), inset 0 1px rgba(0, 0, 0, 0.2);
  }

  .nav-menu li:hover {
    background: rgba(0, 0, 0, 0.2);
    box-shadow: inset 0 -1px rgba(0, 0, 0, 0);
    color: #fff;
  }

  .nav-menu-open .nav-menu {
    visibility: visible;
    -webkit-transform: translate3d(0, 0, 0);
    transform: translate3d(0, 0, 0);
  }

  .nav-menu-open .nav-menu:after {
    width: 0;
    height: 0;
    opacity: 0;
    transition: opacity 0.5s, width 0.1s 0.5s, height 0.1s 0.5s;
  }

  /*
   *  book List
  */
  .toolbar {
    margin-bottom: 30px;
    border-bottom: 1px solid #d2d6d5;
  }

  .toolbar select {
    margin-bottom: 7px;
    color: #313131;
    font-size: 14px;
    border: none;
    background-color: transparent;
  }

  .toolbar .filter-options {
    line-height: 40px;
  }

  .toolbar a.filter-item {
    margin-right: 16px;
    padding-bottom: 18px;
    font-size: 14px;
    color: #313131;
    border-bottom: 0px solid transparent;
    transition: all 250ms ease-out;
  }

  .toolbar a.filter-item:last-child {
    margin-right: 0;
  }

  .toolbar a.filter-item.active {
    padding-bottom: 15px;
    color: #219e9a;
    font-weight: bold;
    border-bottom: 3px solid #219e9a;
  }

  #grid {
    margin-bottom: 60px;
  }

  .book-item {
    margin: 15px 0;
    padding: 15px;
    list-style-type: none;
  }

  .book-item:after {
    content: '';
    position: absolute;
    top: 0;
    right: 15px;
    width: calc(100% - 105px);
    height: 100%;
    border-radius: 3px;
    box-shadow: 0 0 0 0 transparent;
    background-color: rgba(255, 255, 255, 0);
    z-index: -1;
    transition: all 250ms ease-out;
  }

  .book-item:hover:after {
    box-shadow: 0px 1px 1px 1px rgba(210, 214, 213, 0.5);
    background-color: #fff;
  }

  .book-item:hover .item-img img {
    box-shadow: 0px 0px 10px 0px rgba(49, 49, 49, 0.25);
  }

  .book-item:hover a.button {
    background-color: #219e9a;
  }

  .book-item:hover .bk-bookdefault {
    -webkit-transform: rotate3d(0, 1, 0, 25deg);
    transform: rotate3d(0, 1, 0, 25deg);
  }

  .book-item:hover .bk-back {
    opacity: 1;
  }

  .book-item .item-img {
    display: inline-block;
    float: left;
    padding-right: 30px;
  }

  .book-item .item-img img {
    box-shadow: 0 0 0 0 transparent;
    transition: all 250ms ease-out;
  }

  .book-item .item-details {
    padding-right: 30px;
  }

  .book-item h3 {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .book-item p:not(.author) {
    display: block;
    display: -webkit-box;
    height: 63px;
    /* Fallback for non-webkit */
    font-size: 15px;
    line-height: 1.4;
    -webkit-line-clamp: 3;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  /*
   *  Book Rotate
  */
  .bk-img {
    position: relative;
    display: inline-block;
    float: left;
    padding-right: 30px;
    list-style: none;
    /* Individual style & artwork */
  }

  .bk-img .bk-wrapper {
    position: relative;
    width: 150px;
    height: 215px;
    float: left;
    z-index: 1;
    -webkit-perspective: 1400px;
    perspective: 1400px;
  }

  .bk-img .bk-wrapper:last-child {
    margin-right: 0;
  }

  .bk-img .bk-book {
    position: absolute;
    width: 100%;
    height: 215px;
    -webkit-transform-style: preserve-3d;
    transform-style: preserve-3d;
    transition: -webkit-transform .5s;
    transition: transform .5s;
    transition: transform .5s, -webkit-transform .5s;
  }

  .bk-img .bk-book > div,
  .bk-img .bk-front > div {
    display: block;
    position: absolute;
  }

  .bk-img .bk-front {
    -webkit-transform-style: preserve-3d;
    transform-style: preserve-3d;
    -webkit-transform-origin: 0% 50%;
    transform-origin: 0% 50%;
    transition: -webkit-transform .5s;
    transition: transform .5s;
    transition: transform .5s, -webkit-transform .5s;
    -webkit-transform: translate3d(0, 0, 20px);
    transform: translate3d(0, 0, 20px);
    z-index: 10;
  }

  .bk-img .bk-front > div {
    z-index: 1;
    -webkit-backface-visibility: hidden;
    backface-visibility: hidden;
    -webkit-transform-style: preserve-3d;
    transform-style: preserve-3d;
    border-radius: 0 3px 3px 0;
    box-shadow: inset 4px 0 10px rgba(0, 0, 0, 0.1);
  }

  .bk-img .bk-front:after {
    content: '';
    position: absolute;
    top: 1px;
    bottom: 1px;
    left: -1px;
    width: 1px;
  }

  .bk-img .bk-front,
  .bk-img .bk-back,
  .bk-img .bk-front > div {
    width: 150px;
    height: 215px;
  }

  .bk-img .bk-left,
  .bk-img .bk-right {
    width: 40px;
    left: -20px;
  }

  .bk-img .bk-back {
    -webkit-transform: rotate3d(0, 1, 0, -180deg) translate3d(0, 0, 20px);
    transform: rotate3d(0, 1, 0, -180deg) translate3d(0, 0, 20px);
    box-shadow: 5px 7px 15px rgba(0, 0, 0, 0.3);
    border-radius: 3px 0 0 3px;
    opacity: 0;
    transition: opacity 250ms ease-out;
  }

  .bk-img .bk-back:after {
    content: '';
    position: absolute;
    top: 0;
    left: 10px;
    bottom: 0;
    width: 3px;
    background: rgba(0, 0, 0, 0.06);
    box-shadow: 1px 0 3px rgba(255, 255, 255, 0.1);
  }

  .bk-img .bk-left {
    height: 215px;
    -webkit-transform: rotate3d(0, 1, 0, -90deg);
    transform: rotate3d(0, 1, 0, -90deg);
  }

  .bk-img .bk-left h2 {
    width: 215px;
    height: 40px;
    -webkit-transform-origin: 0 0;
    transform-origin: 0 0;
    -webkit-transform: rotate(90deg) translateY(-40px);
    transform: rotate(90deg) translateY(-40px);
  }

  .bk-img .bk-cover {
    /*background-image: url(../images/1.png);*/
    background-repeat: no-repeat;
    background-position: 10px 40px;
  }

  .bk-img .bk-cover:after {
    content: '';
    position: absolute;
    top: 0;
    left: 10px;
    bottom: 0;
    width: 3px;
    background: rgba(0, 0, 0, 0.06);
    box-shadow: 1px 0 3px rgba(255, 255, 255, 0.1);
  }

  .bk-img .bk-cover {
    background-repeat: no-repeat;
    background-position: top left !important;
  }

  .bk-img .bk-front > div,
  .bk-img .bk-left {
    background-color: #219e9a;
  }

  /*
   *  Lightbox
  */
  .main-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: none;
    background-color: rgba(49, 49, 49, 0.65);
  }

  .main-overlay .overlay-full {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
  }

  .main-overlay .overlay-details {
    position: absolute;
    display: block;
    z-index: 1;
    top: 50%;
    left: 50%;
    width: 100%;
    max-width: 800px;
    max-height: 480px;
    padding: 40px;
    overflow: hidden;
    border-radius: 3px;
    background-color: #fff;
    box-shadow: 0px 2px 1px 2px rgba(0, 0, 0, 0.3);
    -webkit-transform: translate(-50%, -50%);
    transform: translate(-50%, -50%);
  }

  .main-overlay .overlay-image {
    display: inline-block;
    margin-right: 30px;
    max-width: 275px;
    vertical-align: top;
  }

  .main-overlay .overlay-image img {
    position: relative;
    display: inline-block;
    z-index: 1;
    box-shadow: -12px 12px 15px -5px rgba(0, 0, 0, 0.3);
  }

  .main-overlay .overlay-image .back-color {
    position: absolute;
    top: 0;
    left: 0;
    width: 275px;
    height: 100%;
    border-top-left-radius: 3px;
    border-bottom-left-radius: 3px;
    background-color: green;
  }

  .main-overlay .close-overlay-btn {
    position: absolute;
    top: 15px;
    right: 15px;
    width: 30px;
    height: 30px;
    opacity: 0.3;
    text-indent: -9999px;
    transition: opacity 250ms ease-out;
  }

  .main-overlay .close-overlay-btn:hover {
    opacity: 1;
  }

  .main-overlay .close-overlay-btn:before {
    content: ' ';
    position: absolute;
    left: 15px;
    width: 2px;
    height: 33px;
    background-color: #313131;
  }

  .main-overlay .close-overlay-btn:after {
    content: ' ';
    position: absolute;
    left: 15px;
    width: 2px;
    height: 33px;
    background-color: #313131;
  }

  .main-overlay .close-overlay-btn:before {
    -webkit-transform: rotate(45deg);
    transform: rotate(45deg);
  }

  .main-overlay .close-overlay-btn:after {
    -webkit-transform: rotate(-45deg);
    transform: rotate(-45deg);
  }

  .main-overlay .back-preview-btn {
    position: absolute;
    top: 7px;
    left: -34px;
    width: 30px;
    height: 30px;
    opacity: 0.3;
    text-indent: -9999px;
    transition: opacity 250ms ease-out;
  }

  .main-overlay .back-preview-btn:hover {
    opacity: 1;
  }

  .main-overlay .back-preview-btn:before {
    content: ' ';
    position: absolute;
    left: 15px;
    width: 2px;
    height: 15px;
    background-color: #313131;
  }

  .main-overlay .back-preview-btn:after {
    content: ' ';
    position: absolute;
    top: 10px;
    left: 15px;
    width: 2px;
    height: 15px;
    background-color: #313131;
  }

  .main-overlay .back-preview-btn:before {
    -webkit-transform: rotate(45deg);
    transform: rotate(45deg);
  }

  .main-overlay .back-preview-btn:after {
    -webkit-transform: rotate(-45deg);
    transform: rotate(-45deg);
  }

  .main-overlay .overlay-desc {
    display: inline-block;
    margin-top: -400px;
    width: calc(100% - 320px);
    vertical-align: top;
    transition: all 500ms ease-out;
  }

  .main-overlay .overlay-desc.activated {
    display: inline-block;
    margin-top: 0;
  }

  .main-overlay .overlay-preview {
    position: relative;
    display: inline-block;
    float: right;
    margin-top: 40px;
    width: calc(100% - 310px);
    vertical-align: top;
    transition: all 500ms ease-out;
  }

  .main-overlay .overlay-preview.activated {
    margin-top: -430px;
  }

  .main-overlay .preview-content {
    padding-right: 24px;
    padding-top: 10px;
    display: block;
    display: -webkit-box;
    height: 360px;
    /* Fallback for non-webkit */
    font-size: 15px;
    line-height: 1.5;
    -webkit-line-clamp: 16;
    overflow: hidden;
    text-overflow: ellipsis;
    overflow: scroll;
    text-overflow: clip;
    font-weight: 400;
  }

  .main-overlay .preview-content h5,
  .main-overlay .preview-content p {
    font-family: 'Roboto Slab', serif;
  }

  .main-overlay .preview-content h5 {
    font-weight: bold;
  }

  .main-overlay .preview-content p {
    font-weight: normal;
  }

  .main-overlay .preview-content:before {
    content: '';
    position: absolute;
    left: 0;
    top: 40px;
    width: 100%;
    height: 30px;
    background: rgba(255, 255, 255, 0);
    background: linear-gradient(to top, rgba(255, 255, 255, 0) 0%, rgba(255, 255, 255, 0) 0%, rgba(255, 255, 255, 0.8) 30%, white 80%);
  }

  .main-overlay .preview-content:after {
    content: '';
    position: absolute;
    left: 0;
    bottom: 0;
    width: 100%;
    height: 50px;
    background: rgba(255, 255, 255, 0);
    background: linear-gradient(to bottom, rgba(255, 255, 255, 0) 0%, rgba(255, 255, 255, 0) 0%, rgba(255, 255, 255, 0.8) 30%, white 80%);
  }

  .overlay-details {
    display: none;
  }

  /*
   *  Footer
  */
  #footer {
    margin-top: 60px;
    padding: 15px 0 20px;
    border-top: 1px solid #fff;
    background-color: rgba(49, 49, 49, 0.5);
  }

  #footer div,
  #footer a,
  #footer p {
    color: #fff;
    font-size: 12px;
    text-align: center;
  }

</style>
<template>
  <mdb-container fluid>
    <div id="main-container" class="main-container nav-effect-1">
      <div class="main clearfix">
        <!-- Main Section -->
        <div class="page-container">
          <section id="book_list">
            <div class="toolbar row">
              <div class="filter-options small-12 medium-9 columns">
                <button class="filter-item active" data-group="all">All Categories</button>
                <button class="filter-item" data-group="fantasy">Fantasy</button>
                <a href="#" class="filter-item" data-group="sci-fi">Sci-Fi</a>
                <a href="#" class="filter-item" data-group="classic">Classics</a>
                <a href="#" class="filter-item" data-group="fairy">Fairy Tale</a>
                <a href="#" class="filter-item" data-group="young">Young Adult</a>
              </div>

              <div class="small-12 medium-3 columns">
                <select class="sort-options">
                  <option value="" disabled selected>Sort by</option>
                  <option value="">Featured</option>
                  <option value="title">Alphabetical</option>
                  <option value="date-created">Published</option>
                </select>
              </div>
            </div>
            <div class="grid-shuffle">
              <ul id="grid" class="row">
                <li class="book-item small-12 medium-6 columns" data-groups='["classic"]'
                    data-date-created='1937' data-title='Of Mice and Men' data-color='#fcc278'>
                  <div class="bk-img">
                    <div class="bk-wrapper">
                      <div class="bk-book bk-bookdefault">
                        <div class="bk-front">
                          <div class="bk-cover"
                               style="background-image: url('http://interactivejoe.com/book-viewer/assets/images/bk_1-small.jpg')"></div>
                        </div>
                        <div class="bk-back"></div>
                        <div class="bk-left"></div>
                      </div>
                    </div>
                  </div>
                  <div class="item-details">
                    <h3 class="book-item_title">Of Mice and Men</h3>
                    <p class="author">by John Steinbeck &bull; 1937</p>
                    <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam tellus nisi, eget
                      pulvinar in, molestie et arcu.</p>
                    <a href="#" class="button ">Details</a>
                  </div>

                  <div class="overlay-details">
                    <a href="#" class="close-overlay-btn">Close</a>
                    <div class="overlay-image">
                      <img src="http://interactivejoe.com/book-viewer/assets/images/bk_1-large.jpg"
                           alt="Book Cover">
                      <div class="back-color"></div>
                    </div>
                    <div class="overlay-desc activated">
                      <h2 class="overlay_title">Of Mice and Men</h2>
                      <p class="author">by John Steinbeck</p>
                      <p class="published">1937</p>
                      <p class="synopsis">Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam
                        luctus convallis sem. Nunc sed tristique augue. Aenean at nulla vel lacus
                        volutpat bibendum vitae ut nibh. Aliquam eu metus et purus rutrum malesuada.
                        Aenean in auctor mauris,
                        non vulputate libero. Nullam auctor, purus ut cursus convallis, lectus tellus
                        dignissim lectus, id tempor ipsum leo ut nulla. Vestibulum vitae elit erat.</p>
                      <a href="#" class="button preview">Preview</a>
                    </div>
                    <div class="overlay-preview">
                      <a href="#" class="back-preview-btn">Back</a>
                      <h4 class="preview-title">Preview</h4>
                      <div class="preview-content">
                        <h5>Chapter 1</h5>
                        <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam luctus
                          convallis sem. Nunc sed tristique augue. Aenean at nulla vel lacus volutpat
                          bibendum vitae ut nibh. Aliquam eu metus et purus rutrum malesuada. Aenean
                          in auctor
                          mauris, non vulputate libero. Nullam auctor, purus ut cursus convallis,
                          lectus tellus dignissim lectus, id tempor ipsum leo ut nulla. Vestibulum
                          vitae elit erat.</p>
                        <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam luctus
                          convallis sem. Nunc sed tristique augue. Aenean at nulla vel lacus volutpat
                          bibendum vitae ut nibh. Aliquam eu metus et purus rutrum malesuada. Aenean
                          in auctor
                          mauris, non vulputate libero. Nullam auctor, purus ut cursus convallis,
                          lectus tellus dignissim lectus, id tempor ipsum leo ut nulla. Vestibulum
                          vitae elit erat.</p>
                        <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam luctus
                          convallis sem. Nunc sed tristique augue. Aenean at nulla vel lacus volutpat
                          bibendum vitae ut nibh. Aliquam eu metus et purus rutrum malesuada. Aenean
                          in auctor
                          mauris, non vulputate libero. Nullam auctor, purus ut cursus convallis,
                          lectus tellus dignissim lectus, id tempor ipsum leo ut nulla. Vestibulum
                          vitae elit erat.<br>&nbsp;<br>&nbsp;</p>
                      </div>
                    </div>
                  </div>
                </li>
              </ul>
            </div>

          </section>

        </div>
      </div>
      <!-- /main -->

      <div class="main-overlay">
        <div class="overlay-full"></div>
      </div>

    </div>
  </mdb-container>
</template>

<script>
  import {mdbCol, mdbContainer, mdbRow,} from 'mdbvue';
  import Shuffle from 'shufflejs'

  export default {
    name: "BookItem",
    components: {
      mdbCol,
      mdbContainer,
      mdbRow
    },
    created() {
      const bookList = function (element) {
        this.element = element;

        this.shuffle = new Shuffle(element, {
          itemSelector: '.book-item',
        });
        this._activeFilters = [];
        this.addFilterButtons();
        this.addSorting();
        this.addSearchFilter();
        this.mode = 'exclusive';
      };


      bookList.prototype.toArray = function (arrayLike) {
        return Array.prototype.slice.call(arrayLike);
      };


      // Catagory Filter Functions
      // Toggle mode for the Catagory filters
      bookList.prototype.toggleMode = function () {
        if (this.mode === 'additive') {
          this.mode = 'exclusive';
        } else {
          this.mode = 'additive';
        }
      };

      // Filter buttons eventlisteners
      bookList.prototype.addFilterButtons = function () {
        var options = document.querySelector('.filter-options');
        if (!options) {
          return;
        }
        var filterButtons = this.toArray(options.children);
        console.log(filterButtons);
        filterButtons.forEach(function (button) {
          button.addEventListener('click', this._handleFilterClick.bind(this), false);
        }, this);
      };

      // Function for the Catagory Filter
      bookList.prototype._handleFilterClick = function (evt) {
        var btn = evt.currentTarget;
        var isActive = btn.classList.contains('active');
        var btnGroup = btn.getAttribute('data-group');

        if (this.mode === 'additive') {
          if (isActive) {
            this._activeFilters.splice(this._activeFilters.indexOf(btnGroup));
          } else {
            this._activeFilters.push(btnGroup);
          }

          btn.classList.toggle('active');
          this.shuffle.filter(this._activeFilters);

        } else {
          this._removeActiveClassFromChildren(btn.parentNode);

          var filterGroup;
          if (isActive) {
            btn.classList.remove('active');
            filterGroup = Shuffle.ALL_ITEMS;
          } else {
            btn.classList.add('active');
            filterGroup = btnGroup;
          }

          this.shuffle.filter(filterGroup);
        }
      };

      // Remove classes for active states
      bookList.prototype._removeActiveClassFromChildren = function (parent) {
        var children = parent.children;
        for (var i = children.length - 1; i >= 0; i--) {
          children[i].classList.remove('active');
        }
      };

      // Sort By
      // Watching for the select box to change to run
      bookList.prototype.addSorting = function () {
        var menu = document.querySelector('.sort-options');
        if (!menu) {
          return;
        }
        menu.addEventListener('change', this._handleSortChange.bind(this));
      };

      // Sort By function Handler runs on change
      bookList.prototype._handleSortChange = function (evt) {
        var value = evt.target.value;
        var options = {};

        function sortByDate(element) {
          return element.getAttribute('data-created');
        }

        function sortByTitle(element) {
          return element.getAttribute('data-title').toLowerCase();
        }

        if (value === 'date-created') {
          options = {
            reverse: true,
            by: sortByDate,
          };
        } else if (value === 'title') {
          options = {
            by: sortByTitle,
          };
        }

        this.shuffle.sort(options);
      };

      // Searching the Data-attribute Title
      // Advanced filtering
      // Waiting for input into the search field
      bookList.prototype.addSearchFilter = function () {
        var searchInput = document.querySelector('.shuffle-search');
        if (!searchInput) {
          return;
        }
        searchInput.addEventListener('keyup', this._handleSearchKeyup.bind(this));
      };

      // Search function Handler to search list
      bookList.prototype._handleSearchKeyup = function (evt) {
        var searchInput = document.querySelector('.shuffle-search');
        var searchText = evt.target.value.toLowerCase();

        // Check if Search input has value to toggle class
        if (searchInput && searchInput.value) {
          $('.catalog-search').addClass('input--filled');
        } else {
          $('.catalog-search').removeClass('input--filled');
        }

        this.shuffle.filter(function (element, shuffle) {

          // If there is a current filter applied, ignore elements that don't match it.
          if (shuffle.group !== Shuffle.ALL_ITEMS) {
            // Get the item's groups.
            var groups = JSON.parse(element.getAttribute('data-groups'));
            var isElementInCurrentGroup = groups.indexOf(shuffle.group) !== -1;

            // Only search elements in the current group
            if (!isElementInCurrentGroup) {
              return false;
            }
          }

          var titleElement = element.querySelector('.book-item_title');
          var titleText = titleElement.textContent.toLowerCase().trim();

          return titleText.indexOf(searchText) !== -1;
        });
      };

      // Wait till dom load to start the Shuffle js funtionality
      document.addEventListener('DOMContentLoaded', function () {
        window.book_list = new bookList(document.getElementById('grid'));
      });
    }
  }
</script>
