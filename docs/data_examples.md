# Ion Channel API Data Examples

## Basic example

## Endpoints used

`GET /v1/vulnerability/getProducts`

`GET /v1/metadata/getLanguages`

`GET /v1/dependency/getLatestVersionForDependency`

`GET /v1/dependency/getLatestVersionsForDependency`

`GET /v1/dependency/getResolvedDependencies`

`GET /v1/search`

`GET /v1/vulnerability/getVulnerabilities`

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


## Dependencies

You can use the Ion Channel API to interact with various dependency related endpoints.  A simple workflow would be to get a specific version for a dependency and then the dependency tree for that version.

The first request would be to get the latest version.  In this example we are using the Express NPM package.

`GET /v1/dependency/getLatestVersionForDependency` this takes two query params `name` and `type`.  In this case `type` represents the language or ecosystem (java, npm, ruby, pypi).  The name is the actual name of the package.

A similar endpoint that will provide all versions for a dependency is `GET /v1/dependency/getLatestVersionsForDependency`

As of this writing the latest version of Express will respond with

```
{
  "meta": {
    "copyright": "Copyright 2019 - Selection Pressure",
    "authors": [
      "tlpinney",
      "Matthew Mayer"
    ],
    "version": "v1",
    "total_count": 1
  },
  "links": {
    "self": "https://api.ionchannel.io/v1/dependency/getLatestVersionForDependency?name=express&type=npm"
  },
  "timestamps": {
    "created": "2019-08-07T22:50:05.046Z",
    "updated": "2019-08-07T22:50:05.103Z"
  },
  "data": {
    "version": "5.0.0-alpha.7"
  }
}
```

The latest version is 5.0.0-alpha.7.  If we want to get the dependency tree for this version we use the `GET /v1/dependency/getResolvedDependencies`  This endpoint requires 3 params `name`, `version`, `type` and if the type is java a `group` param is also required.

The Express version determined above with result in the following response data.

```
{
  "meta": {
    "copyright": "Copyright 2017 - Ion Channel Corp (ionchannel.io)",
    "authors": [
      "tlpinney",
      "Matthew Mayer"
    ],
    "version": "v1",
    "total_count": 9
  },
  "links": {
    "self": "https://api.ionchannel.io/v1/dependency/getResolvedDependencies?name=express&version=5.0.0-alpha.7&type=npm"
  },
  "timestamps": {
    "created": "2019-08-07T22:55:09.687Z",
    "updated": "2019-08-07T22:55:10.150Z"
  },
  "data": {
    "type": "npmjs",
    "scope": "runtime",
    "package": "package",
    "name": "express",
    "org": "expressjs",
    "version": "5.0.0-alpha.7",
    "latest_version": "5.0.0-alpha.7",
    "dependencies": [
      {
        "type": "npmjs",
        "scope": "runtime",
        "package": "package",
        "name": "accepts",
        "org": "jshttp",
        "version": "1.3.7",
        "latest_version": "1.3.7",
        "dependencies": [
          {
            "type": "npmjs",
            "scope": "runtime",
            "package": "package",
            "name": "mime-types",
            "org": "jshttp",
            "version": "2.1.24",
            "latest_version": "2.1.24",
            "dependencies": [
              {
                "type": "npmjs",
                "scope": "runtime",
                "package": "package",
                "name": "mime-db",
                "org": "jshttp",
                "version": "1.40.0",
                "latest_version": "1.40.0",
                "dependencies": [],
                "requirement": "1.40.0"
              }
            ],
            "requirement": "~2.1.24"
          },
          {
            "type": "npmjs",
            "scope": "runtime",
            "package": "package",
            "name": "negotiator",
            "org": "jshttp",
            "version": "0.6.2",
            "latest_version": "0.6.2",
            "dependencies": [],
            "requirement": "0.6.2"
          }
        ],
        "requirement": "~1.3.5"
      },
      {
        "type": "npmjs",
        "scope": "runtime",
        "package": "package",
        "name": "array-flatten",
        "org": "blakeembrey",
        "version": "2.1.1",
        "latest_version": "2.1.2",
        "dependencies": [],
        "requirement": "2.1.1"
      },
      ...
    ],
    "requirement": "5.0.0-alpha.7"
  }
}
```


# Searching for products

If you have a dependency name and version you can use Ion Channel to search for product information.  This can be helpful in determining if a dependency has a vulnerability.

The `GET /v1/search` endpoint takes one param `q` that allows for querying by various attributes.  For example if we would like to know product informations for Express at version 4.4.2 we can supply `expressAND4.4.2`.  This results in a set of products where both express and 4.4.2 are found.


