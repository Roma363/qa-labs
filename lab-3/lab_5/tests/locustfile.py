from locust import HttpUser, task, between, LoadTestShape

class ServerTester(HttpUser):
    wait_time = between(0.1, 1.0)

    @task(2)
    def load_index(self):
        self.client.get("/")

    @task(1)
    def load_page2(self):
        self.client.get("/page2.html")

    @task(1)
    def load_404(self):
        self.client.get("/404")

# ==========================================
# СЦЕНАРІЙ 1: Stress-тестування (Поступове зростання)
# ==========================================
class StepLoadShape(LoadTestShape):
    """
    Поступове зростання навантаження (Stress Test).
    Кожні 10 секунд додається 50 нових користувачів,
    щоб знайти точку, де сервер "ламається" або починає гальмувати.
    """
    step_time = 10
    step_users = 50
    spawn_rate = 10
    time_limit = 600 # Тест триває максимум 10 хвилин

    def tick(self):
        run_time = self.get_run_time()
        if run_time > self.time_limit:
            return None # Зупинка тесту

        current_step = int(run_time / self.step_time) + 1
        return (current_step * self.step_users, self.spawn_rate)


# ==========================================
# СЦЕНАРІЙ 2: Spike-тестування (Різкий стрибок)
# Щоб виконати Spike-тест, ЗАКОМЕНТУЙТЕ клас StepLoadShape вище
# і РОЗКОМЕНТУЙТЕ клас SpikeLoadShape нижче.
# ==========================================
# class SpikeLoadShape(LoadTestShape):
#     """
#     Spike Test: Нормальне навантаження, потім різкий стрибок,
#     потім знову повернення до нормального навантаження.
#     """
#     time_limit = 300 # 5 хвилин
#     spawn_rate = 100
#
#     def tick(self):
#         run_time = self.get_run_time()
#
#         if run_time < 60:
#             # 1-ша хвилина: 50 користувачів (нормальне навантаження)
#             return (50, 10)
#         elif run_time < 120:
#             # 2-га хвилина: РІЗКИЙ СТРИБОК до 1000 користувачів
#             return (1000, self.spawn_rate)
#         elif run_time < 180:
#             # 3-тя хвилина: повернення до 50 користувачів
#             return (50, 10)
#         elif run_time < self.time_limit:
#              return (50, 10)
#
#         return None