import axios from "axios";

// 后端 Gin 订单接口地址（注意端口需与后端一致）
const api = axios.create({ baseURL: "http://localhost:8080/orders" });

export const getOrders = () => api.get("");
export const createOrder = (order: any) => api.post("", order);
export const updateOrder = (order: any) => api.put("", order);
export const deleteOrder = (id: number) => api.delete(`/${id}`);
