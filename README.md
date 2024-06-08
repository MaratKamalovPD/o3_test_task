# O3 Test Task
Тестовое задание для стажера-разработчика

## Подготовка к запуску
```
git clone git@github.com:MaratKamalovPD/o3_test_task.git
```
после чего открыть в терминале корневую директорию

## Выбор способа хранения данных

открываем `o3_test_task\internal\pkg\config\config.yaml` и меняем поле **storage_type**

| Значение        | Описание                                               |
|:----------------|:-------------------------------------------------------|
| postgres        | Хранение данных в Postgresql                           |
| inmemory        | Хранение данных в памяти                               |


![image](https://github.com/MaratKamalovPD/o3_test_task/assets/55689955/3925f839-fd93-4769-8403-5ca2ac636478)

## Запуск 
```
docker-compose down
```
```
docker-compose build main
```
### Запуск с хранением данных в памяти 
```
docker-compose up main
```
### Запуск с хранением данных в Postgresql
```
docker-compose up 
```

## Работа с API
### Работа с поставми
Запросы к постам можно делать по адресу:

```
http://localhost:8080/post
```

**Получения списка постов**
```
{posts{title, content, areCommentsDisabled, userId, createdAt}}
```

**Получение информации о конкретном посте**
```
{post(id:2){title, content, areCommentsDisabled, userId, createdAt}}
```

**Создание поста**
```
mutation _{createPost(title:"post title", content:"post content", userId:1){title, content, userId}}
```

**Отключение\включение комментариев**
```
mutation _{disableComments(postId:1, userId:1)}
```

### Работа с комментариями
Запросы к комментариям можно делать по адресу:

```
http://localhost:8080/comment
```

**Получение комментариев к посту**
```
{commentsByPost(postId:2, limit:20, offset:0){content, postId, parentCommentId}}
```

**Создание комментария**
```
mutation _{createComment(content:"comment content",  postId:3, userId:1){ content, postId, parentCommentId, userId}}
```

или если комментарий оставлен под другим комментарием

```
mutation _{createComment(content:"comment content", parentCommentId:3, postId:3, userId:1){ content, postId, parentCommentId, userId}} 
```

### Тестирование
```
go test --cover  ./...
```
