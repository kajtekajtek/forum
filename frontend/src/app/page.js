"use client";
import useAuth from "./hooks/useAuth";
import { useUser } from "./context/UserContext";

export default function Home() {
    useAuth();

    const { loggedInUser } =  useUser();

    return (
        <h1>Hello, {loggedInUser.name}!</h1> 
    );
}
