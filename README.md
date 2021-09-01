# Validator-info
![CreatePlan](https://img.shields.io/badge/release-v4.0.0-red)
![CreatePlan](https://img.shields.io/badge/go-1.15%2B-blue)
![CreatePlan](https://img.shields.io/badge/license-Apache--2.0-green)  
Exporter for Validator info.

## Introduction
This Prometheus exporter is for monitoring information which is not provided from Tendermint’s basic Prometheus exporter(localhost:26660)

## List of supported chains
Rizon(groot), Terra(bombay)

## Install
```bash
cd $HOME
git clone https://github.com/Mr-K-Validator/Validator-info.git
cd $HOME/Validator-info

go build

./Validator-info version
## Validator-info v4.0.0
```

## Service(ex: rizon)
- **--chain** _string_: Chain name of the monitoring node(rizon | terra)
```bash
## Create a systemd service
sudo tee /etc/systemd/system/Validator-info.service > /dev/null <<EOF
[Unit]
Description=Exporter for Validator info
After=network-online.target

[Service]
User=${USER}
ExecStart=$HOME/Validator-info/Validator-info run \
  --chain "rizon" \
  --oper-addr "rizonvaloper1tqv36eh27pnkta5pmlzrzmcp4zc09qcz7x086z"
Restart=always
RestartSec=3
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=Validator-info

[Install]
WantedBy=multi-user.target
EOF

## Start service
sudo systemctl enable Validator-info
sudo systemctl start Validator-info

## log
journalctl -f -u Validator-info.service
```

## Screenshot
![grafana](https://github.com/Mr-K-Validator/Validator-info/blob/master/screenshot/example_grafana.png)
![prometheus](https://github.com/Mr-K-Validator/Validator-info/blob/master/screenshot/example_prometheus.png)
![log](https://github.com/Mr-K-Validator/Validator-info/blob/master/screenshot/example_log.png)
