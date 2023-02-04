import React from "react";
import {createRoot } from 'react-dom/client';
import 'bootstrap/dist/css/bootstrap.min.css';
import App from "./components/App";
import './styles.css';


const rootElement  = document.getElementById("root");
const root = createRoot(rootElement);

root.render(<App />)