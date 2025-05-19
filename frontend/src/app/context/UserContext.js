// app/context/UserContext.js - user authentication context
"use client";

import React, { createContext, useContext, useState, useEffect } from "react";
import bcrypt from "bcryptjs";
import { useRouter } from "next/navigation";

const UserContext = createContext(null);
export const useUser = () => useContext(UserContext);

// cookie utilities
function setTokenCookie(token) {
    document.cookie = `token=${token}; path=/`;
}

function getTokenFromCookie() {
    if (typeof document === 'undefined') return null;
    const match = document.cookie.match(/(^| )token=([^;]+)/);
    return match ? match[2] : null;
}

function removeTokenCookie() {
    if (typeof document === 'undefined') return;
    document.cookie = 'token=; Max-Age=0; path=/';
}

export default function userProvider({ children }) {
    const router = useRouter();
    const [loggedInUser, setLoggedInUser] = useState(null);
    const [loading, setLoading] = useState(true);

    // read token from cookie and set user on mount
    useEffect(() => {
        const token = getTokenFromCookie();
        if (token) {
            try {
                const user = JSON.parse(atob(token));
                setLoggedInUser(user);
            } catch (e) {
                console.error("Invalid token cookie", e);
                removeTokenCookie();
            }
        }
        setLoading(false);
    }, []);

    // validate, set cookie, context and redirect (mock)
    const login = (email, password) => {
        const users = JSON.parse(localStorage.getItem("users")) || [];

        const user = users.find((u) => u.email == email);
        if (!user) {
            return false;
        }

        const correctPassword = bcrypt.compareSync(password, user.password);
        if (!correctPassword) {
            return false;
        }

        const token = btoa(JSON.stringify({ username: user.username }));
        setTokenCookie(token);
        setLoggedInUser({ username: user.username });
        router.replace("/");
        return true;
    }

    // clear cookie, context and redirect (mock)
    const logout = () => {
        removeTokenCookie();
        setLoggedInUser(null);
        router.replace("/login");
    }

    // store user in localStorage (mock)
    const register = (username, email, password) => {
        const users = JSON.parse(localStorage.getItem("users")) || [];

        if (users.some((user) => user.email === email)) {
            return false;
        }

        const passwordHash = bcrypt.hashSync(password, 10);

        users.push({
            username: username,
            email: email,
            password: passwordHash,
            servers: []
        });

        localStorage.setItem("users", JSON.stringify(users));
        return true;
    }

    return (
        <UserContext.Provider value={{
            loggedInUser, loading, login, logout, register
        }}>
            {children}
        </UserContext.Provider>
    )
}
