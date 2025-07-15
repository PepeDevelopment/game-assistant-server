# game-assistant-server

## Prerequisites
Install both docker and docker compose.

## How to start
- Prepare **.env** file containing env variables defined in *docker-compose.yml* file. Values for variables don't matter at this stage.
- `docker compose up --build`
- Run localhost:8080 in your browser to access **adminer** - lightweight database viewer.
And obviously `docker compose down -v` for teardown, to remove containers and volumes

## Notes
As *schema.sql* is ran only once, if db-data is empty. To drop volumes call **docker compose down -v**.