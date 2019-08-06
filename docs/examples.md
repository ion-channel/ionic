# Ion Channel API Examples

## Basic example

Run the following command at your favorite command prompt.

```
curl https://api.ionchannel.io/v1/vulnerability/getProducts?external_id=cpe:/a:lodash:lodash:4.17.11
```

You should see a response containing some product information for `lodash` at version `4.17.11`.  The vulnerability service does not require Authentication for requests but other services do.  

## Authentication

You can get an API token from the user settings section of the Ion Channel console.  Once you have the token you can supply that in the `Authentication` header in your requests.

```
curl "https://api.ionchannel.io/v1/metadata/getLanguages?text=hello+ion+channel"
```

The curl request above will result in a 401 code with authentication failed.

```
{"message":"authentication failed (API)","fields":{},"code":401}
```

Making the same request with your token will respond with all of the languages detected in your request.

```
curl -H 'Authorization: Bearer yourtokenhere' "https://api.ionchannel.io/v1/metadata/getLanguages?text=hello+ion+channel"
```

with a response something like:

```
{
	"data": [{
		"name": "English",
		"confidence": 1.0
	}],
	"timestamps": {
		"created": "2019-08-06 22:02:43 +0000",
		"updated": "2019-08-06 22:02:43 +0000"
	},
	"links": {
		"self": "https://api.ionchannel.io/v1/metadata/getLanguages?text=hello+this+is+english"
	},
	"meta": {
		"copyright": "Copyright 2019 Selection Pressure LLC www.selectpress.net",
		"authors": ["Ion Channel"],
		"version": "v1"
	}
}
```


Ok!  You are all set to begin `POST`ing, `GET`ting and `DELETE`ing data from Ion Channel.

## Creating a ruleset

Now that you have your token and seen some simple requests for Ion Channel you should be ready to create a
