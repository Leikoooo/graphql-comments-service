# Тестовое задание в озон 20.04.2024

# Задача:

Реализовать систему для добавления и чтения постов и комментариев с использованием GraphQL, аналогичную комментариям к постам на популярных платформах, таких как Хабр или Reddit.

Характеристики системы постов:
•Можно просмотреть список постов.
•Можно просмотреть пост и комментарии под ним.
•Пользователь, написавший пост, может запретить оставление комментариев к своему посту.

Характеристики системы комментариев к постам:
•Комментарии организованы иерархически, позволяя вложенность без ограничений.
•Длина текста комментария ограничена до, например, 2000 символов.
•Система пагинации для получения списка комментариев.

(*) Дополнительные требования для реализации через GraphQL Subscriptions:
•Комментарии к постам должны доставляться асинхронно, т.е. клиенты, подписанные на определенный пост, должны получать уведомления о новых комментариях без необходимости повторного запроса.

Требования к реализации:
•Система должна быть написана на языке Go.
•Использование Docker для распространения сервиса в виде Docker-образа.
•Хранение данных может быть как в памяти (in-memory), так и в PostgreSQL. Выбор хранилища должен быть определяемым параметром при запуске сервиса.
•Покрытие реализованного функционала unit-тестами.

Критерии оценки:
•Как хранятся комментарии и как организована таблица в базе данных/in-memory, включая механизм пагинации.
•Качество и чистота кода, структура проекта и распределение файлов по пакетам.
•Обработка ошибок в различных сценариях использования.
•Удобство и логичность использования системы комментариев.
•Эффективность работы системы при множественном одновременном использовании, сравнимая с популярными сервисами, такими как Хабр.
•В реализации учитываются возможные проблемы с производительностью, такие как проблемы с n+1 запросами и большая вложенность комментариев.


# Запуск

### 1) Настройте .env файл для выбора нужного репозитория postgres / graph<br>
### 2) Запустите Docker-compose

# EndPoints

## Эндпоинты для postgres

GET **localhost:8080/getAllPosts** - возвращает все созданные посты

`{}`
<br><br>
POST **localhost:8080/savePost** - создает пост

`{
"title":"test",
"content":"testContent",
"authorID":"1",
"allowComments": false
}`


<br><br>
POST **localhost:8080/saveComment** - создает комментарий к посту

`{
"postID": "2",
"parentID": "1",
"content": "Good",
"authorID": "1"
}`

<br><br>
GET **localhost:8080/getCommentsByPostId?id=`{id}`&page=`{page}`&pageSize=`{page_size}`**

`{}`

<br><br>
## Эндпоинты для graph

**localhost:8080/query**

**CreatePost**

`mutation CreatePost {
createPost(
input: {
title: "title2"
content: "content"
authorID: "authorID"
allowComments: true
}
) {
id
title
content
authorID
allowComments
createdAt
updatedAt
}
}`

**GetPosts**

`query GetPosts {
getPosts {
id
title
content
authorID
allowComments
createdAt
updatedAt
}
}`

**GetPostByID**

`mutation GetPostByID {
getPostByID(id: "22b542d4-336a-458b-a85b-abec1e230898") {
id
title
content
authorID
allowComments
createdAt
updatedAt
}
}`

**CreateComment**

`mutation CreateComment {
createComment(
input: { postID: "1", content: "test", authorID: "1", parentID: "" }
) {
id
postID
parentID
content
authorID
createdAt
updatedAt
repliesCount
}
}
`

**GetCommentsByPostID**

`query GetCommentsByPostID {
getCommentsByPostID(
postID: "1"
maxDepth: 10
pageSize: 10
page: 1
maxReplies: 10
) {
id
postID
parentID
content
authorID
createdAt
updatedAt
repliesCount
replies {
id
postID
parentID
content
authorID
createdAt
updatedAt
repliesCount
}
}
}`

