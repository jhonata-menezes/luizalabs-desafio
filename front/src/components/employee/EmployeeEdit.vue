<template>
    <div>

        <div>
            <!-- name -->
            <div class="field is-horizontal">
                <div class="field-label">
                    <label class="label">Name:</label>
                </div>

                <div class="field-body">
                    <div class="field">
                        <div class="control">
                            <input class="input" type="text" v-model="form.name">
                        </div>
                    </div>
                </div>
            </div>

            <!-- email -->
            <div class="field is-horizontal">
                <div class="field-label">
                    <label class="label">E-mail:</label>
                </div>
                <div class="field-body">
                    <div class="field">
                        <div class="control">
                            <input class="input" type="email" v-model="form.email">
                        </div>
                    </div>
                </div>
            </div>

            <!-- department -->
            <div class="field is-horizontal">
                <div class="field-label">
                    <label class="label">Departments:</label>
                </div>
                <div class="field-body">
                    <div class="field">
                        <div class="control">
                            <div class="select">
                                <select v-model="form.department">
                                    <option v-for="department of departments" :key="department.id" :value="department.id">{{department.name}}</option>
                                </select>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div v-if="error">
                <article class="message is-danger">
                    <p class="message-body">
                        {{error}}
                    </p>
                </article>
                <br/>
            </div>

            <!-- buttons -->
            <div class="field is-grouped">
                <div class="control">
                    <button class="button is-success" ref="btnSave"
                            :disabled="!form.name || (form.email && form.email.indexOf('@') === -1) || !form.department"
                            @click="addEmployee()">Save</button>
                </div>
                <div class="control">
                    <button class="button is-danger" @click="$emit('close')">Cancel</button>
                </div>
            </div>
        </div>

    </div>
</template>

<script>

export default {
    name: "EmployeeAdd",
    props: ['employee'],
    data () {
        return {
            form: {
                id: 0,
                name: '',
                email: '',
                department: 0
            },

            departments: [],
            error: ''
        }
    },

    methods: {
        getDepartments () {
            fetch(`${process.env.VUE_APP_API_URL}/department`).then(response => {
                if (response.status === 200) {
                    response.json().then(data => {
                        this.departments = data
                        this.mergeDepartment()
                    })
                }
            }).catch(()=>{})
        },

        mergeDepartment () {
            this.form.department = this.departments.find(item => this.employee.department === item.name).id
        },

        addEmployee () {
            let btn = this.$refs.btnSave
            btn.classList.add('is-loading')
            fetch(`${process.env.VUE_APP_API_URL}/employee`, {
                method: 'put',
                body: JSON.stringify(this.form),
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                }
            }).then(response => {
                if (response.status === 200) {
                    response.json().then(data => {
                        if (data.status === 'ok') {
                            this.$emit('close')
                        } else if (data.status === 'error') {
                            this.error = data.message
                        }
                    })
                }
            }).finally(() => btn.classList.remove('is-loading'))
        }
    },

    created () {
        this.getDepartments()
        this.form = Object.assign({}, this.employee)
    }
}
</script>

<style scoped>

</style>