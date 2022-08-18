# blog-app

API для блога.

Для запуска:
    
    git clone https://github.com/Kolyan4ik99/blog-app
    cd blog-app
    make

URL для всех запросов

    http://localhost:8080/

Нужно авторизоваться: создать пользователя

    Request:
    POST /auth/sign-up
    {
    "name": "NameNewUser",
    "password": "14125",
    "email": "mailForNewUser@gmail.com"
    }

    Response:
    status: 201
    user_id: 1


Создать новый пост. Важно: author должен пройти этап sign-up

    POST /api/post/
    {
    "header": "Nikol253",
    "text": "hwerhwkrmbe",
    "author": 1
    }

    Response:
    status: 201
    Post was successful created

Получить все существующие посты:

    GET /api/post/

    Response:
    status: 200
    [{
    "author":1,
    "header":"Nikol253",
    "text":"hwerhwkrmbe"
    }]

Получить пост по id:

    GET /api/post/1

    Response:
    status: 200
    [{
    "author":1,
    "header":"Nikol253",
    "text":"HWEHWEGWERgqgweg"
    }]

    GET /api/post/51235

    Response 
    status: 404
