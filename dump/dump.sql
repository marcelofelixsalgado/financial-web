create table users(id int not null auto_increment, name varchar(255), password varchar(255), primary key(id));

create user 'financial_client'@'localhost' identified by 'financial_client';

grant all privileges on financial_balance_db.* to 'financial_client'@'localhost';