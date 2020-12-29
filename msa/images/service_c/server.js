"use strict";

const express = require('express');
const axios = require('axios');

const PORT = 9999;
const HOST = "0.0.0.0";

const app = express();
app.get('/', (req, res) => {
    const promise = (url) => {
        return axios.get(url).then((res)=> res.data).catch((err)=>`Error from ${url}`)
    }
    let res_a = promise('http://service_a.default:7777');
    let res_b = promise('http://service_b.default:8888');
    
    res.send(`Service C, version 1.0\n${res_a}\n${res_b}`);
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);