# Tar1SincronoSD
Parte sincronica de la tarea de Sistemas distribuidos 2023-2

Cambiar versiones de GO y Compose en 'DockerFile'(opcional) y 'compose.yml'(obligatorio) de acuerdo a sus instalados si lo quieren probar local.

Pensar como montar esto en maquinas virtuales ya que no habra localhost, debemos averiguar como hacer que se comuniquen con ip + puerto. -> esto afectara al compose actual.

Para crear contenedor individual.
    "docker build . -t 'contenedor:lastest'"
Para correr contenedor individual
    "docker run 'contenedor:lastest'"
Para correr Docker compose
    'Docker-compose up'

Para compilar proto
    'protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     <carpeta/file.proto>'


Cosas varias.
    La conexion sincrona actual es request reply. es decir los clientes deben pedir antes de recibir algo.
    Debe haber una opcion de broadcast para el central, asi enviar sin que pidan los clientes.

    Falta logica de escritura en archivos (central).
    falta logica calculo de keys en regionales.

