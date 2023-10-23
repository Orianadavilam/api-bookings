### Microservicio de Reservas

Hola, bienvenido a este repositorio de código fuente de una API escrita en Go (Golang) diseñada para gestionar reservas, en principio hecha para vuelos.

### Responsabilidad Principal

La principal función de este microservicio es permitir crear, modificar y eliminar reservas, y listar todas las reservas hechas por un usuario con el respectivo detalle de las mismas.

### Requerimientos previos

**Instalaciones y configuraciones**

1. **Golang:** Es necesario tener instalado el lenguaje de programación Golang en el entorno donde se quiera ejecutar la aplicación.
2. **PostreSQL:** Es necesario tener instalado el servidor de PostgreSQL para que la aplicación funcione correctamente.

### Instrucciones para levantar el microservicio de manera local

1. Cloná este repo en tu disco local:

```
git clone https://github.com/Orianadavilam/api-bookings.git
```

2. Navegá hasta la carpeta del proyecto que acabas de clonar y abrela con VS Code.
3. Abrí una terminal dentro del proyecto clonado
4. Ejecuta el microservicio:

```
go run *.go
```

El microservicio estará disponible en `http://localhost:9000`

### Autenticación

Actualmente este repositorio no se encuentra autenticado.

## API Documentation

### GET /api/bookings/healthy

Nos permite verificar si la aplicación se encuentra levantada o no. Responde {"UP"} si está levantada.

### Créditos

- [Oriana Dávila][1]

[1]: https://github.com/Orianadavilam
