// src/app/context/UserContext.js - user data context
"use client";

import React, { 
    createContext, 
    useContext, 
    useState, 
    useEffect, 
    useCallback 
} from "react";
import { 
    fetchServers, 
    createServer as apiCreateServer 
} from '../../lib/api/apiClient';
import { useKeycloak } from '../context/KeycloakContext';

const UserContext = createContext(null);
export const useUser = () => useContext(UserContext);

export const UserProvider = ({ children }) => {
    const { keycloak, authenticated } = useKeycloak();

    const [servers, setServers] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError]     = useState(null);

    const [isCreatingServer, setCreatingServer]     = useState(false);
    const [createServerError, setCreateServerError] = useState(null);

    const loadServers = useCallback(async () => {
        if (!authenticated) {
            setServers([]);
            setLoading(false);
            return;
        }

        setLoading(true);
        setError(null);

        try {
            const token = keycloak.token;
            const data  = await fetchServers(token);
            setServers(data);
        } catch (err) {
            console.error(err);
            setError('Failed to load the server list')
        } finally {
            setLoading(false);
        }
    }, [authenticated, keycloak]);

    const createServer = useCallback(async (name) => {
        if (!authenticated) {
            setCreateServerError("Not authorized.");
            return null;
        }

        setCreatingServer(true);
        setCreateServerError(null);

        try {
            const token = keycloak.token;
            const data  = await apiCreateServer(token, name);
            await loadServers();
            return data.server;
        } catch (err) {
            console.error(err);
            setError('Failed to create server')
            return null;
        } finally {
            setCreatingServer(false);
        }
    }, [authenticated, keycloak, loadServers]);

    useEffect(() => {
        loadServers();
    }, [loadServers]);

    return (
        <UserContext.Provider value={{
            servers, 
            loading, 
            error, 
            refreshServers: loadServers,
            createServer,
            isCreatingServer,
            createServerError
        }}>
            {children}
        </UserContext.Provider>
    );
}
