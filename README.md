## Description
A small web service that takes the source and a list of destinations
and returns a list of routes between source and each destination. Both source and
destination are defined as a pair of latitude and longitude. The returned list of routes
should is sorted by driving time and distance (if time is equal).

## Run the application
1. Clone the project to your machine ```$ git clone https://github.com/Nivelian/ingrid-task```
2. Navigate into the directory ```$ cd ingrid-task```
3. Create executable file ```$ go build```
4. Change the executable file rights ```$ chmod +x ./ingrid-task```
5. Start the server on port 8080 ```$ ./ingrid-task```
6. Test routing ```$ curl 'http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219'```
