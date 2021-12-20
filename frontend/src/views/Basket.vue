<template>
  <div><div v-for="product in products" :key="product.id">
      <img class="image" :src="product.image"/>
      Name: {{ product.name }} &emsp;
      Price: {{ product.price }} &emsp;
      Quantity: {{ product.quantity }} &emsp;
      <button class="increment" v-on:click="increment(product.id)"> + </button> &ensp;
      <button class="decrement" v-on:click="decrement(product.id)"> - </button> &ensp;
      <button class="delete" v-on:click="deleteproduct(product.id)">Delete</button>
    </div></div>
</template>

<script>
import axios from "axios"

export default {
  name: "Basket",
  data() {
    return{
      products: []
    }
  },
  methods: {
    async getBasketProducts() {
      await axios.get("http://localhost:3000/api/basket")
        .then((response) => {
          this.products = response.data
        })
    },
    async increment(productId){
      await axios.post("http://localhost:3000/api/basket/products/increment/" + productId)
      this.getBasketProducts()
    },
    async decrement(productId){
      await axios.post("http://localhost:3000/api/basket/products/decrement/" + productId)
      this.getBasketProducts()
    },
    async deleteproduct(productId){
      await axios.post("http://localhost:3000/api/basket/delete/" + productId)
      this.getBasketProducts()
    },
    
  },
  created() {
    this.getBasketProducts()
  }
}
</script>

<style>
.image{
  width: 15rem;
}
</style>