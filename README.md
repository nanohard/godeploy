## godeploy
A self-service webhook deployment server.

You will need 3 things:
- [config.yml](#configyml)
- [.env](#env)
- [systemd file](#libsystemdsystemgodeploy)

### config.yml
```yaml
projects:
  - name: website
    repo_url: git@github.com:nanohard/website
    build:
      type: commands
      git_dir: /var/www/nanohard.net
      commands:
        - command: git reset --hard
        - command: git fetch
        - command: git checkout origin/master
  - name: robonano
    repo_url: git@github.com:nanohard/robonano
    build:
      type: tags
      git_dir: /home/user/robonano
      pre_commands:
        - command: first command before fetch
        - command: second command before fetch
      post_commands:
        - command: go build -o robonano main.go pages.go
        - command: systemctl stop robonano.service
        - command: cp -r robonano static/ tpl/ /var/www/bot.nanohard.net/
        - command: systemctl start robonano.service

```
In this example the two webhooks would be:
- example.com/deploy/website
- example.com/deploy/robonano

<br>

- **name**: name of project; used as the URI for the webhook
- **repo_url**: currently unused
- **type**: `commands` or `tags`
  - **git_dir** is the path where the git dir is located; used as the working directory for all commands
  - `commands` will run the commands in order
  - `tags` will fetch and checkout the latest tag
    - **pre_commands** will run before it fetches latest tag
    - **post_commands** will run after the latest tag is checked out
    
### .env
```ini
SSL=false
HOST=
PORT=4545
CERTFILE=/etc/letsencrypt/live/example.com/fullchain.pem
KEYFILE=/etc/letsencrypt/live/example.com/privkey.pem
```

### /lib/systemd/system/godeploy
```ini
[Unit]
Description=Git project deployment service
After=network.target
[Service]
Type=simple
User=user
Group=www-data
ExecStart=/var/www/godeploy/godeploy
WorkingDirectory=/var/www/godeploy
Restart=on-failure
RestartSec=3
[Install]
WantedBy=multi-user.target
Alias=godeploy.service
```
If `user` and `group` are not set then it will run as `root`.
It's advised to use the systemd file in order to be able to run commands requiring `sudo`, but also set the `user` and `group` so the files modified are not saved as `root`.
