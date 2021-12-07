<template>
  <div><div v-for="product in products" :key="product.id">
      <img :src="product.image" class="image"/>
      {{ product.name }}
      {{ product.price }}
      {{ product.quantity }}
      <button class="increment" v-on:click="increment(product)">+</button>
      <button class="increment" v-on:click="decrement(product)">-</button>
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
      await axios.get("http://localhost:5000/api/basket/products")
        .then((response) => {
          this.products = response.basketProducts
        })
    },
    async increment(product){
      const response = await axios.post("http://localhost:5000/api/basket/products/increment/" + product.id, {quantity: product.quantity})
      this.getBasketProducts()
      return response
    },
    async decrement(product){
      const response = await axios.post("http://localhost:5000/api/basket/products/increment/" + product.id, {quantity: product.quantity})
      this.getBasketProducts()
      return response
    } 
  },
  created() {
    this.getBasketProducts()
  }
}
</script>