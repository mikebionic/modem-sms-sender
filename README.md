# SMS sender API

API for making SMS sending from modem, by sending post request to route

> Request for Send SMS

| url                                   | method |
| ------------------------------------- | :----: |
| 127.0.0.1:8000/api/v1/send-modem-sms  |  POST  |

```json
{
	"phone_number": "+99361509038",
  "message_text": "Mike rocks mecreate"
	//"query_string": "select \"ResName\" from tbl_dk_resource",
}
```
