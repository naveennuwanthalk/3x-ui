<template>
  <div class="login-container">
    <a-card class="login-card" title="X-UI Login">
      <a-form
        :form="form"
        @submit.prevent="handleSubmit"
      >
        <a-form-item>
          <a-input
            v-decorator="[
              'username',
              { rules: [{ required: true, message: 'Please input your username!' }] }
            ]"
            placeholder="Username"
          >
            <a-icon slot="prefix" type="user" />
          </a-input>
        </a-form-item>
        <a-form-item>
          <a-input-password
            v-decorator="[
              'password',
              { rules: [{ required: true, message: 'Please input your password!' }] }
            ]"
            placeholder="Password"
          >
            <a-icon slot="prefix" type="lock" />
          </a-input-password>
        </a-form-item>
        <a-form-item>
          <a-button
            type="primary"
            html-type="submit"
            :loading="loading"
            block
          >
            Log in
          </a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loading: false,
      form: this.$form.createForm(this)
    }
  },
  methods: {
    handleSubmit() {
      this.form.validateFields((err, values) => {
        if (!err) {
          this.loading = true
          this.$http.post('/auth/login', values)
            .then(response => {
              localStorage.setItem('user', JSON.stringify(response.data.user))
              this.$router.push('/')
            })
            .catch(error => {
              this.$message.error(error.response?.data?.message || 'Login failed')
            })
            .finally(() => {
              this.loading = false
            })
        }
      })
    }
  }
}
</script>

<style scoped>
.login-container {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f0f2f5;
}

.login-card {
  width: 100%;
  max-width: 400px;
}
</style>