// src/app/layout.js - global layout
import "../../styles/global.css";
import "../../styles/variables.css";
import 'bootstrap/dist/css/bootstrap.min.css'
import { KeycloakProvider } from "./context/KeycloakContext";
import { UserProvider } from "./context/UserContext";
import Navigation from "./components/Navigation";
import CreateServerForm from "./components/CreateServerForm";
import ServerList from "./components/ServerList";

export const metadata = {
    title: 'forum',
}

export default function RootLayout({ children }) {
    return (
        <html lang="en">
            <body>
                <KeycloakProvider>
                <UserProvider>
                    <div className="d-flex flex-column vh-100">
                        <Navigation/>
                        <CreateServerForm/>
                        <ServerList/>
                        <main> {children} </main>
                    </div>
                </UserProvider>
                </KeycloakProvider>
            </body>
        </html>
    );
}
