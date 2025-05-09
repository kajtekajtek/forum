// src/app/layout.js - global layout
import "./globals.css";
import 'bootstrap/dist/css/bootstrap.min.css'
import UserProvider from "./context/UserContext"

export const metadata = {
    title: 'forum',
}

export default function RootLayout({ children }) {
    return (
        <html lang="en">
            <body>
                <UserProvider>
                    {children}
                </UserProvider>
            </body>
        </html>
    );
}
