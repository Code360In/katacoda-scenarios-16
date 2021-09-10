
# curl



```
echo t4
```{{execute T4}}


Get subscription
```
curl -k -X GET http://127.0.0.1:8000/redfish/v1/EventService/Subscriptions/ 
```{{execute  T4}}


Create subscription
```
curl -k -X POST http://127.0.0.1:8000/redfish/v1/EventService/Subscriptions/  -H 'content-type: application/json'  -d '{"Destination": "http://127.0.0.1:8443/", "Protocol" : "Redfish", "EventTypes": ["Alert","StatusChange","ResourceUpdated","ResourceAdded","ResourceRemoved"], "Context": "Public"}'
```{{execute  T4}}



test alert

```
curl -vvv -k -X POST http://127.0.0.1:8000/redfish/v1/EventService/Actions/EventService.SubmitTestEvent  -H 'content-type: application/json'  -d '{
            "EventId": "Heartbeat",
            "EventTimestamp": "2021-09-09T23:08:42Z",
            "EventType": "Alert",
            "Message": "Heartbeat OK",
            "MessageArgs": [
                "NoAMS",
                "Busy",
                "NoCached"
            ],
            "MessageId": "iLOEvents.2.1.Heartbeat",
            "OriginOfCondition": "/redfish/v1/Systems/4/",
            "Severity": "OK"
}'
```{{execute  T4}}


```
curl -k -X GET http://127.0.0.1:8000/redfish/v1/EventService/Subscriptions/ 
```{{execute  T4}}


```
curl -k -X DELETE http://127.0.0.1:8000/redfish/v1/EventService/Subscriptions/4 
```{{execute  T4}}


