# Go server with stuff I (mostly) like

- [Sass lang](https://sass-lang.com/)
- [Templ](https://templ.guide/)
- MySQL connection (docker compose file included with adminer, if preferred)
- Logging Middleware
- Google Oauth - Implemented as far as authentication only. No sessions, login vs register or authorization are built-in.
- [Htmx](https://htmx.org/) optionally included, for those pesky user needs
- [Air](https://github.com/air-verse/air) live reload

## Setup

- Setup a new project in Google Cloud Console with consent screen & oauth credentials
- Copy .env.sample files at root and in ./api to .env and fill out
- Find and replace "gotemplate" with project name
- Spin-up MySQL db in docker or any other way. Provided docker compose file will init mysql with whatever DB_NAME is in .env
- run air or main.go to confirm things are happy
- watch your sass ```sass --watch scss:static/css```
