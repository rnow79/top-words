# Top Words Demo
Welcome to "Top Words" project. This project has two parts:
- A service that accepts input as text, n as int, and provides a json with the n top used words, and times of occurence.
- A tiny frontend for service testing.

## Endpoints

There is only one effective endpoint:

`POST http://localhost/{n}{text}{ignorecase}`

- **VAR N**: [Optional, default 10, max 100]. You can define the number of results, so n=10 means Top 10 words.
- **VAR IGNORECASE**: [Optional, default 0 (false)]. Tells the service if word count is made ignoring case or not, so ignorecase=1 will be activate this feature.
- **VAR TEXT**: [Required]. The text to process.

There is also a redirecting endoint:

`GET http://localhost -> http://localhost/front/index.htm`.

In the front end page, you can test the service.