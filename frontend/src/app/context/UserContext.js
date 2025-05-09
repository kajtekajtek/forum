// app/context/UserContext.js - user authentication context
"use client";

import React, { createContext, useContext, useState, useEffect } from "react";
import bcrypt from "bcryptjs";
import { useRouter } from "next/navigation";

const UserContext = createContext(null);

export const useUser = () => useContext(UserContext);

export default function userProvider({ children }) {
    const router = useRouter();
    const [loggedInUser, setLoggedInUser] = useState(null);

    useEffect(() => {
        const loggedInUser = JSON.parse(localStorage.getItem("user"));
        setLoggedInUser(loggedInUser || null);
    }, []);

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

        setLoggedInUser(user);
        localStorage.setItem("user", JSON.stringify(user));
        return true;
    }

    const logout = () => {
        setLoggedInUser(null);
        localStorage.removeItem("user");
        router.push("/login");
    }

    const register = (username, email, password) => {
        const users = JSON.parse(localStorage.getItem("users")) || [];

        const userExists = users.some((user) => user.email === email);
        if (userExists) {
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
            loggedInUser, login, logout, register
        }}>
            {children}
        </UserContext.Provider>
    )
}
