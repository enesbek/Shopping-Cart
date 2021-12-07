import { shallowMount } from '@vue/test-utils';
import Basket from "@/components/Basket";
import axios from "axios";

jest.mock("axios")

describe("Implementation test for Basket.vue with succesful HTTP GET", () => {
    let wrapper = null

    beforeEach(() => {
        const responseGet = { basketProducts:
            [
                {
                    name: "iphone",
                    price: 1000,
                    image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/ip-3_large.jpg",
                    quantity: 1
                },
            ]
        }

        axios.get.mockResolvedValue(responseGet)

        wrapper = shallowMount(Basket)
    })

    afterEach(() => {
        jest.resetModules()
        jest.clearAllMocks()
    })

    it("does load the products data when a succesful HTTP GET occurs", () => {
        wrapper.vm.getBasketProducts()
    
        expect(axios.get).toHaveBeenCalledWith('http://localhost:5000/api/basket/products')
        
        wrapper.vm.$nextTick().then(function () {
            const expectedMockList = [
                {
                    name: "iphone",
                    price: 1000,
                    image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/ip-3_large.jpg",
                    quantity: 1
                }
            ]
            expect(wrapper.vm.products).toStrictEqual(expectedMockList)
        })
    })
    
    it("fetches succesfully increment quantity of product", async () => {
        let testId = 1
        const expectedMockList = [
            {
                name: "iphone",
                price: 1000,
                image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/ip-3_large.jpg",
                quantity: 2
            }
        ]
        try{
            axios.post.mockResolvedValue(expectedMockList)
        } catch (error){
            console.log(error)
        }

        const response = await wrapper.vm.increment(testId)
        expect(response).toBe(expectedMockList)
    })

    it("fetches succesfully decrement quantity of product", async () => {
        let testId = 1
        const expectedMockList = [
            {
                name: "iphone",
                price: 1000,
                image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/ip-3_large.jpg",
                quantity: 1
            }
        ]
        try{
            axios.post.mockResolvedValue(expectedMockList)
        } catch (error){
            console.log(error)
        }

        const response = await wrapper.vm.decrement(testId)
        expect(response).toBe(expectedMockList)
    })
    
})
