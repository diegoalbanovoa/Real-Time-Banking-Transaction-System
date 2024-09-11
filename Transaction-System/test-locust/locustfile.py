import random
from locust import HttpUser, TaskSet, task, between

# Define el comportamiento de las transacciones bancarias para los usuarios simulados
class BankTransactionBehavior(TaskSet):

    @task(1)
    def deposit(self):
        """
        Simula una transacción de depósito.
        Se selecciona un ID de cuenta aleatorio entre 1 y 100 y un monto de depósito aleatorio entre 10 y 1000.
        Luego, se envía una solicitud HTTP POST al endpoint /deposit con estos datos en formato JSON.
        """
        account_id = random.randint(1, 100)  # Simula una cuenta con un ID aleatorio entre 1 y 100
        amount = round(random.uniform(10.0, 1000.0), 2)  # Genera un monto de depósito aleatorio entre 10.0 y 1000.0
        self.client.post("/deposit", json={"account_id": account_id, "amount": amount})  # Envía la solicitud POST

    @task(1)
    def withdraw(self):
        """
        Simula una transacción de retiro.
        Se selecciona un ID de cuenta aleatorio entre 1 y 100 y un monto de retiro aleatorio entre 10 y 500.
        Luego, se envía una solicitud HTTP POST al endpoint /withdraw con estos datos en formato JSON.
        """
        account_id = random.randint(1, 100)  # Simula una cuenta con un ID aleatorio entre 1 y 100
        amount = round(random.uniform(10.0, 500.0), 2)  # Genera un monto de retiro aleatorio entre 10.0 y 500.0
        self.client.post("/withdraw", json={"account_id": account_id, "amount": amount})  # Envía la solicitud POST


# Define el usuario virtual que ejecuta las transacciones bancarias
class WebsiteUser(HttpUser):
    """
    Clase que simula un usuario que ejecuta las tareas definidas en BankTransactionBehavior.
    Los usuarios esperan entre 1 y 3 segundos entre cada tarea, simulando un comportamiento más realista.
    """
    tasks = [BankTransactionBehavior]  # Asigna las tareas que el usuario ejecutará
    wait_time = between(1, 3)  # Tiempo de espera aleatorio entre 1 y 3 segundos entre solicitudes
