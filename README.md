Мини-ТЗ: Order Management Service


🎯 Цель проекта
Микросервис для управления клиентами и их заказами с возможностью сегментации клиентов.
🛠 Технологии
Язык: Go
Фреймворк: gorilla/mux (роутинг)
ORM: gorm (Postgres)
База данных: PostgreSQL (поднята через Docker Hub)
Инфраструктура: Docker (контейнеризация сервиса и базы)


📦 Сущности
Customer
ID, Name, Email, Phone, Promocode
Связи:
1 → M Orders
M ↔ M CustomerSegments
Order
ID, CustomerID, Status, Price
Связи:
M → 1 Customer
CustomerSegments
ID, Name
Связи:
M ↔ M Customers
CustomerSegmentsLink (таблица связки Many-to-Many)
CustomerID, CustomerSegmentID



⚙️ Архитектура проекта
repository/ — слой работы с БД (CRUD для каждой сущности, через GORM).
services/ — бизнес-логика (CRUD, вызов repository, возврат DTO).
handler/ — HTTP-слой (GET, POST, PUT, DELETE, вызов services).
dto/ — DTO-модели (например, CustomerDTO без поля Promocode).
mappers/ — преобразование сущностей в DTO.
main.go
Инициализация repository, services, handler.
Настройка роутера (mux.NewRouter()).
Поднятие локального сервера на :8082.




🔗 REST API (эндпоинты)
Примеры:
GET /customers → список клиентов (DTO без Promocode)
POST /customers → создать клиента
PUT /customers/{id} → обновить клиента
DELETE /customers/{id} → удалить клиента
То же самое реализовано для orders и customer_segments.
