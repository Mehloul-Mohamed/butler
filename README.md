<h1 align="center"> Ally </h1>

<p align="center">Helper utility to make playing CTFs a little less messy.</p>

## Usage
```
usage: ally {start,list,attempt,info}
positional arguments:
  {start,list,attempt}
	start			Start a CTF
	list			Show challenge list
	attempt			Attempt a challenge
	info			Show scoreboard & team stats
	fetch			Fetch all challenges
```

- `start`  
This will create a directory inside `~/ctf/` with the name of the CTF.
Inside the directory there is one file:  
`credentials.txt` :  Containing the CTF url and api token
```
usage: ally start name url token

positional arguments:
  name                  CTF name
  url                   CTF url
  token                 Your API token
```
### CTF Commands (must be used within the directory of a CTF):
- `list`  
lists the available challenges
```
usage: ally list
```
- `attempt`  
This will create a directory for the challenge inside the directory of the CTF and will download the files for the challenge into it.
```
usage: ally attempt id

positional arguments:
  id                    Challenge id
```
- `fetch`  
Fetch all challenges
```
usage: ally fetch
```
- `info`:  (Credit to [@shadow1004](https://github.com/shadow1004) for the idea)  
Shows the top three teams on the scoreboard + general stats about the user's team
```
usage: ally info
```
