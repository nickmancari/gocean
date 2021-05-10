![](https://github.com/nickmancari/gocean/blob/master/img/gocean_logo.png)
<h1 align='center'>Digital Ocean  CLI Tool</h1>
<h2 align='center'>Manage Droplets From the Command Line With Digital Ocean APIs</h2>
<br>

![](https://img.shields.io/badge/Code-Go-informational?style=flat&logo=go&logoColor=white&color=00add8)
![](https://img.shields.io/badge/Cloud-DigitalOcean-informational?style=flat&logo=digitalocean&logoColor=white&color=0080ff)

# Using Gocean:

Gocean takes simple flags to manage Droplets from the command line:
```
$ gocean --create testdroplet
```
Control droplets by name from the command line, instead of using Droplet IDs:
```
$ gocean --reboot testdroplet
```
See an overview of your droplets in one clean command line view:
```
$ gocean --droplet ls
```
Even delete a droplet:
```
$ gocean --destroy testdroplet
```
<br>

| Flag | Description | Common Commands
| --- | --- | --- |
|`--token <token>` | Add Digital Ocean API token to gocean | `--token rm` remove token |
|`--droplet <droplet name>` | Show brief overview of a specific droplet | `--droplet ls` list all droplets |
|`--create <droplet name>` | Create droplet with name specified |
|`--destroy <droplet name>` | Delete specified droplet |
|`--reboot <droplet name>` | Reboot specific droplet |
|`--shutdown <droplet name>` | Power off specific droplet |
|`--boot <droplet name>` | Turn on a specific droplet |

<br>

## Prerequisites:

- Digital Ocean Account
- Digital Ocean API Token

<br><br>
# Setup:

Install Gocean:
```
$ git clone https://github.com/nickmancari/gocean.git
$ cd gocean && ./install.sh
```
Install API token:
```
$ gocean --token asdjbvuwefjw143r8f9s8vne9r283hr98654wrg46954b9w8rb41185b9nw84g
```
<br><br>

## Future Functionality:
- [ ] Native SSH into droplet from Gocean ($ gocean --ssh dropletname)
- [ ] Better token storage and handling
- [X] Manage droplets by name, converting the droplet ID
- [X] Format JSON responses
- [ ] Bash script installer
- [ ] More robust control options
- [X] Add color & better organise output
- [ ] Refactor actions into one function
