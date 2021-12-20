import axios from "axios"

exports.getProducts = endpoint => {
  const url = endpoint.url

  return axios
    .request({
      method: "GET",
      baseURL: url,
      url: "/api/allproducts",
      headers: { Accept: "application/json" },
    })
    .then(response => response.data)
}

exports.increment = endpoint => {
  const url = endpoint.url
  return axios
    .request({
      method: "POST",
      baseURL: url,
      url: "/api/basket/products/increment",
      data: {
        quantity: { id: "1", name: "iphone", price: 1000, image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/ip-3_large.jpg", quantity: 1 }
      },
      headers: { Accept: "application/json" },
    })
    .then(response => response.data)
}

exports.decrement = endpoint => {
  const url = endpoint.url
  return axios
    .request({
      method: "POST",
      baseURL: url,
      url: "/api/basket/products/decrement",
      data: {
        quantity: { id: "1", name: "iphone", price: 1000, image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/ip-3_large.jpg", quantity: 1 }
      },
      headers: { Accept: "application/json" },
    })
    .then(response => response.data)
}