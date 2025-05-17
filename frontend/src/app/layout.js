// src/app/layout.js - global layout
import "../../styles/global.css";
import "../../styles/variables.css";
import 'bootstrap/dist/css/bootstrap.min.css'
import UserProvider from "./context/UserContext"
import Navigation from "./components/Navigation";

export const metadata = {
    title: 'forum',
}

export default function RootLayout({ children }) {
    return (
        <html lang="en">
            <body>
                <UserProvider>
                    <div className="d-flex flex-column vh-100">
                        <Navigation/>
                        <main> {children} </main>
                    </div>
                </UserProvider>
            </body>
        </html>
    );
}
