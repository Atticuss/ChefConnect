## ChefConnect

A recipe storage solution built primarily via a graph database (dgraph). Current API POC in Python, MVP will migrated to Golang. Angular front-end.

## Dev Help

API: ec2-34-238-150-16.compute-1.amazonaws.com:4000

Ratel: ec2-34-238-150-16.compute-1.amazonaws.com:8000

The API can be started via a simple `python3 recipy.py`. Long running commands, such as running the API, should be done via `screen`. Current screens can be found via `-ls`:

```
ubuntu@ip-172-31-12-209:~/chefconnect$ screen -ls
There are screens on:
	10780.run	(04/17/20 21:29:35)	(Detached)
	2163.ratel	(04/17/20 13:31:29)	(Detached)
	1958.alpha	(04/17/20 13:17:45)	(Detached)
	1774.zero	(04/17/20 13:16:34)	(Detached)
4 Sockets in /run/screen/S-ubuntu.
```

Attach to a specific screen via `-d -R`, e.g.: `screen -d -R run`. Detach with `ctrl-a, ctrl-d`. The alpha, zero, and ratel screens are for dgraph. THe run screen is for the API. Commands for restarting each, replacing the IP address as necessary:

```
docker run -it -p 5080:5080 --network dgraph_default -p 6080:6080 -v ~/zero:/dgraph dgraph/dgraph:latest dgraph zero --my=172.31.12.209:5080
docker run -it -p 7080:7080 --network dgraph_default -p 8080:8080 -p 9080:9080 -v ~/recipy_data:/dgraph dgraph/dgraph:latest dgraph alpha --lru_mb=1024 --zero=172.31.12.209:5080 --my=172.31.12.209:7080
docker run -it -p 8000:8000 --network dgraph_default dgraph/dgraph:latest dgraph-ratel
```

Currently a single-host deployment. Full production env will migrate to HA. Data can be populated via `python3 manage.py`. Will create the schema and insert two small sample recipes.