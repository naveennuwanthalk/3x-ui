<template>
  <div class="inbounds">
    <a-card>
      <div slot="title">
        <a-space>
          <span>Inbound Management</span>
          <a-button type="primary" @click="showAddModal">
            <a-icon type="plus" /> Add Inbound
          </a-button>
        </a-space>
      </div>

      <a-table
        :columns="columns"
        :data-source="inbounds"
        :loading="loading"
        rowKey="id"
      >
        <template slot="enable" slot-scope="text">
          <a-switch :checked="text" @change="(checked) => handleEnableChange(checked, record)" />
        </template>

        <template slot="protocol" slot-scope="text">
          <a-tag :color="getProtocolColor(text)">{{ text }}</a-tag>
        </template>

        <template slot="traffic" slot-scope="text, record">
          {{ formatTraffic(record.up + record.down) }}
        </template>

        <template slot="action" slot-scope="text, record">
          <a-space>
            <a-button type="link" @click="() => showEditModal(record)">
              <a-icon type="edit" />
            </a-button>
            <a-button type="link" @click="() => showQRCode(record)">
              <a-icon type="qrcode" />
            </a-button>
            <a-popconfirm
              title="Are you sure you want to delete this inbound?"
              @confirm="() => handleDelete(record.id)"
            >
              <a-button type="link" danger>
                <a-icon type="delete" />
              </a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </a-card>

    <!-- Add/Edit Modal -->
    <a-modal
      :title="modalTitle"
      :visible="modalVisible"
      @ok="handleModalOk"
      @cancel="handleModalCancel"
      :confirmLoading="modalLoading"
    >
      <a-form :form="form">
        <a-form-item label="Protocol">
          <a-select
            v-decorator="[
              'protocol',
              { rules: [{ required: true, message: 'Please select protocol!' }] }
            ]"
          >
            <a-select-option value="vmess">VMess</a-select-option>
            <a-select-option value="vless">VLESS</a-select-option>
            <a-select-option value="trojan">Trojan</a-select-option>
            <a-select-option value="shadowsocks">Shadowsocks</a-select-option>
          </a-select>
        </a-form-item>

        <a-form-item label="Port">
          <a-input-number
            v-decorator="[
              'port',
              { rules: [{ required: true, message: 'Please input port!' }] }
            ]"
            :min="1"
            :max="65535"
          />
        </a-form-item>

        <a-form-item label="Remark">
          <a-input
            v-decorator="[
              'remark',
              { rules: [{ required: true, message: 'Please input remark!' }] }
            ]"
          />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loading: false,
      inbounds: [],
      modalVisible: false,
      modalLoading: false,
      modalTitle: 'Add Inbound',
      currentInbound: null,
      form: this.$form.createForm(this),
      columns: [
        {
          title: 'Enable',
          dataIndex: 'enable',
          scopedSlots: { customRender: 'enable' }
        },
        {
          title: 'Remark',
          dataIndex: 'remark'
        },
        {
          title: 'Protocol',
          dataIndex: 'protocol',
          scopedSlots: { customRender: 'protocol' }
        },
        {
          title: 'Port',
          dataIndex: 'port'
        },
        {
          title: 'Traffic',
          scopedSlots: { customRender: 'traffic' }
        },
        {
          title: 'Action',
          scopedSlots: { customRender: 'action' }
        }
      ]
    }
  },
  methods: {
    fetchInbounds() {
      this.loading = true
      this.$http.get('/api/inbounds')
        .then(response => {
          this.inbounds = response.data
        })
        .catch(error => {
          this.$message.error('Failed to fetch inbounds')
          console.error(error)
        })
        .finally(() => {
          this.loading = false
        })
    },
    formatTraffic(bytes) {
      const units = ['B', 'KB', 'MB', 'GB', 'TB']
      let size = bytes
      let unitIndex = 0
      while (size >= 1024 && unitIndex < units.length - 1) {
        size /= 1024
        unitIndex++
      }
      return `${size.toFixed(2)} ${units[unitIndex]}`
    },
    getProtocolColor(protocol) {
      const colors = {
        vmess: 'blue',
        vless: 'green',
        trojan: 'purple',
        shadowsocks: 'orange'
      }
      return colors[protocol] || 'default'
    },
    showAddModal() {
      this.modalTitle = 'Add Inbound'
      this.currentInbound = null
      this.modalVisible = true
      this.form.resetFields()
    },
    showEditModal(record) {
      this.modalTitle = 'Edit Inbound'
      this.currentInbound = record
      this.modalVisible = true
      this.form.setFieldsValue(record)
    },
    handleModalOk() {
      this.form.validateFields((err, values) => {
        if (!err) {
          this.modalLoading = true
          const request = this.currentInbound
            ? this.$http.put(`/api/inbounds/${this.currentInbound.id}`, values)
            : this.$http.post('/api/inbounds', values)

          request
            .then(() => {
              this.$message.success('Operation successful')
              this.modalVisible = false
              this.fetchInbounds()
            })
            .catch(error => {
              this.$message.error(error.response?.data?.message || 'Operation failed')
            })
            .finally(() => {
              this.modalLoading = false
            })
        }
      })
    },
    handleModalCancel() {
      this.modalVisible = false
    },
    handleDelete(id) {
      this.$http.delete(`/api/inbounds/${id}`)
        .then(() => {
          this.$message.success('Deleted successfully')
          this.fetchInbounds()
        })
        .catch(error => {
          this.$message.error('Failed to delete')
          console.error(error)
        })
    },
    handleEnableChange(checked, record) {
      this.$http.put(`/api/inbounds/${record.id}`, { enable: checked })
        .catch(error => {
          this.$message.error('Failed to update status')
          console.error(error)
        })
    }
  },
  mounted() {
    this.fetchInbounds()
  }
}
</script>

<style scoped>
.inbounds {
  padding: 24px;
}
</style>