# Script used to start lite5gc-dev containter
#!/usr/bin/env bash
#set -x
FILE="/etc/docker/daemon.json"
function show_restart_tips()
{
    echo "using the local docker repo must restart docker? (y or n)"
    read can_restart
    if [ $can_restart == 'n' ]; then
        exit 0;
    fi
}

if [ -d /etc/docker ]; then
    sudo chmod 777 -R /etc/docker
fi

if [ -f $FILE ]; then
    cat $FILE |grep insecure-registries &> /dev/null
    if [ $? == 0 ]; then
        cat $FILE |grep insecure-registries|grep "10.18.1.2:5000" &> /dev/null
        if [ $? != 0 ]; then
            show_restart_tips
            line=$(sed -n -e '/insecure-registries/=' $FILE)
            new=$(cat $FILE|grep insecure-registries|awk -F "[" '{print $1 "[\"10.18.1.2:5000\"," $2}')
            sed -i ""$line"c $new" $FILE
        else
            exit 0
        fi
    else
        show_restart_tips
        sudo sed -i '0,/{/{s/{/&\n    "insecure-registries": ["10.18.1.2:5000"],/}' $FILE
    fi
else
    show_restart_tips
    sudo echo '{    "insecure-registries": ["10.18.1.2:5000"]   }' >> $FILE
fi
sudo service docker restart
#set +x