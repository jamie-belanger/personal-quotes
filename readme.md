# Personal Quotes Library
Years ago, I wrote a basic system for managing personal quotes in PHP. This library contains a few hundred quotes I've collected over the years, including inspirational sayings like

> The brick walls you run into in life are not there to keep you out. They are there to stop the people who don't want it badly enough. They are there to stop the <i>other</i> people.
> 
> <center>&mdash; Randy Pausch</center>

and funny quotes like

> A guy came and knocked on my door. He said to me, "They're starving in Africa."
> 
> I said, "Hey, I've been here all day, pal."
> 
> <center>&mdash; John Pinette, Show Me The Buffet (1998)</center>

I've been wanting to rewrite this library in a modern language for a long time now... and actually tried twice but got busy with life. Last week I started teaching myself **[Go](https://go.dev)** and thought about this project again.

So, here we are.


## Getting Started
Right now what's here is very basic. You can start the server with 
```bash
go run ./cmd/web -port 1234
```
and use any port you want (or omit that and use the default, `4000`). It might take a little bit to build the first time.

The application currently only supports SQLite, and it will create a database file under the `./data` folder for you if it's not there.

Then call the endpoints with curl like this:

Create a quote:
```bash
curl -X POST -d "body=test&author=test" http://localhost:4000/api/quotes
```

Get a random quote:
```bash
curl -X GET http://localhost:4000/api/quotes/random
```

Get a specific quote:
```bash
curl -X GET http://localhost:4000/api/quotes/1
```

Update a quote:
```bash
curl -X PUT -d "body=test!&author=test+lots" http://localhost:4000/api/quotes/1
```

Delete a quote:
```bash
curl -X DELETE http://localhost:4000/api/quotes/1
```



## TODO

I'll be committing over the next few days, weeks, or years, starting with the API and database parts. The intention of this project is to one day have:

* API = so you can query this service from anything else in your network, app, or suite
* Docker = will come soon
* Database = storage in simple SQLite, but I also want to support MySQL and PostgreSQL
* UI = eventually I'll make a simple management UI, with OIDC for easier permissions

I have other things I'd love to play with (can you use Swagger from Go? GraphQL?) but this is more than enough for now. I'll get the groundwork in place and then revamp this TODO into a roadmap later.


## License
I haven't decided yet. I'm leaning toward MIT and just giving this to the world. But whatever I choose, I promise I won't enshittify the code with paid features.


## Dependencies
Aside from Go itself, I also used the following packages:
 * **[Blue Monday](https://github.com/microcosm-cc/bluemonday)** (and its dependencies)


## About AI
Big topic these days, so I'll be clear: I'm a programmer. I've been writing code for over 40 years now. Right now, I'm teaching myself Go and I am building this with minimal assistance &mdash; basically I ask AI questions, sometimes, because that helps me learn Go better. At no point am I letting AI write the code for me.
