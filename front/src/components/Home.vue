<template>
  <div>
    <section class="section">
      <div>
        <h1 class="title is-4">Employee Management</h1>
      </div>
      <br/>
      <div>
        <div class="is-grouped field">
          <div class="control">
            <a class="button is-success" @click="addEmployee()">
              New Employee
            </a>
          </div>

          <div class="control">
            <a class="button is-light" @click="components.departmentAdd.active = true">
              New Department
            </a>
          </div>
        </div>
      </div>
      <div class="table-container">
        <table class="table is-striped is-fullwidth">
          <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>E-mail</th>
            <th>Department</th>
            <th>-</th>
            <th>-</th>
          </tr>
          </thead>

          <tbody>
          <tr v-for="employee of employees" :key="employee.id">
            <td>{{employee.id}}</td>
            <td>{{employee.name}}</td>
            <td>{{employee.email}}</td>
            <td>{{employee.department}}</td>
            <td>
              <a class="button is-warning" @click="components.employeeEdit.employee = employee; components.employeeEdit.active = true;">Edit</a>
            </td>
            <td>
              <a class="button is-danger" @click="deleteEmployee(employee.id)">Delete</a>
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </section>

    <!-- modal employee -->
    <div class="modal" :class="{'is-active': components.employeeAdd.active}" v-if="components.employeeAdd.active">
      <div class="modal-background"></div>
      <div class="modal-card">
        <header class="modal-card-head">
          <p class="modal-card-title">New Employee</p>
          <button class="delete" aria-label="close" @click="components.employeeAdd.active = false"></button>
        </header>
        <section class="modal-card-body">
          <div class="">
            <employee-add @close="components.employeeAdd.active = false; getEmployees()"/>
          </div>
        </section>
      </div>
    </div>

    <!-- modal employee edit -->
    <div class="modal" :class="{'is-active': components.employeeEdit.active}" v-if="components.employeeEdit.active">
      <div class="modal-background"></div>
      <div class="modal-card">
        <header class="modal-card-head">
          <p class="modal-card-title">Edit Employee</p>
          <button class="delete" aria-label="close" @click="components.employeeEdit.active = false"></button>
        </header>
        <section class="modal-card-body">
          <div class="">
            <employee-edit @close="components.employeeEdit.active = false; getEmployees()" :employee="components.employeeEdit.employee"/>
          </div>
        </section>
      </div>
    </div>

    <!-- modal department -->
    <div class="modal" :class="{'is-active': components.departmentAdd.active}">
      <div class="modal-background"></div>
      <div class="modal-card">
        <header class="modal-card-head">
          <p class="modal-card-title">New Department</p>
          <button class="delete" aria-label="close" @click="components.departmentAdd.active = false"></button>
        </header>
        <section class="modal-card-body">
          <div class="">
            <department-add @close="components.departmentAdd.active = false"/>
          </div>
        </section>
      </div>
    </div>
  </div>

</template>

<script>
  import EmployeeAdd from '@/components/employee/EmployeeAdd';
  import EmployeeEdit from '@/components/employee/EmployeeEdit';
  import DepartmentAdd from '@/components/department/DepartmentAdd';

export default {
  name: 'Home',
  components: {EmployeeAdd, DepartmentAdd, EmployeeEdit},
  data () {
    return {
      employees: [],

      components: {
        employeeAdd: {
          active: false
        },

        employeeEdit: {
          active: false,
          employee: {}
        },

        departmentAdd: {
          active: false
        }
      }
    }
  },

  methods: {
    getEmployees () {
      fetch('http://localhost:8000/employee').then(response => {
        if (response.status === 200) {
          response.json().then(data => this.employees = data)
        }
      }).catch(err => {
        console.log(err);
      })
    },

    deleteEmployee (id) {
      fetch(`http://localhost:8000/employee/${id}`, {method: 'delete'}).then(response => {
        if (response.status === 200) {
          this.getEmployees()
        }
      }).catch(err => {
        console.log(err);
      })
    },

    addEmployee() {
      this.components.employeeAdd.active = true;
    }
  },

  mounted() {
    this.getEmployees();
  }
}
</script>

<style scoped>
</style>
