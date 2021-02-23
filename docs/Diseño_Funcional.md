# Diseño Funcional

Este documento contiene una toma de requisitos de la aplicación.

## Funcionalidad Básica

- El bot escucha a uno o varios usuarios de telegram, o bien escucha en un canal de telegram.
- Mientras no se interactúa directamente con el bot mediante mensaje, o comando, el bot no almacena información de ningún tipo.
- Al interactur

### Guardado de notas
- El bot almacenará las notas o recordatorios que le envíen los usuarios.

### Recordatorio
- Cuando un usuario le pregunte, el bot le enviará las notas que tiene en memoria.

### Borrado de notas
- El usuario podrá pedir al bot que borre una nota determinada, o varias.
- Sólo podrá borrar las notas el mismo usuario que pidió al bot que crease la nota o recordatorio.

# Fuera de alcance

- Modificación de notas. En lugar de modificar una nota, habrá que borrar la nota anterior y crear una nueva.
- Queda fuera de alcance la persistencia de las notas entre distintas sesiones, reinicio del bot, etc.
- Queda también fuera de alcance todas las cuestiones relativas a seguridad.
