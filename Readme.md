kubectl simplification
Supporting port-forward, exec

Usage
## port-forward
```
gokubectl -action=port-forward -p 8123:8080 -pn=podName
```
## exec
```
go run gokubectl.go -action=exec -pn=log -test
```