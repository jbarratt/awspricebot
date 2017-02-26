# awsprice bot

A slackbot that helps with AWS Pricing.

Powered by [slick](https://github.com/abourget/slick) and [awsprice](https://github.com/jbarratt/awsprice)

To install:

	go get -u github.com/jbarratt/awsprice

To configure:

Edit `.awspricebot.conf` and add your values:


	{
		"Slack": {
			"api_token": "xoxo-yyyyyyyyyyy-mmmmmmmvvvvvvvvvvvvvvvvvvvv",
			"nickname": "awsprice",
			"general_channel": "#general",
			"team_domain": "yourteam.slack.com",
			"join_channels": [
				"#your-favorite-aws-user-channel"
			],
			"web_base_url": "http://github.com/serialized/awsprice",
			"db_path": "./slick.bolt.db"
		}
	}
