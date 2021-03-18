# Pi Resource Monitor API
## Web API returning JSON response for remote monitoring of Pi temperature, uptime and resources

Quick and Simple Golang Webapp for sharing basic statistics via a JSON API. By default this runs on port 8081 -- this can be changed by editing the `main.go` file, then recompiling with `go build -o monitorWebapp main.go`.

Install by running the `install.sh` script as root.
_**Note:** The install script will set the service to run as the root user by default. Feel free to change this as required by editing the `install.sh` script._
