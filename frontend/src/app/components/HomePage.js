// app/components/HomePage.js - Home page component
import React from "react";
import { useKeycloak } from "../context/KeycloakContext";

export default function Home() {
    const { authenticated, userInfo } = useKeycloak();
    
    if (!authenticated) {
        return (
            <h1>welcome to forum.</h1> 
        )
    }

    return (
        <h1>hello, {userInfo?.preferred_username}!</h1> 
    );
};
