![](https://github.com/nickmancari/gocean/blob/master/img/gocean_logo.png)
<h1 align='center'>Digital Ocean  CLI Tool</h1>
<h2 align='center'>Manage Droplets From the Command Line With Digital Ocean APIs</h2>
<br>

# Using Gocean:

Gocean takes simple flags to manage Droplets from the command line:
```
$ gocean --create testdroplet
```
Control droplets by name from the commandline, instead of using Droplet IDs:
```
$ gocean --reboot testdroplet
```
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
- [ ] Natively SSH into droplet from Gocean ($ gocean --ssh dropletname)
- [ ] Better token storage and handling
- [ ] Manage droplets by name, converting the droplet ID
