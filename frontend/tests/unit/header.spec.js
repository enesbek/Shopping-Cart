import { shallowMount } from '@vue/test-utils'
import Header from '@/components/Header.vue'

describe("header component", () => {
    let wrapper = null

    beforeEach(() => {
        wrapper = shallowMount(Header, {
            propsData: {
                title: "SHOPPING CART",
            }
        })
    })

    afterEach(() => {
        wrapper = null;
    })

    it("render component", () => {
        expect(wrapper.vm.$options.name).toMatch("Header")
    });
    
    it("control the routes", () => {
        expect(wrapper.findAll('button').length).toEqual(3)
        expect(wrapper.findAll('button').at(0).text()).toMatch("SHOPPING CART")
        expect(wrapper.findAll('button').at(1).text()).toMatch("CREATE")
        expect(wrapper.findAll('button').at(2).text()).toMatch("BASKET")
    });
});