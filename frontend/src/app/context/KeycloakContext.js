'use client';

import React, { createContext, useContext, useState, useEffect } from "react";
import keycloak from '../../lib/keycloak/keycloak';

const KeycloakContext = createContext(null);
export const useKeycloak = () => useContext(KeycloakContext);

export const KeycloakProvider = ({ children }) => {
    const [kc, setKc] = useState(null);
    const [authenticated, setAuthenticated] = useState(false);
    const [userInfo, setUserInfo] = useState(null);

    useEffect(() => {
        keycloak.init({ onLoad: 'login-required', pkceMethod: 'S256' })
        .then(auth => {
            setAuthenticated(auth);
            setKc(keycloak);

            if (auth) {
                keycloak.loadUserInfo().then(setUserInfo);
            }
        });
    }, []);

    return (
        <KeycloakContext.Provider value={{ keycloak: kc, authenticated, userInfo }}>
            {children}
        </KeycloakContext.Provider>
    );
};
