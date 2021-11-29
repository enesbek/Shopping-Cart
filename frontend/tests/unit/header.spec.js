import { shallowMount } from '@vue/test-utils'
import Header from '@/components/Header.vue'

describe("displays message", () => {
    it("render component", () => {
        // mount() returns a wrapped Vue component we can interact with
        const wrapper = shallowMount(Header, {
            propsData: {
                title: "SHOPPING CART"
            }
        })
        expect(wrapper.vm.$options.name).toMatch("Header")

        expect(wrapper.text()).toMatch("SHOPPING CART")
    });
    
})