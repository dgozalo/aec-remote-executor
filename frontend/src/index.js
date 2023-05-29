/**
 =========================================================
 * Material Dashboard 2 React - v2.2.0
 =========================================================

 * Product Page: https://www.creative-tim.com/product/material-dashboard-react
 * Copyright 2023 Creative Tim (https://www.creative-tim.com)

 Coded by www.creative-tim.com

 =========================================================

 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 */

import React from "react";
import {createRoot} from "react-dom/client";
import {BrowserRouter} from "react-router-dom";
import App from "App";

// Material Dashboard 2 React Context Provider
import {MaterialUIControllerProvider} from "context";
import {ApolloClient, ApolloProvider, InMemoryCache} from "@apollo/client";

const container = document.getElementById("app");
const root = createRoot(container);

const client = new ApolloClient({
    uri: "http://localhost:8080/query",
    cache: new InMemoryCache()
});


root.render(
    <BrowserRouter>
        <MaterialUIControllerProvider>
            <ApolloProvider client={client}>
                <App />
            </ApolloProvider>
        </MaterialUIControllerProvider>
    </BrowserRouter>
);
