import Keycloak from 'keycloak-js';

if (!process.env.NEXT_PUBLIC_KEYCLOAK_URL || 
    !process.env.NEXT_PUBLIC_KEYCLOAK_REALM || 
    !process.env.NEXT_PUBLIC_KEYCLOAK_CLIENT_ID) {
  throw new Error("Missing Keycloak configuration in environment variables.");
}

const keycloak = new Keycloak({
    url: process.env.NEXT_PUBLIC_KEYCLOAK_URL,
    realm: process.env.NEXT_PUBLIC_KEYCLOAK_REALM,
    clientId: process.env.NEXT_PUBLIC_KEYCLOAK_CLIENT_ID,
});

export default keycloak;
