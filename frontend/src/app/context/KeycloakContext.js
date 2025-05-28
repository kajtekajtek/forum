'use client';

import React, { createContext, useContext, useState, useEffect } from "react";
import keycloak from '../../lib/keycloak/keycloak';

const KeycloakContext = createContext(null);
export const useKeycloak = () => useContext(KeycloakContext);

export const KeycloakProvider = ({ children }) => {
    const [initialized, setInitialized] = useState(false);
    const [authenticated, setAuthenticated] = useState(false);
    const [userInfo, setUserInfo] = useState(null);

    useEffect(() => {
        keycloak
            // initialization
            .init({
                onLoad: 'check-sso',
                pkceMethod: 'S256',
                silentCheckSsoRedirectUri: `${window.location.origin}/silent-check-sso.html`,
                checkLoginIframe: false,
            })
            // authentication
            .then((auth) => {
                setAuthenticated(auth);

                if (auth) {
                    return keycloak.loadUserInfo();
                }
            })
            // loading user info
            .then((info) => {
                if (info) setUserInfo(info);
            })
            .catch((err) => {
                console.error('Keycloak init error:', err);
            })
            .finally(() => {
                setInitialized(true);
            });
    }, []);

    const redirectToLogout = () => {
        keycloak.logout({ redirectUri: window.location.origin });
    };

    const redirectToLogin = () => {
        keycloak.login({ redirectUri: window.location.origin });
    };

    if (!initialized) {
        return <div>Initializing Keycloak...</div>;
    }

    return (
        <KeycloakContext.Provider value={{ redirectToLogin, redirectToLogout, keycloak, authenticated, userInfo }}>
            {children}
        </KeycloakContext.Provider>
    );
};
