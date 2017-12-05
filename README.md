# Remote Download Server
[Work in Progress]
a simple remote download server with only purpose is to download from URL provided. 

### Main purpose of this server 
  - Created for Raspberry PI whcih will act as remote downloader 
  - Download process can be started from anywhere. 
 
### Start the server.

```sh
$ cd remote-download-server
$ go run server.go
```
### Start a Download Process
```sh
POST /downloads
{ 
 "url" : "url for Item to be downloaded"
}
```

### Get the List of Download Process 
```sh
GET /downloads
```

### Todos
 - Pause the Download Process
 - Add Torrentz 

License
----
MIT


