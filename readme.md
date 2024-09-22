Build command

`env GOOS=linux GOARCH=amd64 go build`

upload to server

`
scp -i ~/Downloads/ssh-key-2024-09-18.key partner-dummy-env opc@144.25.93.47:~/
`

upload HTML templates

`
scp -i ~/Downloads/ssh-key-2024-09-18.key templates/home.html opc@144.25.93.47:~/templates/home.html
`


ssh to server

`
ssh opc@144.25.93.47 -i ~/Downloads/ssh-key-2024-09-18.key
`


start, restart, stop and check status of service

`
sudo systemctl start partnerenv.service
sudo systemctl restart partnerenv.service
sudo systemctl status partnerenv.service
sudo systemctl stop partnerenv.service
`

Check log of service

`
journalctl -u partnerenv
`
