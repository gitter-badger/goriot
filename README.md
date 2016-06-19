# GoRIOT

GoRIOT is a RiotAPI client for League of Legends that can be used in go. It's extremely simple and very easy to use.

### Usage

All that's needed to form set up a GoRIOT client is to call the `NewRiotClient()` func. For example:

```
client := NewRiotClient("API_KEY_HERE")
```
You can then easily perform requests on the RiotAPI using the above generated client struct. For example:
```
summoners, _ := c.GetSummonerByName("na", "itsViz", "kylokatarn", "faker")
```
The above example should return to you a map of SummonerDTO structs containing information on the provided summoners in the North American region.
