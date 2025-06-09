"use client";

import React from "react";
import Link from "next/link";

export default function ChannelList({ serverId, channels, loading, error }) {
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

    if (!channels || channels.length === 0) {
        <div className="alert alert-info my-3" role="alert">
            No channels found.
        </div>
    }

    return (
        <ul className="list-group mb-4">
            {channels.map((c) => (
                <li
                    key={c.id}
                    className="list-group-item d-flex justify-content-between align-items-center"
                >
                    <Link
                        href={`/servers/${serverId}/channels/${c.id}`}
                        className="me-auto text-decoration-none">
                        {c.name}
                    </Link>
                    <small className="text-muted">
                        {new Date(c.createdAt).toLocaleString()}
                    </small>
                </li>
            ))}
        </ul>
    );
}
