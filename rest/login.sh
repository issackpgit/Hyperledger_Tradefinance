curl -X POST -d @login.json http://127.0.0.1:5000/registrar
sleep 1
curl -X POST -d @login_jim.json http://127.0.0.1:5000/registrar


