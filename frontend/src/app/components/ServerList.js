// app/components/ServerList.js - server list component
"use client";

import React from "react";
import { useUser } from "../context/UserContext";

export default function ServerList() {
    const { loggedInUser } = useUser();

    if (!loggedInUser) {
        return;
    }

    return (
        
    )
}
