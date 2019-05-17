#GoLang Giveaway Bot

This project is maintained (for educational purposes) by 14 yo programmer. I made this bot just for fun, i wanted to learn some aspects of golang, so i decided to make this project. Have fun using and improving it.

To run project u must have:

Name    | Command
--------|--------
[discordgo](https://github.com/bwmarrin/discordgo) | go get github.com/bwmarrin/discordgo
[cron](github.com/robfig/cron) | go get github.com/robfig/cron
[yaml](gopkg.in/yaml.v2) | go get gopkg.in/yaml.v2

you must also config your bot.
1. Rename `config.yml.default` to `config.yml`
2. Set `prefix` to your prefix. Default value: `!`
3. Set `giveawayHours` to hour, when bot should resolve giveaway. Default value: `18`
   *NOTE!* USE VALUES IN 24 HOUR FORMAT
4. Set `giveawayMinutes` to minute, when bot should resolve giveaway. Default value: `0`
5. Set `giveawaySeconds` to second, when bot should resolve giveaway. Default value: `0`
6. Set `botToken` to your bot's token. See [this](https://discordapp.com/developers/applications/).
<p>Giveaway will be resolved defaultly at 18:00:00</p>
   
*Please notice that it's pre-pre-alpha version, feel free to commit and improve code ✨*

######P.S. I don't admit this code xD Its license-free


   