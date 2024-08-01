# Secret-noteAPI-SPA-Nabila
The end result of the build script is a single docker image that had the compiled golang binary, alongside the production built vuejs code. The golang application would serve the static, public entrypoint for the SPA ( index.html ), and then the vuejs application would load in the browser


#run vue project 
npm run dev


on the browser navigate to this url after running serve
http://localhost:8080/

create a production-ready build of the app
npm run build




docker build and run backend

docker build -t backend -f Dockerfile-backend .

docker run -p 8080:8080 backend




json for createnote 

{
  "note_text": "This is a test note.",
  "expiration_date": "2026-01-02T15:04:05Z",
  "max_viewers": 5
}

json fpr create user

{
  "username": "123",
  "password": "password"
}



build docker compose

docker compose up


uuid beeen "" lama ba search 3al note


.value instead of this in composition API

blue color is always overriden

