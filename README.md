# Secret-noteAPI-SPA-Nabila
The end result of the build script is a single docker image that had the compiled golang binary, alongside the production built vuejs code. The golang application would serve the static, public entrypoint for the SPA ( index.html ), and then the vuejs application would load in the browser


#run vue project 
npm run serve


on the browser navigate to this url after running serve
http://localhost:8080/

create a production-ready build of the app
npm run build




docker build and run backend

docker build -t backend -f Dockerfile-backend .

docker run -p 8080:8080 backend