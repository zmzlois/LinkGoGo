## The most overengineered todo list

Curious about how overengineered a Todo list (everyone's loved and hated) can be

I guess it's something similar to LinkTree. Hosted on [contabo](https://contabo.com/)

The idea started with "why tf do you need to pay LinkTree for analytics and fancy features" it's an overbloated next.js app + gatsby site (it has gotten much slower over the years) and the project itself its a 1 person product can be built in 3 days (I meant about 72 hours).

~~And it was.~~

### This project uses:

- [Chi](https://docs.gofiber.io/) for backend.
- [templ](https://templ.guide/) as frontend to give me an experience similar to React (not quite xd)
- [htmx](https://htmx.org/) (just a little bit)
- [Alpine](https://alpinejs.dev/) for interactivity.
- [ent](https://entgo.io/) the only ORM and query builder in Golang feels like [drizzle](https://orm.drizzle.team/) in Golang
- [Sentry](https://docs.gofiber.io/contrib/fibersentry/) for error monitoring.
- [AxiomFM](https://axiom.co) for analytics with [Fiberzap](https://arc.net/l/quote/bifhxfck) as an adapter. (not sure might change later)

oh yeah, and a [docker compose file](docker-compose.yml) to help me (or you) to scaffold database during development.
