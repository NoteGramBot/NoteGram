# Bot sencillo de Telegram usando lenguaje Go.

🚨🚨🚨 Atención, Achtung, Cuidadín: Todos los issues se tienen que cerrar con un commit 🚨🚨🚨

## Problema

Se quiere poder anotar cosas desde telegram y crear notas compartidas en sus
grupos.

## Solución

Se propone una aplicación que, mediante el uso de alguna biblioteca que implemente la
API de bots de telegram, almacene y muestre las notas o anotaciones enviadas por los usuarios.

### Tecnologías

* Logging: Se usa la biblioteca ['log'](https://golang.org/pkg/log).
* Almacenamiento: Se usa [MongoDB](https://www.mongodb.com/es).
	* Se crea una colección de Anotaciones que almacene el id del chat, y la
	anotación en cuestión.
* Configuración: Se usa un fichero JSON que almacene la configuración junto
con la biblioteca ['encoding/json'](https://golang.org/pkg/encoding/json/).

## Instalación/Uso (WIP)

Para utilizar el bot necesitamos un Token.

- Hay que hablar con el [@botfather](https://t.me/botfather) para crear un bot
y que nos dé un token para usar con el bot.
- Ver: [Bots: An introduction for developers](https://core.telegram.org/bots) 
para más detalles.

## Estructura del proyecto:

📁[docs](docs) Documentación

📁[docs/Diseño Funcional](docs/Diseño_Funcional.md) - Diseño funcional de la 
aplicación (alto nivel) 

📁[tests](tests) - test_de_funcionalidad/unitarios/etc.

📁[data](data) - Datos de la aplicación

## Equipo:

* @delightfulagony
* @igponce
* @ILoveYouDrZaius
* @murcian0
