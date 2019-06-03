# Installation Guide
## Windows
1. Download and Install postgres database from https://www.enterprisedb.com/downloads/postgres-postgresql-downloads
2. Open sql shell (psql) from the start menu
3. on `Server [localhost]` press enter
4. on `Database [postgres]` press enter
5. on `Port [port]` press enter
6. on `Username [postgres]` press enter
7. on `Password: ` enter password
8. Run command `postgres#: CREATE USER user_name;`
    * note: the semicolon `;` is important
9. 

## Ubuntu
1. Install postgres using `sudo apt-get update` `sudo apt-get install postgres postgres-contrib`
2. Enter postgres terminal with `sudo -u postgres psql`
3. Enter `CREATE DATABASE classmanager;` Do not forget the semicolon
4. Create user with `CREATE USER myuser WITH PASSWORD 'mypassword';`
5. Grant privalages on database to user with `GRANT ALL PRIVILEGES ON DATABASE classmanager TO myuser;`
6. Quit postgres terminal with `\q`