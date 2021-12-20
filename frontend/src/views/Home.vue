<template>
  <div>
    <div id="products" class="products">
      <div v-for="product in products" :key="product.id">
        <img :src="product.image" class="image"/> <br>
        <p>{{ product.name }}</p>
        <p>${{ product.price }}</p>
        <button v-on:click="addItemToCart(product.id)" class="adButton">Add To Cart</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios"

export default {
  name: "Home",
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
      await axios.get("http://localhost:3000/api/allproducts")
        .then((response) => {
          this.products = response.data
        })
    },
    async addItemToCart(id) {
      await axios.post("http://localhost:3000/api/basket/add/" + id)
    }
  }
}
</script>

<style>


.products{
  display:grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 10px;
  grid-auto-rows: minmax(100px, auto);
}
.image {
  width: 15rem;
}
</style>
