import random
from locust import HttpUser, TaskSet, task, between

class BankTransactionBehavior(TaskSet):

    @task(1)
    def deposit(self):
        # Simular un depósito aleatorio en una cuenta
        account_id = random.randint(1, 100)  # Simulamos cuentas con IDs entre 1 y 100
        amount = round(random.uniform(10.0, 1000.0), 2)  # Montos aleatorios entre 10 y 1000
        self.client.post("/deposit", json={"account_id": account_id, "amount": amount})

    @task(1)
    def withdraw(self):
        # Simular un retiro aleatorio de una cuenta
        account_id = random.randint(1, 100)  # Simulamos cuentas con IDs entre 1 y 100
        amount = round(random.uniform(10.0, 500.0), 2)  # Montos aleatorios entre 10 y 500
        self.client.post("/withdraw", json={"account_id": account_id, "amount": amount})

class WebsiteUser(HttpUser):
    tasks = [BankTransactionBehavior]
    wait_time = between(1, 3)  # Simula una espera de entre 1 y 3 segundos entre solicitudes
