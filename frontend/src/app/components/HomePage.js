// app/components/Home.js - Home page component
import React from "react";
import { useUser } from "../context/UserContext";

export default function Home() {
    const { loggedInUser } = useUser();

    return (
        <h1>Hello, {loggedInUser?.username}</h1> 
    );
};
