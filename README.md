# Secret-noteAPI-SPA-NabilaSherif
A full-stack web application that allows users to securely share self-destructing secret notes. The application features a Go backend using the Gin framework, a Vue.js frontend, and an SQLite database.

## Features

- Users can create an account, login and logout.
- After logging in, user can create note and view his/her notes without incrementing the views. 
- Secret notes are created with a user-defined expiration time and a maximum view limit.
- Notes are permanently deleted after either the expiration time is reached or the maximum view count is exceeded.
- Note views gets incremented when accessing the note via its automatically generated url.
- A 404 error page is displayed when attempting to access a deleted note or any oother unrecognized url.
- Rate limiting is applied to prevent abuse 

## Major Frameworks and Libraries
<img height="50" src="https://user-images.githubusercontent.com/25181517/192109061-e138ca71-337c-4019-8d42-4792fdaa7128.png"> <img height="50" src="https://user-images.githubusercontent.com/25181517/121401671-49102800-c959-11eb-9f6f-74d49a5e1774.png"><img height="50" src="https://user-images.githubusercontent.com/25181517/117448124-a2da9800-af3e-11eb-85d2-bd1b69b65603.png"><img height="50" src="https://user-images.githubusercontent.com/25181517/192149581-88194d20-1a37-4be8-8801-5dc0017ffbbe.png"><img height="50" src="https://github.com/marwin1991/profile-technology-icons/assets/136815194/50c63e54-074f-494b-b786-01eb7870c927"><img height="50" src="https://user-images.githubusercontent.com/25181517/192108374-8da61ba1-99ec-41d7-80b8-fb2f7c0a4948.png"><img height="50" src="https://github.com/marwin1991/profile-technology-icons/assets/136815194/82df4543-236b-4e45-9604-5434e3faab17"> <img height="50" src="https://user-images.githubusercontent.com/25181517/117207330-263ba280-adf4-11eb-9b97-0ac5b40bc3be.png"><img height="50" src="https://user-images.githubusercontent.com/25181517/192158606-7c2ef6bd-6e04-47cf-b5bc-da2797cb5bda.png">

## API Endpoints

- POST /note: Creates a new note. Requires authorization.
- GET /note/:note_url: Retrieves a note by URL.
- POST /users: Creates a new user. 
- POST /login: Logs in a user.
- GET /notes: Retrieves all notes for the logged-in user. Requires authorization.
- POST /logout: Logs out the user. Requires authorization.

## commands to run front-end
```
cd frontend
npm run dev
```

## example commands to run back-end
```
cd backend
go run cmd/main.go -p 8080 -d test.db
```


## on the browser navigate to this url
http://localhost:8080/



## docker build and run backend
```
docker build -t backend -f backend/Dockerfile backend
docker run -p 8080:8080 backend
```
## docker build and run frontend
```
docker build -t frontend -f frontend/Dockerfile frontend
docker run -p 5173:5173 frontend
```



## build and run docker compose
```
docker-compose up --build
```
## run docker compose
```
docker compose up
```
## stop docker compose 
```
docker-compose down
```
