source /etc/profile
export LANG="en_US.UTF-8"
export HISTTIMEFORMAT="%Y-%m-%d %H:%M:%S"
#export APPNAME="za-bill"
#export WEB_PORT=28080
#export PANDORA_PORT=$(($WEB_PORT+1))
#export HSF_SERVER_PORT=$(($WEB_PORT+2))
#export HSF_HTTP_PORT=$(($WEB_PORT+3))
#export DEBUG_PORT=$(($WEB_PORT+4))
export WEB_TERMINAL_PORT=$(($WEB_PORT+5))
export SSH_PORT=$(($WEB_PORT+6))
export SHUTDOWN_PORT=$(($WEB_PORT+7))

export HOST_IP_ADDRESS=$(hostname -i | awk '{print $NF}')

#hsf sar
#/root/default/webapps/taobao-hsf.sar
#logging
LOG_DIRECTORY="/alidata1/admin/$APPNAME/logs"
#rm -rf $LOG_DIRECTORY $HOME/default/logs $HOME/logs
mkdir -p $LOG_DIRECTORY && chmod 777 -R $LOG_DIRECTORY
ln -sv $LOG_DIRECTORY $HOME/default/logs
ln -sv $LOG_DIRECTORY $HOME/logs

##1.config sshd port
sed -i "s/_SSH_PORT_/$SSH_PORT/" /etc/ssh/sshd_config_template
cp -f /etc/ssh/sshd_config_template /etc/ssh/sshd_config
##2.change root's password
echo "root:Test1234" | chpasswd
##3.start ssh
ssh-keygen -q -t rsa -b 2048 -f /etc/ssh/ssh_host_rsa_key -N ''
ssh-keygen -q -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key -N ''
/usr/sbin/sshd &

#JAVA_OPTS
export JAVA_OPTS="-Dserver.port=$WEB_PORT -Dhttp.server.port=default -Dpandora.qos.port=$PANDORA_PORT -Dhsf.server.ip=$HOST_IP_ADDRESS -Dhsf.server.port=$HSF_SERVER_PORT -Dhsf.http.enable=true -Dhsf.http.port=$HSF_HTTP_PORT -Xdebug -Xnoagent -Djava.compiler=NONE -Xrunjdwp:transport=dt_socket,server=y,suspend=n,address=$DEBUG_PORT"
#export JAVA_OPTS="$JAVA_OPTS -Xms1g -Xmx1g -XX:NewSize=512m -XX:MaxNewSize=512m -XX:PermSize=256m -XX:MaxPermSize=256m"
#export JAVA_OPTS="$JAVA_OPTS -Xms2g -Xmx2g -XX:NewSize=1g -XX:MaxNewSize=1g"
NewSize=`expr ${JvmSize} / 2`
export JAVA_OPTS="$JAVA_OPTS -Xms${JvmSize}m -Xmx${JvmSize}m -XX:NewSize=${NewSize}m -XX:MaxNewSize=${NewSize}m -XX:PermSize=256m -XX:MaxPermSize=256m"
export JAVA_OPTS="$JAVA_OPTS -Darguments.flag=dev"

#
#export HOST_IP_ADDRESS=$(hostname -i | awk '{print $NF}')

#web ternimal
#cd $LOG_DIRECTORY
#nohup python /opt/server.py --unsecure --host=${HOST_IP_ADDRESS} --port=${WEB_TERMINAL_PORT} > webssh.log &

#run
cd $HOME/default/webapps;pwd
nohup java ${JAVA_OPTS} -jar ${APPNAME}.jar > nohup.out &

#web ternimal
cd $LOG_DIRECTORY
python /opt/server.py --unsecure --host=${HOST_IP_ADDRESS} --port=${WEB_TERMINAL_PORT}

#run
#HOST_IP_ADDRESS=$(hostname -i | awk '{print $NF}')
#cd $HOME/default/webapps;pwd
#java -Xms1g -Xmx1g -XX:NewSize=512m -XX:MaxNewSize=512m -XX:PermSize=256m -XX:MaxPermSize=256m -Dhsf.server.ip=$HOST_IP_ADDRESS -Dpandora.qos.port=$PANDORA_PORT -Dhsf.server.port=$HSF_SERVER_PORT -Dhsf.http.port=$HSF_HTTP_PORT -Xdebug -Xnoagent -Djava.compiler=NONE -Xrunjdwp:transport=dt_socket,server=y,suspend=n,address=$DEBUG_PORT -Dserver.port=$WEB_PORT -Dhttp.server.port=default -jar ${APPNAME}.jar
