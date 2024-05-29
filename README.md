# Тестовое задание в озон 20.04.2024

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

