# Simple API for government data

## How to run

```bash
go run .
```

You should see a message 

```bash
Starting the webserver on :3000
````

Now the API is turning locally on :3000 port ;)

## Available endpoints

* `http://localhost:3000/admin/upload` - uploads data from https://www.data.gouv.fr/fr/datasets/r/406c6a23-e283-4300-9484-54e78c8ae675 . (Call this method before executing any other requests to be able to operate with data)

* `http://localhost:3000/api/departments/list` - list of departments

* `http://localhost:3000/api/departments/stat/975/2020-05-14` - data available by department
