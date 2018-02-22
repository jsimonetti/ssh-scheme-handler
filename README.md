
# Build
GO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ~/bin/ssh-scheme-handler main.go

# Install

create file: .local/share/applications/ssh-terminal.desktop
with contents:
```
[Desktop Entry]
Version=1.0
Name=SSH Terminal
Type=Application
Icon=utilities-terminal
Exec=/home/jsimonetti/bin/ssh-scheme-handler --url %u
StartupNotify=true
Terminal=false
MimeType=x-scheme-handler/ssh;
```

run: # gio mime x-scheme-handler/ssh ssh-terminal.desktop
