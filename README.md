## GO Z V MAIL

This project will consist in 3 folders:

- `data-embedding`: the folder that automated the process to create and embed the index `emails` to the zincsearch database.
- `web`: the folder that contains the web application.
- `server`: the folder that contains the go server that will handle the requests to the database `zincsearch` an retrieves the results(limited to 200).

## Requirements:

- Go >= 1.22(recommended)
- Docker
- Docker-compose
- Node >= 20.10.0(recommended)
- Graphviz(if you want to generate the profiling graphs)

## Instructions:

To setup the project firs run the following commands

- Give write permission 
```bash
    chmod a+rwx ./data-embedding
```

- Start docker images
```bash
    docker-compose up
```

### data-embedding

> [!NOTE]  
> All the cmd commas are for a linux based os.

- Obtain the data downloaded from [enron_mail](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz)
- Unzip it e.g.
    ```bash
        tar -xvzf enron_mail_20110402.tgz
    ```
- Enter the file and move it to the project director e.g.
    ```bash
        cd enron_mail_20110402
        mv maildir /path/your_project_path/data-embedding
    ```
- Run the script
    ```bash
        cd data-embedding
        go run main.go
    ```
- Gen profiling graphs(**Optional**)

    - CPU profiling
        ```bash
            cd profs
            go tool pprof cpu.prof
            (pprof) pdf
        ```
    - Memory profiling
        ```bash
            cd profs
            go tool pprof mem.prof
            (pprof) pdf
        ```
### server

Right now the server doesn't need any external configuration, just make sure that the 
zincsearch server is running in `localhost:4080` and that the user credentials are the same
as the ones set in `config/credentials.go`

- Start server
    ```bash
        cd ../server
        go run main.go
    ```

- Try it out: You can use any program(Postman, Insomnia, etc..), for simplicity I'm going to use curl.

#### Status request:

```bash
    curl -i -X GET http://localhost:3001/
```
**Result:**

```bash
    HTTP/1.1 200 OK
    Vary: Origin
    Date: Sat, 09 Mar 2024 02:03:11 GMT
    Content-Length: 0 
```

#### Search emails request:

```bash
    curl -X POST http://localhost:3001/emailSearch -H "Content-Type: application/json" --data '{"filter": "manipulated"}'
```
**Result:**

```bash
    {"time": 712,"emails": [{"id": "26rdgTPY702","from": "linda.robertson@enron.com",...]}
```

### web

To start the web server run:

```bash
    cd ../web
    npm i
    npm run lin
    npm run dev
```

Now just open a web browser at `http://localhost:5173/` and use the app.

## Project next steps 

- [ ] Search how to improve data-embedding to not use all the computer CPU, because if the computer has low spects it probably will crash.
- [ ] Add more filter options to the search emails request.
- [ ] Add to web project at least one new feature.
- [ ] Add tests to server project.
- [ ]  ~~Add tests to web project.~~
- [ ] Dockerize sever and web projects.
