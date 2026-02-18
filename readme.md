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


## TODO

I'll be committing over the next few days, weeks, or years, starting with the API and database parts. The intention of this project is to one day have:

* API = so you can query this service from anything else in your network, app, or suite
* Docker = will come soon
* Database = storage in simple SQLite, but I also want to support MySQL and PostgreSQL
* UI = eventually I'll make a simple management UI, with OIDC for easier permissions

I have other things I'd love to play with (can you use Swagger from Go? GraphQL?) but this is more than enough for now. I'll get the groundwork in place and then revamp this TODO into a roadmap later.


## License
I haven't decided yet. I'm leaning toward MIT and just giving this to the world. But whatever I choose, I promise I won't enshittify the code with paid features.


## About AI
Big topic these days, so I'll be clear: I'm a programmer. I've been writing code for over 40 years now. Right now, I'm teaching myself Go and I am building this with minimal assistance &mdash; basically I ask AI questions, sometimes, because that helps me learn Go better. At no point am I letting AI write the code for me.
