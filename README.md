## The most overengineered todo list

I was curious about how overengineered a Todo list (everyone's loved and hated) can be

I guess it's something similar to LinkTree. Hosted on [contabo](https://contabo.com/)

The idea started with "why tf do you need to pay LinkTree for analytics and fancy features" it's an overbloated next.js app + gatsby site (it has gotten much slower over the years) and the project itself its a 1 person product can be built in 3 days (I meant about 72 hours).

~~And it was.~~

### This project uses:

- [Chi](https://go-chi.io/#/) for backend.
- [templ](https://templ.guide/) as frontend to give me an experience similar to React (not quite xd)
- [htmx](https://htmx.org/) (just a little bit)
- [Alpine](https://alpinejs.dev/) for interactivity.
- [ent](https://entgo.io/) the only ORM and query builder in Golang feels like [drizzle](https://orm.drizzle.team/) in Golang
- [Sentry](https://docs.gofiber.io/contrib/fibersentry/) for error monitoring.
- [AxiomFM](https://axiom.co) for analytics (not sure might change later). At the moment see current implementation details in `pkg/monitor/axiom.go`

oh yeah, and a [docker compose file](docker-compose.yml) to help me (or you) to scaffold database during development.


## To run this project

```golang
git clone git@github.com:zmzlois/LinkGoGo.git
cd LinkGoGo
cp .env.example .env 
pnpm run dev
```
see [environment variables](#environment-variables) to find where you can obtain values for this project.

`pnpm run dev` is a built in pipeline to run all modules and update templ generationg concurrently. See `package.json` for what's inside this command

Alternatively if you need to start with a Postgres database you can do a `docker compose up -d`. It will 

1. Start the app 
2. Start a Postgres database (If you wish to disable this, comment out sections according to the `docker-compose.yml` file)


# Highlights

## Bare minimal authentication, but secure 
The implementation details can be seen in [pkg/auth](./pkg/auth) folder and 


## Environment Variables 

To obtain credentials for this project you will need 
- An [Axiom](https://app.axiom.io) account
- A [Sentry](https://sentry.io) account
- A Discord app registered in [development portal](https://discord.com/developers/applications)

<dl>
 <dt> URL
 <dd> The url you are running it. Might be `http://localhost:3000` or `https://linkgogo.domain.com` depends on how you will run it in production.
 <dt> Environment
 <dd> This value is used by Sentry to understand whether the error happened in `production` or `development`. See [Sentry guide](https://docs.sentry.io/platforms/go/configuration/environments/) for how this works.
 <dt>PORT</dt>
 <dd>On convention it's either `:3000` or `:8080`. Remember to put a colon!!!
 <dt>DB_HOST
 <dd>If you are using docker compose the value would be `localhost`
 <dt>DB_PORT, DB_USER, DB_PASSWORD, DB_NAME
 <dd>You can modify these values in docker-compose.yml
 <dt>JWT_SECRET_KEY
 <dd>You can generate one by `openssl rand -hex 32` or other commands. See more about the required secret key standards on [jtw.io](https://jwt.io/introduction)
 <dt>AXIOM_TOKEN 
 <dd>Find this value on your organisation's setting page `API token` tab.
 <dt>AXIOM_ORG_ID
 <dd>Copy paste the string under your organisation's name on setting page (there is a copy pasta button!!!)
 <dt>DISCORD_APPLICATION_ID, DISCORD_PUBLIC_KEY
 <dd>Required. You should find the values on your application's `General Information` page
 <dt>DISCORD_CLIENT_ID
 <dd>Required. Find this on `OAuth2` tab. 
 <dt>DISCORD_CLIENT_SECRET
 <dd> Required. Generate this value on `OAuth2` tab. 
 <dt>DISCORD_REDIRECT_URI
 <dd>Required. If you are developing locally the url would be `http://localhost:port/discord-callback` (please don't add `s` in http it will cause redirection failure). If you are deployed to production with SSL it should be `https://domain.com/discord-callback`. Other than providing it in `.env` file you will also need to configure this value on application's `OAuth2` tab. 
 <dt>DISCORD_OAUTH_URL
 <dd>Optional. You don't need to provide this value because the `Auth` package helps you generate this based on values above and the scope you selected. 
 <dt>Scopes (see in `pkg/auth/entry.go` and `pkg/auth/scope.go`)
 <dd> A list of scopes(string value) seperated by space. 
 <dt>Sentry DSN
 <dd>Once you register your sentry account, sentry will ask you for a language and the dsn value will be provided in the documentation + setting page. 
</dl>



### Things to be improved on 

- OAuth
  - At the moment OAuth package is customised for discord authentication, but it should be rewritten to fit standard OAuth authentication by providing different credentials and scope
- Logging
  - The logging 


