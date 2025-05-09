// src/app/layout.js - global layout
import "./globals.css";
import 'bootstrap/dist/css/bootstrap.min.css'

export const metadata = {
    title: 'forum',
}

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>
        {children}
      </body>
    </html>
  );
}
