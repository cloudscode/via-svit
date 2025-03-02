<template>
  <div>
    <table border="1" width="100%">
      <thead>
      <tr>
        <th>代号</th>
        <th>名称</th>
        <th>价格</th>
        <th>信息1</th>
        <th>信息2</th>
        <th>信息3</th>
        <th>信息4</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="item in items" :key="item.code">
        <td>{{ item.code }}</td>
        <td>{{ item.name }}</td>
        <td>{{ item.price }}</td>
        <td>{{ item.info1 }}</td>
        <td>{{ item.info2 }}</td>
        <td>{{ item.info3 }}</td>
        <td>{{ item.info4 }}</td>
      </tr>
      </tbody>
    </table>
    <div>
      <button @click="prevPage" :disabled="page === 1">上一页</button>
      <button @click="nextPage" :disabled="page === totalPages">下一页</button>
      <span>第 {{ page }} 页 / 共 {{ totalPages }} 页</span>
    </div>
  </div>
</template>

<script>

import {GetItems} from '../../wailsjs/go/backend/App'
import {services} from '../../wailsjs/go/models'
import {useRoute, useRouter} from 'vue-router'

export default {
  data() {
    return {
      items: [],
      page: 1,
      pageSize: 10,
      totalCount: 0,
    };
  },
  computed: {
    totalPages() {
      return Math.ceil(this.totalCount / this.pageSize);
    },
  },
  methods: {
    async loadItems() {
      const result = await GetItems(this.page, this.pageSize);
      this.items = result.items;
      this.totalCount = result.total_count;
    },
    async nextPage() {
      if (this.page < this.totalPages) {
        this.page++;
        await this.loadItems();
      }
    },
    async prevPage() {
      if (this.page > 1) {
        this.page--;
        await this.loadItems();
      }
    },
  },
  mounted() {
    this.loadItems();
  },
};
</script>
