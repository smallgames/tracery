tracery
=======

small game 

报文格式
length|body
通过读length来确定body的长度且只读length的长度，长的多的丢弃，短的不够也没事。
读过的包省下的全丢弃，等待下一个包。


