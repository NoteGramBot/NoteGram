# Bot sencillo de Telegram usando lenguaje Go.

## Problema

Se quiere poder anotar cosas desde telegram y crear notas compartidas en sus
grupos.

## Solución

Se propone una aplicación que, mediante el uso de alguna biblioteca que implemente la
API de bots de telegram, almacene y muestre las notas o anotaciones enviadas por los usuarios.

# Instrucciones

## Compilación

Este proyecto utiliza [Mask](https://github.com/jakedeichert/mask) como runner.
Antes de compilar y testear el código hay que instalar mask (basta con descargarlo desde la [página de releases de Mask](https://github.com/jakedeichert/mask/releases).

Antes de compilar hay que cargar las dependencias con ```mask depend```. Esto obtendrá los paquetes externos de los que depende el bot.

Para ejecutar los test:
```
mask test
```

Para compilar el proyecto ejecutamos ```mask build``` desde el directorio raíz del proyecto.

El ejecutable de salida estará en el la carpeta 📁 /bin del proyecto.

## Instalación/Uso (WIP)

Para utilizar el bot necesitamos un Token.

- Hay que hablar con el [@botfather](https://t.me/botfather) para crear un bot
y que nos dé un token para usar con el bot.
- Ver: [Bots: An introduction for developers](https://core.telegram.org/bots) 
para más detalles.

Notegram necesita un fichero de configuración JSON. 

En [config_sample.json](./config_sample.json) tienes un fichero de ejemplo:


```json
{
"secret": "telegram_secret",
"dbhost":"123.34.45.56",
"dbport": 27017, 
"dbuser": "scott",
"dbpass": "tiger",
"dbcollection": "notegram",
"loglevel": "Debug"
}
```

| Campo | Descripción |
|-------|-------------|
| secret | Secret proporcionado por [el Botfather](https://web.telegram.org/#/im?p=@BotFather) |
| dbhost | Direccion IP o hostname del servidor de bases de datos Mongodb |
| dbport | Puerto en el que escucha el servidor |
| dbuser | Usuario para conectarse a la BBDD |
| dbpass | Clave de acceso de conexión a la BBDD |
| dbcollection | Collection ('tabla') de mongodb en la que almacenamos los mensajes de los usuarios |

# Hacking

Cosas a tener en cuenta para modificar / aportar código al proyecto.

🚨🚨🚨 Atención, Achtung, Cuidadín: Todos los issues se tienen que cerrar con un commit 🚨🚨🚨

## Estructura del proyecto:

- 📁 [docs](docs) Documentación
- 📁[src](src) Código de la aplicación
- 📁[bin](bin) - Ficheros ejecutables de la aplicacion

## Tecnologías

* Logging: Se usa la biblioteca ['log'](https://golang.org/pkg/log).
* Almacenamiento: Se usa [MongoDB](https://www.mongodb.com/es).
	* Se crea una colección de Anotaciones que almacene el id del chat, y la
	anotación en cuestión.
* Configuración: Se usa un fichero JSON que almacene la configuración junto
con la biblioteca ['encoding/json'](https://golang.org/pkg/encoding/json/).


## Equipo:

* @delightfulagony
* [@igponce](https://github.com/igponce)
* @ILoveYouDrZaius
* @murcian0

```:qw```
