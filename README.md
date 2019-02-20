# PSNapi
PlaystationApi wrapper written in go
handlers.Login(refresh_token) - will return oauthcode
handlers.UserGames(oauth, username) - to get user played games
handlers.GetGameTrophieTitles(oauth, game.titleid) - will return trophyTitle and NpCommunicationId
handlers.GetGameTrophies(oauth, NpCommunicationId, username) - will return all trophies for game if username is not "" it will mark earned trophies
handlers.GetGameTrophieGroups(oauth, NpCommunicationId) - will return trophie groups for game
handlers.GetGameTrophieData(oauth, NpCommunicationId, groupId) - will return trophie data for given group of game, usefull for earned rate
TODO
MSGS
?