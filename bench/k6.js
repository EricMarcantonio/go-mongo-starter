import http from 'k6/http';

export default function() {
    http.get('http://localhost:8080/');
}

//k6 run .\k6.js --vus 1000 --duration 1s