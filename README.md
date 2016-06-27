# goriot

[![Join the chat at https://gitter.im/sdwolfe32/goriot](https://badges.gitter.im/sdwolfe32/goriot.svg)](https://gitter.im/sdwolfe32/goriot?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

![alt text](/etc/goriotsmall.png "goriot client")

goriot is a simple set of API bindings for the Riot API written in Go.

### Installing

```
go get github.com/sdwolfe32/goriot
```

### Usage

All that's needed to form set up a goriot client is to call the `NewRiotClient()` func. For example:

```
client := NewRiotClient("API_KEY_HERE")
```
You can then easily perform requests on the Riot API using the above generated client struct. For example:
```
summoners, _ := client.GetSummonersByName("na", "itsViz", "kylokatarn", "faker")
```
The above example should return to you a map of SummonerDTO structs containing information on the provided summoners in the North American region.

## Full example

```
package main

import (
	"fmt"
	"log"

  "github.com/sdwolfe32/goriot"
)

var sumName = "itsViz"

func main() {
	c := NewRiotClient("API_KEY_HERE")

	summoners := c.GetSummonersByName("na", sumName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Summoner ID : ", summoners[sumName].ID)
	fmt.Println("Summoner Level : ", summoners[sumName].SummonerLevel)
}
```

## Disclaimer
goriot isn’t endorsed by Riot Games and doesn’t reflect the views or opinions of Riot Games or anyone officially involved in producing or managing League of Legends. League of Legends and Riot Games are trademarks or registered trademarks of Riot Games, Inc. League of Legends © Riot Games, Inc.

The MIT License (MIT)
=====================

Copyright © 2016 Steven Wolfe

Permission is hereby granted, free of charge, to any person
obtaining a copy of this software and associated documentation
files (the “Software”), to deal in the Software without
restriction, including without limitation the rights to use,
copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following
conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.
