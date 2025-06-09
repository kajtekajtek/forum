"use client";

import React, { useEffect, useState, useCallback } from "react";
import { useKeycloak } from "../../context/KeycloakContext";
import { fetchServerChannels } from "../../../lib/api/apiClient";
import ChannelList from "../../components/ChannelList";
import CreateChannelForm from "../../components/CreateChannelForm";

export default function ServerPage({ params }) {
    const { serverId } = params;
    const { keycloak, authenticated } = useKeycloak();

    const [channels, setChannels] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    const loadChannels = useCallback(async () => {
        if (!authenticated) return;

        setLoading(true);
        setError(null);

        try {
            const data = await fetchServerChannels(keycloak.token, serverId);
            setChannels(data);
        } catch (err) {
            console.error(err);
            setError("Failed to load channels");
        } finally {
            setLoading(false);
        }
    }, [authenticated, keycloak, serverId]);

    useEffect(() => {
        loadChannels();
    }, [loadChannels]);

    const onChannelCreation = () => {
        loadChannels();
    };

    if (!authenticated) return null;

    return (
        <div className="container mt-3">
            <h2>Server Channels</h2>
            <CreateChannelForm serverId={serverId} onCreated={onChannelCreation}/>
            <ChannelList
                serverId={serverId}
                channels={channels}
                loading={loading}
                error={error}
            />
        </div>
    );
}
