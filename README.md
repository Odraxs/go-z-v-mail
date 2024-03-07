## GO Z V MAIL

This project will consist in 3 folders:

- `data-embedding`: the folder that automated the process to create the index `emails` to the zincsearch database.
- `web`: the folder that will contain the web application.
- `server`: the folder that will contain the go server that send request to the database an retrieves the results.

## Requirements:

- go 1.22(recommended)
- docker
- docker-compose
- node 20.10.0(recommended)

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

**Note!** all the cmd commas are for a linux od
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
