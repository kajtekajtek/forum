// app/components/Navigation.js - navigation component
"use client";

import React from "react";
import { useKeycloak } from '../context/KeycloakContext';

export default function Navigation() {
    const { redirectToLogout, authenticated  } = useKeycloak();

    return (
        <div className="navbar-nav-expand-lg">
            <div className="container-fluid">
                <div className="navbar-nav ms-auto">
                    {/* if user is logged in, show logout button,
                        else show login and register buttons */}
                    {authenticated ? (
                        <button className="btn btn-primary" 
                                onClick={redirectToLogout}>
                            Logout
                        </button>
                    ) : (
                        <>
                            {/* TODO: login & register buttons */}
                        </>
                    )}
                </div>
            </div>
        </div>
    )
}
