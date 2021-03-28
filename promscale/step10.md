
# Django

https://www.djangoproject.com/

More examples to add metrics to application:

https://docs.timescale.com/latest/tutorials/tutorial-howto-monitor-django-prometheus

### Install Django
```
python3 -m pip install Django
```{{execute}}


### Create a template project

```
django-admin startproject mysite
```{{execute}}

### Verify that Django is working

```
cd mysite/
python3 manage.py runserver 0:8000
```{{execute}}



In your project settings.py file,set ALLOWED_HOSTS to this (line 28) :

```
ALLOWED_HOSTS = ['[[HOST_SUBDOMAIN]]-8000-[[KATACODA_HOST]].environments.katacoda.com']
```{{copy}}


update:
```
python3 manage.py migrate
```{{execute}}



### Export prometheus-style monitoring metrics from your Django application

```
python3 -m pip install django-prometheus
```{{execute}}


Modify settings.py and urls.py
In settings.py, add:

<pre class="file">

INSTALLED_APPS = [
   ...
   'django_prometheus',
   ...
]

MIDDLEWARE = [
    'django_prometheus.middleware.PrometheusBeforeMiddleware',
    # All your other middlewares go here, including the default
    # middlewares like SessionMiddleware, CommonMiddleware,
    # CsrfViewmiddleware, SecurityMiddleware, etc.
    'django_prometheus.middleware.PrometheusAfterMiddleware',
]

</pre>



In urls.py, make sure you have this in the header:

<pre class="file">

from django.conf.urls import include

</pre>


Then add this under urlpatterns:

<pre class="file">

urlpatterns = [
    ...
    path('', include('django_prometheus.urls')),
]

</pre>




update:
```
python3 manage.py migrate
```{{execute}}

results:

<pre class="file">

$ python3 manage.py migrate
Operations to perform:
  Apply all migrations: admin, auth, contenttypes, sessions
Running migrations:
  Applying contenttypes.0001_initial... OK
  Applying auth.0001_initial... OK
  Applying admin.0001_initial... OK
  Applying admin.0002_logentry_remove_auto_add... OK
  Applying admin.0003_logentry_add_action_flag_choices... OK
  Applying contenttypes.0002_remove_content_type_name... OK
  Applying auth.0002_alter_permission_name_max_length... OK
  Applying auth.0003_alter_user_email_max_length... OK
  Applying auth.0004_alter_user_username_opts... OK
  Applying auth.0005_alter_user_last_login_null... OK
  Applying auth.0006_require_contenttypes_0002... OK
  Applying auth.0007_alter_validators_add_error_messages... OK
  Applying auth.0008_alter_user_username_max_length... OK
  Applying auth.0009_alter_user_last_name_max_length... OK
  Applying auth.0010_alter_group_name_max_length... OK
  Applying auth.0011_update_proxy_permissions... OK
  Applying auth.0012_alter_user_first_name_max_length... OK
  Applying sessions.0001_initial... OK
$ 

</pre>

Verify that metrics are being exported
Restart the application and curl the /metrics endpoint:

```
python3 manage.py runserver 0:8000
```{{execute}}

```
curl localhost:8000/metrics
```{{execute}}


Update prometheus.yml. Under scrape_configs:, add:

<pre class="file">

- job_name: django
  scrape_interval: 10s
  static_configs:
  - targets:
    - localhost:8000

</pre>



Verify:

https://[[HOST_SUBDOMAIN]]-8000-[[KATACODA_HOST]].environments.katacoda.com/metrics


Admin:
```
python3 manage.py createsuperuser
```{{execute}}

https://[[HOST_SUBDOMAIN]]-8000-[[KATACODA_HOST]].environments.katacoda.com/admin



add to promethues.yml

```
  - job_name: 'django'
    metrics_path: /metrics
    scheme: https
    static_configs:
      - targets: ['[[HOST_SUBDOMAIN]]-8000-[[KATACODA_HOST]].environments.katacoda.com']
        labels:
          group: 'django'
```{{copy}}


restart prometheus:

```
docker restart prometheus
```{{execute}}


Verify new target is up:

https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com/targets



Grafana dashboard:

https://grafana.com/grafana/dashboards/7996



# Vegeta

```
go get -u github.com/tsenart/vegeta
```{{execute}}

```
echo "GET https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com/targets" | vegeta attack -rate=10 -duration=30s | vegeta report
```{{execute}}


results:

<pre class="file">

$ echo "GET https://2886795276-8000-simba06b.environments.katacoda.com/metrics" | vegeta attack -rate=10 -duration=30s | vegeta report
Requests      [total, rate, throughput]         300, 10.03, 10.02
Duration      [total, attack, wait]             29.935s, 29.9s, 34.928ms
Latencies     [min, mean, 50, 90, 95, 99, max]  30.376ms, 36.883ms, 36.776ms, 39.625ms, 40.623ms, 43.931ms, 112.719ms
Bytes In      [total, mean]                     11889724, 39632.41
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:300  
Error Set:
$ 


</pre>


# H2LOAD 10k

```
h2load https://2886795276-8000-simba06b.environments.katacoda.com/metrics -n 10000 -t 2 -c4 -T 10
```{{execute}}