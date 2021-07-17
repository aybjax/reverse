# reverse proxy / load balancing in apache2
```
sudo a2enmod proxy
sudo a2enmod proxy_http
sudo a2enmod proxy_balancer
sudo a2enmod lbmethod_byrequests
```
`
sudo systemctl restart apache2
`

# apache conf
```
<VirtualHost *:80>
  ProxyPreserveHost On
  ProxyRequests On
  ServerName aybjax.com
  ServerAlias www.aybjax.com
  ProxyPass / http://127.0.0.1:8080/
  ProxyPassReverse / http://127.0.0.1:8080/

  ErrorLog /var/log/apache2/aybjax.log

</VirtualHost>
```
