<template>
  <div class="settings">
    <a-card title="System Settings">
      <a-tabs>
        <a-tab-pane key="1" tab="General Settings">
          <a-form :form="form" @submit.prevent="handleSubmit">
            <a-form-item label="Panel Port">
              <a-input-number
                v-decorator="[
                  'port',
                  { rules: [{ required: true, message: 'Please input panel port!' }] }
                ]"
                :min="1"
                :max="65535"
              />
            </a-form-item>

            <a-form-item label="Username">
              <a-input
                v-decorator="[
                  'username',
                  { rules: [{ required: true, message: 'Please input username!' }] }
                ]"
              />
            </a-form-item>

            <a-form-item label="New Password">
              <a-input-password
                v-decorator="[
                  'password',
                  { rules: [{ required: false, message: 'Please input new password!' }] }
                ]"
                placeholder="Leave blank to keep current password"
              />
            </a-form-item>

            <a-form-item label="Theme">
              <a-radio-group
                v-decorator="[
                  'theme',
                  { initialValue: 'light' }
                ]"
              >
                <a-radio-button value="light">Light</a-radio-button>
                <a-radio-button value="dark">Dark</a-radio-button>
              </a-radio-group>
            </a-form-item>

            <a-form-item>
              <a-button type="primary" html-type="submit" :loading="loading">
                Save Changes
              </a-button>
            </a-form-item>
          </a-form>
        </a-tab-pane>

        <a-tab-pane key="2" tab="Backup & Restore">
          <a-space direction="vertical" style="width: 100%">
            <a-card title="Database Backup" size="small">
              <a-button type="primary" @click="handleBackup">
                Download Backup
              </a-button>
            </a-card>

            <a-card title="Database Restore" size="small">
              <a-upload
                name="database"
                :action="'/api/settings/restore'"
                :headers="{
                  authorization: 'authorization-text',
                }"
                @change="handleRestore"
              >
                <a-button>
                  <a-icon type="upload" /> Upload Backup
                </a-button>
              </a-upload>
            </a-card>
          </a-space>
        </a-tab-pane>
      </a-tabs>
    </a-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: this.$form.createForm(this),
      loading: false
    }
  },
  methods: {
    handleSubmit() {
      this.form.validateFields((err, values) => {
        if (!err) {
          this.loading = true
          this.$http.post('/api/settings', values)
            .then(() => {
              this.$message.success('Settings updated successfully')
            })
            .catch(error => {
              this.$message.error(error.response?.data?.message || 'Failed to update settings')
            })
            .finally(() => {
              this.loading = false
            })
        }
      })
    },
    handleBackup() {
      window.location.href = '/api/settings/backup'
    },
    handleRestore(info) {
      if (info.file.status === 'done') {
        this.$message.success('Database restored successfully')
      } else if (info.file.status === 'error') {
        this.$message.error('Failed to restore database')
      }
    }
  },
  mounted() {
    // Fetch current settings
    this.$http.get('/api/settings')
      .then(response => {
        const settings = response.data
        this.form.setFieldsValue(settings)
      })
      .catch(error => {
        this.$message.error('Failed to fetch settings')
        console.error(error)
      })
  }
}
</script>

<style scoped>
.settings {
  padding: 24px;
}
</style>