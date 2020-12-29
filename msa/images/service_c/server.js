"use strict";

const express = require('express');
const axios = require('axios');

const PORT = 9999;
const HOST = "0.0.0.0";

const app = express();
app.get('/', (req, res) => {
    let res_a = axios.get('http://service_a:7777');
    let res_b = axios.get('http://service_b:8888');
    
    res.send(`Service C, version 1.0\n${res_a}\n${res_b}`);
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);