```
{
  "data": [
    {
      "id": 0,
      "name": "express",
      "org": "expressjs",
      "version": "4.4.2",
      "up": "",
      "edition": "",
      "aliases": null,
      "created_at": "2017-08-20T05:53:59.711Z",
      "updated_at": "2018-05-25T07:07:45.181Z",
      "title": "Expressjs Express 4.4.2",
      "references": [
        {
          "Change Log": "https://github.com/expressjs/express/blob/master/History.md#4140--2016-06-16"
        }
      ],
      "part": "/a",
      "language": "",
      "external_id": "cpe:/a:expressjs:express:4.4.2",
      "source": null,
      "confidence": 0.991032
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

# Is there a vulnerability?

Now that you have product information you can query for related vulnerabilities.  `GET /v1/vulnerability/getVulnerabilities` with `product` and `version` query params will respond with any linked vulnerabilities.

```
{
  "data": [
    {
      "id": 267877816,
      "external_id": "CVE-2014-6393",
      "title": "CVE-2014-6393",
      "summary": "The Express web framework before 3.11 and 4.x before 4.5 for Node.js does not provide a charset field in HTTP Content-Type headers in 400 level responses, which might allow remote attackers to conduct cross-site scripting (XSS) attacks via characters in a non-standard encoding.",
      "score": "6.1",
      "score_version": "3.0",
      "score_system": "CVSS",
      "score_details": {
        "cvssv2": {
          "version": "2.0",
          "vectorString": "AV:N/AC:M/Au:N/C:N/I:P/A:N",
          "accessVector": "NETWORK",
          "accessComplexity": "MEDIUM",
          "authentication": "NONE",
          "confidentialityImpact": "NONE",
          "integrityImpact": "PARTIAL",
          "availabilityImpact": "NONE",
          "baseScore": 4.3
        },
        "cvssv3": {
          "version": "3.0",
          "vectorString": "CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N",
          "attackVector": "NETWORK",
          "attackComplexity": "LOW",
          "privilegesRequired": "NONE",
          "userInteraction": "REQUIRED",
          "scope": "CHANGED",
          "confidentialityImpact": "LOW",
          "integrityImpact": "LOW",
          "availabilityImpact": "NONE",
          "baseScore": 6.1,
          "baseSeverity": "MEDIUM"
        }
      },
      "vector": null,
      "access_complexity": null,
      "vulnerability_authentication": null,
      "confidentiality_impact": null,
      "integrity_impact": null,
      "availability_impact": null,
      "vulnerability_source": null,
      "assessment_check": null,
      "scanner": null,
      "recommendation": null,
      "modified_at": "2017-08-18T12:18:00.000Z",
      "published_at": "2017-08-09T18:29:00.000Z",
      "created_at": "2018-03-09T03:56:29.547Z",
      "updated_at": "2018-12-19T08:56:08.194Z",
      "source": [
        {
          "id": 1,
          "name": "NVD",
          "description": "National Vulnerability Database",
          "created_at": "2017-02-09T20:18:35.385Z",
          "updated_at": "2017-02-13T20:12:05.342Z",
          "attribution": "Copyright © 1999–2017, The MITRE Corporation. CVE and the CVE logo are registered trademarks and CVE-Compatible is a trademark of The MITRE Corporation.",
          "license": "Submissions: For all materials you submit to the Common Vulnerabilities and Exposures (CVE®), you hereby grant to The MITRE Corporation (MITRE) and all CVE Numbering Authorities (CNAs) a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable copyright license to reproduce, prepare derivative works of, publicly display, publicly perform, sublicense, and distribute such materials and derivative works. Unless required by applicable law or agreed to in writing, you provide such materials on an \"AS IS\" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied, including, without limitation, any warranties or conditions of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A PARTICULAR PURPOSE.\n\nCVE Usage: MITRE hereby grants you a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable copyright license to reproduce, prepare derivative works of, publicly display, publicly perform, sublicense, and distribute Common Vulnerabilities and Exposures (CVE®). Any copy you make for such purposes is authorized provided that you reproduce MITRE's copyright designation and this license in any such copy.\n",
          "copyright_url": "http://cve.mitre.org/about/termsofuse.html"
        }
      ]
    }
  ],
  "meta": {
    "copyright": "Copyright 2018 Selection Pressure LLC www.selectpress.net",
    "authors": [
      "Ion Channel Dev Team"
    ],
    "version": "v1",
    "last_update": "2018-10-11T21:23:06.164Z",
    "total_count": 1,
    "limit": 10,
    "offset": 0
  },
  "links": {
    "self": "https://api.ionchannel.io/v1/vulnerability/getVulnerabilities?product=express&version=4.4.2"
  }
}
```

The response for express version 4.4.2 has one vulnerability with a medium score of 6.1.
