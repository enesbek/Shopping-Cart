import { shallowMount } from '@vue/test-utils';
import Products from "@/components/Products";
import axios from "axios";

jest.mock('axios')

describe("Implementation test for Products.vue with succesful HTTP GET", () => {
  var wrapper = null

  beforeEach(() => {
    const responseGet = { products: 
      [
        {
          name: "iphone",
          price: 1000,
          image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/ip-3_large.jpg"
        },
        {
        name: "samsung",
        price: 900,
        image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/samsung/thumb/121859-4_large.jpg"
        }
      ]
    }

    axios.get.mockResolvedValue(responseGet)

    wrapper = shallowMount(Products)
  })

  afterEach(() => {
    jest.resetModules()
    jest.clearAllMocks()
  })

  it("does load the products data when a succesful HTTP GET occurs", () => {
    //The unit test starts by calling the getAllProducts() function
    wrapper.vm.getAllProducts()

    expect(axios.get).toHaveBeenCalledWith('http://localhost:5000/api/allproducts')

    wrapper.vm.$nextTick().then(function () {
      const expectedMockList = [
        {
          name: "iphone",
          price: 1000,
          image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/ip-3_large.jpg"
        },
        {
        name: "samsung",
        price: 900,
        image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/samsung/thumb/121859-4_large.jpg"
        }
      ]
      expect(wrapper.vm.products).toStrictEqual(expectedMockList)
    })
  })
})