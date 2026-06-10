# Lab 5

## Запуск

```powershell
cd ../lab_5
cmake -S . -B build
cmake --build build
.\build\lab_5_server.exe
```
Open: http://localhost:8080/ or http://localhost:8080/second_page.html

## Тестування

Тестування проводиться через locust (Python):

```powershell
cd ../tests
locust -f locustfile.py
```