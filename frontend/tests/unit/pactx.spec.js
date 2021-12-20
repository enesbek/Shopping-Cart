"use-strict"

import { pactWith } from "jest-pact"
import { Matchers } from "@pact-foundation/pact"
import { getProducts, increment, decrement} from "@/../services/shoppingCart.service.js"
const { eachLike } = Matchers

pactWith(
    { consumer: "ShoppingCartConsumer", provider: "ShoppingCartProvider", cors: true, },
    provider => {
        describe("Get Products", () => {
            const PRODUCT = 
            {  
                id: "1", 
                name: "iphone", 
                price: 1000, 
                image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/ip-3_large.jpg", 
                quantity: 1,
                basket: false 
            }

            const getProdResponse = {
                status: 200,
                headers: { "Content-Type": "application/json"},
                body: [PRODUCT]
            }

            const getProdRequest = {
                uponReceiving: "a request for products",
                withRequest: {
                    method: "GET",
                    path: "/api/allproducts",
                    headers: {
                        Accept: "application/json",
                    },
                },

            }

            beforeEach(() => {
                const interaction = {
                    state: "i have products",
                    ...getProdRequest,
                    willRespondWith: getProdResponse,
                }
                return provider.addInteraction(interaction)
            })

            it("returns a succesful body", () => {
                return getProducts({
                    url: provider.mockService.baseUrl,
                }).then(response => {
                    expect(response).toEqual([PRODUCT])
                })
            })
        })

        describe("increment quantity of product", () => {
            const INCREMENT = { id: 1, 
                                name: "iphone", 
                                price: 1000, 
                                image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/ip-3_large.jpg", 
                                quantity: 2 
                            }

            const incrementResponse = {
                status: 200,
                headers: { "Content-Type": "application/json" },
                body: INCREMENT
            }

            const incrementRequest = {
                uponReceiving: "a request increment quantity of product",
                withRequest: {
                    method: "POST",
                    path: "/api/basket/products/increment",
                    data: { id: 1 },
                    headers: {
                        Accept: "application/json",
                    },
                },
            }

            beforeEach(() => {
                return provider.addInteraction({
                    state: "increment product",
                    ...incrementRequest,
                    willRespondWith: incrementResponse
                })
            })

            it("returns a succesful body", () => {
                return increment({
                    url: provider.mockService.baseUrl
                }).then( response => {
                    expect(response).toEqual(INCREMENT)
                })
            })
        })

        describe("decrement quantity of product", () => {
            const DECREMENT = { id: 1, 
                                name: "iphone", 
                                price: 1000, 
                                image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/ip-3_large.jpg", 
                                quantity: 1 
                            }

            const decrementResponse = {
                status: 200,
                headers: { "Content-Type": "application/json" },
                body: DECREMENT
            }

            const decrementRequest = {
                uponReceiving: "a request decrement quantity of product",
                withRequest: {
                    method: "POST",
                    path: "/api/basket/products/decrement",
                    data: { id: 1 },
                    headers: {
                        Accept: "application/json",
                    },
                },
            }

            beforeEach(() => {
                return provider.addInteraction({
                    state: "decrement product",
                    ...decrementRequest,
                    willRespondWith: decrementResponse
                })
            })

            it("returns a succesful body", () => {
                return decrement({
                    url: provider.mockService.baseUrl
                }).then(response => {
                    expect(response).toEqual(DECREMENT)
                })
            })
        })
    }
)