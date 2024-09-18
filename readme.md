Build command

`env GOOS=linux GOARCH=amd64 go build` 

upload to server

```scp -i ~/Downloads/ssh-key-2024-09-18.key templates/home.html opc@144.25.93.47:~/  ```

run on server 

`./partner-dummy-env`
