# tasksbot
  
Simple Telegram bot to organize small tasks between people. Created on golang, using [telebot](https://github.com/tucnak/telebot) library.

To store temporary data (not-sended tasks), it needs key-value stoage. With [gokv](https://github.com/philippgille/gokv) you can use any database, but here I implemented redis and simple go-map, if you can't use redis. To use redis, just fill redis enviromnemt variables.

## Environments

* BOT_TOKEN --- token of your bot (get it from BotFather);
* BOT_CHANNEL_ID --- id of the channek, which you want to use for your tasks;
* REDIS_ADDRESS (optional) --- address of the Redis server, including the port;
* REDIS_PASSWORD (optional) --- password for the Redis server;
* REDIS_DB (optional) --- redis database (number);

## To do

* Localization files for messages.

## License

Distributed under MIT, feel free to use.
