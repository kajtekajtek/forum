// app/components/Navigation.js - navigation component
"use client";

import React from "react";
import AuthButtons from '../components/AuthButtons';

export default function Navigation() {
    return (
        <div className="navbar-nav-expand-lg">
            <div className="container-fluid">
                <div className="navbar-nav ms-auto">
                    <AuthButtons/>
                </div>
            </div>
        </div>
    )
}
