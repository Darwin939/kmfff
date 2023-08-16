
1. run -> docker compose up --build
2. check it in localhost:8000/proxy


3. curl example
    curl --location 'localhost:8000/proxy' \
    --header 'Content-Type: application/json' \
    --data '{
        "method": "GET",
        "url": "https://google.com",
        "headers": {
            "Authentication": "Basic bG9naW46cGFzc3dvcmQ="
        }
    }
    '
