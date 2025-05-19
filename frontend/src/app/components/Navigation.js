// app/components/Navigation.js - navigation component
"use client";

import React from "react";
import { useUser } from "../context/UserContext";
import Link from "next/link";

export default function Navigation() {
    const { loggedInUser, logout } = useUser();

    return (
        <div className="navbar-nav-expand-lg">
            <div className="container-fluid">
                <div className="navbar-nav ms-auto">
                    {/* if user is logged in, show logout button,
                        else show login and register buttons */}
                    {loggedInUser ? (
                        <button className="btn btn-primary" onClick={logout}>
                            Logout
                        </button>
                    ) : (
                        <>
                            <Link href="/login">
                                <button className="btn btn-primary me-2">Login</button>
                            </Link>
                            <Link href="/register">
                                <button className="btn btn-primary">Register</button>
                            </Link>
                        </>
                    )}
                </div>
            </div>
        </div>
    )
}
