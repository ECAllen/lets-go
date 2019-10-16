#!/bin/bash

#sql="\"INSERT INTO keys ('date','chan','key','name','desc','ser','ep','cat') VALUES('xxx','xxx','xxxx','xxxx','xxxxx.','xxx','xxx','xxxx');\""      
#echo $sql
#sqlite3 mydb.db "$sql"

index_sql="CREATE INDEX idx_memories_created ON memories(created);"

create_sql="CREATE TABLE IF NOT EXISTS memories(id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, content TEXT NOT NULL, created TEST NOT NULL);"

data_sql="INSERT INTO memories (title, content, created) VALUES (    'An old silent pond',    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō', '2019-10-10');

INSERT INTO memories (title, content, created) VALUES (    'Over the wintry forest',    'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',  '2019-10-10' );

INSERT INTO memories (title, content, created) VALUES (    'First autumn morning',    'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo', '2019-10-10'   );"

sqlite3 memories.db "$create_sql"
sqlite3 memories.db "$index_sql"
sqlite3 memories.db "$data_sql"
