# Password Manager Service (Golang)

## How to run app locally
Run command -> `air`

## How to deploy
- Open instance SSH
- Check PID of running app -> `ps aux | grep password-manager-service`
- Kill running app using PID -> `kill {PID}`
- Run app on background -> `nohup ./password-manager-service &`
- Exit SSH -> `exit`

## TODO (deploy.yml)
- checkout to virtual env
- install go
- set environments
- build
- push build to build folder
- connect to vm
- stop currently running apache2 process
- go to build branch and git pull
- run recently pulled build using apache2