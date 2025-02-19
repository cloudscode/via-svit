<template>
  <div>
    <h1>订单管理</h1>
    <ul>
      <li v-for="order in orders" :key="order.id">
        {{ order.product }} - {{ order.amount }}
        <!-- 根据需求可添加编辑和删除按钮 -->
      </li>
    </ul>
  </div>
</template>
<script lang="ts" setup>

import { ref, onMounted } from "vue";
import { getOrders } from "../api/order";

const orders = ref<Array<{ id: number; product: string; amount: number; created_at: string }>>([]);

const fetchOrders = async () => {
  try {
    const { data } = await getOrders();
    orders.value = data;
  } catch (err) {
    console.error("获取订单失败", err);
  }
};

onMounted(fetchOrders);
</script>
