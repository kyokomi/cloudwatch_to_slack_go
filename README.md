# cloudwatch_to_slack_go
cloudwatch_to_slack_go is `apex` lambda function for golang(Go)

![image](https://cloud.githubusercontent.com/assets/1456047/20602592/5db5fac4-b2a1-11e6-810c-82433dab64ed.png)

## Lambda Environment

![2016-11-24 23 55 35](https://cloud.githubusercontent.com/assets/1456047/20602715/cbbccb2e-b2a1-11e6-853d-6da35ab57418.png)

- SLACK_WEBHOOK_URL: https://api.slack.com/incoming-webhooks
- TIMEZONE_NAME: https://github.com/tkuchiki/go-timezone/blob/master/timezone.go#L184 shortName

## Build

### Project directory

```
<Your Apex Project Root>
├── functions
│   └── cloudwatch_to_slack_go
└── vendor
    └── github.com
        ├── apex
        │   └── go-apex
        └── tkuchiki
            └── go-timezone
```

### Run deploy or build

```
apex deploy cloudwatch_to_slack_go
```

```
apex build cloudwatch_to_slack_go > cloudwatch_to_slack_go
```

## License

MIT License
