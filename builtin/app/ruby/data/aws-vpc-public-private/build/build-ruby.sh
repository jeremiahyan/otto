#!/bin/bash

set -o nounset -o errexit -o pipefail -o errtrace

error() {
   local sourcefile=$1
   local lineno=$2
   echo "ERROR at ${sourcefile}:${lineno}; Last logs:"
   grep otto /var/log/syslog | tail -n 20
}
trap 'error "${BASH_SOURCE}" "${LINENO}"' ERR

oe() { "$@" 2>&1 | logger -t otto > /dev/null; }
ol() { echo "[otto] $@"; }

# cloud-config can interfere with apt commands if it's still in progress
ol "Waiting for cloud-config to complete..."
until [[ -f /var/lib/cloud/instance/boot-finished ]]; do
  sleep 0.5
done

ol "Adding apt respositories and updating..."
oe sudo apt-get update
oe sudo apt-get install -y software-properties-common
oe sudo add-apt-repository -y ppa:chris-lea/node.js
oe sudo apt-add-repository -y ppa:brightbox/ruby-ng
oe sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 561F9B9CAC40B2F7
echo 'deb https://oss-binaries.phusionpassenger.com/apt/passenger trusty main' | sudo tee /etc/apt/sources.list.d/passenger.list > /dev/null
oe sudo apt-get update

# TODO: parameterize ruby version as input
export RUBY_VERSION="{{ ruby_version }}"

ol "Installing Ruby, Passenger, Nginx, and other packages..."
export DEBIAN_FRONTEND=noninteractive
oe sudo apt-get install -y bzr git mercurial build-essential \
  libpq-dev zlib1g-dev software-properties-common \
  apt-transport-https \
  nodejs \
  ruby$RUBY_VERSION ruby$RUBY_VERSION-dev \
  nginx-extras passenger

ol "Installing Bundler..."
oe sudo gem install bundler --no-ri --no-rdoc

ol "Extracting app..."
sudo mkdir -p /srv/otto-app
sudo tar zxf /tmp/otto-app.tgz -C /srv/otto-app

ol "Adding application user..."
oe sudo adduser --disabled-password --gecos "" otto-app

ol "Setting permissions..."
oe sudo chown -R otto-app: /srv/otto-app

ol "Configuring nginx..."

# This is required for passenger to get a reasonable environment where it can
# find executables like /usr/bin/env, /usr/bin/curl, etc. It also apparently
# needs to occur high in the config. Appending it is insufficient. :-|
sudo sed -i '1s/^/env PATH;\n/' /etc/nginx/nginx.conf
sudo sed -i '1s/^/# Otto: set PATH so passenger can see binaries\n/' /etc/nginx/nginx.conf

# Need to remove this so nginx reads our site
sudo rm /etc/nginx/sites-enabled/default

# These lines are present as comments in the passenger-packaged nginx.conf, but
# it's easier to drop a separate config than to sed out an uncomment.
cat <<NGINXCONF | sudo tee /etc/nginx/conf.d/passenger.conf > /dev/null
# Generated by Otto
passenger_root /usr/lib/ruby/vendor_ruby/phusion_passenger/locations.ini;
passenger_ruby /usr/bin/passenger_free_ruby;
NGINXCONF

cat <<NGINXCONF | sudo tee /etc/nginx/sites-enabled/otto-app.conf > /dev/null
# Generated by Otto
server {
    listen 80;
    root /srv/otto-app/public;
    passenger_enabled on;
}
NGINXCONF

ol "Bundle installing the app..."
sudo -u otto-app -i /bin/bash -lc "cd /srv/otto-app && bundle install --deployment --without development test"

ol "...done!"
