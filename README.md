# Password Manager Service (Golang)

## How to run app locally
Run command -> `air`

## How to deploy
- Open instance SSH
- Check PID of running app -> `ps aux | grep password-manager-service`
- Kill running app using PID -> `kill {PID}`
- Run app on background -> `nohup ./password-manager-service &`
- Exit SSH -> `exit`