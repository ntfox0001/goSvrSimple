
echo package main > version.go
echo const ( >> version.go
echo GitVersion = "%1" >> version.go
echo ) >> version.go