// app/components/HomePage.js - Home page component
import React, { useEffect } from "react";
import { useKeycloak } from "../context/KeycloakContext";

export default function Home() {
    const { redirectToLogin, authenticated, userInfo } = useKeycloak();
    
    if (!authenticated) {
        redirectToLogin();
        return <p>Redirecting to Keycloak...</p>;
    }

    return (
        <h1>Hello, {userInfo.preferred_username}!</h1> 
    );
};
