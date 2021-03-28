
# Django

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
python3 manage.py runserver
```{{execute}}



In your project settings.py file,set ALLOWED_HOSTS to this (line 28) :

```
ALLOWED_HOSTS = ['https://[[HOST_SUBDOMAIN]]-8000-[[KATACODA_HOST]].environments.katacoda.com']
```{{copy}}


update:
```
python3 manage.py migrate
```{{execute}}


Verify:

https://[[HOST_SUBDOMAIN]]-8000-[[KATACODA_HOST]].environments.katacoda.com


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
python3 manage.py runserver
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


