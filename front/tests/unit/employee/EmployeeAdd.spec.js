import EmployeeAdd from '@/components/employee/EmployeeAdd.vue';
import { shallowMount, createLocalVue } from '@vue/test-utils';

const localVue = createLocalVue();

describe('EmployeeAdd.vue', () => {
    let wrapper;
    let dataResponse = [{id:1, name:'support 1'}, {id: 2, name: 'support 2'}]
    let dataForm = {name: 'Luiza Labs', email: 'vemserfeliz@luizalabs.com', department: 1}

    beforeEach(() => {
        fetch.mockResponseOnce(JSON.stringify(dataResponse))
        wrapper = shallowMount(EmployeeAdd, {
            localVue,
        });
        wrapper.setData({departments: dataResponse})
    });

    it('test data', () => {
        expect(wrapper.vm.departments).toEqual(dataResponse)
        expect(wrapper.vm.$el.textContent).toMatch(dataResponse[0].name)
    })

    it('test form', () => {
        wrapper.setData({form: dataForm})
        expect(wrapper.vm.form).toEqual(dataForm)
    })
})