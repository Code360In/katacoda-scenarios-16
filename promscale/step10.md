
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
python3 manage.py runserver
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

from django.conf.urls import url

</pre>


Then add this under urlpatterns:

<pre class="file">

urlpatterns = [
    ...
    url('', include('django_prometheus.urls')),
]

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
