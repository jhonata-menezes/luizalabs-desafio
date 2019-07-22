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

            <!-- buttons -->
            <div class="field is-grouped">
                <div class="control">
                    <button class="button is-success" ref="btnSave"
                            :disabled="!form.name"
                            @click="addDepartment()">Save</button>
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
        name: "DepartmentAdd",
        data () {
            return {
                form: {
                    name: '',
                },
            }
        },

        methods: {
            addDepartment () {
                let btn = this.$refs.btnSave
                btn.classList.add('is-loading')
                fetch(`${process.env.VUE_APP_API_URL}/department`, {
                    method: 'put',
                    body: JSON.stringify(this.form),
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json'
                    }
                }).then(response => {
                    if (response.status === 200) {
                        btn.classList.remove('is-loading')
                        this.$emit('close')
                        this.form.name = ''
                    }
                })
            }
        }
    }
</script>

<style scoped>

</style>