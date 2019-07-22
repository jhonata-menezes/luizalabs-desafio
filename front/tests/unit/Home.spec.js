import Home from '@/components/Home.vue';
import { shallowMount, createLocalVue } from '@vue/test-utils';

const localVue = createLocalVue();

describe('Home.vue', () => {
    let wrapper;
    let dataResponse = [{id:1, name:'LuizaLabs', department:'support', email:'vemserfeliz@luizalabs.com'}]

    beforeEach(() => {
        fetch.mockResponseOnce(JSON.stringify(dataResponse))
        wrapper = shallowMount(Home, {
            localVue,
        });
        wrapper.setData({employees: dataResponse})
    });

    it('test data', () => {
        expect(wrapper.vm.employees).toEqual(dataResponse)
        expect(wrapper.vm.$el.textContent).toMatch(dataResponse[0].email)
    })
})