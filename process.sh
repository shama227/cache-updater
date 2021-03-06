#!/usr/sh

# /etc/systemd/system/mydaemon.service

# first arg - service name
service_name="neutrino-cache-daemon"

build_only () {
  file="$service_name.service"
  echo "
    [Unit]
    Description=Neutrino Cache Updater
    After=network.target

    [Service]
    Type=simple
    Restart=on-failure
    RestartSec=10
    startLimitIntervalSec=60

    WorkingDirectory=$PWD
    ExecStart=$PWD/cache-updater
    PermissionsStartOnly=true
    StandardOutput=syslog
    StandardError=syslog
    SyslogIdentifier=$service_name
    [Install]
    WantedBy=multi-user.target
  " > "$file"

  mv "$file" /etc/systemd/system/
}

build_n_start () {
  build_only
  start_service
}
start_service () {
  systemctl start "$service_name"
}

main () {
  while [ -n "$1" ]
  do
    case "$1" in
      --build-only ) build_only ;;
      --build-n-start ) build_n_start ;;
      --start ) start_service ;;
      --script ) daemon_fl=$2 ;;
      --service ) service_name=$2 ;;
    esac
    shift
  done
}

# shellcheck disable=SC2068
main $@