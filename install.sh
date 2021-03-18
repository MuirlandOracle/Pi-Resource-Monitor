#!/bin/bash

if [[ $EUID -ne 0 ]]; then
	printf "\033[01;31m[-] Script is not running as root\033[0m\n"
	exit
fi

cp monitorWebapp /usr/bin/monitorWebapp

cat << EOF > /etc/systemd/system/monitorWebapp.service
[Unit]
Description=JSON API for Pi Stats
[Service]
User=root
Group=root
WorkingDirectory=/usr/bin
ExecStart=/usr/bin/monitorWebapp
Restart=always
RestartSec=5
[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload >/dev/null
systemctl enable monitorWebapp.service >/dev/null 2>&1
systemctl start monitorWebapp.service >/dev/null

printf "\033[01;32m[+] Service Installed\033[0m\n"
