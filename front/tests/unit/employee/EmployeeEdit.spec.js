import EmployeeEdit from '@/components/employee/EmployeeEdit.vue';
import { shallowMount, createLocalVue } from '@vue/test-utils';

const localVue = createLocalVue();

describe('EmployeeEdit.vue', () => {
    let wrapper;
    let dataResponse = [{id:1, name:'support 1'}, {id: 2, name: 'support 2'}]
    let dataEmployee = {name: 'Luiza Labs', email: 'vemserfeliz@luizalabs.com', department: 'support 1'}

    beforeEach(() => {
        fetch.mockResponseOnce(JSON.stringify(dataResponse))
        wrapper = shallowMount(EmployeeEdit, {
            localVue,
        });
        wrapper.setData({form: dataEmployee})
        wrapper.setData({departments: dataResponse})
        wrapper.setProps({employee: dataEmployee})
    });

    it('test data', () => {
        expect(wrapper.vm.departments).toEqual(dataResponse)
        expect(wrapper.vm.$el.textContent).toMatch(dataResponse[0].name)
    })

    it('test form', () => {
        wrapper.vm.mergeDepartment()
        dataEmployee.department = 1
        expect(wrapper.vm.form).toEqual(dataEmployee)
    })
})