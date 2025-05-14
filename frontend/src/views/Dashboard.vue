<template>
  <div class="dashboard">
    <a-row :gutter="16">
      <a-col :span="8">
        <a-card>
          <a-statistic
            title="Total Traffic"
            :value="formatTraffic(totalTraffic)"
            :precision="2"
          >
            <template #prefix>
              <a-icon type="arrow-up" />
              <a-icon type="arrow-down" />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      <a-col :span="8">
        <a-card>
          <a-statistic
            title="Active Users"
            :value="activeUsers"
          >
            <template #prefix>
              <a-icon type="user" />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      <a-col :span="8">
        <a-card>
          <a-statistic
            title="CPU Usage"
            :value="cpuUsage"
            :precision="2"
            suffix="%"
          >
            <template #prefix>
              <a-icon type="dashboard" />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
    </a-row>

    <a-card style="margin-top: 16px">
      <a-tabs default-active-key="1">
        <a-tab-pane key="1" tab="Traffic History">
          <!-- Traffic chart will go here -->
        </a-tab-pane>
        <a-tab-pane key="2" tab="System Info">
          <a-descriptions bordered>
            <a-descriptions-item label="Xray Version">
              {{ systemInfo.xrayVersion }}
            </a-descriptions-item>
            <a-descriptions-item label="Operating System">
              {{ systemInfo.os }}
            </a-descriptions-item>
            <a-descriptions-item label="Memory Usage">
              {{ systemInfo.memoryUsage }}%
            </a-descriptions-item>
          </a-descriptions>
        </a-tab-pane>
      </a-tabs>
    </a-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      totalTraffic: 0,
      activeUsers: 0,
      cpuUsage: 0,
      systemInfo: {
        xrayVersion: '',
        os: '',
        memoryUsage: 0
      }
    }
  },
  methods: {
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
    fetchDashboardData() {
      this.$http.get('/api/dashboard')
        .then(response => {
          const data = response.data
          this.totalTraffic = data.totalTraffic
          this.activeUsers = data.activeUsers
          this.cpuUsage = data.cpuUsage
          this.systemInfo = data.systemInfo
        })
        .catch(error => {
          this.$message.error('Failed to fetch dashboard data')
          console.error(error)
        })
    }
  },
  mounted() {
    this.fetchDashboardData()
    // Refresh data every 30 seconds
    setInterval(this.fetchDashboardData, 30000)
  },
  beforeDestroy() {
    clearInterval(this.interval)
  }
}
</script>

<style scoped>
.dashboard {
  padding: 24px;
}
</style>