// app/components/ServerList.js - server list component
"use client";

import React from "react";
import { useKeycloak } from '../context/KeycloakContext';
import { useUser } from '../context/UserContext';

export default function ServerList() {
    const { authenticated} = useKeycloak();
    const { servers, loading, error } = useUser();

    if (!authenticated) {
        return;
    }

    if (loading) {
        return (
            <div className="d-flex justify-content-center my-4">
                <div className="spinner-border text-secondary" role="status">
                    <span className="visually-hidden">Loading...</span>
                </div>
            </div>
        );
    }

    if (error) {
        return (
            <div className="alert alert-danger my-3" role="alert">
                {error}
            </div>
        );
    }

    if (!servers) {
        return (
            <div className="alert alert-info my-3" role="alert">
                Your server list is empty. Join a server or create one.
            </div>
        );
    }

    return (
        <ul className="list-group mb-4">
            {servers.map((s) => (
                <li 
                    className="list-group-item d-flex justify-content-between align-items-center"
                    key={s.id}>
                    <span>{s.name}</span>
                    <small className="text-muted">
                        {new Date(s.createdAt).toLocaleString()}
                    </small>
                </li>
            ))}
        </ul>
    )
}
