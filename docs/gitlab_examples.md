# Ion Channel API Gitlab Integration Examples

## Endpoints used

`POST /v1/gitlab/createProjectByUrl`

`POST /v1/gitlab/analyzeProjectByUrl`

## Really basic example

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

with a response similar to:

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

## Gitlab

Ion Channel supports integrations from either Github or Gitlab.  In this example we will see a
simple Gitlab integration flow.  You can create a project in Ion Channel using the `/v1/gitlab/createProjectByUrl` this endpoint takes a body with a single field of `url` similar to:

```
{
  "url":"git@gitlab.tld:org/repo.git"
}
```

After posting the data you will receive `201` on success of creation or a `400` if there
is an issue with the data.

Once you have created the project you can simply call `/v1/gitlab/analyzeProjectByUrl`.  
This endpoint takes a similar data set to the previous but adds a git `reference` as a
required field.

```
{
  "url":"git@gitlab.tld:org/repo.git",
  "reference":"branch/tag/commit"
}
```


This will ensure that the analysis is performed on the required code set.  The endpoint returns
`201` on analysis creation and `404` on project not found or `400` on an other data
issues.
