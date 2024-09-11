# Real-Time Banking Transaction System

## Descripción del Proyecto
Este es un sistema de transacciones bancarias en tiempo real, diseñado siguiendo los principios de **Domain-Driven Design (DDD)**. El sistema permite realizar depósitos y retiros en cuentas bancarias de forma segura y eficiente, utilizando la arquitectura **hexagonal** para separar la lógica del dominio de la infraestructura. El proyecto está desarrollado en **Go** para el backend, con **MySQL** como base de datos, y está preparado para pruebas de carga con **Locust**.

## Características
- **Transacciones bancarias en tiempo real**: Depósitos y retiros en cuentas de usuario.
- **Arquitectura DDD**: Código organizado en capas siguiendo los principios de Domain-Driven Design, separando la lógica de dominio de las capas de infraestructura y aplicación.
- **Perfiles de rendimiento**: Monitoreo de la aplicación utilizando `pprof` y `trace` para identificar problemas de rendimiento.
- **Pruebas de carga**: Pruebas de carga automatizadas con Locust.
- **Base de datos MySQL**: El sistema utiliza MySQL para almacenar las cuentas y transacciones, con contenedores Docker para fácil configuración.

## Estructura del Proyecto
Aquí tienes el texto formateado como un archivo README.md para Git:

```bash
Transaction-System/
│
├── cmd/
│   └── bankservice/         # Punto de entrada principal del servicio de transacciones bancarias
├── internal/
│   ├── application/         # Servicios de aplicación (orquestación de la lógica de dominio)
│   ├── domain/              # Lógica del dominio (entidades y repositorios)
│   │   ├── account/         # Lógica relacionada con cuentas bancarias
│   │   ├── transaction/     # Lógica relacionada con transacciones bancarias
│   └── infrastructure/      # Implementaciones de infraestructura (repositorios y controladores HTTP)
│       ├── database/        # Implementaciones de repositorios basados en MySQL
│       └── http-conection/  # Controladores HTTP para la API REST
├── tests/                   # Pruebas unitarias y de integración
└── scripts/
    └── locust/              # Pruebas de carga con Locust

```
## Requisitos Previos
Asegúrate de tener instalados los siguientes componentes:
- **Go** (versión 1.16 o superior)
- **Docker** (para levantar la base de datos MySQL)
- **MySQL** (si no usas Docker para la base de datos)
- **Locust** (para pruebas de carga)

## Configuración del Proyecto

### Paso 1: Clonar el repositorio
```bash
git clone https://github.com/diegoalbanovoa/Real-Time-Banking-Transaction-System.git
cd transaction-system
```
### Paso 2: Configuración de la base de datos MySQL con Docker
El sistema está configurado para utilizar MySQL como base de datos, la cual se puede configurar fácilmente con Docker. El archivo docker-compose.yml ya está incluido para facilitar el levantamiento del servicio.

Levanta el contenedor de MySQL usando Docker:
```bash
docker-compose up -d
```
Verifica que el contenedor esté corriendo:
```bash
docker ps
```

La base de datos MySQL estará disponible en localhost:3306 con las siguientes credenciales (también puedes modificarlas en docker-compose.yml):

- Usuario: bankuser
- Contraseña: bankpassword
- Base de datos: bankdb

### Paso 3: Migración de tablas

Una vez que la base de datos esté corriendo, asegúrate de crear las tablas necesarias para las cuentas y las transacciones. Si estás utilizando MySQL directamente, ejecuta las siguientes consultas SQL:

```bash
CREATE TABLE IF NOT EXISTS accounts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    account_number VARCHAR(20) NOT NULL,
    balance DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    account_id INT NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    transaction_type ENUM('deposit', 'withdrawal') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES accounts(id)
);
```

### Paso 4: Ejecutar el servicio
Para ejecutar el servicio, navega hasta el directorio cmd/bankservice y utiliza el siguiente comando:

```bash
go run .
```
El servidor debería estar corriendo en http://localhost:8080 y pprof en http://localhost:6060.

### Endpoints de la API
- POST /deposit
  Realiza un depósito en una cuenta.

    Solicitud:
    ```bash
    {
  "account_id": 1,
  "amount": 500.00
    }
    ```
    Respuesta:
    ```bash
    Depósito exitoso
    ```
- POST /withdraw
  Realiza un retiro de una cuenta.
  Solicitud:
    ```bash
    {
  "account_id": 1,
  "amount": 200.00
    }
    ```
  Respuesta:
    ```bash
   Retiro exitoso
    ```
  
### Pruebas de Carga con Locust
El proyecto incluye pruebas de carga utilizando Locust. Para ejecutar estas pruebas:

Paso 1: Instalar Locust
```bash
pip install locust
```
Paso 2: Ejecutar pruebas de carga
Navega al directorio de pruebas de carga y ejecuta:
```bash
locust -f scripts/locust/locustfile.py
```

Luego, abre un navegador y dirígete a http://localhost:8089. Allí podrás configurar el número de usuarios simulados y la tasa de solicitudes.

### Pruebas Unitarias
El sistema incluye pruebas unitarias para la lógica del dominio. Para ejecutar todas las pruebas, usa el siguiente comando:

```bash
go test ./...
```

Esto ejecutará las pruebas en todos los paquetes del proyecto.
