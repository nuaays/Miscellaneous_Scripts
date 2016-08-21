---------------------------------------------------------
Ubuntu Init
sudo apt-get update
sudo apt-get install -y htop mutt mailutils curl wget gawk vim git subversion expect mercurial bison build-essential gawk gcc g++ libc6-dev make autoconf automake python-dev python-pip libssl-dev apache2-utils git-core

sudo apt-get install -y aptitude openssh-server flex gperf dos2unix
sudo apt-get install -y libz-dev libncurses5-dev ia32-libs lib32z-dev lib32ncurses5-dev  lib32ncurses5 ia32-libs lib32stdc++6



=====JAVA 7
brower http://www.oracle.com/technetwork/cn/java/javase/downloads/jdk7-downloads-1880260.html
http://download.oracle.com/otn-pub/java/jdk/7u79-b15/jdk-7u79-linux-x64.tar.gz
vim /etc/profile
export JAVA_HOME=/opt/jdk1.7.0_79
export JRE_HOME=${JAVA_HOME}/jre
export PATH=${JAVA_HOME}/bin:$PATH


====Golang
http://www.golangtc.com/static/go/1.7rc6/go1.7rc6.linux-amd64.tar.gz


sudo mkdir -p /localdisk && sudo chmod 777 -R /localdisk



====sublime3
sudo add-apt-repository ppa:webupd8team/sublime-text-3
sudo apt-get update
sudo apt-get install sublime-text-installer
import urllib.request,os; pf = 'Package Control.sublime-package'; ipp = sublime.installed_packages_path(); urllib.request.install_opener( urllib.request.build_opener( urllib.request.ProxyHandler()) ); open(os.path.join(ipp, pf), 'wb').write(urllib.request.urlopen( 'http://sublime.wbond.net/' + pf.replace(' ','%20')).read())

=====nodejs
apt-get install -y python-software-properties software-properties-common  
add-apt-repository ppa:chris-lea/node.js  
apt-get update  
apt-get install nodejs  
http://nodejs.org/dist/v0.12.9/node-v0.12.9.tar.gz
http://nodejs.org/dist/v0.12.9/node-v0.12.9-linux-x64.tar.gz
wget http://nodejs.org/dist/v0.12.9/node-v0.12.9.tar.gz && tar -zxvf node-v0.12.9.tar.gz && cd node-v0.12.9/ && ./configure && make && make install
npm install express mongojs nodemailer



sudo add-apt-repository ppa:gophers/go
sudo apt-get update
sudo apt-get install golang-stable



====docker
sudo apt-get update
sudo apt-get install apt-transport-https ca-certificates
sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D
echo "deb https://apt.dockerproject.org/repo ubuntu-trusty main" > /etc/apt/sources.list.d/docker.list
sudo apt-get update
sudo apt-get purge lxc-docker
apt-cache policy docker-engine
apt-get install -y docker-engine=1.12.0-0~precise
sudo groupadd docker
sudo usermod -aG docker $USER
echo 'DOCKER_OPTS="--dns 8.8.8.8"' > /etc/default/docker
sudo service docker restart
#upgrade
sudo apt-get upgrade docker-engine
#uninstall
sudo apt-get autoremove --purge docker-engine && rm -rf /var/lib/docker



====github
git config --global user.email "nuaays@gmail.com"
git config --global user.name "Yang Sen"
ssh-keygen -t rsa && cat ~/.ssh/id_rsa.pub 



=====docker-compose
pip install -U docker-compose








-------------------------------------------------------------------------------
CentOS Setup
yum install htop curl wget gawk subversion git expect bison python-dev

