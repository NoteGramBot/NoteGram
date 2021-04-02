# Tareas para Notegram

Tareas para construif el bot.
De momento "solo" tenemos un build de que compila el bot en el directorio ./bin del proyecto.

## build

> Compila y crea el ejecutable del proyecto

~~~sh
echo "Building project (output in bin)"
BINDIR=$PWD/bin
cd src ; go build -o $BINDIR
~~~

## check

> Pasa el linter al proyecto
~~~sh
echo "Building project (output in bin)"
BINDIR=$PWD/bin
cd src ; go vet Notegram Notegram/core Notegram/tg Notegram/data
~~~

## depend

> Instala dependencias

~~~sh
echo "Instalando dependencias"
BINDIR=$PWD/bin
cd src ; go get github.com/go-telegram-bot-api/telegram-bot-api
~~~


## test

> Pasa los test del proyecto
~~~sh
cd src ; go test ./...
~~~


## hello

> Just say jellow

~~~sh
echo Jellow
~~~


## yaml
> Check Yaml files
> Comprueba que los ficheros yaml están bien.
> Prueba a cargar agil.yaml y .travis.yml

~~~python
import yaml

files = [ '.travis.yml','agil.yaml' ] 

for ff in files:
   print(f"Cargando {ff}")
   with open(ff, 'r') as fp:
      yml = fp.read()

   print(f"Parseando {ff}")
   out = yaml.load(yml)

   print(f"Resultado:\n{out}")
~~~
