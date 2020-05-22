# ChatShare

This is the ChatShare server that is used on the Madhouse Miners network.
Usage is open to anyone to use, but we don't provide support or warranty.

### Environment Variables

The following environment variables are required:

###### chatsharePSK
The Pre-shared Key used by the server and client to authenticate

###### discordToken
The discord token for the bot user to be used

###### discordChannel
The channel ID for the channel to monitor and post to

### Connections
The server needs to have an outgoing connection to the discordapp.com domain.
It starts a websocket server on port 8080
