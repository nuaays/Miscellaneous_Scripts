1.sudo docker pull mysql:latest
2.sudo mkdir -p /mysql/data  && sudo chmod 777 -R /mysql/data/
3.sudo docker run -v /mysql/data:/var/lib/mysql --name mysqldb -e MYSQL_DATABASE='mysqldb' -e MYSQL_USER='mysql' -e MYSQL_PASSWORD='mysql' -e MYSQL_ALLOW_EMPTY_PASSWORD='yes'  -e MYSQL_ROOT_PASSWORD='' -d mysql:latest
4.sudo docker ps
5.sudo docker exec -it mysqldb bash
	mysql
		show databases;
		use mysqldb;
		CREATE TABLE Catalog(CatalogId INTEGER PRIMARY KEY,Journal VARCHAR(25),Publisher VARCHAR(25),Edition VARCHAR(25),Title VARCHAR(45),Author VARCHAR(25));
		INSERT INTO Catalog VALUES('1','Oracle Magazine','Oracle Publishing','November December 2013','Engineering as a Service','David A. Kelly');
		SELECT * FROM Catalog;
		show databases;
	exit

6.sudo docker run --name mysqldb2 -e MYSQL_ROOT_PASSWORD=mysql -d mysql


7.sudo docker run --name mysqldb2 -e MYSQL_ROOT_PASSWORD=mysql -d mysql
8.docker exec -it mysqldb2 bash
	mysql -uroot -pmysql
	show databases;
	show tables;
	exit
	
