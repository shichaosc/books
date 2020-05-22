<!--suppress ALL -->
<style scoped>
</style>
<template>


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

  export default {
    name: 'Search',
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