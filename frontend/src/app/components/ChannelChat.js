"use client";

import React, { useState, useEffect, useRef } from "react";
import { useKeycloak } from "../context/KeycloakContext";
import {
    fetchChannelMessages,
    sendChannelMessage,
} from "../../lib/api/apiClient";

export default function ChannelChat({ serverId, channelId }) {
    const { keycloak, authenticated } = useKeycloak();
    const [messages, setMessages] = useState([]);
    const [input, setInput] = useState("");
    const eventRef = useRef(null);

    useEffect(() => {
        if (!authenticated) return;

        const token = keycloak.token;

        fetchChannelMessages(token, serverId, channelId)
            .then((data) => setMessages(data))
            .catch((err) => console.error(err))

        const url = `${process.env.NEXT_PUBLIC_API_URL}/servers/${serverId}/channels/${channelId}/stream?token=${token}`;
        const es = new EventSource(url);
        eventRef.current = es;

        es.addEventListener("message", (e) => {
            try {
                const msg = JSON.parse(e.data);
                setMessages((prev) => [...prev, msg]);
            } catch (err) {
                console.error("Invalid SSE message", err);
            }
        });

        es.onerror = (e) => {
            console.error("SSE error", e);
        };

        return () => {
            es.close();
        };
    }, [authenticated, serverId, channelId, keycloak]);

    const onSubmit = async (e) => {
        e.preventDefault();
        const content = input.trim();
        if (!content) return;
        try {
            await sendChannelMessage(keycloak.token, serverId, channelId, content)
            setInput("");
        } catch (err) {
            console.error(err);
        }
    };

    return (
        <div>
            <ul className="list-group mb-3">
            {messages.map((m) => (
                <li className="list-group-item" key={m.id}>
                    <strong>{m.userId}:</strong> {m.content}
                </li>
            ))}
            </ul>
            <form onSubmit={onSubmit} className="input-group">
                <input
                    className="form-control"
                    value={input}
                    onChange={(e) => setInput(e.target.value)}
                    placeholder="Type a message..."
                />
                <button className="btn btn-primary" type="submit">
                    Send
                </button>
            </form>
        </div>
    );
}
