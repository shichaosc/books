<!--suppress ALL -->
<style scoped>

  .row {
    padding-bottom: 15px;
  }

  .card-deck .card {
    margin: 10px;
  }

  .container {
    margin-top: 10px;
  }

  .bg {
    /* The image used */
    /*background-image: url("https://qncweb.ktvsky.com/20190323/leimeng/7be0d6d7436808e629074c4b072556de.jpg");*/
    /*background-color: #fffcc8;*/
    /* Full height */
    height: 100%;

    /* Center and scale the image nicely */
    background-position: center;
    background-repeat: no-repeat;
    background-size: cover;
  }


</style>
<template>
  <div class="bg">
    <mdb-container>
      <mdb-row v-for="tag in tags">
        <mdb-card-group deck>
          <mdb-card class="hoverable" v-for="t in tag">
            <mdb-view hover>
              <a href="#!">
                <mdb-card-image class=".justify-content-center" :src="t.Url"
                                alt="Card image cap">
                </mdb-card-image>
                <mdb-mask flex-center waves overlay="white-slight"></mdb-mask>
              </a>
            </mdb-view>
            <mdb-card-body>
              <mdb-card-title>{{t.Name}}</mdb-card-title>
              <mdb-card-text>
                {{t.Description}}
              </mdb-card-text>
              <mdb-btn color="primary" v-on:click="goBook(t.Id)">选 择 分 类</mdb-btn>
            </mdb-card-body>
          </mdb-card>
        </mdb-card-group>
      </mdb-row>
    </mdb-container>
  </div>

</template>
<script>
  import {
    mdbBtn,
    mdbCard,
    mdbCardAvatar,
    mdbCardBody,
    mdbCardFooter,
    mdbCardGroup,
    mdbCardHeader,
    mdbCardImage,
    mdbCardText,
    mdbCardTitle,
    mdbCardUp,
    mdbCol,
    mdbContainer,
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
    mdbRow,
    mdbView
  } from 'mdbvue';
  import ajax from "../utils/requests";

  export default {
    name: 'Home',
    data() {
      return {
        tags: [],
        total: 0
      }
    },
    methods: {
      goBook(bookId) {
        this.$router.push({path: `/book/list/${bookId}`})
      }
    },
    mounted() {
      ajax.get('/tag', {}).then(
        res => {
          const data = res.Data;
          let inner = []
          let first = undefined
          for (let i in data) {
            if (i == 0) {
              first = data[i]
            } else {
              inner.push(data[i])
            }
            if (i != 0 && i % 3 == 0) {
              this.tags.push(inner)
              inner = []
            }
          }
          if (inner) {
            inner.push(first)
            this.tags.push(inner)
          }
          this.total = res.Total;
        }
      )
    },
    components: {
      mdbContainer,
      mdbRow,
      mdbCol,
      mdbCard,
      mdbCardImage,
      mdbCardHeader,
      mdbCardBody,
      mdbCardTitle,
      mdbCardText,
      mdbCardFooter,
      mdbCardUp,
      mdbCardAvatar,
      mdbCardGroup,
      mdbBtn,
      mdbView,
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
      mdbDropdownItem
    }
  }

</script>