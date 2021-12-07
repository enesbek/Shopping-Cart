<template>
  <div id="products" class="content">
    <div v-for="product in products" :key="product.id">
      <img :src="product.image" class="image"/>
      {{ product.name }}
      {{ product.price }}
      <button v-on:click="addItemToCart(product.name)" class="adButton">Add To Cart</button>
    </div>
  </div>
</template>

<script>
import axios from "axios"

export default {
  name: "Products",
  data() {
    return{
      products: []
    }
  },
  created() {
    this.products = this.getAllProducts()
  },
  methods: {
    async getAllProducts() {
      await axios.get("http://localhost:5000/api/allproducts")
        .then((response) => {
          this.products = response.products
        })
    },
    async addItemToCart(name) {
      await fetch("http://localhost:5000/api/addToCart/" + name)
    }
  }
}
</script>