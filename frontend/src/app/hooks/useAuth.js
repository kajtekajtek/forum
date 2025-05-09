// app/hooks/useAuth.js - custom hook for user authentication check
import { useEffect } from "react";
import { useRouter } from "next/navigation";

export default function useAuth() {
    const router = useRouter();

    useEffect(() => {
        // check whether the user is authenticated
        if (!JSON.parse(localStorage.getItem("user"))) {
            // if not, redirect to login page
            router.push("/login");
        }
    }, [router]);
}
