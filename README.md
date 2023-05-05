## Antispam API

### Go port of [Antispam API](https://github.com/thehamkercat/telegram-antispam-rs)

#### Notes
- The classifier works in aggressive mode, it can sometimes classify non-spam messages/emails as spam (when the input is too small)
- The dataset provided may contain some NSFW texts or personal info, it's not thoroughly checked.
- I've included a docker-based example, but you can run it without docker as well.

### With Docker compose

```sh
$ git clone https://github.com/axrav/AntispamAPI
$ cd AntispamAPI
$ docker-compose build
$ docker-compose up
```

## Endpoints:

```http
POST /predict HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
  "message": "Hello there how is it going"
}

HTTP/1.1 200 OK
content-length: 59
content-type: application/json

{
  "HamPercent": 97,
  "SpamPercent": 3,
  "isSpam": false,
  "message": "Hello there how is it going"
}
```
### For more reference about usage checkout [This](https://github.com/thehamkercat/telegram-antispam-rs)