import axios from "axios";

if (!process.env.NEXT_PUBLIC_API_URL) {
  throw new Error("Missing API configuration in environment variables.");
}

const apiClient = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL,
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
    const response = await apiClient.post('/servers',
        { name },
        { headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },}
    );

    return response.data;
};

/*
    fetches user's server list
    returns array of servers
*/
export const fetchServers = async (token) => {
    const response = await apiClient.get('/servers',
        { headers: { Authorization: `Bearer ${token}` }}
    );

    return response.data.servers;
};

/*
    fetch messages from a channel
    returns array of messages
*/
export const fetchChannelMessages = async (token, serverId, channelId) => {
    const response = await apiClient.get(
        `/servers/${serverId}/channels/${channelId}/messages`,
        {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        },
    );

    return response.data.messages;
};

/*
    send a message to a channel
    returns created message
*/
export const sendChannelMessage = async (
    token,
    serverId,
    channelId,
    content,
) => {
    const response = await apiClient.post(
        `/servers/${serverId}/channels/${channelId}/messages`,
        { content },
        { headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        }}
    );

    return response.data.message;
};

export default apiClient;
