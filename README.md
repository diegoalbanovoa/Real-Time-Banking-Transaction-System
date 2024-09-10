# Real-Time Banking Transaction System

## Descripción del proyecto
Este proyecto simula un sistema de transacciones bancarias en tiempo real. Está diseñado para crear cuentas bancarias, generar transacciones aleatorias (depósitos y retiros) y almacenarlas en una base de datos MySQL.

### El programa incluye:
- **Generación de cuentas bancarias**: Se generan cuentas con un número único y un balance inicial aleatorio.
- **Generación de transacciones**: Se crean transacciones aleatorias para cada cuenta, ya sea depósitos o retiros, y se registran en la base de datos.
- **Simulación de carga**: El sistema puede ser utilizado para generar grandes cantidades de cuentas y transacciones, ideal para simular escenarios de carga y realizar pruebas de rendimiento.

## Requisitos
- Go (Golang) instalado en tu sistema.
- Docker y Docker Compose para crear y gestionar el entorno de la base de datos MySQL.

## Configuración del entorno Docker
El proyecto incluye un archivo de configuración Docker Compose que te permitirá iniciar fácilmente un contenedor con MySQL.

### Configuración de `docker-compose.yml`
El siguiente archivo `docker-compose.yml` define un contenedor con MySQL:

```yaml
version: '3'
services:
  mysql:
    image: mysql:8.0
    container_name: mysql-container
    environment:
      MYSQL_ROOT_PASSWORD: rootpassword
      MYSQL_DATABASE: bankdb
      MYSQL_USER: bankuser
      MYSQL_PASSWORD: bankpassword
    ports:
      - "3306:3306"
    volumes:
      - mysql-data:/var/lib/mysql
      - ./db-init:/docker-entrypoint-initdb.d
    networks:
      - bank-network

volumes:
  mysql-data:

networks:
  bank-network:
```

### Instrucciones para iniciar Docker

1. Clona este repositorio en tu máquina local.

```yaml
git clone https://github.com/diegoalbanovoa/Real-Time-Banking-Transaction-System.git
cd real-time-banking
```

2. Inicializa Docker: Asegúrate de tener Docker y Docker Compose instalados.

3. Inicia el contenedor MySQL:
```yaml
docker-compose up -d
```
Esto descargará la imagen de MySQL, creará un contenedor con la base de datos bankdb y lo expondrá en el puerto 3306. Además, el contenedor será inicializado con un usuario bankuser y una contraseña bankpassword.