import axios from 'axios';

const apiClient = axios.create({
    baseURL: process.env.BACKEND_URL || 'http://localhost:8080/api',
    timeout: 10000,
});

apiClient.interceptors.request.use((config) => {
    return config;
})

/* 
    create server with given name
    returns { server, membership } object
*/
export const createServer = async (token, name) => {
    const response = await apiClient.post(
        '/servers',
        { name },
        {
            headers: {
                Authorization: `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
        }
    );
    return response.data;
};

/*
    fetches user's server list
    returns array of servers
*/
export const fetchServers = async (token) => {
    const response = await apiClient.get(
        '/servers',
        {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        }
    );
    return response.data.servers;
};

export default apiClient;
