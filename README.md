![](https://github.com/nickmancari/gocean/blob/master/img/gocean_logo.png)
<h1 align='center'>Digital Ocean  CLI Tool</h1>
<h2 align='center'>Manage Droplets From the Command Line With Digital Ocean APIs</h2>
<br>

![](https://img.shields.io/badge/OS-Linux-informational?style=flat&logo=Linux&logoColor=white&color=2bbc8a)
![](https://img.shields.io/badge/Code-Go-informational?style=flat&logo=go&logoColor=white&color=00add8)
![](https://img.shields.io/badge/Cloud-DigitalOcean-informational?style=flat&logo=digitalocean&logoColor=white&color=0080ff)

# Using Gocean:

Control droplets by name from the command line, instead of using Droplet IDs:
```
$ gocean --action testdroplet reboot
```
See an overview of your droplets in one clean command line view:
```
$ gocean --droplet ls
```
Even delete a droplet:
```
$ gocean --destroy testdroplet
```
Create a Droplet from the command line with interface:
```
$ gocean --create -i
```
Start an SSH Session into a Droplet with a setup interface:
```
$ gocean --connect -i
```
<br>

| Flag | Description | Common Commands |
| --- | --- | --- |
|`--token <token>` | Add Digital Ocean API Token to Gocean | `--token rm` remove token |
|`--droplet <droplet name>` | Show Brief Overview of a Specific Droplet | `--droplet ls` list all droplets |
|`--create -i` | Create Droplet with Simple Interface |
|`--connect -i` | SSH Session to Droplet with Interface |
|`--destroy <droplet name>` | Delete Specified Droplet |
|`--action <droplet name> <action>` | Control Droplet with Actions | `power_off` `power_on` `reboot` `shutdown`|

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
- [X] Native SSH into droplet from Gocean ($ gocean --ssh dropletname)
- [ ] Better token storage and handling
- [X] Format JSON responses
- [ ] Bash script installer
