from locust import HttpUser, task, between, LoadTestShape, events

@events.init_command_line_parser.add_listener
def _(parser):
    parser.add_argument(
        "--test-type",
        type=str,
        choices=["stress", "spike"],
        default="stress",
        help="Оберіть тип тесту: stress або spike"
    )

class ServerTester(HttpUser):
    # =======================================================================
    # ЗМІНА НА МІЙ ВИБІР: Реалістична поведінка користувачів
    # =======================================================================
    # Раніше було between(0.1, 1.0) — це змушувало кожного юзера спамити сервер
    # як кулемет (до 10 запитів на секунду!). Через це сервер ліг уже на 50 юзерах.
    # Змінюємо на 1.0 - 3.0 секунди. Тепер юзери роблять паузи, як живі люди.
    # Це дозволить серверу витримати БІЛЬШЕ користувачів, а графік буде красивим.
    wait_time = between(1.0, 3.0)

    @task(2)
    def load_index(self):
        self.client.get("/")

    @task(1)
    def load_page2(self):
        self.client.get("/page2.html")

    @task(1)
    def load_404(self):
        self.client.get("/404")

class CustomLoadShape(LoadTestShape):
    def tick(self):
        run_time = self.get_run_time()
        current_test_type = self.runner.environment.parsed_options.test_type

        if current_test_type == "spike":
            # =======================================================================
            # СЦЕНАРІЙ: Оптимізований Spike-тест (Компактний і чіткий)
            # =======================================================================
            # Скорочуємо фази, щоб увесь тест тривав 3 хвилини (180 сек).
            # На короткому проміжку часу "скачок" на графіку виглядає ефектніше.

            if run_time < 30:
                # 1. Стартове плато: 30 секунд тримаємо 20 юзерів (база)
                return (20, 20)

            elif run_time < 90:
                # 2. СТРИБОК: Зменшуємо пік з 1000 до 400 юзерів (оскільки сервер слабкий),
                # але піднімаємо spawn_rate до 100! Організуємо стрибок за 4 секунди.
                return (400, 100)

            elif run_time < 180:
                # 3. ВІДНОВЛЕННЯ: Повертаємося до 20 юзерів.
                # Ставимо RATE = 100, щоб Locust миттєво (за ті ж 4 секунди) вбив зайвих юзерів.
                return (20, 100)

            return None

        else:
            # Сценарій Stress-тесту (теж адаптований під новий wait_time)
            step_time = 15
            step_users = 30
            spawn_rate = 10
            time_limit = 300

            if run_time > time_limit:
                return None

            current_step = int(run_time / step_time) + 1
            return (current_step * step_users, spawn_rate)