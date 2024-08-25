# hayai
An EEW system for Linux using JMA data provided by the Wolfx Project.

This software comes with zero guarantee. This software could fail at any time. I made this since there were no better free and open source alternatives for Linux.

# Installation
## Arch
Install the `ontake-hayai-git` package from the AUR.

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

# Configuration
You can configure the software by modifying the values in `~/.config/ontake/hayai/config.yml` and then restarting `hayai`.