# tripatra-api
- Built with domain design pattern and runs on **docker compose**
- jwt-go management with redis to track which users are still logged in and have logged out

# Getting started
- Go to the application project directory then run the command "docker-compose up --build -d"
- Import json data collection "tripatra-api.postman_collection.json" to test with postman
- To access mailhog, open "http://localhost:8025" in a web browser
- Stop running application and remove container, run the command "docker-compose down"