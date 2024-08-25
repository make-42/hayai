# hayai
An EEW system for Linux using JMA data provided by the Wolfx Project.

This software comes with zero guarantee. This software could fail at any time. I made this since there were no better free and open source alternatives for Linux.

# Installation
## Arch
Install the `ontake-hayai-git` package from the AUR.

Enable the service by running: `sudo systemctl enable --now hayai@[username]`

# Configuration
You can configure the software by modifying the values in `~/.config/ontake/hayai/config.yml` and then restarting the service.