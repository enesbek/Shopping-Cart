<template>
  <header>
    <button>{{ title }}</button>
    <button>CREATE</button>
    <button>BASKET</button>
  </header>
</template>

<script>
export default{
  name: "Header",
  props: {
    title: String
  }
}
</script>

<style scoped>
header {
  margin: auto;
}

header h1 {
  font-size: 2.5em;
}
</style>
