// src/app/components/AuthButtons.js - login/register/logout buttons
'use client';

import React from "react";
import { useKeycloak } from "../context/KeycloakContext";

export default function AuthButtons() {
    const { keycloak, authenticated } = useKeycloak();

    if (authenticated) {
        return (
            <button
                className="btn btn-primary"
                onClick={() => 
                    keycloak.logout({ redirectUri: window.location.origin })
            }>
                Logout
            </button>
        );
    }

    return (
        <div>
            <button
                className="btn btn-primary"
                onClick={() =>
                    keycloak.login({ redirectUri: window.location.origin })
            }>
                Login
            </button>
            <button
                className="btn btn-primary"
                onClick={() =>
                    keycloak.register({ redirectUri: window.location.origin })
            }>
                Register
            </button>
        </div>
    );
}
