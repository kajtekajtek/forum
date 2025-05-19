import axios from 'axios';

const apiClient = axios.create({
    baseURL: process.env.BACKEND_URL || 'http://localhost:8080/api',
    timeout: 10000,
});

apiClient.interceptors.request.use((config) => {
    return config;
})

export default apiClient;
