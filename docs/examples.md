# Ion Channel API Examples

## Basic example

## Endpoints used

`GET /v1/vulnerability/getProducts`

`GET /v1/metadata/getLanguages`

`GET /v1/teams/getTeams`

`GET /v1/ruleset/getRules`

`POST /v1/ruleset/createRuleset`

`POST /v1/project/createProject`

`POST /v1/scanner/analyzeProject`

`GET /v1/scanner/getAnalysisStatus`

## Really basic example

Run the following command at your favorite command prompt.

```
curl https://api.ionchannel.io/v1/vulnerability/getProducts?external_id=cpe:/a:lodash:lodash:4.17.11
```

You should see a response containing some product information for `lodash` at version `4.17.11`.  The vulnerability service does not require Authentication for requests but other services do.  

## Authentication

You can get an API token from the user settings section of the Ion Channel console.  Once you have the token you can supply that in the `Authentication` header in your requests.

```
curl "https://api.ionchannel.io/v1/metadata/getLanguages?text=hello+world"
```

The curl request above will result in a 401 code with authentication failed.

```
{"message":"authentication failed (API)","fields":{},"code":401}
```

Making the same request with your token will respond with all of the languages detected in your request.

```
curl -H 'Authorization: Bearer yourtokenhere' "https://api.ionchannel.io/v1/metadata/getLanguages?text=hello+world"
```

with a response similar to:

```
{
  "data": [
    {
      "name": "English",
      "confidence": 0.67913777
    }
  ],
  "meta": {
    "copyright": "Copyright 2018 Selection Pressure LLC www.selectpress.net",
    "authors": [
      "Ion Channel Dev Team"
    ],
    "version": "v1",
    "total_count": 0,
    "offset": 0
  }
}
```


Ok!  You are all set to begin `POST`ing, `GET`ting and `DELETE`ing data from Ion Channel.

## Creating a ruleset

Now that you have your token and seen some simple requests for Ion Channel you should be ready to get underway with a git repository analysis workflow.

Assuming you have an API token from above you should also have a team.  You can view your teams with the `GET /v1/teams/getTeams` call.

You response should look similar to:

```
{
  "data": [
    {
      "id": "5fc9af6e-ddc2-42e1-8646-3c3bf5efcc5b",
      "created_at": "2019-08-06T22:20:49.54772Z",
      "updated_at": "2019-08-06T22:20:49.54772Z",
      "name": "yourteam",
      "delivering": false,
      "sys_admin": false,
      "poc_name": "you",
      "poc_email": "email@gmail.com",
      "role": "admin",
      "status": "active"
    }
  ],
  "meta": {
    "copyright": "Copyright 2018 Selection Pressure LLC www.selectpress.net",
    "authors": [
      "Ion Channel Dev Team"
    ],
    "version": "v1",
    "total_count": 1,
    "offset": 0
  }
}
```

The team's `id` above is needed for creating a ruleset as it will be attached and only usable on the given team.

You can view all of the rules by making a request to `GET /v1/ruleset/getRules`.  

```
{
  "data": [
    {
      "id": "2981e1b0-0c8f-0137-8fe7-186590d3c755",
      "scan_type": "community",
      "name": "Has more than one committer",
      "description": "The project must have more than 1 committer.",
      "category": "Community",
      "created_at": "2019-02-07T19:29:14.691Z",
      "updated_at": "2019-07-09T14:43:46.712Z"
    },
    {
      "id": "d928de6b-9aa0-2b98-4663-17c23d68efc3",
      "scan_type": "external_coverage",
      "name": "Code Coverage > 70%",
      "description": "The project must have code coverage greater than the threshold set here. (Code coverage is performed by 3rd party tools and is configured separately)",
      "category": "Code Coverage",
      "created_at": "2016-08-23T20:50:15.652Z",
      "updated_at": "2019-07-09T14:43:46.684Z"
    },
    {
      "id": "786adcff-70d7-3f4e-cee2-385068ae0ed1",
      "scan_type": "external_coverage",
      "name": "Code Coverage > 80%",
      "description": "The project must have code coverage greater than the threshold set here. (Code coverage is performed by 3rd party tools and is configured separately)",
      "category": "Code Coverage",
      "created_at": "2016-08-23T20:50:15.658Z",
      "updated_at": "2019-07-09T14:43:46.689Z"
    },
    ...
```

Putting together the rule ids and team id you can create a Ruleset using `POST /v1/ruleset/createRuleset` using a similar payload to the following:

```
{
  "name": "Test Ruleset",
  "description": "Fail on vulnerability",
  "team_id": "5fc9af6e-ddc2-42e1-8646-3c3bf5efcc5b",
  "rule_ids": [
    "00be1862-959c-45d8-8fb5-2b748fe854d6"
  ]
}
```

Grab the ruleset id from the response data object.  We will use this to create our project.


## Creating a project

The `POST /v1/project/createProject` endpoint is where you send requests for project creation.  In this case we are creating a project called `test project` with the ruleset we defined above.  It will analyze the `git` repository located at `git@github.com/ion-channel/ionic.git` defaulting to the branch called `master`.  In this case we are going to analyze the Ion Channel Go SDK.

```
{
  "team_id": "5fc9af6e-ddc2-42e1-8646-3c3bf5efcc5b",
  "ruleset_id": "99b9a7e4-65dd-4e55-9382-00c97d2d819b",
  "name": "test project",
  "description": "this is a new project",
  "type": "git",
  "source": "git@github.com:ion-channel/ionic.git",
  "branch": "master",
  "active": true
}
```

## Analyzing a project

You can request analyzing of a project using `POST /v1/scanner/analyzeProject` and providing the `project_id` and `team_id` as json in the body.  You should receive a response with the status `queued`.

```
{
  "data": {
    "id": "155969be-cea9-4c48-a3a4-044a3d26bb4e",
    "team_id": "5fc9af6e-ddc2-42e1-8646-3c3bf5efcc5b",
    "project_id": "428baa9f-4c5c-4366-984a-b8f6b3c6e89e",
    "build_number": null,
    "status": "accepted",
    "message": "Request for analysis 155969be-cea9-4c48-a3a4-044a3d26bb4e on test project has been accepted.",
    "created_at": "2019-08-07T17:57:06.963Z",
    "updated_at": "2019-08-07T17:57:06.963Z",
    "branch": "master"
  },
  "meta": {
    "copyright": "Copyright 2018 Selection Pressure LLC www.selectpress.net",
    "authors": [
      "Ion Channel Dev Team"
    ],
    "version": "v1"
  },
  "links": {
    "created": "https://api.ionchannel.io/v1/scanner/analyzeProject?id=155969be-cea9-4c48-a3a4-044a3d26bb4e"
  },
  "timestamps": {
    "created": "2019-08-07T17:57:06.963Z",
    "updated": "2019-08-07T17:57:06.963Z"
  }
}
```

You can request status from `GET /v1/scanner/getAnalysisStatus` while supplying the `team_id`, `project_id` and the `id` from the analysis status above.  At this point you have a project with governance under risk analysis.
