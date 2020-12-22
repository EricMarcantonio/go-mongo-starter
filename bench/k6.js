import http from 'k6/http';

export default function() {
    var url = 'http://localhost:8080/user'
    var payload = JSON.stringify({
        test: "lol"
    })
    var params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params)
}

//k6 run .\k6.js --vus 1000 --duration 1s