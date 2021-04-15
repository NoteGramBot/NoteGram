# Tareas para Notegram

Tareas para construif el bot.
De momento "solo" tenemos un build de que compila el bot en el directorio ./bin del proyecto.

## build

> Compila y crea el ejecutable del proyecto

~~~sh
echo "Building project (output in bin)"
BINDIR=$PWD/bin
go build -o $BINDIR
~~~

## check

> Pasa el linter al proyecto
~~~sh
echo "Building project (output in bin)"
BINDIR=$PWD/bin
go vet Notegram Notegram/core Notegram/tg Notegram/data
~~~

## depend

> Instala dependencias

~~~sh
echo "Instalando dependencias"
BINDIR=$PWD/bin
go get github.com/go-telegram-bot-api/telegram-bot-api
~~~


## test

> Pasa los test del proyecto (sin cobertura)
~~~sh
go test ./...
~~~

## coverage
> Pasa los test del proyecto y produce un fichero de cobertura
~~~sh
go test ./... -race -coverprofile=coverage.txt -covermode=atomic
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
import pprint

files = [ '.travis.yml','agil.yaml' ] 
pp = pprint.PrettyPrinter(indent=2)

for ff in files:
   print(f"Cargando {ff}")
   with open(ff, 'r') as fp:
      yml = fp.read()

   print(f"Parseando {ff}")
   out = yaml.load(yml)

   print(f"Resultado:\n")
   pp.pprint(out)
~~~

## gofmt

> Formatea (todo) el código con gofmt

~~~sh
for i in $(find . -name '*.go') ; do
   gofmt -s -w $i
done
~~~

## setupmongodb

> Crea entorno de test/desarrollo en mongodb

El entorno contiene:
 - Usuario scott para crear / acceder a notas.
 - Base de datos 'notegram'
 - Collection "notegram.notas"l
 - Crea una primera nota de ejemplo para que db.notas.find()

~~~sh
mongo -host localhost <<MONGO_EOF
use notegram

try {
   db.createUser({user:"scott", 
               pwd:"tiger",
               roles: [ { role:"readWrite", db: "notegram"}] });
} catch (err) { print(err);}

try {
   db.createCollection("notas");
} catch (err) { print(err);}

try {
   db.notas.insert({
      user: "test123",
      content: "Tercera nota (con cacteres unicode 💩💩💩 )",
      content_type: "test/plain",
      content_encoding: "utf8"
   });
} catch (err) { print(err);}

print("Setup terminado");

MONGO_EOF
~~~

## cleanmongodb

> Borra contenido de la BBDD mongodb

~~~sh
mongo -host localhost <<MONGO_EOF
use notegram

db.notas.deleteMany({});

MONGO_EOF
~~~

## dumpmongodb
> Que carajo esta guardado en MONGODB ????

~~~sh
mongo -host localhost <<MONGO_EOF
use notegram

db.notas.find({});


