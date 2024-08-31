# hayai
An EEW system for Linux using JMA data provided by the Wolfx Project.

This software comes with zero guarantee. This software could fail at any time. I made this since there were no better free and open source alternatives for Linux.

# Installation
## Linux
Install the `ontake-hayai-git` package from the AUR if on Arch. For other distros you can just build `hayai` from source with `go build` (only `go` is required).

You can now use your desktop environment to set the `hayai` command to run automatically at login.

For KDE that means creating the `.config/autostart/hayai.desktop` file containing:
```
[Desktop Entry]
Exec=/usr/bin/hayai
Name=Hayai
Comment=An EEW system for Linux using JMA data provided by the Wolfx Project.
Type=Application
X-KDE-AutostartScript=true
```

## Windows
Clone this repo.

`cd` into it.

Build this project with `go build` (`go` is required).

Set the `hayai.exe` to automatically start with your session.

# Configuration
You can configure the software by modifying the values in `~/.config/ontake/hayai/config.yml` (on Linux) and `C:\Users\<User>\AppData\Roaming\ontake\hayai\config.yml` (on Windows) and then restarting `hayai`.
