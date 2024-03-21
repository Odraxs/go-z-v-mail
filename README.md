## GO Z V MAIL

This project consist in 4 folders:

- `data-embedding`: the folder that automated the process to create and embed the index `emails` to the zincsearch database.
- `web`: the folder that contains the web application.
- `server`: the folder that contains the go server that will handle the requests to the database `zincsearch` and retrieves the results(limited to 200).
- `docker`: the folder that contains the docker-compose file that can be used to run the entire project.

## Requirements:

- Go == 1.22
- Docker
- Docker-compose
- Node >= 20.10.0(recommended)
- Graphviz(if you want to generate the profiling graphs)

## I just want to see the project running!

- Give permissions:
    ```bash
       chmod +x envs.sh
       chmod a+rwx ./data-embedding
    ```
- Only follow the instructions of [data-embedding](#data-embedding)
- Then run the following commands:
    ```bash
        . envs.sh
        cd docker
        docker-compose up
    ```
> [!NOTE] 
> Remember, if you make changes in web or server you should remove the images so docker compose remounts them.

## Development Instructions:

To setup the project first run the following commands:

- Set up env variables(you can modify the file to set the values you want)
```bash
    chmod +x envs.sh
    . envs.sh
```

- Give write permission 
```bash
    chmod a+rwx ./data-embedding
```

- Start zincsearch docker image
```bash
    docker-compose up
```

### data-embedding

> [!NOTE]  
> All the cmd commands were made from a linux based os.

- Obtain the data downloaded from [enron_mail](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz)
- Unzip it e.g.
    ```bash
        tar -xvzf enron_mail_20110402.tgz
    ```
- Enter the folder and move it to the project directory e.g.
    ```bash
        cd enron_mail_20110402
        mv maildir /path/your_project_path/data-embedding
    ```
- Run the script
    ```bash
        cd data-embedding
        go run main.go
    ```
> [!WARNING]  
> This step consumes a lot of CPU recourses, I recommend to run it with everything else closed.

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

- Change `zincsearchEndpoint` to `http://localhost:4080/api` in the `zincsearch_repo.go` file.

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
    curl -X POST http://localhost:3001/emailSearch -H "Content-Type: application/json" --data '{"term": "manipulated", "max_results": 10, "field": "content", "sort_fields": []}'
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
    npm run lint
    npm run dev
```

Now just open a web browser at `http://localhost:5173/` and use the app.

## Project next steps 

- [x] Add additional options to the search emails request.
- [x] Add to web project at least one new feature.
- [x] Dockerize sever and web projects.
- [x] Add tests to server project.
- [ ] Search how to improve data-embedding to not use all the computer CPU, because if the computer has low spects it probably will crash.
- [ ]  ~~Add tests to web project.~~
