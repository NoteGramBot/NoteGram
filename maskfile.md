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

## hello

> Just say jellow

~~~sh
echo Jellow
~~~

